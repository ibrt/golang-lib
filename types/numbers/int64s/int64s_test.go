package int64s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int64s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(int64s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := int64s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int64s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int64s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int64s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int64s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int64(0), int64s.Val(nil))
	require.Equal(t, int64(0), int64s.Val(int64s.Ptr(0)))
	require.Equal(t, int64(1), int64s.Val(int64s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int64(1), int64s.ValDef(nil, 1))
	require.Equal(t, int64(0), int64s.ValDef(int64s.Ptr(0), 1))
	require.Equal(t, int64(1), int64s.ValDef(int64s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := int64s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, int64(10), v)
	v, err = int64s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, int64(Max), v)
	_, err = int64s.ParseDec("")
	require.Error(t, err)
	_, err = int64s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := int64s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, int64(0x20), v)
	v, err = int64s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, int64(Max), v)
	v, err = int64s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, int64(Max), v)
	_, err = int64s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", int64s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", int64s.StrHex(0x10))
	require.Equal(t, "a", int64s.StrHex(0xA))
	require.Equal(t, "11", int64s.StrHex(0x11))
}
