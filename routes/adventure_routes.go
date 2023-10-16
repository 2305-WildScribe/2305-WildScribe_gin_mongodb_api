package routes

import (
				"gin-mongo-api/controllers"
				"github.com/gin-gonic/gin"
)

func AdventureRoute(router *gin.Engine)  {

	router.POST("/api/v0/adventure", controllers.CreateAdventure())
	router.DELETE("/api/v0/adventure", controllers.DeleteAdventure())
	router.POST("/api/v0/user/adventure", controllers.GetAnAdventure())
	router.POST("/api/v0/user/adventures", controllers.GetAdventuresForUser())
}