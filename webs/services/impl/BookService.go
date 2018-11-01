package impl

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	bookBL "iamcc.cn/doubanbookapi/webs/services/bll/book"
	buyrecordBL "iamcc.cn/doubanbookapi/webs/services/bll/buyrecord"
	"strings"
	"time"
)

type BookService struct {
}

func NewBookService() services.IBookService {
	var result services.IBookService = &BookService{}
	return result
}

func (this *BookService) GetRankInIsbn(isbnList *arraylist.List) ([]*entities.BookInfo, error) {
	var (
		result []*entities.BookInfo = []*entities.BookInfo{}
		err    error
	)

	// 参数
	if nil == isbnList || isbnList.Empty() {
		return result, err
	}

	bookList, err := bookBL.GetBookListByIsbn(isbnList)
	// 排序
	result = bookBL.Sort(bookList)

	return result, err
}

func (this *BookService) AddOrUpdateBook(bookInfo *entities.BookInfo) (*entities.BookInfo, error) {
	var (
		result *entities.BookInfo
		err    error
	)

	// 参数
	if nil == bookInfo {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	result, err = bookBL.AddOrUpdateBook(bookInfo)

	return result, err
}
func (this *BookService) GetBook(id string) (*entities.BookInfo, error) {
	// 参数
	if "" == id {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.GetBook(id)

	return data, err
}

func (this *BookService) GetBookByIsbn(isbn string) (*entities.BookInfo, error) {
	// 参数
	if "" == isbn {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.GetBookByIsbn(isbn)

	// 如果本地不存在，调用 douban 接口获取，并缓存在本地
	if nil == err && nil == data {
		data, err = NewDoubanApiService().GetBookByIsbn(isbn)
	}

	return data, err
}

func (this *BookService) GetBookByAuthor(author string) (*entities.BookInfo, error) {
	// 参数
	if 0 == strings.Compare("", author) {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.GetBookAuthor(author)

	return data, err
}

func (this *BookService) GetBuyRecord(criteriaMap *hashmap.Map) ([]*entities.BuyRecord, error) {
	var (
		result []*entities.BuyRecord
		err    error
	)

	// 参数
	result, err = buyrecordBL.GetList(criteriaMap)

	return result, err
}

func (this *BookService) AddBuyRecord(buyRecord *entities.BuyRecord) (*entities.BuyRecord, error) {
	// 参数
	if 0 == strings.Compare("", buyRecord.Isbn) {
		return nil, errs.ERR_INVALID_PARAMETERS
	} else if 0 == strings.Compare("", buyRecord.Store) {
		return nil, errs.ERR_INVALID_PARAMETERS
	} else if buyRecord.Quantity < 1 {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 修改实体
	buyRecord.CreateTime = time.Now()
	buyRecord.UpdateTime = utils.ZeroTime()

	// 添加
	err := buyrecordBL.Add(buyRecord)

	return buyRecord, err
}
