package impl

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	bll "iamcc.cn/doubanbookapi/webs/services/bll/calisapi"
	"log"
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

func (this *CalisApiService) UpdateLocalBookCip() ([]*entities.Book, error) {
	var (
		bookService services.IBookService = NewBookService()

		result []*entities.Book = make([]*entities.Book, 0)
		err    error
	)

	// 查找本地记录
	criteriaMap := hashmap.New()
	criteriaMap.Put("cipsExists", false)
	ds, err := bookService.GetBookBy(criteriaMap)

	// 准备返回的数据
	// 调用 Calid 服务，更新 cip
	dataList := ds.DataList.([]*entities.Book)
	for _, book := range dataList {
		// 获取 cip
		cipArray, err := this.GetCipByIsbn(book.Isbn13)
		if nil == err {
			book.Cips = cipArray
		}

		book, err = bookService.AddOrUpdateBook(book)
		if nil == err {
			log.Printf("%s <<%s>> --------> %v", book.Isbn13, book.Title, book.Cips)
			result = append(result, book)
		}
	}

	return result, err
}
