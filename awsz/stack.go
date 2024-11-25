package awsz

import (
	"sync"

	cf "github.com/awslabs/goformation/v7/cloudformation"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/errorz"
)

var (
	_ GenericStack = (*Stack)(nil)
)

// GenericStack describes the public parts of a generic stack.
type GenericStack interface {
	MustMaterialize()
	MustGetOutputs() *cfz.Outputs
}

// GetTemplateFunc generates a stack template.
type GetTemplateFunc func() *cf.Template

// Stack manages a stack.
type Stack struct {
	stackName string
	ops       *Operations
	f         GetTemplateFunc
	m         *sync.Mutex
	outputs   *cfz.Outputs
}

// NewStack initializes a new stack.
func NewStack(baseStackName string, ops *Operations, f GetTemplateFunc) *Stack {
	return &Stack{
		stackName: ops.GetTemplateUtils().GetStackName(baseStackName),
		ops:       ops,
		f:         f,
		m:         &sync.Mutex{},
		outputs:   nil,
	}
}

// GetName returns the full stack name.
func (s *Stack) GetName() string {
	return s.stackName
}

// GetOperations returns the Stack's underlying *Operations.
func (s *Stack) GetOperations() *Operations {
	return s.ops
}

// MustMaterialize materializes the stack.
func (s *Stack) MustMaterialize() {
	s.m.Lock()
	defer s.m.Unlock()

	buf, err := s.f().JSON()
	errorz.MaybeMustWrap(err)

	tags := map[string]string{
		"Stack":     s.stackName,
		"AppPrefix": s.ops.GetAppPrefix(),
	}

	if stagePrefix := s.ops.GetStagePrefix(); stagePrefix != "" {
		tags["StagePrefix"] = stagePrefix
	}

	s.outputs = cfz.NewOutputs(s.ops.MustUpsertCFStack(s.stackName, string(buf), tags))
}

// MustGetOutputs returns the stack outputs, describing the stack if needed.
func (s *Stack) MustGetOutputs() *cfz.Outputs {
	s.m.Lock()
	defer s.m.Unlock()

	if s.outputs == nil {
		s.outputs = cfz.NewOutputs(s.ops.MustDescribeCFStack(s.stackName))
	}

	return s.outputs
}
