package search

import (
	"encoding/json"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/olivere/elastic"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	elasticsearchDAL "iamcc.cn/doubanbookapi/frameworks/services/impl/dal/elasticsearch"
	"iamcc.cn/doubanbookapi/webs/entities"
	"log"
	"strings"
)

func SearchBook(keyword string, criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	// 查询条件
	query := elastic.NewBoolQuery()
	// 关键词
	query = query.Must(elastic.NewQueryStringQuery(keyword))
	// 分支条件
	if nil != criteriaMap && !criteriaMap.Empty() {
		fieldList := criteriaMap.Keys()
		for _, fieldObj := range fieldList {
			field := fieldObj.(string)
			if valueObj, found := criteriaMap.Get(fieldObj); found {
				value := valueObj.(string)
				// ---
				if 0 == strings.Compare("isbn13", strings.ToLower(field)) {
					query = query.Filter(elastic.NewWildcardQuery(field, "*"+value+"*"))
				} else if 0 == strings.Compare("tags", strings.ToLower(field)) {
					query = query.Filter(elastic.NewMatchPhraseQuery(field+".title", value))
				} else if strings.Contains(strings.ToLower(field), "rating") {
					if strings.Contains(strings.ToLower(field), strings.ToLower("numRaters")) {
						// numRaters
						arr := strings.Split(value, ",")
						if 0 < len(arr) {
							q := elastic.NewRangeQuery("rating.numRaters")
							if 0 != strings.Compare("", strings.TrimSpace(arr[0])) {
								q = q.From(arr[0])
							} else {
								q = q.From(nil)
							}

							if 1 < len(arr) {
								if 0 != strings.Compare("", strings.TrimSpace(arr[1])) {
									q = q.Lte(arr[1])
								} else {
									q = q.To(nil)
								}
							}

							query = query.Filter(q)
						}
					} else if strings.Contains(strings.ToLower(field), strings.ToLower("average")) {
						// average
						arr := strings.Split(value, ",")
						if 0 < len(arr) {
							q := elastic.NewRangeQuery("rating.average")
							if 0 != strings.Compare("", strings.TrimSpace(arr[0])) {
								q = q.From(arr[0])
							} else {
								q = q.From(nil)
							}

							if 1 < len(arr) {
								if 0 != strings.Compare("", strings.TrimSpace(arr[1])) {
									q = q.Lte(arr[1])
								} else {
									q = q.To(nil)
								}
							}

							query = query.Filter(q)
						}
					}
				} else {
					query = query.Filter(elastic.NewSimpleQueryStringQuery(value).Field(field))
				}
				// ---
			}
		}
	}

	s, e := query.Source()
	log.Printf("[SEARCH] Search source: %s | %v", s, e)

	return SearchBookByQuery(query, pageSize, pageNo)
}

func SearchBookByQuery(query elastic.Query, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	var (
		result *datasources.PagedDataSource
		err    error
	)

	// 参数
	if nil == query {
		query = elastic.NewSimpleQueryStringQuery("")
	}

	indexName := "sl_book_new"
	typeName := "book"

	jsonList, totalRecordCount, err := elasticsearchDAL.Search(indexName, typeName, query, pageSize, pageNo)
	pagination := datasources.NewPagination(totalRecordCount, pageSize, pageNo)
	if nil != err {
		return result, err
	}

	// 转码 json -> obj
	dataList, err := parseBookList(jsonList)

	// PagedDataSource
	result = datasources.NewPagedDataSource(pagination, dataList)

	log.Printf("[SEARCH] Found %d record(s).", len(dataList))

	return result, err
}

func DeleteAllBook() (int64, error) {
	indexName := "sl_book_new"
	typeName := "book"

	return elasticsearchDAL.DeleteAll(indexName, typeName)
}

func AddBook(book *entities.Book) (bool, error) {
	res, err := BatchAddBook([]*entities.Book{book})
	return 0 < res, err
}

func BatchAddBook(bookList []*entities.Book) (int64, error) {
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

func parseBook(data *json.RawMessage) (*entities.Book, error) {
	var (
		result entities.Book
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

func parseBookList(dataList []*json.RawMessage) ([]*entities.Book, error) {
	var (
		result []*entities.Book = make([]*entities.Book, 0)
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
