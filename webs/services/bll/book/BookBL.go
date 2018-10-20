package book

import (
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/services/impl/dal"
	"iamcc.cn/doubanbookapi/webs/entities"
	"time"
)

func Add(bookInfo *entities.BookInfo) error {
	var (
		err error
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
			return err
		}

		if nil != foundValue { // 已经存在，更新现有记录
			oldBookInfo := foundValue.(entities.BookInfo)
			oldBookInfo.UpdateTime = time.Now()

			err = dal.Update(colName, selector, oldBookInfo)
		} else { // 不存在，新增记录
			bookInfo.CreateTime = time.Now()

			err = dal.Insert(colName, bookInfo)
		}
	}

	return err
}

func Get(id string) (*entities.BookInfo, error) {
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
			result = entities.ToBookInfo(val)
		}
	}

	return result, err
}

func GetAuthor(author string) (*entities.BookInfo, error) {
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
			result = entities.ToBookInfo(val)
		}
	}

	return result, err
}

func GetByIsbn(isbn string) (*entities.BookInfo, error) {
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
			result = entities.ToBookInfo(val)
		}
	}

	return result, err
}
