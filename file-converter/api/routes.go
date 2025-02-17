package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/convert/text", ConvertTextHandler)
	r.POST("/convert/image", ConvertImageHandler)
	r.POST("/compress", CompressHandler)

	return r
}
