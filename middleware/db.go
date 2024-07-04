package middleware

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func DBMiddleware(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", engine)
		c.Next()
	}
}
