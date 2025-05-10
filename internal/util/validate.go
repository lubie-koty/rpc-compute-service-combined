package util

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidate() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidatePort(portValue string) (int, error) {
	intPortValue, err := strconv.Atoi(portValue)
	if err != nil {
		return 0, err
	}
	err = Validate.Var(intPortValue, "required,gt=1,lt=65535")
	if err != nil {
		return 0, err
	}
	return intPortValue, nil
}

func ValidateHost(hostValue string) (string, error) {
	err := Validate.Var(hostValue, "required")
	if err != nil {
		return "", err
	}
	return hostValue, nil
}

func ValidateAppMode(appModeValue string) (string, error) {
	err := Validate.Var(appModeValue, "required,oneof=grpc rest")
	if err != nil {
		return "", err
	}
	return appModeValue, nil
}
