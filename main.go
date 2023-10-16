package main

import (
	// "gin-mongo-api/collections"
	"gin-mongo-api/configs"
	"gin-mongo-api/controllers"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"

	// "log"
	"os"
	// "github.com/joho/godotenv"
)

func main() {
        router := gin.Default()

		configs.ConnectDB()
		routes.UserRoute(router)
		adventure_collection := configs.GetCollection(configs.DB, "adventures")
		controllers.SetAdventureCollection(adventure_collection)

        routes.AdventureRoute(router)
				if os.Getenv("PROD_ENV") == "production" {
				port := os.Getenv("PORT")
				router.Run(":"+port)
				} else {
					router.Run("localhost:6000")
				}    
}