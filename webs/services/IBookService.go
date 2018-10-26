package services

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	// Book
	AddBook(bookInfo *entities.BookInfo) (*entities.BookInfo, error)
	GetBook(id string) (*entities.BookInfo, error)
	GetBookByIsbn(isbn string) (*entities.BookInfo, error)
	GetBookByAuthor(author string) (*entities.BookInfo, error)

	// Buy record
	AddBuyRecord(*entities.BuyRecord) (*entities.BuyRecord, error)
	GetBuyRecord(criteriaMap *hashmap.Map) ([]*entities.BuyRecord, error)
}
