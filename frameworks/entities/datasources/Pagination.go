package datasources

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/datasources"
	"math"
	"strconv"
)

type Pagination struct {
	CurrentPage int
	PrevPage    int
	NextPage    int

	PageSize  int
	PageCount int
	//RecordCount int64
	TotalRecordCount int64

	IsFirst bool
	IsLast  bool
	HasNext bool
	HasPrev bool
}

func ParsePageArgs(ctx *gin.Context) (pageNo int, pageSize int) {
	pageNoStr := ctx.DefaultQuery("pageNo", strconv.Itoa(datasources.DEFAULT_PAGE_NO))
	pageSizeStr := ctx.DefaultQuery("pageSize", strconv.Itoa(datasources.DEFAULT_PAGE_SIZE))

	pageNo, err := strconv.Atoi(pageNoStr)
	if nil != err || pageNo < 1 {
		pageNo = 1
	}
	pageSize, err = strconv.Atoi(pageSizeStr)
	if nil != err || pageSize < 1 {
		pageSize = 10
	}

	return
}

func NewPaginationByContext(ctx *gin.Context, totalRecordCount int64) *Pagination {
	// Args
	currentPage, pageSize := ParsePageArgs(ctx)

	return NewPagination(totalRecordCount, pageSize, currentPage)
}

func NewPagination( /*dataList1 []interface{}, */ totalRecordCount int64, pageSize int, currentPage int) *Pagination {
	// RecordCount
	//var recordCount int = 0
	//if nil != dataList {
	//	recordCount = len(dataList)
	//}
	// PageCount
	var pageCount int = 1
	if 0 < totalRecordCount {
		pageCount = int(math.Ceil(float64(totalRecordCount) / float64(pageSize)))
	}

	// CurrentPage
	if currentPage < 1 {
		currentPage = 1
	} else if pageCount < currentPage {
		currentPage = pageCount
	}
	// PrevPage
	var prevPage int = currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	// NextPage
	var nextPage int = currentPage + 1
	if pageCount < nextPage {
		nextPage = pageCount
	}

	// Prev / Next
	var (
		isFirst bool = 1 == currentPage
		isLast  bool = pageCount == currentPage
		hasPrev bool = !isFirst
		hasNext bool = !isLast
	)

	return &Pagination{
		CurrentPage: currentPage,
		PrevPage:    prevPage,
		NextPage:    nextPage,

		PageSize:         pageSize,
		PageCount:        pageCount,
		TotalRecordCount: totalRecordCount,

		IsFirst: isFirst,
		IsLast:  isLast,
		HasPrev: hasPrev,
		HasNext: hasNext,
	}
}

//
//
//func (this *Pagination) BuildPageInfo() (*PageInfo) {
//	return &PageInfo{
//		DataList: this.dataList,
//
//		CurrentPage: this.currentPage,
//		PrevPage:    this.prevPage(),
//		NextPage:    this.nextPage(),
//
//		PageSize:    this.pageSize,
//		PageCount:   this.pageCount(),
//		RecordCount: this.recordCount(),
//
//		IsFirst: this.isFirst(),
//		IsLast:  this.isLast(),
//		HasPrev: this.hasPrev(),
//		HasNext: this.hasNext(),
//	}
//}
