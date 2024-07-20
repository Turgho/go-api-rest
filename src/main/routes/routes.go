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
	r.PUT("/trips/:id/confirm", controllers.UpdateTrip)
	r.GET("trips/:id", controllers.FindTrip)

	// Participants Routes
	r.POST("trips/:id/invite", controllers.InviteParticipant)
	r.GET("participants/:tripID", controllers.FindParticipants)
}
