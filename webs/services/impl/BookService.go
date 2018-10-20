package impl

import (
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/webs/entities"
	"iamcc.cn/doubanbookapi/webs/services"
	bookBL "iamcc.cn/doubanbookapi/webs/services/bll/book"
)

type BookService struct {
}

func NewBookService() services.IBookService {
	var result services.IBookService = &BookService{}
	return result
}

func (this *BookService) Get(id string) (*entities.BookInfo, error) {
	// 参数
	if "" == id {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.Get(id)

	return data, err
}

func (this *BookService) GetByIsbn(isbn string) (*entities.BookInfo, error) {
	// 参数
	if "" == isbn {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.GetByIsbn(isbn)

	// 如果本地没有结果，去豆瓣查
	if nil == err && nil == data {
		data, err = NewDoubanApiService().GetBookByIsbn(isbn)

		// 保存
		err = bookBL.Add(data)
	}

	return data, err
}

func (this *BookService) GetByAuthor(author string) (*entities.BookInfo, error) {
	// 参数
	if "" == author {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 调用
	data, err := bookBL.GetAuthor(author)

	return data, err
}
