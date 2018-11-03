package search

import (
	"encoding/json"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	elasticsearchDAL "iamcc.cn/doubanbookapi/frameworks/services/impl/dal/elasticsearch"
	"iamcc.cn/doubanbookapi/webs/entities"
)

func SearchBook(keyword string, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	var (
		result *datasources.PagedDataSource
		err    error
	)

	indexName := "sl_book_new"
	typeName := "book"
	jsonList, totalRecordCount, err := elasticsearchDAL.Search(indexName, typeName, keyword, pageSize, pageNo)
	pagination := datasources.NewPagination(totalRecordCount, pageSize, pageNo)
	if nil != err {
		return result, err
	}

	// 转码 json -> obj
	dataList, err := parseBookList(jsonList)

	// PagedDataSource
	result = datasources.NewPagedDataSource(pagination, dataList)

	return result, err
}

func DeleteAllBook() (int64, error) {
	indexName := "sl_book_new"
	typeName := "book"

	return elasticsearchDAL.DeleteAll(indexName, typeName)
}

func AddBook(book *entities.BookInfo) (bool, error) {
	res, err := BatchAddBook([]*entities.BookInfo{book})
	return 0 < res, err
}

func BatchAddBook(bookList []*entities.BookInfo) (int64, error) {
	var (
		result int64
		err    error
	)

	indexName := "sl_book_new"
	typeName := "book"

	// objList -> interfaceList
	var itfList []interface{} = make([]interface{}, 0)
	for _, book := range bookList {
		itfList = append(itfList, book)
	}

	result, err = elasticsearchDAL.BatchAdd(indexName, typeName, itfList)

	return result, err
}

func parseBook(data *json.RawMessage) (*entities.BookInfo, error) {
	var (
		result entities.BookInfo
		err    error
	)

	// 参数
	if nil == data {
		err = errs.ERR_INVALID_PARAMETERS
		return &result, err
	}

	// json -> obj
	err = json.Unmarshal(*data, &result)
	//err = json.Unmarshal([]byte(jsonString), &obj)
	if nil != err {
		return &result, err
	}

	return &result, err
}

func parseBookList(dataList []*json.RawMessage) ([]*entities.BookInfo, error) {
	var (
		result []*entities.BookInfo = make([]*entities.BookInfo, 0)
		err    error
	)

	for _, cur := range dataList {
		if nil != cur {
			obj, err := parseBook(cur)
			if nil != err {
				break
			}
			result = append(result, obj)
		}
	}

	return result, err
}
