package controllers

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/entities/responses"
	"net/http"
)

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
