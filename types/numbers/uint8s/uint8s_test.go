package uint8s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint8s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(uint8s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := uint8s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint8s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint8s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint8s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint8s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint8(0), uint8s.Val(nil))
	require.Equal(t, uint8(0), uint8s.Val(uint8s.Ptr(0)))
	require.Equal(t, uint8(1), uint8s.Val(uint8s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint8(1), uint8s.ValDef(nil, 1))
	require.Equal(t, uint8(0), uint8s.ValDef(uint8s.Ptr(0), 1))
	require.Equal(t, uint8(1), uint8s.ValDef(uint8s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := uint8s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, uint8(10), v)
	v, err = uint8s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, uint8(Max), v)
	_, err = uint8s.ParseDec("")
	require.Error(t, err)
	_, err = uint8s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := uint8s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, uint8(0x20), v)
	v, err = uint8s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, uint8(Max), v)
	v, err = uint8s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, uint8(Max), v)
	_, err = uint8s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", uint8s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", uint8s.StrHex(0x10))
	require.Equal(t, "a", uint8s.StrHex(0xA))
	require.Equal(t, "11", uint8s.StrHex(0x11))
}
