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
	"net/http"
	// "github.com/joho/godotenv"
)

func main() {
	// config.AllowOrigins = []string{"http://localhost:3000"}
        router := gin.Default()

		configs.ConnectDB()
		routes.UserRoute(router)
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AddAllowMethods("POST", "PUT", "DELETE", "OPTIONS")
		router.Use(cors.New(config))
		router.Run()
		corsMiddleware := cors.New(config)
		router.Use(corsMiddleware)
        routes.AdventureRoute(router)
		router.OPTIONS("/*path", func(c *gin.Context) {
			// Handle CORS preflight OPTIONS requests here
			if c.Writer.Status() == http.StatusNoContent {
				// Handle errors here (e.g., return a custom error response)
				c.JSON(http.StatusForbidden, gin.H{"error": "CORS preflight request denied"})
				c.Abort()
			}
		})
		if os.Getenv("PROD_ENV") == "production" {
			port := os.Getenv("PORT")
			router.Run(":"+port)
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