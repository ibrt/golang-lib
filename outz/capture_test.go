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
	outz.MustBeginCapturing(outz.SetupStandardStreams)
	defer outz.MustClearCapturing()

	g.Expect(fmt.Fprint(os.Stdout, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(os.Stderr, "<err>")).Error().To(Succeed())

	outBuf, outErr := outz.MustEndCapturing()

	g.Expect(outBuf).To(Equal("<out>"))
	g.Expect(outErr).To(Equal("<err>"))

	g.Expect(func() {
		outz.MustClearCapturing()
	}).ToNot(Panic())
}

func (*CaptureSuite) TestMustClearCapturing(g *WithT) {
	outz.MustBeginCapturing(outz.SetupStandardStreams)
	defer outz.MustClearCapturing()
	fmt.Println("hidden")
}
