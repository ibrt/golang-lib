package cfz

import (
	"fmt"

	awscft "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/ibrt/golang-lib/errorz"
)

// Outputs describes a set of outputs.
type Outputs struct {
	stack *awscft.Stack
}

// NewOutputs initializes a new Outputs.
func NewOutputs(stack *awscft.Stack) *Outputs {
	return &Outputs{
		stack: stack,
	}
}

// Ref gets the value of a reference output.
func (e *Outputs) Ref(res Resource) string {
	outputKey := fmt.Sprintf("%vRef", res.Logical())

	for _, output := range e.stack.Outputs {
		if *output.OutputKey == outputKey {
			return *output.OutputValue
		}
	}

	panic(errorz.Errorf("no such export: ref for %v", res))
}

// Att gets the value of an attribute output.
func (e *Outputs) Att(res Resource, att Attribute) string {
	outputKey := fmt.Sprintf("%vAtt%v", res.Logical(), att.Logical())

	for _, output := range e.stack.Outputs {
		if *output.OutputKey == outputKey {
			return *output.OutputValue
		}
	}

	panic(errorz.Errorf("no such export: att %v for %v", att, res))
}
