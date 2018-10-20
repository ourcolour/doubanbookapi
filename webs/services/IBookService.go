package services

import (
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	//Add(c *gin.Context) (*entities.BookInfo, error)
	Get(id string) (*entities.BookInfo, error)
	GetByIsbn(isbn string) (*entities.BookInfo, error)
	GetByAuthor(author string) (*entities.BookInfo, error)
}
