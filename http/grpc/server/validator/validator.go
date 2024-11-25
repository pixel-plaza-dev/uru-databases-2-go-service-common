package validator

import (
	commonvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/validator"
	commonvalidatorerror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/validator/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type (
	// Validator interface
	Validator interface {
		ValidateEmail(emailField string, email string, validations *map[string][]error)
		ValidateBirthdate(birthdateField string, birthdate *timestamppb.Timestamp, validations *map[string][]error)
		ValidateNonEmptyStringFields(request interface{}, fieldsToValidate *map[string]string) *map[string][]error
		CheckValidations(validations *map[string][]error, code codes.Code) error
	}

	// DefaultValidator struct
	DefaultValidator struct{}
)

// NewDefaultValidator creates a new default validator
func NewDefaultValidator() DefaultValidator {
	return DefaultValidator{}
}

// ValidateEmail validates the email address field
func (d DefaultValidator) ValidateEmail(emailField string, email string, validations *map[string][]error) {
	if _, err := commonvalidator.ValidMailAddress(email); err != nil {
		(*validations)[emailField] = append(
			(*validations)[emailField],
			commonvalidator.InvalidMailAddressError,
		)
	}
}

// ValidateBirthdate validates the birthdate field
func (d DefaultValidator) ValidateBirthdate(birthdateField string, birthdate *timestamppb.Timestamp, validations *map[string][]error) {
	if birthdate == nil || birthdate.AsTime().After(time.Now()) {
		(*validations)[birthdateField] = append(
			(*validations)[birthdateField],
			commonvalidator.InvalidBirthdateError,
		)
	}
}

// ValidateNonEmptyStringFields validates the non-empty string fields
func (d DefaultValidator) ValidateNonEmptyStringFields(request interface{}, fieldsToValidate *map[string]string) *map[string][]error {
	// Validation variables
	validations := make(map[string][]error)

	// Check if the required string fields are empty
	commonvalidator.ValidNonEmptyStringFields(
		&validations,
		request,
		fieldsToValidate,
	)

	return &validations
}

// CheckValidations checks if there are any validations
func (d DefaultValidator) CheckValidations(validations *map[string][]error, code codes.Code) error {
	if len(*validations) > 0 {
		err := commonvalidatorerror.FailedValidationError{FieldsErrors: validations}
		return status.Error(code, err.Error())
	}
	return nil
}
