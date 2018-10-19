package services

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IDoubanApiService interface {
	GetBookByIsbn(c *gin.Context) (*entities.BookInfo, error)
}
