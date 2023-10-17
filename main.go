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

		// adventure_collection :
		// same as
		// config := cors.DefaultConfig()
		// config.AllowAllOrigins = true
		// router.Use(cors.New(config))
		router.Use(cors.Default())
		router.Run()
        routes.AdventureRoute(router)
				if os.Getenv("PROD_ENV") == "production" {
				port := os.Getenv("PORT")
				router.Run(":"+port)
				} else {
					router.Run("localhost:6000")
				}    
}