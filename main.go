package main

import (
	// "gin-mongo-api/collections"
	"gin-mongo-api/configs"
	// "gin-mongo-api/controllers"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"

	// "log"
	"os"
	"net/http"
	// "github.com/rs/cors"
	// "github.com/joho/godotenv"
)

func main() {
        router := gin.Default()

		configs.ConnectDB()
		routes.UserRoute(router)

		// adventure_collection :
		// controllers.SetAdventureCollection(adventure_collection)
 // Create an HTTP server
 	// c := cors.New(cors.Options{
    // AllowedOrigins:   []string{"http://localhost:3000"},
    // AllowCredentials: true,
    // AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// })

	// handler := c.Handler(mux)
	server := http.NewServeMux()

	// Define your HTTP handler function
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the necessary headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*") // Change * to your specific origin(s)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		// Your API logic here
	})

	http.ListenAndServe(":8080", server)

        routes.AdventureRoute(router)
				if os.Getenv("PROD_ENV") == "production" {
				port := os.Getenv("PORT")
				router.Run(":"+port)
				} else {
					router.Run("localhost:6000")
				}    
}