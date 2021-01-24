package int64s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int64s"
	"github.com/stretchr/testify/require"
)

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

func TestParse(t *testing.T) {
	v, err := int64s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int64(10), v)
	_, err = int64s.Parse("")
	require.Error(t, err)
	_, err = int64s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int64{2, 0, 3, 1, 4}
	require.False(t, int64s.Slice(s).IsSorted())
	int64s.Slice(s).Sort()
	require.Equal(t, []int64{0, 1, 2, 3, 4}, s)
	require.True(t, int64s.Slice(s).IsSorted())
}
