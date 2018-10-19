package impl

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/webs/services"
)

type DefaultService struct {
}

func NewDefaultService() services.IDefaultService {
	var result services.IDefaultService = &DefaultService{}
	return result
}

func (this *DefaultService) Version(c *gin.Context) string {
	return configs.SERVICE_VERSION
}
