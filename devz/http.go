package devz

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cheggaaa/pb/v3"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/ioz"
)

// DownloadError describes a download error.
type DownloadError struct {
	url        string
	statusCode int
	message    string
}

// NewDownloadError initializes a new download error.
func NewDownloadError(url string, statusCode int, message string) *DownloadError {
	return &DownloadError{
		url:        url,
		statusCode: statusCode,
		message:    message,
	}
}

// GetURL returns the URL.
func (e *DownloadError) GetURL() string {
	return e.url
}

// GetStatusCode returns the status code.
func (e *DownloadError) GetStatusCode() int {
	return e.statusCode
}

// Error implements the error interface.
func (e *DownloadError) Error() string {
	if e.message != "" {
		return fmt.Sprintf("download error for URL \"%v\": HTTP %v: %v", e.url, e.statusCode, e.message)
	}
	return fmt.Sprintf("download error for URL \"%v\": HTTP %v", e.url, e.statusCode)
}

// MustDownloadFile downloads a file.
func MustDownloadFile(url, outFilePath string) {
	consolez.DefaultCLI.Notice("download-file", url, filez.MustRelForDisplay(outFilePath))

	resp, err := http.Get(url)
	errorz.MaybeMustWrap(err)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		defer errorz.IgnoreClose(resp.Body)
		buf, _ := io.ReadAll(resp.Body)
		errorz.MustWrap(NewDownloadError(url, resp.StatusCode, string(buf)))
	}

	r := resp.Body

	if resp.ContentLength >= 0 {
		bar := pb.New64(resp.ContentLength).
			SetTemplate(pb.Full).
			Set(pb.Bytes, true).
			SetRefreshRate(50 * time.Millisecond).
			Start()
		r = bar.NewProxyReader(r)
		defer bar.Finish()
	}

	buf := ioz.MustReadAllAndClose(r)
	filez.MustWriteFile(outFilePath, 0777, 0666, buf)
}
