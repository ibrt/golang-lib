package devz_test

import (
	"os/exec"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/shellz"
	"github.com/ibrt/golang-lib/shellz/tshellz"
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

func (*Suite) TestRunGoChecks(ctrl *gomock.Controller) {
	devz.GoToolGolint.GetVersion()      // warm up
	devz.GoToolStaticCheck.GetVersion() // warm up

	m := tshellz.NewMockExecutor(ctrl)

	shellz.DefaultExecutor = m
	defer shellz.RestoreDefaultExecutor()

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "mod", "tidy"})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "generate", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "fmt", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "build", "-v", "-tags=t1,t2", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "run", devz.GoToolGolint.GetArgument(), "-set_exit_status", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "vet", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "run", devz.GoToolStaticCheck.GetArgument(), "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "mod", "tidy"})
		})).
		Times(1).
		Return(nil)

	devz.MustRunGoChecks(&devz.GoChecksParams{
		AllPackages:   []string{"./..."},
		BuildTags:     []string{"t1", "t2"},
		PrintNotices:  true,
		PrintCommands: true,
	})
}
