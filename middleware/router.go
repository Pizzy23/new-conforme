package middleware

import (
	_ "conforme/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title
// @version 1.0
// @description API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/add-questions")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ResponseHandler())
	//Use response, but not Token
	r.GET("/token", generateTokenHandler)

	auth := r.Group("/api")
	auth.Use(authMiddleware)
	//Response and token service
	auth.PUT("/exemp")

	return r
}
