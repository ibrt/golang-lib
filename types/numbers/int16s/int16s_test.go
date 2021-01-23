package int16s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int16s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(int16s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := int16s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int16(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int16s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int16s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int16s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int16s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int16(0), int16s.Val(nil))
	require.Equal(t, int16(0), int16s.Val(int16s.Ptr(0)))
	require.Equal(t, int16(1), int16s.Val(int16s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int16(1), int16s.ValDef(nil, 1))
	require.Equal(t, int16(0), int16s.ValDef(int16s.Ptr(0), 1))
	require.Equal(t, int16(1), int16s.ValDef(int16s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := int16s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, int16(10), v)
	v, err = int16s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, int16(Max), v)
	_, err = int16s.ParseDec("")
	require.Error(t, err)
	_, err = int16s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := int16s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, int16(0x20), v)
	v, err = int16s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, int16(Max), v)
	v, err = int16s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, int16(Max), v)
	_, err = int16s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", int16s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", int16s.StrHex(0x10))
	require.Equal(t, "a", int16s.StrHex(0xA))
	require.Equal(t, "11", int16s.StrHex(0x11))
}
