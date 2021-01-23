package ints_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/ints"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(ints.BitSize-1) - 1

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

func TestParseDec(t *testing.T) {
	v, err := ints.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, int(10), v)
	v, err = ints.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, int(Max), v)
	_, err = ints.ParseDec("")
	require.Error(t, err)
	_, err = ints.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := ints.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, int(0x20), v)
	v, err = ints.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, int(Max), v)
	v, err = ints.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, int(Max), v)
	_, err = ints.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", ints.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", ints.StrHex(0x10))
	require.Equal(t, "a", ints.StrHex(0xA))
	require.Equal(t, "11", ints.StrHex(0x11))
}
