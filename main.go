package main

import (
	// "gin-mongo-api/collections"
	"gin-mongo-api/configs"
	// "gin-mongo-api/controllers"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
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

		router.Use(CORS())
		// router.Use(cors.Default())
		router.Run()
        routes.AdventureRoute(router)
		if os.Getenv("PROD_ENV") == "production" {
			port := os.Getenv("PORT")
			router.Run(":"+port)
		} else {
			router.Run("localhost:6000")
	}    
}
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}