package outz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/outz"
)

type ColorsSuite struct {
	// intentionally empty
}

func TestColorsSuite(t *testing.T) {
	fixturez.RunSuite(t, &ColorsSuite{})
}

func (s *ColorsSuite) TestColors(g *WithT) {
	outz.MustStartCapturing(outz.SetupColorStreams)
	defer outz.MustResetCapturing()

	g.Expect(outz.GetColorDefault().Print("default")).Error().To(Succeed())
	g.Expect(outz.GetColorHighlight().Print("highlight")).Error().To(Succeed())
	g.Expect(outz.GetColorSecondaryHighlight().Print("secondaryHighlight")).Error().To(Succeed())
	g.Expect(outz.GetColorSecondary().Print("secondary")).Error().To(Succeed())
	g.Expect(outz.GetColorInfo().Print("info")).Error().To(Succeed())
	g.Expect(outz.GetColorSuccess().Print("success")).Error().To(Succeed())
	g.Expect(outz.GetColorWarning().Print("warning")).Error().To(Succeed())
	g.Expect(outz.GetColorError().Print("error")).Error().To(Succeed())

	outBuf, errBuf := outz.MustStopCapturing()
	g.Expect(outBuf).To(Equal("\x1b[0mdefault\x1b[0m\x1b[1mhighlight\x1b[0m\x1b[1;2msecondaryHighlight\x1b[0m\x1b[2msecondary\x1b[0m\x1b[36minfo\x1b[0m\x1b[32msuccess\x1b[0m\x1b[33mwarning\x1b[0m\x1b[91merror\x1b[0m"))
	g.Expect(errBuf).To(Equal(""))
}
