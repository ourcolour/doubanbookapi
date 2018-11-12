package v1

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/webs/entities"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"regexp"
	"strings"
)

var (
	searchService = NewSearchService()
)

func SearchBookController_ActionDispatcher(c *gin.Context) {
	action := c.Param("action")
	keyword := "*"

	// 1. /search/:field/:keyword
	m1Regex := regexp.MustCompile(`^/([^/]+)$`)

	if m1Regex.MatchString(action) {
		keyword = m1Regex.FindStringSubmatch(action)[1]
	}

	getBookByField(c, keyword)
}

func getBookByField(c *gin.Context, keyword string) {
	// 参数
	pageNo, pageSize := datasources.ParsePageArgs(c)

	var criteriaMap *hashmap.Map = hashmap.New()
	criteriaMap.FromJSON(MustGetRequestBody(c))

	isbnList := fetchIsbnList(searchService.SearchBook(keyword, criteriaMap, pageSize, pageNo))
	Json(c, isbnList, nil)
}

func fetchIsbnList(pds *datasources.PagedDataSource, err error) []string {
	var result []string = make([]string, 0)

	if nil == pds || nil == pds.DataList || nil != err {
		return result
	}

	var dataList []*entities.Book = pds.DataList.([]*entities.Book)
	for _, cur := range dataList {
		item := cur
		isbn13 := strings.Join([]string{item.Isbn13, item.Title}, "|")
		result = append(result, isbn13)
	}

	return result
}

func SearchBookController_SyncBook(c *gin.Context) {
	delCount, addCount, err := searchService.SyncBook()
	Json(
		c,
		map[string]int64{
			"deleted":  delCount,
			"inserted": addCount,
		},
		err)
}
