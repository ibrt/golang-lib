package consolez_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/kong"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/outz"
)

type CLISuite struct {
	// intentionally empty
}

func TestCLISuite(t *testing.T) {
	fixturez.RunSuite(t, &CLISuite{})
}

func (*CLISuite) TestTool(g *WithT) {
	type CLI struct {
		Command struct {
			Arg string `arg:"" help:"Positional argument."`
		} `cmd:"" help:"Command"`
		Flag string `flag:"" help:"Flag."`
	}

	k, err := kong.New(&CLI{})
	g.Expect(err).To(Succeed())

	kCtx, err := k.Parse([]string{"--flag=flag-value", "command", "arg-value"})
	g.Expect(err).To(Succeed())

	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Tool("Tool", kCtx)

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("┌─────────────────┐\n│ %v \x1b[1mTool\x1b[0m command │\n└─────────────────┘\n\n\x1b[1mInput          Value       \n\x1b[22m\x1b[33m--flag=STRING  \x1b[0mflag-value  \n\x1b[33m<arg>          \x1b[0marg-value   \n", consolez.IconRocket)))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestBanner(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Banner("Title", "tagline")

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("┌──────────────────┐\n│ %v \x1b[1mTitle\x1b[0m tagline │\n└──────────────────┘\n", consolez.IconRocket)))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestHeader_AddSpaceBeforeHeadersTrue(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	c := consolez.NewCLI(true, os.Exit)

	f1 := c.Header("H1 %v", 1)
	f2 := c.Header("H2 %v", 1)
	f3 := c.Header("H3 %v", 1)
	f4 := c.Header("H3 %v", 1)

	f4()
	f3()
	f3()
	f2()
	f1()

	f1 = c.Header("H1 %v", 2)

	f1()

	outBuf, errBuf := outz.MustStopCapturing()

	g.Expect(outBuf).To(Equal(strings.Join(
		[]string{
			fmt.Sprintf("\n%v \x1b[1mH1 1\x1b[0m\n", consolez.IconHighVoltage),
			fmt.Sprintf("\n%v H2 1\n", consolez.IconBackhandIndexPointingRight),
			"\x1b[1;2m—— \x1b[0m\x1b[1;2mH3 1\x1b[0m\n",
			"\x1b[1;2m—— \x1b[0m\x1b[1;2mH3 1\x1b[0m\n",
			fmt.Sprintf("\n%v \x1b[1mH1 2\x1b[0m\n", consolez.IconHighVoltage),
		}, "")))

	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestHeader_AddSpaceBeforeHeadersFalse(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	c := consolez.NewCLI(false, os.Exit)

	f1 := c.Header("H1 %v", 1)
	f2 := c.Header("H2 %v", 1)
	f3 := c.Header("H3 %v", 1)
	f4 := c.Header("H3 %v", 1)

	f4()
	f3()
	f3()
	f2()
	f1()

	f1 = c.Header("H1 %v", 2)

	f1()

	outBuf, errBuf := outz.MustStopCapturing()

	g.Expect(outBuf).To(Equal(strings.Join(
		[]string{
			fmt.Sprintf("%v \x1b[1mH1 1\x1b[0m\n", consolez.IconHighVoltage),
			fmt.Sprintf("%v H2 1\n", consolez.IconBackhandIndexPointingRight),
			"\x1b[1;2m—— \x1b[0m\x1b[1;2mH3 1\x1b[0m\n",
			"\x1b[1;2m—— \x1b[0m\x1b[1;2mH3 1\x1b[0m\n",
			fmt.Sprintf("%v \x1b[1mH1 2\x1b[0m\n", consolez.IconHighVoltage),
		}, "")))

	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestNotice(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Notice("scope", "p1", "p2", "p3")

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal("\x1b[2m[...................scope]\x1b[0m\x1b[0m p1\x1b[0m\x1b[2m p2\x1b[0m\x1b[2m p3\x1b[0m\n"))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestCommand_Rel(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Command("cmd", "p1", "p2")

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("%v cmd \x1b[2mp1 p2\x1b[0m\n", consolez.IconRunner)))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestCommand_Abs(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Command(filez.MustAbs("cmd"), "p1", "p2")

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("%v cmd \x1b[2mp1 p2\x1b[0m\n", consolez.IconRunner)))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestNewTable(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.NewTable("A", "B").AddRow("a", "b").Print()

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal("\x1b[1mA  B  \n\x1b[22m\x1b[33ma  \x1b[0mb  \n"))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestError_DebugFalse(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Error(fmt.Errorf("test error"), false)

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("\n%v \x1b[1mError\x1b[22m\n\x1b[91mtest error\x1b[0m\n", consolez.IconCollision)))
	g.Expect(errBuf).To(Equal(""))
}

func (*CLISuite) TestError_DebugTrue(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	consolez.DefaultCLI.Error(fmt.Errorf("test error"), true)

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(HavePrefix(fmt.Sprintf("\n%v \x1b[1mError\x1b[22m\n\x1b[91mtest error\x1b[0m\n(errorz.dump)", consolez.IconCollision)))
	g.Expect(errBuf).To(Equal(""))
}
