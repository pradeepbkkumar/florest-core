package service

import (
	"github.com/jabong/florest-core/src/common/constants"
)

const (
	SWAGGER_ALLOWED_HEADERS = constants.SESSION_ID + ", Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, " + constants.TRANSACTION_ID + ", " + constants.USER_ID
)
