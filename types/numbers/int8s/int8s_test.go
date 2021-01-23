package int8s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int8s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(int8s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := int8s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int8s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int8s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int8s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int8s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int8(0), int8s.Val(nil))
	require.Equal(t, int8(0), int8s.Val(int8s.Ptr(0)))
	require.Equal(t, int8(1), int8s.Val(int8s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int8(1), int8s.ValDef(nil, 1))
	require.Equal(t, int8(0), int8s.ValDef(int8s.Ptr(0), 1))
	require.Equal(t, int8(1), int8s.ValDef(int8s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := int8s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, int8(10), v)
	v, err = int8s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, int8(Max), v)
	_, err = int8s.ParseDec("")
	require.Error(t, err)
	_, err = int8s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := int8s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, int8(0x20), v)
	v, err = int8s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, int8(Max), v)
	v, err = int8s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, int8(Max), v)
	_, err = int8s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", int8s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", int8s.StrHex(0x10))
	require.Equal(t, "a", int8s.StrHex(0xA))
	require.Equal(t, "11", int8s.StrHex(0x11))
}
