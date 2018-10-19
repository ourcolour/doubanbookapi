package controllers

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/entities/responses"
	"net/http"
)

func Json(c *gin.Context, data interface{}, err error) {
	resp := MakeJsonResponse(data, err)

	c.JSON(http.StatusOK, resp)
}
