package inject_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/inject/internal/injectmocks"
	"github.com/stretchr/testify/require"
)

func TestMustInitialize_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	mockInjector := injectmocks.NewMockTestInjector(ctrl)
	mockReleaser := injectmocks.NewMockTestReleaser(ctrl)

	mockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(mockInjector.Inject, mockReleaser.Release, nil)
	mockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(context.Background(), nil)
	mockReleaser.EXPECT().Release()

	injector, releaser := inject.MustInitialize(context.Background(), mockInitializer.Initialize)
	_, _ = injector(context.Background())
	releaser()
}

func TestMustInitialize_Panic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	mockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(nil, nil, errors.Errorf("test"))

	require.PanicsWithError(t, "test", func() {
		inject.MustInitialize(context.Background(), mockInitializer.Initialize)
	})
}

func TestBootstrap_NoInitializers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	injector, releaser, err := inject.Bootstrap(context.Background())
	fixtures.RequireNoError(t, err)

	ctx, err := injector(context.Background())
	fixtures.RequireNoError(t, err)
	require.Equal(t, context.Background(), ctx)
	require.NotPanics(t, func() { releaser() })
}

func TestBootstrap_MultipleInitializers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	firstMockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	firstMockInjector := injectmocks.NewMockTestInjector(ctrl)
	firstMockReleaser := injectmocks.NewMockTestReleaser(ctrl)

	firstMockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(firstMockInjector.Inject, firstMockReleaser.Release, nil)
	firstMockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(context.Background(), nil)
	firstMockReleaser.EXPECT().Release()

	secondMockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	secondMockInjector := injectmocks.NewMockTestInjector(ctrl)
	secondMockReleaser := injectmocks.NewMockTestReleaser(ctrl)

	secondMockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(secondMockInjector.Inject, secondMockReleaser.Release, nil)
	secondMockInjector.EXPECT().Inject(gomock.Eq(context.Background())).Return(context.Background(), nil)
	secondMockReleaser.EXPECT().Release()

	injector, releaser, err := inject.Bootstrap(context.Background(), firstMockInitializer.Initialize, secondMockInitializer.Initialize)
	fixtures.RequireNoError(t, err)

	ctx, err := injector(context.Background())
	fixtures.RequireNoError(t, err)
	require.Equal(t, context.Background(), ctx)
	require.NotPanics(t, func() { releaser() })
}

func TestBootstrap_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	firstMockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	firstMockInjector := injectmocks.NewMockTestInjector(ctrl)
	firstMockReleaser := injectmocks.NewMockTestReleaser(ctrl)

	firstMockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(firstMockInjector.Inject, firstMockReleaser.Release, nil)
	firstMockReleaser.EXPECT().Release()

	secondMockInitializer := injectmocks.NewMockTestInitializer(ctrl)
	secondMockInitializer.EXPECT().Initialize(gomock.Eq(context.Background())).Return(nil, nil, errors.Errorf("test"))

	injector, releaser, err := inject.Bootstrap(context.Background(), firstMockInitializer.Initialize, secondMockInitializer.Initialize)
	require.EqualError(t, err, "test")
	require.Nil(t, injector)
	require.NotPanics(t, func() { releaser() })
}
