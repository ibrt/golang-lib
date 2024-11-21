package devz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/fixturez"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestGoTool(g *WithT) {
	devz.GoToolGolint.MustRun(".")

	g.Expect(devz.NewGoTool("a", "b", "c").GetPackage()).To(Equal("a"))
	g.Expect(devz.NewGoTool("a", "b", "c").GetArgument()).To(Equal("a/b@c"))
	g.Expect(devz.NewGoTool("a", "b", "c").GetVersion()).To(Equal("c"))
	g.Expect(devz.NewGoTool("github.com/axw/gocov", "", "unused").GetVersion()).To(Equal("v1.2.1"))

	gt := devz.NewGoTool("a", "b", "")
	g.Expect(gt.GetVersion()).To(Equal("latest"))
	g.Expect(gt.GetVersion()).To(Equal("latest"))
}
