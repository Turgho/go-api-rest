package routes

import (
	"github.com/Turgho/Go-Api-Rest/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Root
	r.GET("/api/", controllers.Index)

	// Trips Routes
	r.POST("/api/trips", controllers.CreateTrip)
	r.PUT("/api/trips/:id/confirm", controllers.UpdateTrip)
	r.GET("/api/trips/:id", controllers.FindTrip)
	r.POST("/api/trips/:id/invite", controllers.InviteParticipant)

	// Participants Routes
	r.GET("/api/participants/:tripID", controllers.FindParticipants)
	r.PUT("/api/participants/:tripID/:participantID", controllers.ConfirmParticipant)

	// Links Routes
	r.POST("/api/trips/:id/links", controllers.RegistryLinks)
	r.GET("/api/trips/:id/links", controllers.FindLinks)

	// Activities Routes
	r.POST("/api/trips/:id/activities", controllers.RegistryActivity)
	r.GET("/api/trips/:id/activities", controllers.FindAcitivies)
}
