package gtz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/gtz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestGoTool(g *WithT) {
	gtz.GoToolGoLint.MustRun(".")

	g.Expect(gtz.NewGoTool("a", "b", "c").GetPackage()).To(Equal("a"))
	g.Expect(gtz.NewGoTool("a", "b", "c").GetArgument()).To(Equal("a/b@c"))
	g.Expect(gtz.NewGoTool("a", "b", "c").GetVersion()).To(Equal("c"))
	g.Expect(gtz.NewGoTool("github.com/axw/gocov", "", "unused").GetVersion()).To(Equal("v1.2.1"))

	gt := gtz.NewGoTool("a", "b", "")
	g.Expect(gt.GetVersion()).To(Equal("latest"))
	g.Expect(gt.GetVersion()).To(Equal("latest"))
}
