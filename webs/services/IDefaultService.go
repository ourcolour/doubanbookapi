package services

import (
	"github.com/gin-gonic/gin"
)

type IDefaultService interface {
	Version(c *gin.Context) string
}
