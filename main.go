package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()

        // router.GET("/", func(c *gin.Context) {
        //         c.JSON(200, gin.H{
        //                 "data": "WILDSCRIBE 2305",
        //         })
        // })
				configs.ConnectDB()
				
				routes.UserRoute(router)


        router.Run("localhost:6000") 
}