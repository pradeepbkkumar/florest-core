package cache

import (
	"errors"
	"github.com/jabong/florest-core/src/common/logger"
)

const (
	ERR_NO_PLATFORM          = "Platform not found"
	ERR_INITIALIZATION       = "Initialization failed"
	ERR_GET_FAILURE          = "Failure in Get() method"
	ERR_SET_FAILURE          = "Failure in Set() method"
	ERR_GET_BATCH_FAILURE    = "Failure in GetBatch() method"
	ERR_DELETE_FAILURE       = "Failure in Delete() method"
	ERR_DELETE_BATCH_FAILURE = "Failure in DeleteBatch() method"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) (ret error) {
	errorString := "ErrorCode: " + errCode + ", developerMessage : " + developerMessage
	logger.Error(errorString)
	ret = errors.New(errorString)
	return ret
}
