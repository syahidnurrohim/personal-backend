package routes

import (
	"personal-backend/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/journal/list", controllers.GetAllJournal)
	}

	return router
}
