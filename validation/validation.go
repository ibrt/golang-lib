package validation

import "github.com/ibrt/golang-lib/errors"

// SimpleValidator describes a type that can validate itself, returning only a "valid" bool.
type SimpleValidator interface {
	Valid() bool
}

// Validator describes a type that can validate itself.
type Validator interface {
	Validate() error
}

// Validate calls Valid and/or Validate if the given value implements SimpleValidator or Validator.
func Validate(v interface{}) error {
	if v, ok := v.(SimpleValidator); ok {
		if !v.Valid() {
			return errors.Errorf("invalid", errors.Skip())
		}
	}

	if v, ok := v.(Validator); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}
