package impl

import (
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/webs/services"
	"iamcc.cn/doubanbookapi/webs/services/bll/default"
)

type DefaultService struct {
}

func NewDefaultService() services.IDefaultService {
	return services.IDefaultService(&DefaultService{})
}

func (this *DefaultService) Version(c *gin.Context) string {
	return configs.SERVICE_VERSION
}

func (this *DefaultService) VerifyMongoDB() error {
	return _default.VerifyDB()
}
