package stringz_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/stringz"
	"github.com/stretchr/testify/require"
)

func TestCollector(t *testing.T) {
	c := stringz.NewCollector()
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.Add("")
	require.Equal(t, []string{""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.Add("")
	require.Equal(t, []string{"", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.AddPtr(stringz.Ptr(""))
	require.Equal(t, []string{"", "", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.AddPtr(nil)
	require.Equal(t, []string{"", "", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c = stringz.NewCollector(stringz.InitialCap(10))
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, 10, cap(c.Slice()))
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.Add("")
	require.Equal(t, []string{""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.Add("")
	require.Equal(t, []string{"", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.AddPtr(stringz.Ptr(""))
	require.Equal(t, []string{"", "", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.AddPtr(nil)
	require.Equal(t, []string{"", "", ""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c = stringz.NewCollector(stringz.SkipDuplicates)
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.Add("")
	require.Equal(t, []string{""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.Add("")
	require.Equal(t, []string{""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c.AddPtr(stringz.Ptr(""))
	require.Equal(t, []string{""}, c.Slice())
	require.Equal(t, map[string]struct{}{"": {}}, c.Map())

	c = stringz.NewCollector(stringz.SkipEmpties)
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.Add("")
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.AddPtr(stringz.Ptr(""))
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c = stringz.NewCollector()
	require.Equal(t, []string{}, c.Slice())
	require.Equal(t, map[string]struct{}{}, c.Map())

	c.Add("a", "b", "c")
	require.Equal(t, []string{"a", "b", "c"}, c.Slice())
	require.Equal(t, "a, b, c", c.Join(", "))

}
