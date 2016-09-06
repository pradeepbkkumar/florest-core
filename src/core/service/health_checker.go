package service

import (
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/core/common/env"
)

type ServiceHealthCheck struct {
}

func (n ServiceHealthCheck) GetName() string {
	return "service"
}

func (n ServiceHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"version":              config.GlobalAppConfig.AppVersion,
		"environmentVariables": env.GetOsEnviron().GetAll(),
	}
}
