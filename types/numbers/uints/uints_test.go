package uints_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uints"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uints.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uints.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uints.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uints.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uints.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint(0), uints.Val(nil))
	require.Equal(t, uint(0), uints.Val(uints.Ptr(0)))
	require.Equal(t, uint(1), uints.Val(uints.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint(1), uints.ValDef(nil, 1))
	require.Equal(t, uint(0), uints.ValDef(uints.Ptr(0), 1))
	require.Equal(t, uint(1), uints.ValDef(uints.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uints.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint(10), v)
	_, err = uints.Parse("")
	require.Error(t, err)
	_, err = uints.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []uint{2, 0, 3, 1, 4}
	require.False(t, uints.Slice(s).IsSorted())
	uints.Slice(s).Sort()
	require.Equal(t, []uint{0, 1, 2, 3, 4}, s)
	require.True(t, uints.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint]struct{}{}, uints.SliceToMap(nil))
	require.Equal(t, map[uint]struct{}{}, uints.SliceToMap([]uint{}))
	require.Equal(t, map[uint]struct{}{1: {}}, uints.SliceToMap([]uint{1}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uints.SliceToMap([]uint{1, 2}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uints.SliceToMap([]uint{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint{}, uints.MapToSlice(nil))
	require.Equal(t, []uint{}, uints.MapToSlice(map[uint]struct{}{}))
	require.Equal(t, []uint{1}, uints.MapToSlice(map[uint]struct{}{1: {}}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uints.SliceToMap(uints.MapToSlice(map[uint]struct{}{1: {}, 2: {}})))
}
