package routes

import (
				"gin-mongo-api/controllers"
				"github.com/gin-gonic/gin"
)

func AdventureRoute(router *gin.Engine)  {

	router.POST("/adventures", controllers.CreateAdventure())
	router.POST("/adventure", controllers.GetAnAdventure())
	router.GET("/user/adventures", controllers.GetAdventuresForUser())

}