package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/memz"
)

type PointersSuite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &PointersSuite{})
}

func (*PointersSuite) TestPtr(g *WithT) {
	v := ""
	g.Expect(memz.Ptr(v)).To(Equal(&v))

	v = "a"
	g.Expect(memz.Ptr(v)).To(Equal(&v))
}

func (*PointersSuite) TestPtrIfTrue(g *WithT) {
	v := "a"
	g.Expect(memz.PtrIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrIfTrue(false, v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrIfTrue(false, v)).To(BeNil())
}

func (*PointersSuite) TestPtrZeroToNil(g *WithT) {
	v := ""
	g.Expect(memz.PtrZeroToNil(v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrZeroToNil(v)).To(Equal(&v))
}

func (*PointersSuite) TestPtrZeroToNilIfTrue(g *WithT) {
	v := ""
	g.Expect(memz.PtrZeroToNilIfTrue(true, v)).To(BeNil())
	g.Expect(memz.PtrZeroToNilIfTrue(false, v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrZeroToNilIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrZeroToNilIfTrue(false, v)).To(BeNil())
}

func (*PointersSuite) TestValNilToZero(g *WithT) {
	g.Expect(memz.ValNilToZero(memz.Ptr(""))).To(Equal(""))
	g.Expect(memz.ValNilToZero(memz.Ptr("a"))).To(Equal("a"))
	g.Expect(memz.ValNilToZero[string](nil)).To(BeEmpty())
}

func (*PointersSuite) TestValNilToDef(g *WithT) {
	g.Expect(memz.ValNilToDef(memz.Ptr(""), "d")).To(Equal(""))
	g.Expect(memz.ValNilToDef(memz.Ptr("a"), "d")).To(Equal("a"))
	g.Expect(memz.ValNilToDef(nil, "d")).To(Equal("d"))
}
