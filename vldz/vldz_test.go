package vldz_test

import (
	"regexp"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/vldz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (s *Suite) TestKindStructOrStructPtr(g *WithT) {
	type validatableStruct struct {
		Value any `json:"value" validate:"kind-struct-or-struct-ptr"`
	}

	for _, v := range []any{
		nil,
		(*struct{})(nil),
		"",
		100,
		[]string{},
		map[string]string{},
		make(chan struct{}),
	} {
		err := vldz.ValidateStruct(&validatableStruct{Value: v})
		g.Expect(err).ToNot(Succeed())

		vErr, ok := errorz.As[*vldz.ValidationError](err)
		g.Expect(ok).To(BeTrue())

		g.Expect(vErr.MaybeGetFieldsSummary()).To(Equal(map[string]any{"value": "kind-struct-or-struct-ptr"}))
		g.Expect(vErr.Error()).To(Equal(strings.Join([]string{
			"validation error(s):",
			"- Key: 'validatableStruct.value' Error:Field validation for 'value' failed on the 'kind-struct-or-struct-ptr' tag",
		}, "\n")))
		g.Expect(vErr.Unwrap()).ToNot(BeNil())
	}

	g.Expect(vldz.ValidateStruct(&validatableStruct{Value: struct{}{}})).To(Succeed())
	g.Expect(vldz.ValidateStruct(&validatableStruct{Value: &struct{}{}})).To(Succeed())
}

func (*Suite) TestValidateStruct(g *WithT) {
	vldz.MustRegisterValidator("custom-validator", vldz.RegexpValidatorFactory(regexp.MustCompile("^valid$")))

	type validatableStruct struct {
		First  string `json:"first" validate:"required"`
		Second string `validate:"custom-validator"`
	}

	{
		err := vldz.ValidateStruct(&validatableStruct{})
		g.Expect(err).ToNot(Succeed())

		vErr, ok := errorz.As[*vldz.ValidationError](err)
		g.Expect(ok).To(BeTrue())

		g.Expect(vErr.MaybeGetFieldsSummary()).To(Equal(map[string]any{
			"first":  "required",
			"Second": "custom-validator",
		}))
		g.Expect(vErr.Error()).To(Equal(strings.Join([]string{
			"validation error(s):",
			"- Key: 'validatableStruct.first' Error:Field validation for 'first' failed on the 'required' tag",
			"- Key: 'validatableStruct.Second' Error:Field validation for 'Second' failed on the 'custom-validator' tag",
		}, "\n")))
	}

	{
		err := vldz.ValidateStruct(&validatableStruct{
			First:  "required",
			Second: "valid",
		})
		g.Expect(err).To(Succeed())
	}

	{
		err := vldz.ValidateStruct("")
		g.Expect(err).ToNot(Succeed())

		vErr, ok := errorz.As[*vldz.ValidationError](err)
		g.Expect(ok).To(BeTrue())

		g.Expect(vErr.MaybeGetFieldsSummary()).To(BeNil())
		g.Expect(vErr.Error()).To(Equal("validation error: validator: (nil string)"))
	}

	g.Expect(func() { vldz.MustValidateStruct(&validatableStruct{}) }).To(Panic())
	g.Expect(func() { vldz.MustValidateStruct(&validatableStruct{First: "required", Second: "valid"}) }).ToNot(Panic())
}