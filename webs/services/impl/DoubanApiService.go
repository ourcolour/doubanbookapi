package impl

import (
	"github.com/gin-gonic/gin"
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

func (this *DoubanApiService) GetBookByIsbn(c *gin.Context) (*entities.BookInfo, error) {
	// 参数
	isbn := c.Param("isbn")

	// 调用
	data, err := doubanapiBL.GetBookByIsbn(isbn)

	return data, err
}
