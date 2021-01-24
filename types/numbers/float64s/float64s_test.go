package float64s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/float64s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := float64s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, float64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := float64s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = float64s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, float64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := float64s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = float64s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, float64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, float64(0), float64s.Val(nil))
	require.Equal(t, float64(0), float64s.Val(float64s.Ptr(0)))
	require.Equal(t, float64(1), float64s.Val(float64s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, float64(1), float64s.ValDef(nil, 1))
	require.Equal(t, float64(0), float64s.ValDef(float64s.Ptr(0), 1))
	require.Equal(t, float64(1), float64s.ValDef(float64s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := float64s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, float64(10), v)
	_, err = float64s.Parse("")
	require.Error(t, err)
	_, err = float64s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []float64{2, 0, 3, 1, 4}
	require.False(t, float64s.Slice(s).IsSorted())
	float64s.Slice(s).Sort()
	require.Equal(t, []float64{0, 1, 2, 3, 4}, s)
	require.True(t, float64s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[float64]struct{}{}, float64s.SliceToMap(nil))
	require.Equal(t, map[float64]struct{}{}, float64s.SliceToMap([]float64{}))
	require.Equal(t, map[float64]struct{}{1: {}}, float64s.SliceToMap([]float64{1}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64s.SliceToMap([]float64{1, 2}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64s.SliceToMap([]float64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []float64{}, float64s.MapToSlice(nil))
	require.Equal(t, []float64{}, float64s.MapToSlice(map[float64]struct{}{}))
	require.Equal(t, []float64{1}, float64s.MapToSlice(map[float64]struct{}{1: {}}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64s.SliceToMap(float64s.MapToSlice(map[float64]struct{}{1: {}, 2: {}})))
}
