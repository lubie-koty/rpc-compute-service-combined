package util

import (
	"fmt"
	"strconv"
	"strings"

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

func ValidateURL(addressValue string) (string, error) {
	splitAddress := strings.Split(addressValue, ":")
	if len(splitAddress) != 2 {
		return "", fmt.Errorf("error while parsing url: %s", addressValue)
	}
	if _, err := ValidateHost(splitAddress[0]); err != nil {
		return "", err
	}
	if _, err := ValidatePort(splitAddress[1]); err != nil {
		return "", err
	}
	return addressValue, nil
}

func ValidateAppMode(appModeValue string) (string, error) {
	err := Validate.Var(appModeValue, "required,oneof=grpc rest")
	if err != nil {
		return "", err
	}
	return appModeValue, nil
}
