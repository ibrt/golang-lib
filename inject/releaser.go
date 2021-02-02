package inject

// Releaser releases an initialized resource.
type Releaser func()

// CompoundReleaser combines multiple releasers into one (invoking them in reverse order).
func CompoundReleaser(releasers ...Releaser) Releaser {
	return func() {
		for i := len(releasers) - 1; i >= 0; i-- {
			SafeRelease(releasers[i])
		}
	}
}

// SafeRelease calls Release on the Releaser if not nil.
// Note: this is mainly needed because a Releaser can be returned by Bootstrap also in case of error.
func SafeRelease(releaser Releaser) {
	if releaser != nil {
		releaser()
	}
}
