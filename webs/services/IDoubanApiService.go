package services

import (
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IDoubanApiService interface {
	GetBookByIsbn(isbn string) (*entities.BookInfo, error)
}
