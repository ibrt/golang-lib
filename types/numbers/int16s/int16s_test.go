package int16s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int16s"
	"github.com/stretchr/testify/require"
)

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

func TestParse(t *testing.T) {
	v, err := int16s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int16(10), v)
	_, err = int16s.Parse("")
	require.Error(t, err)
	_, err = int16s.Parse("A")
	require.Error(t, err)
}
