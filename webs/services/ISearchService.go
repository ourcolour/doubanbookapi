package services

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/olivere/elastic"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
)

type ISearchService interface {
	SearchBook(keyword string, criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error)
	SearchBookByQuery(query elastic.Query, pageSize int, pageNo int) (*datasources.PagedDataSource, error)
	DeleteAllBook() (int64, error)
	SyncBook() (int64, int64, error)
}
