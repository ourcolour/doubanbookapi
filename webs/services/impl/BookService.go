package impl

import (
	"github.com/gin-gonic/gin"
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

func (this *BookService) Get(c *gin.Context) (*entities.BookInfo, error) {
	// 参数
	id := c.Param("id")

	// 调用
	data, err := bookBL.Get(id)

	return data, err
}

func (this *BookService) GetByIsbn(c *gin.Context) (*entities.BookInfo, error) {
	// 参数
	isbn := c.Param("isbn")

	// 调用
	data, err := bookBL.GetByIsbn(isbn)

	// 如果本地没有结果，去豆瓣查
	if nil == err && nil == data {
		data, err = NewDoubanApiService().GetBookByIsbn(c)

		// 保存
		err = bookBL.Add(data)
	}

	return data, err
}
