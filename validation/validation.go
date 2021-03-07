package validation

// SimpleValidator describes a type that can validate itself, returning only a "valid" bool.
type SimpleValidator interface {
	Valid() bool
}

// Validator describes a type that can validate itself.
type Validator interface {
	Validate() error
}
