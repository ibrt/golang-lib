package fixtures

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/ibrt/golang-lib/errors"
)

// BeforeSuite describes a method invoked before starting a test suite.
type BeforeSuite interface {
	BeforeSuite(context.Context) context.Context
}

// AfterSuite represents a method invoked after completing a test suite.
type AfterSuite interface {
	AfterSuite(context.Context)
}

// BeforeTest represents a method invoked before each test method in a suite.
type BeforeTest interface {
	BeforeTest(context.Context, *testing.T) context.Context
}

// AfterTest represents a method invoked after each test method in a suite.
type AfterTest interface {
	AfterTest(context.Context)
}

// RunSuite runs a test suite.
func RunSuite(t *testing.T, s interface{}) {
	sV := reflect.ValueOf(s)
	sT := sV.Type()
	sVE := reflect.Indirect(sV)
	rootCtx := context.Background()

	// before suite
	for i := 0; i < sVE.NumField(); i++ {
		fV := sVE.Field(i)

		switch fV.Interface().(type) {
		case BeforeSuite, AfterSuite, BeforeTest, AfterTest:
			fV.Set(reflect.New(fV.Type().Elem()))
		}

		if f, ok := fV.Interface().(BeforeSuite); ok {
			rootCtx = f.BeforeSuite(rootCtx)
		}
	}

	for i := 0; i < sT.NumMethod(); i++ {
		mT := sT.Method(i)

		if strings.HasPrefix(mT.Name, "Test") {
			t.Run(mT.Name, func(t *testing.T) {
				t.Helper()
				testCtx := rootCtx

				fmt.Printf(">> %v\n", mT.Name)
				defer fmt.Printf("<< %v\n", mT.Name)

				// before test
				for j := 0; j < sVE.NumField(); j++ {
					if f, ok := sVE.Field(j).Interface().(BeforeTest); ok {
						testCtx = f.BeforeTest(testCtx, t)
					}
				}

				// test
				func() {
					defer func() {
						if err := errors.MaybeWrapRecover(recover(), errors.Prefix("panic"), errors.Skip()); err != nil {
							RequireNoError(t, err)
						}
					}()
					sV.Method(i).Call([]reflect.Value{reflect.ValueOf(testCtx), reflect.ValueOf(t)})
				}()

				// after test
				for j := 0; j < sVE.NumField(); j++ {
					if f, ok := sVE.Field(j).Interface().(AfterTest); ok {
						f.AfterTest(testCtx)
					}
				}
			})
		}
	}

	// after suite
	for i := 0; i < sVE.NumField(); i++ {
		if f, ok := sVE.Field(i).Interface().(AfterSuite); ok {
			f.AfterSuite(rootCtx)
		}
	}
}
