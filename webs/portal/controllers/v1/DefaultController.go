package v1

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	. "iamcc.cn/doubanbookapi/frameworks/controllers"
	"iamcc.cn/doubanbookapi/webs/services"
	"iamcc.cn/doubanbookapi/webs/services/impl"
	"net/http"
)

func DefaultController_Version(c *gin.Context) {
	var defaultService services.IDefaultService = impl.NewDefaultService()

	data := defaultService.Version(c)

	Json(c, data, nil)
}

func DefaultController_50XError(c *gin.Context) {
	JsonWithStatusCode(
		c,
		nil,
		errs.ERR_SERVICE_UNAVAILABLE,
		http.StatusServiceUnavailable,
	)
}

func DefaultController_404Error(c *gin.Context) {
	JsonWithStatusCode(
		c,
		nil,
		errs.ERR_NOT_FOUND,
		http.StatusNotFound,
	)
}

func DefaultController_Unauthorized(c *gin.Context) {
	JsonWithStatusCode(
		c,
		nil,
		errs.ERR_UNAUTHORIZED,
		http.StatusUnauthorized,
	)
}
