package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/memz"
)

type MapsSuite struct {
	// intentionally empty
}

func TestMapsSuite(t *testing.T) {
	fixturez.RunSuite(t, &MapsSuite{})
}

func (*MapsSuite) TestMergeMaps(g *WithT) {
	g.Expect(memz.MergeMaps[int, string]()).To(Equal(map[int]string{}))
	g.Expect(memz.MergeMaps[int, string](nil, nil)).To(Equal(map[int]string{}))
	g.Expect(memz.MergeMaps(nil, map[int]string{})).To(Equal(map[int]string{}))
	g.Expect(memz.MergeMaps(map[int]string{}, nil)).To(Equal(map[int]string{}))
	g.Expect(memz.MergeMaps(map[int]string{}, map[int]string{})).To(Equal(map[int]string{}))
	g.Expect(memz.MergeMaps(map[int]string{1: "a", 2: "b"}, map[int]string{1: "x", 3: "c"})).To(Equal(map[int]string{1: "x", 2: "b", 3: "c"}))
	g.Expect(memz.MergeMaps(map[int]string{1: "a", 2: "b"}, nil)).To(Equal(map[int]string{1: "a", 2: "b"}))
}

func (*MapsSuite) TestCopyMap(g *WithT) {
	g.Expect(memz.ShallowCopyMap[int, string](nil)).To(BeNil())
	g.Expect(memz.ShallowCopyMap(map[int]string{})).To(Equal(map[int]string{}))
	g.Expect(memz.ShallowCopyMap(map[int]string{1: "a", 2: "b"})).To(Equal(map[int]string{1: "a", 2: "b"}))
}

func (*MapsSuite) TestFilterMap(g *WithT) {
	g.Expect(memz.FilterMap[int, string](nil, func(_ int, _ string) bool { return true })).To(BeNil())
	g.Expect(memz.FilterMap[int, string](map[int]string{}, func(_ int, _ string) bool { return true })).To(Equal(map[int]string{}))
	g.Expect(memz.FilterMap[int, string](map[int]string{1: "a", 2: "b"}, func(k int, _ string) bool { return k == 1 })).To(Equal(map[int]string{1: "a"}))
}
