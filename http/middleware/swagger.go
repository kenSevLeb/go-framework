package middleware

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerHandler
func SwaggerHandler(open bool) gin.HandlerFunc {
	if open { // 判断是否开启了swagger
		return ginSwagger.WrapHandler(swaggerFiles.Handler)
	}

	return func(c *gin.Context) {
		// Simulate behavior when route unspecified and
		// return 404 HTTP code
		c.String(404, "")
	}
}
