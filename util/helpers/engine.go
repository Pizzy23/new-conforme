package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func GetDBEngineFromContext(c *gin.Context) (*xorm.Engine, bool) {
	engine, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return nil, false
	}
	return engine.(*xorm.Engine), true
}
