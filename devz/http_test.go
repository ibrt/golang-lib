package devz_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
)

type HTTPSuite struct {
	// intentionally empty
}

func TestHTTPSuite(t *testing.T) {
	fixturez.RunSuite(t, &HTTPSuite{})
}

func (*HTTPSuite) TestDownloadError(g *WithT) {
	err := devz.NewDownloadError("http://example.com", http.StatusNotFound, "content")
	g.Expect(err).To(MatchError("download error for URL \"http://example.com\": HTTP 404: content"))
	g.Expect(err.GetURL()).To(Equal("http://example.com"))
	g.Expect(err.GetStatusCode()).To(Equal(http.StatusNotFound))

	err = devz.NewDownloadError("http://example.com", http.StatusNotFound, "")
	g.Expect(err).To(MatchError("download error for URL \"http://example.com\": HTTP 404"))

}

func (*HTTPSuite) TestMustDownloadFile_Success(g *WithT) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/file").
		Reply(http.StatusOK).
		BodyString("content")

	filePath := filez.MustCreateTempFileString("")
	defer filez.MustRemoveAll(filePath)

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(true), fixturez.OutputSetupTable)
	defer fixturez.ResetOutputCapture()

	g.Expect(func() { devz.MustDownloadFile("http://example.com/file", filePath) }).ToNot(Panic())

	outBuf, errBuf := fixturez.MustEndOutputCapture()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("[...........download-file] http://example.com/file %v\n", filePath)))
	g.Expect(errBuf).To(HavePrefix("7 B / 7 B ["))

	g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))
}

func (*HTTPSuite) TestMustDownloadFile_Error(g *WithT) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/file").
		Reply(http.StatusNotFound).
		BodyString("content")

	filePath := filez.MustCreateTempFileString("")
	defer filez.MustRemoveAll(filePath)

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(true), fixturez.OutputSetupTable)
	defer fixturez.ResetOutputCapture()

	g.Expect(
		func() {
			devz.MustDownloadFile("http://example.com/file", filePath)
		}).
		To(PanicWith(MatchError("download error for URL \"http://example.com/file\": HTTP 404: content")))

	outBuf, errBuf := fixturez.MustEndOutputCapture()
	g.Expect(outBuf).To(Equal(fmt.Sprintf("[...........download-file] http://example.com/file %v\n", filePath)))
	g.Expect(errBuf).To(Equal(""))

	g.Expect(filez.MustReadFileString(filePath)).To(Equal(""))
}
