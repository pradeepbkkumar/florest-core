package servicetest

import (
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/core/service"
)

var globalConfigJson = `
    {
	   "AppName":"florest",
	   "AppVersion":"1.0.0",
	   "ServerPort":"8080",
	   "LogConfFile":"conf/logger.json",
	   "MonitorConfig":{
	      "AppName":"florest",
	      "Platform":"DatadogAgent",
	      "AgentServer":"datadog:8125",
	      "Verbose":false,
	      "Enabled":true,
	      "MetricsServer":"datadog:8065"
	   },
	   "Performance":{
	      "UseCorePercentage":100,
	      "GCPercentage":1000
	   },
	   "ApplicationConfig":{
	   "Hello":{
	      "ResponseHeaders":{
	         "CacheControl":{
	            "ResponseType":"public",
	            "NoCache":false,
	            "NoStore":false,
	            "MaxAgeInSeconds":300
	         }
	      }
	   }
	}
   }`

var appConfigJson = `
	{
	   "Hello":{
	      "ResponseHeaders":{
	         "CacheControl":{
	            "ResponseType":"public",
	            "NoCache":false,
	            "NoStore":false,
	            "MaxAgeInSeconds":300
	         }
	      }
	   }
	}
`

func initTestConfig() {
	cm := new(service.ConfigManager)
	cm.InitializeGlobalConfigFromJson(globalConfigJson)
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")
}
