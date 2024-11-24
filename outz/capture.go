package outz

import (
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/ibrt/golang-lib/errorz"
)

var (
	m                      = &sync.Mutex{}
	isCapturing            = false
	outR, outW, errR, errW *os.File
	restoreFuncs           []func()
)

// SetupFunc describes a function that replaces streams with the ones for capturing.
type SetupFunc func(outW *os.File, errW *os.File) RestoreFunc

// RestoreFunc describe a function that restores the original streams.
type RestoreFunc func()

// SetupStandardStreams is a SetupFunc that configures the stdout/stderr streams.
func SetupStandardStreams(outW *os.File, errW *os.File) RestoreFunc {
	origOut := os.Stdout
	origErr := os.Stderr

	os.Stdout = outW
	os.Stderr = errW

	return func() {
		os.Stdout = origOut
		os.Stderr = origErr
	}
}

// GetSetupColorStreams returns a SetupFunc that configures the color streams.
func GetSetupColorStreams(noColor bool) SetupFunc {
	return func(outW *os.File, errW *os.File) RestoreFunc {
		origNoColor := color.NoColor
		origOut := color.Output
		origErr := color.Error

		color.NoColor = noColor
		color.Output = outW
		color.Error = errW

		return func() {
			color.NoColor = origNoColor
			color.Output = origOut
			color.Error = origErr
		}
	}
}

// SetupTableStreams is a SetupFunc that configures the table streams.
func SetupTableStreams(outW *os.File, _ *os.File) RestoreFunc {
	origOut := table.DefaultWriter
	table.DefaultWriter = outW

	return func() {
		table.DefaultWriter = origOut
	}
}

// MustStartCapturing sets up the streams and starts capturing.
func MustStartCapturing(setupFunc SetupFunc, additionalSetupFuncs ...SetupFunc) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(!isCapturing, "capturing already in progress")
	isCapturing = true
	var err error

	outR, outW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	errR, errW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	for _, f := range append([]SetupFunc{setupFunc}, additionalSetupFuncs...) {
		restoreFuncs = append(restoreFuncs, f(outW, errW))
	}
}

// MustStopCapturing restores the original streams and returns the captured stdout/stderr.
func MustStopCapturing() (string, string) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(isCapturing, "capturing not in progress")
	return mustStop()
}

// MustResetCapturing ensures the capturing is stopped (e.g. after a failing test).
func MustResetCapturing() {
	m.Lock()
	defer m.Unlock()

	if isCapturing {
		mustStop()
	}
}

func mustStop() (string, string) {
	defer func() {
		isCapturing = false
		outR = nil
		outW = nil
		errR = nil
		errW = nil
		restoreFuncs = nil
	}()

	for i := len(restoreFuncs) - 1; i >= 0; i-- {
		restoreFuncs[i]()
	}

	errorz.MaybeMustWrap(outW.Close())
	errorz.MaybeMustWrap(errW.Close())

	outBuf, err := io.ReadAll(outR)
	errorz.MaybeMustWrap(err)

	errBuf, err := io.ReadAll(errR)
	errorz.MaybeMustWrap(err)

	return string(outBuf), string(errBuf)
}
