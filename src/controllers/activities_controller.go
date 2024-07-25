package controllers

import (
	"fmt"
	"log"
	"net/http"

	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var activitiesRepo *repo.ActivitiesRepository

func InitializeActivitiesRepo(repo *repo.ActivitiesRepository) {
	activitiesRepo = repo
}

func RegistryActivity(c *gin.Context) {
	tripID := c.Param("id")

	if _, err := uuid.Parse(tripID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viagem inválido"})
		return
	}

	var activity repo.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity.ID = uuid.New().String()
	activity.TripID = tripID

	if err := activitiesRepo.CreateActivity(&activity); err != nil {
		fmt.Printf("%v\n", activity)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o ID da atividade criada
	log.Print("Atividade criada.")
	c.JSON(http.StatusOK, gin.H{"activity_id": activity.ID})
}

func FindAcitivies(c *gin.Context) {
	tripID := c.Param("id")

	activities, err := activitiesRepo.FindActivityFromTrip(tripID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Print("Atividade encontrada.")
	c.JSON(http.StatusOK, activities)
}
