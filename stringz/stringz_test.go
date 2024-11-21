package stringz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/stringz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestAlignRight(g *WithT) {
	g.Expect(stringz.AlignRight("abcd", 4)).To(Equal("abcd"))
	g.Expect(stringz.AlignRight("cd", 2)).To(Equal("..cd"))
	g.Expect(stringz.AlignRight("abcd", 6)).To(Equal("..abcd"))
	g.Expect(stringz.AlignRight("abcdef", 4)).To(Equal("...f"))
}

func (*Suite) TestEnsurePrefix(g *WithT) {
	g.Expect(stringz.EnsurePrefix("ab", "a")).To(Equal("ab"))
	g.Expect(stringz.EnsurePrefix("b", "a")).To(Equal("ab"))
}

func (*Suite) TestEnsureSuffix(g *WithT) {
	g.Expect(stringz.EnsureSuffix("ab", "b")).To(Equal("ab"))
	g.Expect(stringz.EnsureSuffix("a", "b")).To(Equal("ab"))
}
