package ints_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/ints"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := ints.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := ints.PtrZeroToNil(0)
	require.Nil(t, p)
	p = ints.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := ints.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = ints.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int(0), ints.Val(nil))
	require.Equal(t, int(0), ints.Val(ints.Ptr(0)))
	require.Equal(t, int(1), ints.Val(ints.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int(1), ints.ValDef(nil, 1))
	require.Equal(t, int(0), ints.ValDef(ints.Ptr(0), 1))
	require.Equal(t, int(1), ints.ValDef(ints.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := ints.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int(10), v)
	_, err = ints.Parse("")
	require.Error(t, err)
	_, err = ints.Parse("A")
	require.Error(t, err)
}
