package cfz

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// Attribute describes an attribute.
type Attribute string

// Name returns the attribute name.
func (a Attribute) Name() string {
	return string(a)
}

// Logical returns a logical name.
func (a Attribute) Logical() string {
	return strings.ReplaceAll(string(a), ".", "")
}

// Physical returns a physical name.
func (a Attribute) Physical() string {
	return strcase.ToKebab(string(a))
}
