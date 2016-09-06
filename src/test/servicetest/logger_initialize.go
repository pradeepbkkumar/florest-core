package servicetest

import (
	"github.com/jabong/florest-core/src/common/logger"
)

func initTestLogger() {
	configLogger := `
	{
	    "ProfilerEnabled": false,
	    "LogLevel": 1,
	    "DefaultLogType": "dummyLoggerDefault",
	    "AppName" : "florest-test-",
	    "FileLogger": [
	        {
	            "Key": "fileLoggerDefault",
	            "Path": "/tmp/",
	            "FileNamePrefix": "search-suggest-"            
	        }
	    ],
	    "DummyLogger": [
	        {
	            "Key": "dummyLoggerDefault"            
	        }        
    	]
	}`

	logger.InitialiseFromJson(configLogger)
}
