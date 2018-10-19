package services

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	//Add(c *gin.Context) (*entities.BookInfo, error)
	Get(c *gin.Context) (*entities.BookInfo, error)
	GetByIsbn(c *gin.Context) (*entities.BookInfo, error)
}
