package v1

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"regexp"
)

func BookController_Query(c *gin.Context) {
	action := c.Param("action")

	// 1. /book/:isbn
	m1Regex := regexp.MustCompile(`^/(\d{13})$`)
	// 2. /book/isbn/:isbn
	m2Regex := regexp.MustCompile(`^/isbn/(\d{13})$`)
	// 3. /book/id/:id
	m3Regex := regexp.MustCompile(`^/id/(\d+)$`)
	// 4. /book/author/:author
	m4Regex := regexp.MustCompile(`^/author/([^/]+)$`)

	bookService := NewBookService()
	if m1Regex.MatchString(action) {
		isbn := m1Regex.FindStringSubmatch(action)[1]
		data, err := bookService.GetBookByIsbn(isbn)
		Json(c, data, err)
	} else if m2Regex.MatchString(action) {
		isbn := m2Regex.FindStringSubmatch(action)[1]
		data, err := bookService.GetBookByIsbn(isbn)
		Json(c, data, err)
	} else if m3Regex.MatchString(action) {
		id := m3Regex.FindStringSubmatch(action)[1]
		data, err := bookService.GetBook(id)
		Json(c, data, err)
	} else if m4Regex.MatchString(action) {
		author := m4Regex.FindStringSubmatch(action)[1]
		data, err := bookService.GetBookByAuthor(author)
		Json(c, data, err)
	} else {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
	}
}
