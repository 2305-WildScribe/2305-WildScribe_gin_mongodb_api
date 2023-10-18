package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "github.com/rs/cors"
	// "net/http"
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
	
	router.Run()
	if os.Getenv("PROD_ENV") == "production" {
		configs.ConnectDB()
		port := os.Getenv("PORT")
		router.Run(":" + port)
	} else {
		router.Run("localhost:6000")
	}
}
// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Content-Type", "application/json")
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
//         c.Writer.Header().Set("Access-Control-Max-Age", "86400")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//         } else {
//             c.Next()
//         }
//     }
// }