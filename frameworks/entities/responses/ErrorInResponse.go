package responses

import (
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"log"
)

type ErrorInResponse struct {
	Error   error  `json:"error" bson:"error"`
	Message string `json:"message" bson:"message"`
}

func NewErrorInResponseByError(err error) *ErrorInResponse {
	if nil == err {
		log.Panicln(errs.ERR_INVALID_PARAMETERS)
	}

	return &ErrorInResponse{
		Error:   err,
		Message: err.Error(),
	}
}

func NewErrorInResponse(err error, message string) *ErrorInResponse {
	return &ErrorInResponse{
		Error:   err,
		Message: message,
	}
}
