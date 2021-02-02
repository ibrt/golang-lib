package errors

var (
	_ Option = Args{}
)

// A is a shorthand builder for args.
func A(a ...interface{}) Args {
	return a
}

// Args describes a list of args used for formatting an error message.
type Args []interface{}

// Apply implements the Option interface.
func (a Args) Apply(err error) {
	// intentionally empty - args are used only be Errorf and extracted separately
}
