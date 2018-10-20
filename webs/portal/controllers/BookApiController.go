package controllers

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"log"
)

func BookApiController_Query(c *gin.Context) {
	action := c.Param("action")
	params := c.Param("params")
	log.Printf("A: %s P: %s", action, params)

	// 1. /douban/:isbn
	// 2. /douban/isbn/:isbn
	// 3. /douban/id/:id
	var mode int = -1
	if "" != action {
		if "isbn" == action {
			mode = 2
		} else if "id" == action {
			mode = 3
		} else {
			mode = 1
		}
	}

	doubanService := NewDoubanApiService()
	switch mode {
	case 1:
		data, err := doubanService.GetBookByIsbn(action)
		Json(c, data, err)
		break
	case 2:
		data, err := doubanService.GetBookByIsbn(params)
		Json(c, data, err)
		break
	case 3:
		Json(c, nil, errs.ERR_NOT_IMPLEMENTED)
		break
	default:
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
		break
	}
}
