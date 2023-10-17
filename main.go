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
		router.Use(CORSMiddleware())
		router.Run()
        routes.AdventureRoute(router)
		if os.Getenv("PROD_ENV") == "production" {
			port := os.Getenv("PORT")
			router.Run(":"+port)
		} else {
			router.Run("localhost:6000")
	}    
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
        } else {
            c.Next()
        }
    }
}