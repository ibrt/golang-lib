package runners_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/runners"
	"github.com/ibrt/golang-lib/runners/internal/runnersmocks"
	"github.com/stretchr/testify/require"
)

func TestSafe_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSafe := runnersmocks.NewMockTestFunc(ctrl)
	mockSafe.EXPECT().Func().Return(nil)

	require.NotPanics(t, func() {
		fixtures.RequireNoError(t, runners.Safe(mockSafe.Func)())
	})
}

func TestSafe_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSafe := runnersmocks.NewMockTestFunc(ctrl)
	mockSafe.EXPECT().Func().Return(errors.Errorf("test"))

	require.NotPanics(t, func() {
		require.EqualError(t, runners.Safe(mockSafe.Func)(), "test")
	})
}

func TestSafe_Panic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSafe := runnersmocks.NewMockTestFunc(ctrl)
	mockSafe.EXPECT().Func().DoAndReturn(func() error {
		errors.MustErrorf("test")
		return nil
	})

	require.NotPanics(t, func() {
		require.EqualError(t, runners.Safe(mockSafe.Func)(), "test")
	})
}

func TestParallel_Unlimited(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	start := time.Now()

	mockSafe := runnersmocks.NewMockTestFunc(ctrl)
	mockSafe.EXPECT().Func().Times(5).DoAndReturn(func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	})

	fixtures.RequireNoError(t, runners.Parallel(0, mockSafe.Func, mockSafe.Func, mockSafe.Func, mockSafe.Func, mockSafe.Func))
	require.Less(t, time.Since(start), 150*time.Millisecond)
}

func TestParallel_Limited(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	start := time.Now()

	mockSafe := runnersmocks.NewMockTestFunc(ctrl)
	mockSafe.EXPECT().Func().Times(5).DoAndReturn(func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	})

	fixtures.RequireNoError(t, runners.Parallel(2, mockSafe.Func, mockSafe.Func, mockSafe.Func, mockSafe.Func, mockSafe.Func))
	require.Less(t, time.Since(start), 350*time.Millisecond)
}
