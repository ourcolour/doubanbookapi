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
	"strconv"
)

func IsbnController_Draw(c *gin.Context) {
	// 参数
	var (
		isbn      string = c.Param("isbn")
		imageType string = "image/" + c.DefaultQuery("imageType", "png")
		width     int
		height    int

		err error
	)
	width, err = strconv.Atoi(c.DefaultQuery("width", "200"))
	height, err = strconv.Atoi(c.DefaultQuery("height", "40"))

	// 校验
	if 10 != len(isbn) && 13 != len(isbn) {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
		return
	} else if width < 1 || 999 < width {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
		return
	} else if height < 1 || 999 < height {
		Json(c, nil, errs.ERR_INVALID_PARAMETERS)
		return
	}

	bi, err := ean.Encode(isbn)
	if nil != err {
		Json(c, "", err)
		return
	}

	bc, err := barcode.Scale(bi, width, height)
	if nil != err {
		Json(c, "", err)
		return
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, bc)
	if nil != err {
		Json(c, "", err)
		return
	}

	c.Data(http.StatusOK, imageType, buf.Bytes())
}

func IsbnController_Convert(c *gin.Context) {
	// 参数
	var isbn string = c.Param("isbn")
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
