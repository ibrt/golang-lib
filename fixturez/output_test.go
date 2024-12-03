package fixturez_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
)

type OutputSuite struct {
	// intentionally empty
}

func TestOutputSuite(t *testing.T) {
	fixturez.RunSuite(t, &OutputSuite{})
}

func (*OutputSuite) TestOutputCapture(g *WithT) {
	defer fixturez.ResetOutputCapture()
	defer fixturez.ResetOutputCapture()

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(false), fixturez.OutputSetupTable)

	g.Expect(fmt.Fprint(os.Stdout, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(os.Stderr, "<err>")).Error().To(Succeed())

	outBuf, errBuf := fixturez.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("<out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		fixturez.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestResetOutputCapture(_ *WithT) {
	defer fixturez.ResetOutputCapture()
	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(false), fixturez.OutputSetupTable)
	fmt.Println("ignored")
}
