package validation_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/validation"
	"github.com/ibrt/golang-lib/validation/internal/validationmocks"
	"github.com/stretchr/testify/require"
)

func TestValidate_NoValidation(t *testing.T) {
	fixtures.RequireNoError(t, validation.Validate(nil))
}

func TestValidate_SimpleValidator_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSimpleValidator := validationmocks.NewMockSimpleValidator(ctrl)
	mockSimpleValidator.EXPECT().Valid().Return(true)

	fixtures.RequireNoError(t, validation.Validate(mockSimpleValidator))
}

func TestValidate_SimpleValidator_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSimpleValidator := validationmocks.NewMockSimpleValidator(ctrl)
	mockSimpleValidator.EXPECT().Valid().Return(false)

	require.Error(t, validation.Validate(mockSimpleValidator))
}

func TestValidate_Validator_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidator := validationmocks.NewMockValidator(ctrl)
	mockValidator.EXPECT().Validate().Return(nil)

	fixtures.RequireNoError(t, validation.Validate(mockValidator))
}

func TestValidate_Validator_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidator := validationmocks.NewMockValidator(ctrl)
	mockValidator.EXPECT().Validate().Return(errors.Errorf("test"))

	require.Error(t, validation.Validate(mockValidator))
}
