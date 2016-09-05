package service

import (
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type ApiInterface interface {
	GetVersion() versionmanager.Version

	GetOrchestrator() orchestrator.Orchestrator

	GetHealthCheck() healthcheck.HealthCheckInterface

	Init()
}
