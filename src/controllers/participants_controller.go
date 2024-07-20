package controllers

import (
	"log"
	"net/http"

	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var participantsRepo *repo.ParticipantsRepository
var emailsRepo *repo.EmailsRepository

func InitializePartiRepo(repoParticipants *repo.ParticipantsRepository, repoEmails *repo.EmailsRepository) {
	participantsRepo = repoParticipants
	emailsRepo = repoEmails
}

func InviteParticipant(c *gin.Context) {
	var requestBody struct {
		Names  []string `json:"names"`
		Emails []string `json:"emails"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	tripID := c.Param("id")

	// Verifica se os slices têm o mesmo tamanho
	if len(requestBody.Names) != len(requestBody.Emails) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Names and Emails arrays must have the same length"})
		return
	}

	for i := 0; i < len(requestBody.Names); i++ {
		participantName := requestBody.Names[i]
		participantEmail := requestBody.Emails[i]

		email := repo.EmailToInvite{
			ID:     uuid.New().String(),
			TripId: tripID,
			Email:  participantEmail,
		}

		participant := repo.Participants{
			ID:               uuid.New().String(),
			TripId:           tripID,
			EmailsToInviteID: email.ID,
			Name:             participantName,
			IsConfirmed:      0, // Define como não confirmado inicialmente
		}

		if err := emailsRepo.EmailsToInvite(&email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := participantsRepo.CreateParticipants(&participant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	log.Print("Participante criado.")
	c.JSON(http.StatusOK, gin.H{"message": "Participantes criados"})
}

func FindParticipants(c *gin.Context) {
	tripID := c.Param("id")

	if _, err := uuid.Parse(tripID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viagem inválido"})
		return
	}

	participant, err := participantsRepo.FindParticipantsFromTrip(tripID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Print("Participantes encontrado.")
	c.JSON(http.StatusOK, participant)
}
