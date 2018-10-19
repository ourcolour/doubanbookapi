package controllers

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/webs/services"
	"iamcc.cn/doubanbookapi/webs/services/impl"
)

func BookController_QueryById(c *gin.Context) {
	var BookService services.IBookService = impl.NewBookService()

	data, err := BookService.Get(c)

	Json(c, data, err)
}

func BookController_QueryByIsbn(c *gin.Context) {
	data, err := impl.NewBookService().GetByIsbn(c)

	Json(c, data, err)
}
