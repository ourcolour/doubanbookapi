package buyrecord

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/services/impl/dal"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"reflect"
	"time"
)

func GetList(criteriaMap *hashmap.Map) ([]*entities.BuyRecord, error) {
	var (
		result []*entities.BuyRecord = make([]*entities.BuyRecord, 0)
		err    error
	)

	// Build query criterials
	var (
		selector           bson.M   = bson.M{}
		innerSelectorArray []bson.M = []bson.M{}
		val                interface{}
		found              bool
	)
	if val, found = criteriaMap.Get("fromDate"); found {
		q := bson.M{"fromDate": bson.M{"$gt": utils.MustParseDatetime(val.(string))}}
		innerSelectorArray = append(innerSelectorArray, q)
	}
	if val, found = criteriaMap.Get("toDate"); found {
		q := bson.M{"toDate": bson.M{"$lt": utils.MustParseDatetime(val.(string))}}
		innerSelectorArray = append(innerSelectorArray, q)
	}
	if val, found = criteriaMap.Get("store"); found {
		q := bson.M{"store": bson.M{"$regex": val}}
		innerSelectorArray = append(innerSelectorArray, q)
	}
	// paginations
	var pageNo int = 0
	if val, found = criteriaMap.Get("pageNo"); found {
		pageNo = val.(int)
	}
	var pageSize int = 0
	if val, found = criteriaMap.Get("pageSize"); found {
		pageSize = val.(int)
	}

	// Execute query
	colName := "sl_buy_record"
	if len(innerSelectorArray) > 1 {
		selector = bson.M{"$and": innerSelectorArray}
	} else if len(innerSelectorArray) > 0 {
		selector = innerSelectorArray[0]
	}

	// 转数组
	//totalRecordCount := 100
	//currentPage := 1
	//pageSize := 10
	//pagination := datasources.NewPagination(totalRecordCount, currentPage, pageSize)

	skip := pageSize * (pageNo - 1)
	limit := pageSize

	dataList, err := dal.FindList(colName, selector, reflect.TypeOf(entities.BuyRecord{}), skip, limit)
	if nil == err {
		for _, cur := range dataList {
			val := cur.(*entities.BuyRecord)
			result = append(result, val)
		}
	}

	return result, err
}

func Add(buyRecord *entities.BuyRecord) error {
	var (
		err error
	)

	// 参数
	if nil == buyRecord {
		err = errs.ERR_INVALID_PARAMETERS
	}

	// 检查是否存在
	found, err := Exists(buyRecord)
	if nil == err {
		if found {
			err = errs.ERR_DUPLICATED
		} else {
			// 如果不存在，则添加
			buyRecord.CreateTime = time.Now()
			err = dal.Insert("sl_buy_record", buyRecord)
		}
	}

	return err
}

func Exists(buyRecord *entities.BuyRecord) (bool, error) {
	var (
		result bool
		err    error
	)

	if nil == buyRecord {
		err = errs.ERR_INVALID_PARAMETERS
		result = false
	} else {
		colName := "sl_buy_record"
		selector := bson.M{
			"$or": []bson.M{
				bson.M{"isbn": buyRecord.Isbn},
				bson.M{"store": buyRecord.Store},
				bson.M{"buy_date": buyRecord.BuyDate},
			},
		}

		// 检查是否已经存在相同记录
		if count, err := dal.Count(colName, selector); nil != err {
			result = false
		} else {
			result = 0 < count
		}
	}

	return result, err
}
