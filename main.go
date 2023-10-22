package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
		configs.ConnectDB()
		port := os.Getenv("PORT")
		router.Run(":" + port)
	} else {
		router.Run("localhost:6000")
	}
}