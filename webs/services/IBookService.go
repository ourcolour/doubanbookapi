package services

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	// Book
	AddOrUpdateBook(bookInfo *entities.BookInfo) (*entities.BookInfo, error)
	GetBook(id string) (*entities.BookInfo, error)
	GetBookByIsbn(isbn string) (*entities.BookInfo, error)
	GetBookByAuthor(author string) (*entities.BookInfo, error)

	// Buy record
	AddBuyRecord(*entities.BuyRecord) (*entities.BuyRecord, error)
	GetBuyRecord(criteriaMap *hashmap.Map) ([]*entities.BuyRecord, error)

	// Rank
	GetRankInIsbn(isbnList *arraylist.List) ([]*entities.BookInfo, error)
}
