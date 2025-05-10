package config

import (
	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

type Config struct {
	Host           string
	Port           int
	AppMode        string
	GRPCClientHost string
	GRPCClientPort int
	HTTPClientHost string
	HTTPClientPort int
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
	gClientPort, err := util.ValidatePort(envVars["GRPC_CLIENT_PORT"])
	if err != nil {
		return err
	}
	gClientHost, err := util.ValidateHost(envVars["GRPC_CLIENT_HOST"])
	if err != nil {
		return err
	}
	hClientPort, err := util.ValidatePort(envVars["HTTP_CLIENT_PORT"])
	if err != nil {
		return err
	}
	hClientHost, err := util.ValidateHost(envVars["HTTP_CLIENT_HOST"])
	if err != nil {
		return err
	}
	AppConfig = Config{
		Host:           hostValue,
		Port:           portValue,
		AppMode:        appModeValue,
		GRPCClientPort: gClientPort,
		GRPCClientHost: gClientHost,
		HTTPClientPort: hClientPort,
		HTTPClientHost: hClientHost,
	}
	return nil
}
