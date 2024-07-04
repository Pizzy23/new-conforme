package middleware

import (
	_ "conforme/docs"
	sector "conforme/internal/Sector/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"conforme/db"
)

// @title
// @version 1.0
// @description API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	// Set the Gin mode to release
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine
	r := gin.New()

	// Connect to the database
	db.ConnectDatabase()
	engine := db.Repo

	r.Use(DBMiddleware(engine))

	// Use the Logger and Recovery middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Routes without authentication
	r.GET("/token", generateTokenHandler)
	r.POST("/create-painels-test", sector.CreatePainelTest)
	r.GET("/all-find-painel", sector.SearchAllTestPainel)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(authMiddleware)

	return r
}