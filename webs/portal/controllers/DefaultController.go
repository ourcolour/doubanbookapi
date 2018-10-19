package controllers

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/webs/services"
	"iamcc.cn/doubanbookapi/webs/services/impl"
)

func DefaultController_Version(c *gin.Context) {
	var defaultService services.IDefaultService = impl.NewDefaultService()

	data := defaultService.Version(c)

	Json(c, data, nil)
}
