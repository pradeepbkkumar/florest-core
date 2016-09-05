package http

import (
	"github.com/jabong/florest-core/src/common/constants"
	"net/http"
	"strconv"
)

type RequestHeader struct {
	ContentType   string
	Accept        string
	UserId        string
	SessionId     string
	AuthToken     string
	TransactionId string
	RequestId     string
	Timestamp     string
	UserAgent     string
	Referrer      string
	BucketsList   string
	Debug         bool
	ClientAppId   string
}

func GetReqHeader(req *http.Request) RequestHeader {
	return RequestHeader{
		ContentType:   req.Header.Get("Content-Type"),
		Accept:        req.Header.Get("Accept"),
		UserId:        req.Header.Get(constants.USER_ID),
		SessionId:     req.Header.Get(constants.SESSION_ID),
		AuthToken:     req.Header.Get(constants.AUTHTOKEN),
		TransactionId: req.Header.Get(constants.TRANSACTION_ID),
		RequestId:     req.Header.Get(constants.REQUEST_ID),
		Timestamp:     req.Header.Get("ts"),
		UserAgent:     req.Header.Get("User-Agent"),
		Referrer:      req.Header.Get("Referer"),
		BucketsList:   req.Header.Get("bucket"),
		Debug:         getBoolHeaderField(req, constants.DEBUG),
		ClientAppId:   req.Header.Get(constants.APPID),
	}
}

func getBoolHeaderField(req *http.Request, headerKey string) bool {
	value, err := strconv.ParseBool(req.Header.Get(headerKey))
	if err != nil {
		return false
	}
	return value
}
