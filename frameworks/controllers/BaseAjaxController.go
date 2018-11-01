package controllers

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/entities/responses"
	"io/ioutil"
	"net/http"
)

func JsonWithStatusCode(c *gin.Context, data interface{}, err error, httpStatusCode int) {
	resp := MakeJsonResponse(data, err)

	c.JSON(httpStatusCode, resp)
}

func Json(c *gin.Context, data interface{}, err error) {
	resp := MakeJsonResponse(data, err)

	var httpStatusCode int
	switch resp.Code {
	case MSG_CODE_WARNING:
		httpStatusCode = http.StatusAccepted
		break
	case MSG_CODE_ERROR:
		httpStatusCode = http.StatusBadRequest
		break
	default:
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, resp)
}

func MustGetRequestBody(c *gin.Context) []byte {
	result, err := GetRequestBody(c)

	if nil != err {
		Json(c, result, err)
		return nil
	}

	return result
}

func GetRequestBody(c *gin.Context) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	// 参数
	result, err = ioutil.ReadAll(c.Request.Body)

	return result, err
}
