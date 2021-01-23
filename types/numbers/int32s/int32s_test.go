package int32s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int32s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(int32s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := int32s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int32s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int32s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int32s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int32s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int32(0), int32s.Val(nil))
	require.Equal(t, int32(0), int32s.Val(int32s.Ptr(0)))
	require.Equal(t, int32(1), int32s.Val(int32s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int32(1), int32s.ValDef(nil, 1))
	require.Equal(t, int32(0), int32s.ValDef(int32s.Ptr(0), 1))
	require.Equal(t, int32(1), int32s.ValDef(int32s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := int32s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, int32(10), v)
	v, err = int32s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, int32(Max), v)
	_, err = int32s.ParseDec("")
	require.Error(t, err)
	_, err = int32s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := int32s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, int32(0x20), v)
	v, err = int32s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, int32(Max), v)
	v, err = int32s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, int32(Max), v)
	_, err = int32s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", int32s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", int32s.StrHex(0x10))
	require.Equal(t, "a", int32s.StrHex(0xA))
	require.Equal(t, "11", int32s.StrHex(0x11))
}
