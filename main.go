package main

import (
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gin-mongo-api/db"
	"os"
)

func main() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())
	
	routes.AdventureRoute(router)
	routes.UserRoute(router)
	
	// router.Run()
	if os.Getenv("PROD_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
		db.ConnectToMongoDB()
		port := os.Getenv("PORT")
		router.Run(":" + port)
	} else {
		db.ConnectToMongoDB()
		router.Run("localhost:6000")
	}
}
