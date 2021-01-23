package uints_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uints"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(uints.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := uints.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uints.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uints.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uints.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uints.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint(0), uints.Val(nil))
	require.Equal(t, uint(0), uints.Val(uints.Ptr(0)))
	require.Equal(t, uint(1), uints.Val(uints.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint(1), uints.ValDef(nil, 1))
	require.Equal(t, uint(0), uints.ValDef(uints.Ptr(0), 1))
	require.Equal(t, uint(1), uints.ValDef(uints.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := uints.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, uint(10), v)
	v, err = uints.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, uint(Max), v)
	_, err = uints.ParseDec("")
	require.Error(t, err)
	_, err = uints.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := uints.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, uint(0x20), v)
	v, err = uints.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, uint(Max), v)
	v, err = uints.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, uint(Max), v)
	_, err = uints.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", uints.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", uints.StrHex(0x10))
	require.Equal(t, "a", uints.StrHex(0xA))
	require.Equal(t, "11", uints.StrHex(0x11))
}
