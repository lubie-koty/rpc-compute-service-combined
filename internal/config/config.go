package config

import (
	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

type Config struct {
	Host    string
	Port    int
	AppMode string
}

var AppConfig Config

func InitConfig(envVars map[string]string) error {
	portValue, err := util.ValidatePort(envVars["PORT"])
	if err != nil {
		return err
	}
	hostValue, err := util.ValidateHost(envVars["HOST"])
	if err != nil {
		return err
	}
	appModeValue, err := util.ValidateAppMode(envVars["APP_MODE"])
	if err != nil {
		return err
	}
	AppConfig = Config{
		Host:    hostValue,
		Port:    portValue,
		AppMode: appModeValue,
	}
	return nil
}
