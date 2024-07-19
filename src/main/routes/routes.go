package routes

import (
	"github.com/Turgho/Go-Api-Rest/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Root
	r.GET("/", controllers.Index)

	// Trips Routes
	r.POST("/trips", controllers.CreateTrip)
	r.PUT("/trips/:id", controllers.UpdateTrip)
	r.GET("trips/:id", controllers.FindTrip)

	// Emails Routes
	r.POST("/trips/:id/invite", controllers.InviteEmails)
}
