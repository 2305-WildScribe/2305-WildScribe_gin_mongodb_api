package main

import (
	// "gin-mongo-api/collections"
	"gin-mongo-api/configs"
	// "gin-mongo-api/controllers"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "log"
	"os"
	// "net/http"
	// "github.com/joho/godotenv"
)

func main() {
        router := gin.Default()

		configs.ConnectDB()
		routes.UserRoute(router)
		config := cors.DefaultConfig()
		// config.AllowAllOrigins = true
		config.AllowOrigins = []string{"http://localhost:3000"}
		router.Use(cors.New(config))
		router.Run()
        routes.AdventureRoute(router)
		if os.Getenv("PROD_ENV") == "production" {
			port := os.Getenv("PORT")
			router.Run(":"+port)
		} else {
			router.Run("localhost:6000")
	}    
}
