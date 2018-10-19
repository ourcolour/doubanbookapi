package responses

import (
	"encoding/json"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"log"
	"strings"
)

type MsgCode int

const (
	MSG_CODE_OK      MsgCode = 0
	MSG_CODE_WARNING MsgCode = -1
	MSG_CODE_ERROR   MsgCode = -2
)

const (
	MSG_TEXT_OK      = "执行成功。"
	MSG_TEXT_WARNING = "已经执行，但似乎出现了点问题。"
	MSG_TEXT_ERROR   = "服务器繁忙，请稍后再试。"
)

type JsonResponse struct {
	Code    MsgCode `json:"code"`
	Message string  `json:"message"`

	Data  interface{}      `json:"data"`
	Error *ErrorInResponse `json:"error"`
}

func NewJsonResponse(code MsgCode, message string, data interface{}, err error) *JsonResponse {
	var (
		result *JsonResponse = &JsonResponse{}
	)

	if MSG_CODE_OK != code && MSG_CODE_WARNING != code && MSG_CODE_ERROR != code {
		log.Panicln(errs.ERR_INVALID_PARAMETERS.Error())
	}

	// 参数：code
	result.Code = code

	// 参数：message
	if 0 == strings.Compare("", message) { // 如果没有指定 message
		if nil != err { // 如果遇到错误
			result.Message = err.Error()
		} else { // 如果没有错误
			switch code {
			case MSG_CODE_WARNING:
				result.Message = MSG_TEXT_WARNING
				break
			case MSG_CODE_ERROR:
				result.Message = MSG_TEXT_ERROR
				break
			default:
				result.Message = MSG_TEXT_OK
				break
			}
		}
	} else {
		result.Message = message
	}

	// 参数：data
	if nil != data {
		result.Data = data
	}

	// 参数：err
	if nil != err {
		result.Error = NewErrorInResponseByError(err)
	}

	return result
}

func NewJsonResponse_OK(data interface{}) *JsonResponse {
	return NewJsonResponse(MSG_CODE_OK, MSG_TEXT_OK, data, nil)
}

func NewJsonResponse_Warning(message string, data interface{}, error error) *JsonResponse {
	return NewJsonResponse(MSG_CODE_WARNING, message, data, error)
}

func NewJsonResponse_Error(error error) *JsonResponse {
	return NewJsonResponse_ErrorWithMessage("", error)
}
func NewJsonResponse_ErrorWithMessage(message string, error error) *JsonResponse {
	return NewJsonResponse(MSG_CODE_ERROR, message, nil, error)
}

func (this *JsonResponse) ToJsonString() (string, error) {
	var (
		result string
		err    error
	)

	bytes, err := json.Marshal(this)
	if nil == err {
		result = string(bytes)
	}

	return result, err
}

func MakeJsonResponse(data interface{}, err error) *JsonResponse {
	var resp *JsonResponse

	if nil != err {
		resp = NewJsonResponse_Error(err)
	} else {
		resp = NewJsonResponse_OK(data)
	}

	return resp
}
