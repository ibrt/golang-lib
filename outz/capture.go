package outz

import (
	"io"
	"os"
	"sync"

	"github.com/ibrt/golang-lib/errorz"
)

var (
	m              = &sync.Mutex{}
	isCapturing    = false
	captureStdoutR *os.File
	captureStdoutW *os.File
	captureStderrR *os.File
	captureStderrW *os.File
	restoreFuncs   []func()
)

// RestoreFunc describe a function that restores the original streams.
type RestoreFunc func()

// SetupFunc describes a function that replaces streams with the ones for capturing.
type SetupFunc func(out *os.File, err *os.File) RestoreFunc

// SetupStandardStreams is a SetupFunc that configures the stdout/stderr streams.
func SetupStandardStreams(out *os.File, err *os.File) RestoreFunc {
	origStdout := os.Stdout
	origStderr := os.Stderr

	os.Stdout = out
	os.Stderr = err

	return func() {
		os.Stdout = origStdout
		os.Stderr = origStderr
	}
}

// MustBeginCapturing sets up the streams and begins capturing.
func MustBeginCapturing(setupFunc SetupFunc, additionalSetupFuncs ...SetupFunc) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(!isCapturing, "capturing already in progress")
	isCapturing = true
	var err error

	captureStdoutR, captureStdoutW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	captureStderrR, captureStderrW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	for _, f := range append([]SetupFunc{setupFunc}, additionalSetupFuncs...) {
		restoreFuncs = append(restoreFuncs, f(captureStdoutW, captureStderrW))
	}
}

// MustEndCapturing restores the original streams and returns the captured stdout/stderr.
func MustEndCapturing() (string, string) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(isCapturing, "capturing not in progress")
	return mustEndCapture()
}

// MustClearCapturing ensure the capturing is cleared (e.g. after a failing test).
func MustClearCapturing() {
	m.Lock()
	defer m.Unlock()

	if isCapturing {
		mustEndCapture()
	}
}

func mustEndCapture() (string, string) {
	defer func() {
		isCapturing = false
		captureStdoutR = nil
		captureStdoutW = nil
		captureStderrR = nil
		captureStderrW = nil
		restoreFuncs = nil
	}()

	for i := len(restoreFuncs) - 1; i >= 0; i-- {
		restoreFuncs[i]()
	}

	errorz.MaybeMustWrap(captureStdoutW.Close())
	errorz.MaybeMustWrap(captureStderrW.Close())

	outBuf, err := io.ReadAll(captureStdoutR)
	errorz.MaybeMustWrap(err)

	errBuf, err := io.ReadAll(captureStderrR)
	errorz.MaybeMustWrap(err)

	return string(outBuf), string(errBuf)
}
