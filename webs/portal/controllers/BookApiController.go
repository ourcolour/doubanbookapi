package controllers

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/webs/services"
	"iamcc.cn/doubanbookapi/webs/services/impl"
)

func BookApiController_QueryById(c *gin.Context) {
	Json(c, nil, errs.ERR_NOT_IMPLEMENTED)
}

func BookApiController_QueryByIsbn(c *gin.Context) {
	var DoubanApiService services.IDoubanApiService = impl.NewDoubanApiService()

	data, err := DoubanApiService.GetBookByIsbn(c)

	Json(c, data, err)
}
