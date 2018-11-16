package impl

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	bll "iamcc.cn/doubanbookapi/webs/services/bll/calisapi"
	"strings"
)

type CalisApiService struct {
}

func NewCalisApiService() services.ICalisApiService {
	return services.ICalisApiService(&CalisApiService{})
}

func (this *CalisApiService) GetCipByIsbn(isbn string) ([]string, error) {
	// 定义
	var (
		isbnParsed string
		result     []string

		err error
	)

	// 参数
	if "" == strings.TrimSpace(isbn) {
		return nil, err
	}

	// 解析
	isbnParsed, err = bll.ParseIsbn(isbn, 0)
	if nil != err {
		return nil, err
	}
	result, err = bll.GetCipByIsbn(isbnParsed)
	if nil != err {
		return nil, err
	}

	return result, err
}

func (this *CalisApiService) UpdateLocalBookCip() (map[string][]string, error) {
	// 参数
	var (
		bookService services.IBookService = NewBookService()
		result      map[string][]string
		err         error
	)

	// 查找本地记录
	criteriaMap := hashmap.New()
	criteriaMap.Put("cipsExists", false)
	//criteriaMap.Put("cipsHasElements", 0)
	ds, err := bookService.GetBookListBy(criteriaMap)
	if nil != err {
		return nil, err
	}

	// 准备返回的数据
	succeededList := make([]string, 0)
	failedList := make([]string, 0)

	// 调用 Calid 服务，更新 cip
	dataList := ds.DataList.([]*entities.Book)
	for _, book := range dataList {
		// 获取 cip
		cipArray, err := this.GetCipByIsbn(book.Isbn13)
		if nil == err {
			book.Cips = cipArray
		}

		// 更新本地 book 的 cip 字段
		newBook, err := bookService.AddOrUpdateBook(book)
		oldObjectId := book.ObjectId.Hex()
		newObjectId := newBook.ObjectId.Hex()

		if nil != err {
			failedList = append(failedList, oldObjectId)
		} else {
			succeededList = append(succeededList, newObjectId)
		}
	}

	// 返回结果
	result = map[string][]string{
		"succeeded": succeededList,
		"failed":    failedList,
	}

	return result, err
}
