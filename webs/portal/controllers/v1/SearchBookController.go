package v1

import (
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
)

func SearchBookController_GetBookByTitle(c *gin.Context) {
	// 参数
	keyword := c.DefaultQuery("k", "*")
	pageNo, pageSize := datasources.ParsePageArgs(c)

	searchService := NewSearchService()
	pagedDataSource, err := searchService.SearchBook(keyword, pageSize, pageNo)
	Json(c, pagedDataSource, err)

}

func SearchBookController_SyncBook(c *gin.Context) {
	searchService := NewSearchService()
	delCount, addCount, err := searchService.SyncBook()
	Json(
		c,
		map[string]int64{
			"deleted":  delCount,
			"inserted": addCount,
		},
		err)
}
