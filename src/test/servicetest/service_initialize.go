package servicetest

import (
	"github.com/jabong/florest-core/src/core/common/env"
	"github.com/jabong/florest-core/src/core/service"
)

func InitializeTestService() {

	//apiservice.Register()

	env.GetOsEnviron()

	initTestConfig()

	initTestLogger()

	service.InitDBAdapterManager()

	service.InitVersionManager()

	service.InitCustomApiInit()

	service.InitApis()

	service.InitHealthCheck()

	initialiseTestWebServer()

}

func PurgeTestService() {

}
