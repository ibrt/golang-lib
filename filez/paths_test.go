package filez_test

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
)

type PathsSuite struct {
	// intentionally empty
}

func TestPathsSuite(t *testing.T) {
	fixturez.RunSuite(t, &PathsSuite{})
}

func (*PathsSuite) TestMustIsChild(g *WithT) {
	g.Expect(filez.MustIsChild("a", "a")).To(BeTrue())
	g.Expect(filez.MustIsChild("", "a")).To(BeTrue())
	g.Expect(filez.MustIsChild("a", filepath.Join("a", "b"))).To(BeTrue())
	g.Expect(filez.MustIsChild("a", "")).To(BeFalse())
	g.Expect(filez.MustIsChild(filepath.Join("a", "b"), "a")).To(BeFalse())
	g.Expect(filez.MustIsChild("a", "b")).To(BeFalse())
}

func (*PathsSuite) TestMustRelForDisplay(g *WithT) {
	g.Expect(filez.MustRelForDisplay("a")).To(Equal("a"))
	g.Expect(filez.MustRelForDisplay(filez.MustAbs("a"))).To(Equal("a"))
	g.Expect(filez.MustRelForDisplay(filez.MustAbs(filepath.Join("..", "a")))).To(Equal(filez.MustAbs(filepath.Join("..", "a"))))
}
