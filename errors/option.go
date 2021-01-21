package errors

var (
	_ Option = OptionFunc(nil)
)

// Option describes an option which can be applied to an error.
type Option interface {
	Apply(firstWrap bool, err error)
}

// OptionFunc describes an option which can be applied to an error.
type OptionFunc func(firstWrap bool, err error)

// Apply implements the Option interface.
func (f OptionFunc) Apply(firstWrap bool, err error) {
	f(firstWrap, err)
}
