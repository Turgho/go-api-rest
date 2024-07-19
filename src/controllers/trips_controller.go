package controllers

import (
	"log"
	"net/http"

	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var tripRepo *repo.TripRepository

// Função para inicializar o repositório de viagens
func InitializeTripRepo(repo *repo.TripRepository) {
	tripRepo = repo
}

// CreateTrip - handler para a criação de viagens
func CreateTrip(c *gin.Context) {
	var trip repo.Trip

	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gera um UUID para a viagem
	trip.ID = uuid.New().String()
	trip.Status = 0 // Define status como 0 por padrão

	if err := tripRepo.CreateTrip(&trip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o ID da viagem criada
	log.Print("Viagem criada.")
	c.JSON(http.StatusOK, gin.H{"id": trip.ID})
}

func FindTrip(c *gin.Context) {
	tripID := c.Param("id")

	trip, err := tripRepo.FindTripByID(tripID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Print("Viagem encontrada.")
	c.JSON(http.StatusOK, trip)
}

// UpdateTrip - handler para atualizar uma viagem
func UpdateTrip(c *gin.Context) {
	tripID := c.Param("id")

	// Verifica se o ID é um UUID válido
	if _, err := uuid.Parse(tripID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viagem inválido"})
		return
	}

	// Atualiza a viagem no repositório
	if err := tripRepo.UpdateTrip(tripID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna a viagem atualizada
	log.Print("Viagem atualizada.")
	c.JSON(http.StatusOK, gin.H{"id": tripID})
}
