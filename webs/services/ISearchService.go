package services

import (
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
)

type ISearchService interface {
	SearchBook(keyword string, pageSize int, pageNo int) (*datasources.PagedDataSource, error)
	SyncBook() (int64, int64, error)
}
