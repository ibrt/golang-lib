package inject_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/inject/internal/injectmocks"
	"github.com/stretchr/testify/require"
)

func TestMustInject_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInjector := injectmocks.NewMockTestInjector(ctrl)
	mockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(context.Background(), nil)

	ctx := inject.MustInject(context.Background(), mockInjector.Inject)
	require.Equal(t, context.Background(), ctx)
}

func TestMustInject_Panic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInjector := injectmocks.NewMockTestInjector(ctrl)
	mockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(nil, errors.Errorf("test"))

	require.PanicsWithError(t, "test", func() {
		inject.MustInject(context.Background(), mockInjector.Inject)
	})
}

func TestSingletonInjectorFactory(t *testing.T) {
	type contextKeyType int
	const contextKey contextKeyType = 0

	injector := inject.SingletonInjectorFactory(contextKey, "value")

	ctx, err := injector(context.Background())
	fixtures.RequireNoError(t, err)
	require.Equal(t, "value", ctx.Value(contextKey))
}

func TestCompoundInjectorFactory_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	firstMockInjector := injectmocks.NewMockTestInjector(ctrl)
	firstMockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(context.Background(), nil)
	secondMockInjector := injectmocks.NewMockTestInjector(ctrl)
	secondMockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(nil, errors.Errorf("test"))
	thirdMockInjector := injectmocks.NewMockTestInjector(ctrl)

	injector := inject.CompoundInjectorFactory(firstMockInjector.Inject, secondMockInjector.Inject, thirdMockInjector.Inject)
	ctx, err := injector(context.Background())
	require.EqualError(t, err, "test")
	require.Nil(t, ctx)
}

func TestInjectorMiddlewareFactory_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInjector := injectmocks.NewMockTestInjector(ctrl)
	mockInjector.EXPECT().Inject(gomock.Any()).Return(context.Background(), nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	middleware := inject.InjectorMiddlewareFactory(mockInjector.Inject)

	srv := httptest.NewServer(middleware(handler))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	fixtures.RequireNoError(t, err)
	defer errors.IgnoreClose(resp.Body)
	require.EqualValues(t, http.StatusOK, resp.StatusCode)
}
