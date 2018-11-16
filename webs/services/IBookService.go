package services

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/webs/entities"
)

type IBookService interface {
	// Book
	AddOrUpdateBook(bookInfo *entities.Book) (*entities.Book, error)
	GetBook(id string) (*entities.Book, error)
	GetBookByIsbn(isbn string) (*entities.Book, error)
	GetBookByAuthor(author string) (*entities.Book, error)
	GetBookByTitle(title string) ([]*entities.Book, error)

	GetBookListBy(criteriaMap *hashmap.Map) (*datasources.DataSource, error)
	PagedGetBookListBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error)

	// Buy record
	AddBuyRecord(*entities.BuyRecord) (*entities.BuyRecord, error)
	GetBuyRecordBy(criteriaMap *hashmap.Map) (*datasources.DataSource, error)
	PagedGetBuyRecordBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error)

	// Rank
	GetRankInIsbn(isbnList *arraylist.List) ([]*entities.Book, error)

	RemoveDuplicate() error
}
