package inject

import "io"

// Releaser releases an initialized resource.
type Releaser func()

// SafeRelease calls Release on the Releaser if not nil.
// Note: this is mainly needed because a Releaser can be returned by Bootstrap also in case of error.
func SafeRelease(releaser Releaser) {
	if releaser != nil {
		releaser()
	}
}

// CloseReleaserFactory is a releaser that calls Close.
func CloseReleaserFactory(closer io.Closer) Releaser {
	return func() {
		if closer != nil {
			_ = closer.Close()
		}
	}
}

// CompoundReleaserFactory combines multiple releasers into one (invoking them in reverse order).
func CompoundReleaserFactory(releasers ...Releaser) Releaser {
	return func() {
		for i := len(releasers) - 1; i >= 0; i-- {
			SafeRelease(releasers[i])
		}
	}
}
