package book

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	mongoDAL "iamcc.cn/doubanbookapi/frameworks/services/impl/dal/mongodb"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"reflect"
	"sort"
	"time"
)

func AddOrUpdateBook(bookInfo *entities.Book) (*entities.Book, error) {
	var (
		result *entities.Book
		err    error
	)

	if nil == bookInfo {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"$or": []bson.M{
				bson.M{"isbn10": bookInfo.Isbn10},
				bson.M{"isbn13": bookInfo.Isbn13},
				bson.M{"id": bookInfo.Id},
			},
		}

		// 检查是否已经存在相同记录
		foundValue, err := mongoDAL.FindOne(colName, selector)
		if nil != err {
			return result, err
		}

		if nil != foundValue { // 已经存在，更新现有记录
			result = bookInfo
			result.UpdateTime = time.Now()

			book := entities.NewBookByInterface(foundValue)
			err = mongoDAL.UpdateId(colName, book.ObjectId, result)
		} else { // 不存在，新增记录
			result = bookInfo
			result.CreateTime = time.Now()

			err = mongoDAL.Insert(colName, result)
		}
	}

	return result, err
}

func GetBook(id string) (*entities.Book, error) {
	var (
		val interface{}

		result *entities.Book
		err    error
	)

	if "" == id {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"id": id,
		}

		val, err = mongoDAL.FindOne(colName, selector)

		if nil == err && nil != val {
			jsonStr, err := utils.ToJsonString(val)
			if nil != err {
				return result, err
			}
			result = entities.NewBookByJson(jsonStr)
		}
	}

	return result, err
}

func GetBookAuthor(author string) (*entities.Book, error) {
	var (
		val interface{}

		result *entities.Book
		err    error
	)

	if "" == author {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"author": bson.RegEx{
				Pattern: author,
				Options: "i",
			},
		}

		val, err = mongoDAL.FindOne(colName, selector)
		if nil == err && nil != val {
			jsonStr := utils.MustToJsonString(val)
			result = entities.NewBookByJson(jsonStr)
		}
	}

	return result, err
}

func GetBookListBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
	var (
		result *datasources.PagedDataSource
		err    error
	)

	// Build query criterials
	var (
		selector           bson.M   = bson.M{}
		innerSelectorArray []bson.M = []bson.M{}
		val                interface{}
		found              bool
	)
	if nil != criteriaMap {
		if val, found = criteriaMap.Get("title"); found {
			q := bson.M{"title": bson.M{"$regex": val}}
			innerSelectorArray = append(innerSelectorArray, q)
		}
		if val, found = criteriaMap.Get("subtitle"); found {
			q := bson.M{"subtitle": bson.M{"$regex": val}}
			innerSelectorArray = append(innerSelectorArray, q)
		}
		if val, found = criteriaMap.Get("cipsExists"); found {
			q := bson.M{"cips": bson.M{"$exists": val}}
			innerSelectorArray = append(innerSelectorArray, q)
		}
	}
	// paginations
	if pageNo < 1 {
		pageNo = 1
	}
	usePagination := true
	if pageSize < 1 {
		usePagination = false
	}

	// Execute query
	colName := "sl_book_new"
	if len(innerSelectorArray) > 1 {
		selector = bson.M{"$and": innerSelectorArray}
	} else if len(innerSelectorArray) > 0 {
		selector = innerSelectorArray[0]
	}

	skip := pageSize * (pageNo - 1)
	limit := pageSize

	// 查询
	var (
		itfList          []interface{}
		totalRecordCount int64 = 0
	)
	if usePagination { // 分页
		itfList, totalRecordCount, err = mongoDAL.FindList(colName, selector, reflect.TypeOf(entities.Book{}), skip, limit)
	} else { // 不分页
		itfList, err = mongoDAL.FindAll(colName, selector, reflect.TypeOf(entities.Book{}))
	}
	var dataList []*entities.Book = make([]*entities.Book, 0)
	if nil == err {
		for _, cur := range itfList {
			val := cur.(*entities.Book)
			dataList = append(dataList, val)
		}
	}

	// 数据源
	result = &datasources.PagedDataSource{
		DataList:   dataList,
		Pagination: datasources.NewPagination(totalRecordCount, pageSize, pageNo),
	}

	return result, err
}

func GetBookTitle(title string) ([]*entities.Book, error) {
	var (
		result []*entities.Book
		err    error
	)

	if "" == title {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{"$or": []bson.M{
			bson.M{"title": bson.RegEx{Pattern: title, Options: "i"}},
			bson.M{"subtitle": bson.RegEx{Pattern: title, Options: "i"}},
		}}

		typ := reflect.TypeOf(entities.Book{})
		dataList, err := mongoDAL.FindAll(colName, selector, typ)
		if nil == err {
			for _, cur := range dataList {
				val := cur.(*entities.Book)
				if nil != val {
					result = append(result, val)
				}
			}
		}
	}

	return result, err
}

func GetBookListByIsbn(isbnList *arraylist.List) ([]*entities.Book, error) {
	var (
		result []*entities.Book
		err    error
	)

	if nil == isbnList || isbnList.Empty() {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		isbnSelectorArray := []bson.M{}
		itr := isbnList.Iterator()
		for itr.Next() {
			isbn := itr.Value().(string)
			curSelector := bson.M{"isbn13": isbn}
			isbnSelectorArray = append(isbnSelectorArray, curSelector)
		}

		// 查询
		colName := "sl_book_new"
		selector := bson.M{"$or": isbnSelectorArray}

		typ := reflect.TypeOf(entities.Book{})
		dataList, err := mongoDAL.FindAll(colName, selector, typ)
		if nil == err {
			for _, cur := range dataList {
				val := cur.(*entities.Book)
				if nil != val {
					result = append(result, val)
				}
			}
		}
	}

	return result, err
}

func Sort(bookArray []*entities.Book) []*entities.Book {
	// 排序
	score := func(b1, b2 *entities.Book) bool {
		score1 := float64(b1.Rating.NumRaters) * b1.Rating.Average
		score2 := float64(b2.Rating.NumRaters) * b2.Rating.Average
		return score1 > score2
	}
	SortBy(score).Sort(bookArray)

	return bookArray
}

type SortBy func(b1 *entities.Book, b2 *entities.Book) bool

func (by SortBy) Sort(books []*entities.Book) {
	bookSlice := &BookScoreSorter{
		Books: books,
		By:    by,
	}
	sort.Sort(bookSlice)
}

type BookScoreSorter struct {
	Books []*entities.Book
	By    func(b1 *entities.Book, b2 *entities.Book) bool
}

func (this *BookScoreSorter) Len() int {
	return len(this.Books)
}
func (this *BookScoreSorter) Swap(b1, b2 int) {
	this.Books[b1], this.Books[b2] = this.Books[b2], this.Books[b1]
}
func (this *BookScoreSorter) Less(b1, b2 int) bool {
	return this.By(this.Books[b1], this.Books[b2])
}

func GetBookByIsbn(isbn string) (*entities.Book, error) {
	var (
		val interface{}

		result *entities.Book
		err    error
	)

	if "" == isbn {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"isbn13": isbn,
		}

		val, err = mongoDAL.FindOne(colName, selector)
		if nil == err && nil != val {
			jsonStr := utils.MustToJsonString(val)
			result = entities.NewBookByJson(jsonStr)
		}
	}

	return result, err
}
