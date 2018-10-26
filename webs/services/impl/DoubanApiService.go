package impl

import (
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	doubanapiBL "iamcc.cn/doubanbookapi/webs/services/bll/doubanapi"
)

type DoubanApiService struct {
}

func NewDoubanApiService() services.IDoubanApiService {
	var result services.IDoubanApiService = &DoubanApiService{}
	return result
}

func (this *DoubanApiService) GetBookByIsbn(isbn string) (*entities.BookInfo, error) {
	// 参数
	if "" == isbn {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := doubanapiBL.GetBookByIsbn(isbn)

	// 保存到本地
	if nil == err && nil != data {
		data, err = NewBookService().AddBook(data)
	}

	return data, err
}
