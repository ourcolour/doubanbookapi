package services

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	// Book
	AddOrUpdateBook(bookInfo *entities.BookInfo) (*entities.BookInfo, error)
	GetBook(id string) (*entities.BookInfo, error)
	GetBookByIsbn(isbn string) (*entities.BookInfo, error)
	GetBookByAuthor(author string) (*entities.BookInfo, error)
	GetBookByTitle(title string) ([]*entities.BookInfo, error)

	GetBookBy(criteriaMap *hashmap.Map) (*datasources.DataSource, error)
	PagedGetBookBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error)

	// Buy record
	AddBuyRecord(*entities.BuyRecord) (*entities.BuyRecord, error)
	GetBuyRecordBy(criteriaMap *hashmap.Map) (*datasources.DataSource, error)
	PagedGetBuyRecordBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error)

	// Rank
	GetRankInIsbn(isbnList *arraylist.List) ([]*entities.BookInfo, error)
}
