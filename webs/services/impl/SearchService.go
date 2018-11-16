package impl

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/olivere/elastic"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	searchBL "iamcc.cn/doubanbookapi/webs/services/bll/search"
)

type SearchService struct {
}

func NewSearchService() services.ISearchService {
	return services.ISearchService(&SearchService{})
}

func (this *SearchService) SearchBook(keyword string, criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	return searchBL.SearchBook(keyword, criteriaMap, pageSize, pageNo)
}

func (this *SearchService) SearchBookByQuery(query elastic.Query, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	return searchBL.SearchBookByQuery(query, pageSize, pageNo)
}

func (this *SearchService) DeleteAllBook() (int64, error) {
	return searchBL.DeleteAllBook()
}

func (this *SearchService) SyncBook() (int64, int64, error) {
	var (
		addCount int64 = 0
		delCount int64 = 0
		err      error
	)

	// 查询已有记录
	bookService := NewBookService()
	ds, err := bookService.GetBookListBy(nil)
	if nil != err {
		return delCount, addCount, err
	}
	bookList := ds.DataList.([]*entities.Book)

	if nil == bookList || len(bookList) < 1 {
		return delCount, addCount, err
	}

	// 清理数据
	delCount, err = this.DeleteAllBook()
	if nil != err {
		return delCount, addCount, err
	}

	// 同步到es
	addCount, err = searchBL.BatchAddBook(bookList)

	return delCount, addCount, err
}
