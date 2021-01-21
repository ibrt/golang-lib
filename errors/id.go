package errors

var (
	_ Option = ID("")
)

// ID describes an error id.
type ID string

// String implements the fmt.Stringer interface.
func (id ID) String() string {
	return string(id)
}

// In returns true if the given error has this ID.
func (id ID) In(err error) bool {
	return id == GetID(err)
}

// Apply implements the Option interface.
func (id ID) Apply(_ bool, err error) {
	if e, ok := err.(*wrappedError); ok {
		e.id = id
	}
}

// GetID gets the id from the error, or an empty id if not set.
func GetID(err error) ID {
	if e, ok := err.(*wrappedError); ok {
		return e.id
	}
	return ""
}
