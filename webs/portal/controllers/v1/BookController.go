package v1

import (
	"bytes"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	. "iamcc.cn/doubanbookapi/webs/services/impl"
	"image/jpeg"
	"image/png"
	"net/http"
	"regexp"
	"strconv"
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
	// 5. /book/author/:author
	m5Regex := regexp.MustCompile(`^/title/([^/]+)$`)
	// 6. /rank
	m6Regex := regexp.MustCompile(`^/rank/?$`)
	// 7. /thumb
	m7Regex := regexp.MustCompile(`^/thumb/(\d{13})(?:/([123]?))$`)
	// 8. /cip/:isbn
	m8Regex := regexp.MustCompile(`^/cip/([^/]{13})$`)

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
		title := m5Regex.FindStringSubmatch(action)[1]
		BookController_GetBookByTitle(c, title)
	} else if m6Regex.MatchString(action) {
		BookController_Rank(c)
	} else if m7Regex.MatchString(action) {
		isbn := m7Regex.FindStringSubmatch(action)[1]
		size, _ := strconv.Atoi(m7Regex.FindStringSubmatch(action)[2])
		BookController_DrawThumbnailByIsbn(c, isbn, uint(size))
	} else if m8Regex.MatchString(action) {
		isbn := m8Regex.FindStringSubmatch(action)[1]
		BookController_GetCipByIsbn(c, isbn)
	} else {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
	}
}

func BookController_DrawThumbnailByIsbn(c *gin.Context, isbn string, size uint) {
	bookService := NewBookService()
	data, err := bookService.GetBookByIsbn(isbn)
	if nil != err {
		Json(c, data, err)
	} else {
		var imageUrl string
		switch size {
		case 1:
			imageUrl = data.Images.Small
			break
		case 2:
			imageUrl = data.Images.Medium
			break
		case 3:
			imageUrl = data.Images.Large
			break
		default:
			imageUrl = data.Image
		}

		resp, _ := http.Get(imageUrl)
		defer resp.Body.Close()
		img, _ := jpeg.Decode(resp.Body)

		// Jpeg -> PNG
		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		c.Data(http.StatusOK, "image/png", buf.Bytes())
	}
}

func BookController_UpdateLocalBookCip(c *gin.Context) {
	calisApiService := NewCalisApiService()
	data, err := calisApiService.UpdateLocalBookCip()
	Json(c, data, err)
}

func BookController_GetCipByIsbn(c *gin.Context, isbn string) {
	calisService := NewCalisApiService()
	data, err := calisService.GetCipByIsbn(isbn)
	Json(c, data, err)
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

func BookController_GetBookByTitle(c *gin.Context, title string) {
	bookService := NewBookService()
	data, err := bookService.GetBookByTitle(title)
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
