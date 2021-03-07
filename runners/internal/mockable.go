package internal

// TestFunc is a helper interface to mock Func.
type TestFunc interface {
	Func() error
}
