package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.UserRoute(router)
	routes.AdventureRoute(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
	})

	router.Use(func(context *gin.Context) {
		c.HandlerFunc(context.Writer, context.Request)
	})

	// Handle CORS preflight OPTIONS requests
	router.OPTIONS("/*path", func(c *gin.Context) {
		// Just a dummy response to satisfy the preflight request
		if c.Writer.Status() == http.StatusNoContent {
			c.JSON(http.StatusOK, gin.H{})
			c.Abort()
		}
	})

	if os.Getenv("PROD_ENV") == "production" {
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