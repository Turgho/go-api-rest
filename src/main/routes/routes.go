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
	r.GET("/trips/:id", controllers.FindTrip)
	r.POST("/trips/:id/invite", controllers.InviteParticipant)

	// Participants Routes
	r.GET("/participants/:tripID", controllers.FindParticipants)
	r.PUT("/participants/:tripID/:participantID", controllers.ConfirmParticipant)

	// Links Routes
	r.POST("/trips/:id/links", controllers.RegistryLinks)
	r.GET("/trips/:id/links", controllers.FindLinks)
}
