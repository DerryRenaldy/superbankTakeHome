package validation

import (
	cError "accountDashboardService/pkgs/errors"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(inputStruct any) error {
	var errBag []cError.ValidationErrorField
	validate := validator.New()

	dataByte, _ := json.Marshal(inputStruct)

	fmt.Printf("Data : %s\n", string(dataByte))

	err := validate.Struct(inputStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errItems := cError.ValidationErrorField{
				Field:   err.Field(),
				Message: getMsgForTag(err),
			}
			errBag = append(errBag, errItems)
		}

		return cError.GetErrorValidation(cError.InvalidRequestError, errors.New("invalid request body"), errBag)
	}
	return nil
}

func getMsgForTag(error validator.FieldError) string {
	switch error.Tag() {
	case "required":
		return fmt.Sprintf("The '%s' field is required", error.Field())
	case "email":
		return fmt.Sprintf("The '%s' field has invalid format", error.Field())
	}
	return fmt.Sprintf("The '%s' tag is not implemented yet", error.Tag())
}

func registerCustomValidators(validate *validator.Validate, customValidators map[string]validator.Func) error {
	for key, value := range customValidators {
		err := validate.RegisterValidation(key, value)
		if err != nil {
			return fmt.Errorf("failed to register '%s' custom validator: %v", key, err)
		}
	}
	return nil
}
