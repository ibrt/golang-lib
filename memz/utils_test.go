package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/memz"
)

type UtilsSuite struct {
	// intentionally empty
}

func TestUtilsSuite(t *testing.T) {
	fixturez.RunSuite(t, &UtilsSuite{})
}

func (s *UtilsSuite) TestMin(g *WithT) {
	g.Expect(memz.Min(1, 2)).To(BeNumerically("==", 1))
	g.Expect(memz.Min(2, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Min(3, 2)).To(BeNumerically("==", 2))
}

func (s *UtilsSuite) TestMax(g *WithT) {
	g.Expect(memz.Max(1, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Max(2, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Max(3, 2)).To(BeNumerically("==", 3))
}

func (s *UtilsSuite) TestPredicateIsZeroValue(g *WithT) {
	g.Expect(memz.FilterSlice([]int{0, 1, 2}, memz.PredicateIsZeroValue[int])).To(Equal([]int{0}))
	g.Expect(memz.FilterSlice([]string{"", "1", "2"}, memz.PredicateIsZeroValue[string])).To(Equal([]string{""}))
}

func (s *UtilsSuite) TestTransformSprintf(g *WithT) {
	g.Expect(memz.TransformSprintf(1)).To(Equal("1"))
}
