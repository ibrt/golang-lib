package errorz

import (
	"github.com/davecgh/go-spew/spew"
)

var (
	spewConfig = &spew.ConfigState{
		Indent:                  "  ",
		DisableMethods:          true,
		DisablePointerMethods:   true,
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
		SpewKeys:                true,
	}
)

// SDump converts the error to a string representation for debug purposes.
func SDump(err error) string {
	if err == nil {
		return "<nil>"
	}

	if e, ok := err.(*wrappedError); ok {
		type dump struct {
			Message string
			Debug   []error
			Frames  []string
		}

		return spewConfig.Sdump(dump{
			Message: e.Error(),
			Debug:   e.errs,
			Frames:  e.frames.ToSummaries(),
		})
	}

	type dump struct {
		Message string
		Debug   []error
	}

	return spewConfig.Sdump(dump{
		Message: err.Error(),
		Debug:   []error{err},
	})
}
