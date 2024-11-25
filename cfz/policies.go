package cfz

import (
	cf "github.com/awslabs/goformation/v7/cloudformation"

	"github.com/ibrt/golang-lib/errorz"
)

// NewAssumeRolePolicyDocument generates a new assume role policy document.
func NewAssumeRolePolicyDocument(service string) any {
	return map[string]any{
		"Version": "2012-10-17",
		"Statement": []map[string]any{
			{
				"Effect": "Allow",
				"Principal": map[string]any{
					"Service": service,
				},
				"Action": "sts:AssumeRole",
			},
		},
	}
}

// NewPolicyDocument generates a new policy document.
func NewPolicyDocument(statements ...*PolicyStatement) any {
	materializedStatements := make([]any, 0)
	for _, statement := range statements {
		materializedStatements = append(materializedStatements, statement.Build())
	}

	return map[string]any{
		"Version":   "2012-10-17",
		"Statement": materializedStatements,
	}
}

// PolicyStatement describes a policy statement.
type PolicyStatement struct {
	Actions   []string
	Resources []string
	Principal any
}

// NewPolicyStatement initializes a new PolicyStatement.
func NewPolicyStatement() *PolicyStatement {
	return &PolicyStatement{
		Actions:   make([]string, 0),
		Resources: make([]string, 0),
	}
}

// AddActions adds actions to the policy statement.
func (s *PolicyStatement) AddActions(actions ...string) *PolicyStatement {
	s.Actions = append(s.Actions, actions...)
	return s
}

// AddResources adds resources to the policy statement.
func (s *PolicyStatement) AddResources(resources ...string) *PolicyStatement {
	s.Resources = append(s.Resources, resources...)
	return s
}

// SetCurrentRootAccountPrincipal sets the current root account as principal on the policy statement.
func (s *PolicyStatement) SetCurrentRootAccountPrincipal() *PolicyStatement {
	s.Principal = map[string]any{
		"AWS": cf.Sub("arn:aws:iam::${AWS::AccountId}:root"),
	}
	return s
}

// SetWildcardPrincipal sets a wildcard as principal on the policy statement.
func (s *PolicyStatement) SetWildcardPrincipal() *PolicyStatement {
	s.Principal = "*"
	return s
}

// SetServicePrincipal sets a service as principal on the policy statement.
func (s *PolicyStatement) SetServicePrincipal(service string) *PolicyStatement {
	s.Principal = map[string]any{
		"Service": service,
	}
	return s
}

// SetAnyRootAccountPrincipal sets any root account as principal on the policy statement.
func (s *PolicyStatement) SetAnyRootAccountPrincipal() *PolicyStatement {
	s.Principal = map[string]any{
		"AWS": "*",
	}
	return s
}

// Build builds the policy statement.
func (s *PolicyStatement) Build() any {
	errorz.Assertf(len(s.Actions) > 0, "actions unexpectedly empty")
	errorz.Assertf(len(s.Resources) > 0, "resources unexpectedly empty")

	m := map[string]any{
		"Effect":   "Allow",
		"Action":   s.Actions,
		"Resource": s.Resources,
	}

	if s.Principal != nil {
		m["Principal"] = s.Principal
	}

	return m
}
