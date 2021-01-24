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
