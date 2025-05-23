package util

import (
	"os"
)

func GetEnvVariables() map[string]string {
	return map[string]string{
		"PORT":                os.Getenv("PORT"),
		"HOST":                os.Getenv("HOST"),
		"APP_MODE":            os.Getenv("APP_MODE"),
		"SIMPLE_SERVICE_URL":  os.Getenv("SIMPLE_SERVICE_URL"),
		"COMPLEX_SERVICE_URL": os.Getenv("COMPLEX_SERVICE_URL"),
	}
}
