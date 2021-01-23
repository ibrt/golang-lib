package uint64s_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint64s"
	"github.com/stretchr/testify/require"
)

const Max = 1<<(uint64s.BitSize-1) - 1

func TestPtr(t *testing.T) {
	p := uint64s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint64s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint64s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint64s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint64s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint64(0), uint64s.Val(nil))
	require.Equal(t, uint64(0), uint64s.Val(uint64s.Ptr(0)))
	require.Equal(t, uint64(1), uint64s.Val(uint64s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint64(1), uint64s.ValDef(nil, 1))
	require.Equal(t, uint64(0), uint64s.ValDef(uint64s.Ptr(0), 1))
	require.Equal(t, uint64(1), uint64s.ValDef(uint64s.Ptr(1), 1))
}

func TestParseDec(t *testing.T) {
	v, err := uint64s.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, uint64(10), v)
	v, err = uint64s.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, uint64(Max), v)
	_, err = uint64s.ParseDec("")
	require.Error(t, err)
	_, err = uint64s.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := uint64s.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, uint64(0x20), v)
	v, err = uint64s.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, uint64(Max), v)
	v, err = uint64s.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, uint64(Max), v)
	_, err = uint64s.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", uint64s.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", uint64s.StrHex(0x10))
	require.Equal(t, "a", uint64s.StrHex(0xA))
	require.Equal(t, "11", uint64s.StrHex(0x11))
}
