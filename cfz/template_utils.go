package cfz

import (
	"fmt"

	cf "github.com/awslabs/goformation/v7/cloudformation"
	cftags "github.com/awslabs/goformation/v7/cloudformation/tags"
)

// TemplateUtils provides utils for templates.
type TemplateUtils struct {
	namespace string
}

// NewTemplateUtils initializes a new TemplateUtils.
func NewTemplateUtils(namespace string) *TemplateUtils {
	return &TemplateUtils{
		namespace: namespace,
	}
}

// GetStackName generate a stack name.
func (u *TemplateUtils) GetStackName(name string) string {
	return fmt.Sprintf("%v-%v", u.namespace, name)
}

// GetResourceName generates a resource name.
func (u *TemplateUtils) GetResourceName(res Resource) string {
	return fmt.Sprintf("%v-%v", u.namespace, res.Physical())
}

// GenerateResourceTags generates resource tags.
func (u *TemplateUtils) GenerateResourceTags(res Resource) []cftags.Tag {
	return []cftags.Tag{
		{
			Key:   "Name",
			Value: u.GetResourceName(res),
		},
	}
}

// AddOutputRef adds a ref output to a template.
func (u *TemplateUtils) AddOutputRef(tpl *cf.Template, res Resource) {
	tpl.Outputs[fmt.Sprintf("%vRef", res.Logical())] = cf.Output{
		Value: res.Ref(),
	}
}

// AddOutputAtt adds an att output to a template.
func (u *TemplateUtils) AddOutputAtt(tpl *cf.Template, res Resource, att Attribute) {
	tpl.Outputs[fmt.Sprintf("%vAtt%v", res.Logical(), att.Logical())] = cf.Output{
		Value: res.GetAtt(att),
	}
}
