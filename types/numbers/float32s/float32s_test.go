package float32s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/float32s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := float32s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, float32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := float32s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = float32s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, float32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := float32s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = float32s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, float32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, float32(0), float32s.Val(nil))
	require.Equal(t, float32(0), float32s.Val(float32s.Ptr(0)))
	require.Equal(t, float32(1), float32s.Val(float32s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, float32(1), float32s.ValDef(nil, 1))
	require.Equal(t, float32(0), float32s.ValDef(float32s.Ptr(0), 1))
	require.Equal(t, float32(1), float32s.ValDef(float32s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := float32s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, float32(10), v)
	_, err = float32s.Parse("")
	require.Error(t, err)
	_, err = float32s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []float32{2, 0, 3, 1, 4}
	require.False(t, float32s.Slice(s).IsSorted())
	float32s.Slice(s).Sort()
	require.Equal(t, []float32{0, 1, 2, 3, 4}, s)
	require.True(t, float32s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[float32]struct{}{}, float32s.SliceToMap(nil))
	require.Equal(t, map[float32]struct{}{}, float32s.SliceToMap([]float32{}))
	require.Equal(t, map[float32]struct{}{1: {}}, float32s.SliceToMap([]float32{1}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32s.SliceToMap([]float32{1, 2}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32s.SliceToMap([]float32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []float32{}, float32s.MapToSlice(nil))
	require.Equal(t, []float32{}, float32s.MapToSlice(map[float32]struct{}{}))
	require.Equal(t, []float32{1}, float32s.MapToSlice(map[float32]struct{}{1: {}}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32s.SliceToMap(float32s.MapToSlice(map[float32]struct{}{1: {}, 2: {}})))
}

func TestSwapMap(t *testing.T) {
	swap, err := float32s.SwapMap(map[float32]float32{
		1: 2,
		3: 4,
	})
	require.NoError(t, err)
	require.Equal(t,
		map[float32]float32{
			2: 1,
			4: 3,
		}, swap)

	swap, err = float32s.SwapMap(map[float32]float32{})
	require.NoError(t, err)
	require.Equal(t, map[float32]float32{}, swap)

	swap, err = float32s.SwapMap(nil)
	require.NoError(t, err)
	require.Equal(t, map[float32]float32{}, swap)

	_, err = float32s.SwapMap(map[float32]float32{
		1: 3,
		2: 3,
	})
	require.EqualError(t, err, "duplicate value: 3")
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, float32(0), float32s.SafeIndex(nil, 0))
	require.Equal(t, float32(0), float32s.SafeIndex(nil, 1))
	require.Equal(t, float32(0), float32s.SafeIndex(nil, -1))
	require.Equal(t, float32(0), float32s.SafeIndex([]float32{}, 0))
	require.Equal(t, float32(0), float32s.SafeIndex([]float32{}, 1))
	require.Equal(t, float32(0), float32s.SafeIndex([]float32{}, -1))
	require.Equal(t, float32(1), float32s.SafeIndex([]float32{1}, 0))
	require.Equal(t, float32(0), float32s.SafeIndex([]float32{1}, 1))
	require.Equal(t, float32(0), float32s.SafeIndex([]float32{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, float32s.SafeIndexPtr(nil, 0))
	require.Nil(t, float32s.SafeIndexPtr(nil, 1))
	require.Nil(t, float32s.SafeIndexPtr(nil, -1))
	require.Nil(t, float32s.SafeIndexPtr([]float32{}, 0))
	require.Nil(t, float32s.SafeIndexPtr([]float32{}, 1))
	require.Nil(t, float32s.SafeIndexPtr([]float32{}, -1))
	require.Equal(t, float32s.Ptr(1), float32s.SafeIndexPtr([]float32{1}, 0))
	require.Nil(t, float32s.SafeIndexPtr([]float32{1}, 1))
	require.Nil(t, float32s.SafeIndexPtr([]float32{1}, -1))
}
