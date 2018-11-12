package buyrecord

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	mongoDAL "iamcc.cn/doubanbookapi/frameworks/services/impl/dal/mongodb"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"reflect"
	"time"
)

func GetList(criteriaMap *hashmap.Map) (*datasources.DataSource, error) {
	pds, err := GetBuyRecordListBy(criteriaMap, 0, 0)
	return datasources.FromPagedDataSource(pds), err
}

func GetBuyRecordListBy(criteriaMap *hashmap.Map, pageSize int, pageNo int) (*datasources.PagedDataSource, error) {
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
	if pageNo < 1 {
		pageNo = 1
	}
	usePagination := true
	if pageSize < 1 {
		usePagination = false
	}

	// Execute query
	colName := "sl_buy_record"
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
		itfList, totalRecordCount, err = mongoDAL.FindList(colName, selector, reflect.TypeOf(entities.BuyRecord{}), skip, limit)
	} else { // 不分页
		itfList, err = mongoDAL.FindAll(colName, selector, reflect.TypeOf(entities.BuyRecord{}))
	}
	var dataList []*entities.BuyRecord = make([]*entities.BuyRecord, 0)
	if nil == err {
		for _, cur := range itfList {
			val := cur.(*entities.BuyRecord)
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

func AddBuyRecord(buyRecord *entities.BuyRecord) error {
	var (
		err error
	)

	// 参数
	if nil == buyRecord {
		err = errs.ERR_INVALID_PARAMETERS
	}

	// 检查是否存在
	criteriaMap := hashmap.New()
	criteriaMap.Put("isbn", buyRecord.Isbn)

	found, err := ExistsBuyRecordBy(criteriaMap)
	if nil == err {
		if found {
			err = errs.ERR_DUPLICATED
		} else {
			// 如果不存在，则添加
			buyRecord.CreateTime = time.Now()
			err = mongoDAL.Insert("sl_buy_record", buyRecord)
		}
	}

	return err
}

func ExistsBuyRecordBy(criteriaMap *hashmap.Map) (bool, error) {
	var (
		result bool
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
		if val, found = criteriaMap.Get("isbn"); found {
			q := bson.M{"isbn": bson.M{"$regex": val}}
			innerSelectorArray = append(innerSelectorArray, q)
		}
	}

	// Execute query
	if len(innerSelectorArray) > 1 {
		selector = bson.M{"$and": innerSelectorArray}
	} else if len(innerSelectorArray) > 0 {
		selector = innerSelectorArray[0]
	}

	// 检查是否已经存在相同记录
	colName := "sl_buy_record"
	if count, err := mongoDAL.Count(colName, selector); nil != err {
		result = false
	} else {
		result = 0 < count
	}

	return result, err
}
