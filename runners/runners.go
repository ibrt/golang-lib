package runners

import (
	"github.com/ibrt/golang-lib/errors"
	"go4.org/syncutil"
)

// Func describes a runnable function.
type Func func() error

// Safe runs the given Func, catching any panics and returning them as errors.
func Safe(f Func) Func {
	return func() (err error) {
		defer func() {
			if rErr := errors.MaybeWrapRecover(recover()); rErr != nil {
				err = rErr
			}
		}()

		return f()
	}
}

// Parallel runs the given functions in parallel, using the given maxConcurrency if > 0.
func Parallel(maxConcurrency int, fs ...func() error) error {
	group := &syncutil.Group{}
	var gate *syncutil.Gate

	if maxConcurrency > 0 {
		gate = syncutil.NewGate(maxConcurrency)
	}

	for _, f := range fs {
		f := f
		if maxConcurrency > 0 {
			gate.Start()
		}
		group.Go(func() error {
			if maxConcurrency > 0 {
				defer gate.Done()
			}
			return Safe(f)()
		})
	}

	return errors.MaybeWrap(group.Err())
}
