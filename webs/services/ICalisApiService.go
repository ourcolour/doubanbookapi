package services

import "iamcc.cn/doubanbookapi/webs/entities"

type ICalisApiService interface {
	GetCipByIsbn(isbn string) ([]string, error)
	UpdateLocalBookCip() ([]*entities.Book, error)
}
