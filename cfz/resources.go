package cfz

import (
	cf "github.com/awslabs/goformation/v7/cloudformation"
	"github.com/iancoleman/strcase"
)

// Resource describes a resource.
type Resource string

// Ref returns a Ref.
func (r Resource) Ref() string {
	return cf.Ref(r.Logical())
}

// GetAtt returns a GetAtt.
func (r Resource) GetAtt(att Attribute) string {
	return cf.GetAtt(r.Logical(), att.Name())
}

// Logical returns a logical name.
func (r Resource) Logical() string {
	return strcase.ToCamel(string(r))
}

// Physical returns a physical name.
func (r Resource) Physical() string {
	return string(r)
}
