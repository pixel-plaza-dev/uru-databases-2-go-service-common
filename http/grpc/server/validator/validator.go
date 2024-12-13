package validator

import (
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonvalidatorfields "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/validator/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type (
	// Validator interface
	Validator interface {
		ValidateEmail(
			emailField string,
			email string,
			structFieldsValidations *commonvalidatorfields.StructFieldsValidations,
		)
		ValidateBirthdate(
			birthdateField string,
			birthdate *timestamppb.Timestamp,
			structFieldsValidations *commonvalidatorfields.StructFieldsValidations,
		)
		ValidateNilFields(request interface{}, structFieldsToValidate *commonvalidatorfields.StructFieldsToValidate) (
			*commonvalidatorfields.StructFieldsValidations,
			error,
		)
		CheckValidations(structFieldsValidations *commonvalidatorfields.StructFieldsValidations, code codes.Code) error
	}

	// DefaultValidator struct
	DefaultValidator struct {
		mode *commonflag.ModeFlag
	}
)

// NewDefaultValidator creates a new default validator
func NewDefaultValidator(mode *commonflag.ModeFlag) *DefaultValidator {
	return &DefaultValidator{
		mode: mode,
	}
}

// ValidateEmail validates the email address field
func (d *DefaultValidator) ValidateEmail(
	emailField string,
	email string,
	structFieldsValidations *commonvalidatorfields.StructFieldsValidations,
) {
	if _, err := commonvalidatorfields.ValidMailAddress(email); err != nil {
		structFieldsValidations.AddFailedFieldValidationError(emailField, commonvalidatorfields.InvalidMailAddressError)
	}
}

// ValidateBirthdate validates the birthdate field
func (d *DefaultValidator) ValidateBirthdate(
	birthdateField string,
	birthdate *timestamppb.Timestamp,
	structFieldsValidations *commonvalidatorfields.StructFieldsValidations,
) {
	if birthdate == nil || birthdate.AsTime().After(time.Now()) {
		structFieldsValidations.AddFailedFieldValidationError(
			birthdateField,
			commonvalidatorfields.InvalidBirthdateError,
		)
	}
}

// ValidateNilFields validates the nil fields
func (d *DefaultValidator) ValidateNilFields(
	request interface{},
	structFieldsToValidate *commonvalidatorfields.StructFieldsToValidate,
) (*commonvalidatorfields.StructFieldsValidations, error) {
	return commonvalidatorfields.ValidateNilFields(
		request,
		structFieldsToValidate,
		d.mode,
	)
}

// CheckValidations checks if there are any validations
func (d *DefaultValidator) CheckValidations(
	structFieldsValidations *commonvalidatorfields.StructFieldsValidations,
	code codes.Code,
) error {
	if structFieldsValidations.HasFailed() {
		return status.Error(code, structFieldsValidations.String())
	}
	return nil
}
