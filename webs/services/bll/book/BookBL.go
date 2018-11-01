package book

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/services/impl/dal"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"reflect"
	"sort"
	"time"
)

func AddOrUpdateBook(bookInfo *entities.BookInfo) (*entities.BookInfo, error) {
	var (
		result *entities.BookInfo
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
		foundValue, err := dal.FindOne(colName, selector)
		if nil != err {
			return result, err
		}

		if nil != foundValue { // 已经存在，更新现有记录
			jsonStr, err := utils.ToJsonString(foundValue)
			if nil != err {
				return result, err
			}
			result = entities.NewBookInfoByJson(jsonStr)
			result.UpdateTime = time.Now()

			err = dal.UpdateId(colName, result.Id, result)
		} else { // 不存在，新增记录
			result = bookInfo
			result.CreateTime = time.Now()

			err = dal.Insert(colName, result)
		}
	}

	return result, err
}

func GetBook(id string) (*entities.BookInfo, error) {
	var (
		val interface{}

		result *entities.BookInfo
		err    error
	)

	if "" == id {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"id": id,
		}

		val, err = dal.FindOne(colName, selector)

		if nil == err && nil != val {
			jsonStr, err := utils.ToJsonString(val)
			if nil != err {
				return result, err
			}
			result = entities.NewBookInfoByJson(jsonStr)
		}
	}

	return result, err
}

func GetBookAuthor(author string) (*entities.BookInfo, error) {
	var (
		val interface{}

		result *entities.BookInfo
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

		val, err = dal.FindOne(colName, selector)
		if nil == err && nil != val {
			jsonStr := utils.MustToJsonString(val)
			result = entities.NewBookInfoByJson(jsonStr)
		}
	}

	return result, err
}

func GetBookListByIsbn(isbnList *arraylist.List) ([]*entities.BookInfo, error) {
	var (
		result []*entities.BookInfo
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
		typ := reflect.TypeOf(entities.BookInfo{})
		dataList, err := dal.FindAll(colName, selector, typ)
		if nil == err {
			for _, cur := range dataList {
				val := cur.(*entities.BookInfo)
				if nil != val {
					result = append(result, val)
				}
			}
		}
	}

	return result, err
}

func Sort(bookArray []*entities.BookInfo) []*entities.BookInfo {
	// 排序
	score := func(b1, b2 *entities.BookInfo) bool {
		score1 := float64(b1.Rating.NumRaters) * b1.Rating.Average
		score2 := float64(b2.Rating.NumRaters) * b2.Rating.Average
		return score1 > score2
	}
	SortBy(score).Sort(bookArray)

	return bookArray
}

type SortBy func(b1 *entities.BookInfo, b2 *entities.BookInfo) bool

func (by SortBy) Sort(books []*entities.BookInfo) {
	bookSlice := &BookScoreSorter{
		Books: books,
		By:    by,
	}
	sort.Sort(bookSlice)
}

type BookScoreSorter struct {
	Books []*entities.BookInfo
	By    func(b1 *entities.BookInfo, b2 *entities.BookInfo) bool
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

func GetBookByIsbn(isbn string) (*entities.BookInfo, error) {
	var (
		val interface{}

		result *entities.BookInfo
		err    error
	)

	if "" == isbn {
		err = errs.ERR_INVALID_PARAMETERS
	} else {
		colName := "sl_book_new"
		selector := bson.M{
			"isbn13": isbn,
		}

		val, err = dal.FindOne(colName, selector)
		if nil == err && nil != val {
			jsonStr := utils.MustToJsonString(val)
			result = entities.NewBookInfoByJson(jsonStr)
		}
	}

	return result, err
}
