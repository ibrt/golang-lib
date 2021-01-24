package uint16s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint16s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint16s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint16(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint16s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint16s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint16(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint16s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint16s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint16(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint16(0), uint16s.Val(nil))
	require.Equal(t, uint16(0), uint16s.Val(uint16s.Ptr(0)))
	require.Equal(t, uint16(1), uint16s.Val(uint16s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint16(1), uint16s.ValDef(nil, 1))
	require.Equal(t, uint16(0), uint16s.ValDef(uint16s.Ptr(0), 1))
	require.Equal(t, uint16(1), uint16s.ValDef(uint16s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint16s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint16(10), v)
	_, err = uint16s.Parse("")
	require.Error(t, err)
	_, err = uint16s.Parse("A")
	require.Error(t, err)
}
