package helpers

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func GetDBEngineFromContext(c *gin.Context) (*xorm.Engine, bool) {
	engine, exists := c.Get("db")
	if !exists {
		return nil, false
	}
	return engine.(*xorm.Engine), true
}
