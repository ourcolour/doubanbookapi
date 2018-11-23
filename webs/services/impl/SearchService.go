package impl

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/olivere/elastic"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/entities/es"
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

	// 转换  Book -> ES_Book
	bookList := ds.DataList.([]*entities.Book)
	var esBookList []*es.ES_Book = make([]*es.ES_Book, 0)
	for _, book := range bookList {
		if nil == book {
			continue
		}
		esBook := es.NewESBookByBook(book)
		esBookList = append(esBookList, esBook)
	}

	if nil == esBookList || len(esBookList) < 1 {
		return delCount, addCount, err
	}

	// 清理数据
	delCount, err = this.DeleteAllBook()
	if nil != err {
		return delCount, addCount, err
	}

	// 同步到es
	addCount, err = searchBL.BatchAddBook(esBookList)

	return delCount, addCount, err
}
