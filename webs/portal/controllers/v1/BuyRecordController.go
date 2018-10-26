package v1

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/frameworks/entities/datasources"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"io/ioutil"
	"net/http"
)

func BuyRecordController_Add(c *gin.Context) {
	// 参数
	type QueryInfo struct {
		Isbn     string `json:"isbn" bson:"isbn"`
		Store    string `json:"store" bson:"store"`
		Quantity int    `json:"quantity" bson:"quantity"`

		BuyDate    string `json:"buy_date" bson:"buy_date"`
		CreateTime string `json:"create_time" bson:"create_time"`
		UpdateTime string `json:"update_time" bson:"update_time"`
	}
	queryInfo := QueryInfo{}
	err := c.BindJSON(&queryInfo)
	if nil != err {
		Json(c, nil, err)
		return
	}
	buyDate, err := utils.ParseDatetime(queryInfo.BuyDate)
	if nil != err {
		Json(c, nil, err)
		return
	}

	buyRecord := &entities.BuyRecord{
		Isbn:     queryInfo.Isbn,
		Store:    queryInfo.Store,
		Quantity: queryInfo.Quantity,
		BuyDate:  buyDate,
	}

	// 调用
	bookService := NewBookService()
	resultInfo, err := bookService.AddBuyRecord(buyRecord)

	Json(c, resultInfo, err)
}

func BuyRecordController_Query(c *gin.Context) {
	// 参数
	data, err := ioutil.ReadAll(c.Request.Body)
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		var criteriaMap *hashmap.Map = hashmap.New()
		err = criteriaMap.FromJSON(data)
		if nil != err {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			// 分页
			pageNo, pageSize := datasources.ParsePageArgs(c)
			criteriaMap.Put("pageNo", pageNo)
			criteriaMap.Put("pageSize", pageSize)

			// 调用
			resultInfo, err := NewBookService().GetBuyRecord(criteriaMap)
			Json(c, resultInfo, err)
		}
	}

}
