package fixtures_test

import (
	"testing"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/fixtures"
	"github.com/stretchr/testify/require"
)

type fakeT struct {
	isHelper    bool
	logs        [][]interface{}
	isFailed    bool
	isFailedNow bool
}

func (t *fakeT) Helper() {
	t.isHelper = true
}

func (t *fakeT) Log(args ...interface{}) {
	t.logs = append(t.logs, args)
}

func (t *fakeT) Fail() {
	t.isFailed = true
}

func (t *fakeT) FailNow() {
	t.isFailedNow = true
}

func TestAssertNoError(t *testing.T) {
	ft := &fakeT{}
	fixtures.RequireNoError(ft, nil)
	require.True(t, ft.isHelper)
	require.Len(t, ft.logs, 0)
	require.False(t, ft.isFailed)
	require.False(t, ft.isFailedNow)

	ft = &fakeT{}
	fixtures.RequireNoError(ft, errors.Errorf("test error"))
	require.True(t, ft.isHelper)
	require.Len(t, ft.logs, 1)
	require.False(t, ft.isFailed)
	require.True(t, ft.isFailedNow)

	ft = &fakeT{}
	fixtures.AssertNoError(ft, nil)
	require.True(t, ft.isHelper)
	require.Len(t, ft.logs, 0)
	require.False(t, ft.isFailed)
	require.False(t, ft.isFailedNow)

	ft = &fakeT{}
	fixtures.AssertNoError(ft, errors.Errorf("test error"))
	require.True(t, ft.isHelper)
	require.Len(t, ft.logs, 1)
	require.True(t, ft.isFailed)
	require.False(t, ft.isFailedNow)
}
