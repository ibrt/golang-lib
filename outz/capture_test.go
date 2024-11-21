package outz_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/outz"
)

type CaptureSuite struct {
	// intentionally empty
}

func TestCaptureSuite(t *testing.T) {
	fixturez.RunSuite(t, &CaptureSuite{})
}

func (*CaptureSuite) TestCapturing(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams, outz.SetupColorStreams, outz.SetupTableStreams)
	defer outz.MustResetCapturing()

	g.Expect(fmt.Fprint(os.Stdout, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(os.Stderr, "<err>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustStopCapturing()

	g.Expect(outBuf).To(Equal("<out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		outz.MustResetCapturing()
	}).ToNot(Panic())
}

func (*CaptureSuite) TestMustResetCapturing(g *WithT) {
	outz.MustStartCapturing(outz.SetupStandardStreams)
	defer outz.MustResetCapturing()
	fmt.Println("hidden")
}
