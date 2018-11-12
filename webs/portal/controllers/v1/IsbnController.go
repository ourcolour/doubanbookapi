package v1

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/ean"
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	isbnUtil "iamcc.cn/doubanbookapi/utils/isbn"
	"image/png"
	"net/http"
)

func IsbnController_Draw(c *gin.Context) {
	// 参数
	var isbn string = c.DefaultQuery("isbn", "")
	if 10 != len(isbn) && 13 != len(isbn) {
		Json(c, "", errs.ERR_INVALID_PARAMETERS)
		return
	}

	bi, err := ean.Encode(isbn)
	if nil != err {
		Json(c, "", err)
	}
	bc, err := barcode.Scale(bi, 200, 50)
	if nil != err {
		Json(c, "", err)
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, bc)
	if nil != err {
		Json(c, "", err)
	}

	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func IsbnController_Convert(c *gin.Context) {
	// 参数
	var isbn string = c.DefaultQuery("isbn", "")
	if 10 != len(isbn) && 13 != len(isbn) {
		Json(c, "", errs.ERR_INVALID_PARAMETERS)
		return
	}

	var (
		result string
		err    error
	)

	if 10 == len(isbn) {
		result, err = isbnUtil.ConvertToIsbn13(isbn)
	} else if 13 == len(isbn) {
		result, err = isbnUtil.ConvertToIsbn10(isbn)
	}

	Json(c, result, err)
}
