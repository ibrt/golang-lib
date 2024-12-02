package filez_test

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
)

type MustSuite struct {
	// intentionally empty
}

func TestMustSuite(t *testing.T) {
	fixturez.RunSuite(t, &MustSuite{})
}

func (*MustSuite) TestMustAbs(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustAbs("path")).ToNot(Equal("path"))
	}).ToNot(Panic())
}

func (*MustSuite) TestMustRel(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustRel(filepath.Join("a", "b", "c"), filepath.Join("a", "b"))).To(Equal(".."))
		g.Expect(filez.MustRel(filepath.Join("a"), filepath.Join("a", "b"))).To(Equal("b"))
		g.Expect(filez.MustRel("", filepath.Join("a", "b"))).To(Equal(filepath.Join("a", "b")))
		g.Expect(filez.MustRel(filepath.Join("a", "b"), "")).To(Equal(filepath.Join("..", "..")))
	}).ToNot(Panic())
}

func (s *MustSuite) TestGetwdAndChdir(g *WithT) {
	g.Expect(func() {
		wd1 := filez.MustGetwd()
		defer filez.MustChdir(wd1)

		wd2 := filepath.Dir(wd1)
		filez.MustChdir(wd2)
		g.Expect(filez.MustGetwd()).To(Equal(wd2))
	}).ToNot(Panic())
}

func (*MustSuite) TestMustUserHomeDir(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustUserHomeDir()).ToNot(BeEmpty())
	}).ToNot(Panic())
}
