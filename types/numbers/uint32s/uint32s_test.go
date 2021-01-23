package uint32s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint32s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(uint32s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := uint32s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint32s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint32s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint32s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint32s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint32(0), uint32s.Val(nil))
	require.Equal(t, uint32(0), uint32s.Val(uint32s.Ptr(0)))
	require.Equal(t, uint32(1), uint32s.Val(uint32s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint32(1), uint32s.ValDef(nil, 1))
	require.Equal(t, uint32(0), uint32s.ValDef(uint32s.Ptr(0), 1))
	require.Equal(t, uint32(1), uint32s.ValDef(uint32s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := uint32s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, uint32(10), v)
	v, err = uint32s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, uint32(Max), v)
	_, err = uint32s.ParseDec("")
	require.Error(t, err)
	_, err = uint32s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := uint32s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, uint32(0x20), v)
	v, err = uint32s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, uint32(Max), v)
	v, err = uint32s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, uint32(Max), v)
	_, err = uint32s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", uint32s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", uint32s.StrHex(0x10))
	require.Equal(t, "a", uint32s.StrHex(0xA))
	require.Equal(t, "11", uint32s.StrHex(0x11))
}
