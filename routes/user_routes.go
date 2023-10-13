package routes

import (
				"gin-mongo-api/controllers"
				"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
						"data": "WILDSCRIBE 2305",
		})
})

	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())

}