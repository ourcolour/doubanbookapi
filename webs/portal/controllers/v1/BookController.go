package v1

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"regexp"
)

func BookController_ActionDispatcher(c *gin.Context) {
	action := c.Param("action")

	// 1. /book/:isbn
	m1Regex := regexp.MustCompile(`^/(\d{13})$`)
	// 2. /book/isbn/:isbn
	m2Regex := regexp.MustCompile(`^/isbn/(\d{13})$`)
	// 3. /book/id/:id
	m3Regex := regexp.MustCompile(`^/id/(\d+)$`)
	// 4. /book/author/:author
	m4Regex := regexp.MustCompile(`^/author/([^/]+)$`)
	// 5. /rank
	m5Regex := regexp.MustCompile(`^/rank/?$`)

	if m1Regex.MatchString(action) {
		isbn := m1Regex.FindStringSubmatch(action)[1]
		BookController_GetBookByIsbn(c, isbn)
	} else if m2Regex.MatchString(action) {
		isbn := m2Regex.FindStringSubmatch(action)[1]
		BookController_GetBookByIsbn(c, isbn)
	} else if m3Regex.MatchString(action) {
		id := m3Regex.FindStringSubmatch(action)[1]
		BookController_GetBookById(c, id)
	} else if m4Regex.MatchString(action) {
		author := m4Regex.FindStringSubmatch(action)[1]
		BookController_GetBookByAuthor(c, author)
	} else if m5Regex.MatchString(action) {
		BookController_Rank(c)
	} else {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
	}
}

func BookController_GetBookByIsbn(c *gin.Context, isbn string) {
	bookService := NewBookService()
	data, err := bookService.GetBookByIsbn(isbn)
	Json(c, data, err)
}

func BookController_GetBookById(c *gin.Context, id string) {
	bookService := NewBookService()
	data, err := bookService.GetBook(id)
	Json(c, data, err)
}

func BookController_GetBookByAuthor(c *gin.Context, author string) {
	bookService := NewBookService()
	data, err := bookService.GetBookByAuthor(author)
	Json(c, data, err)
}

func BookController_Rank(c *gin.Context) {
	// 参数
	var isbnList arraylist.List
	jsonData := MustGetRequestBody(c)
	err := isbnList.FromJSON(jsonData)
	if nil != err {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
		return
	}

	bookService := NewBookService()
	bookList, err := bookService.GetRankInIsbn(&isbnList)
	Json(c, bookList, err)
}
