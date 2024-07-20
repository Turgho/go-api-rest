package controllers

import (
	"log"
	"net/http"

	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var linksRepo *repo.LinksRepository

// Função para inicializar o repositório de viagens
func InitializeLinksRepo(repo *repo.LinksRepository) {
	linksRepo = repo
}

func RegistryLinks(c *gin.Context) {
	var requestBody struct {
		Urls   []string `json:"urls"`
		Titles []string `json:"titles"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	tripID := c.Param("id")

	// Verifica se os slices têm o mesmo tamanho
	if len(requestBody.Urls) != len(requestBody.Titles) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Urls and Titles arrays must have the same length"})
		return
	}

	for i := 0; i < len(requestBody.Urls); i++ {
		url := requestBody.Urls[i]
		title := requestBody.Titles[i]

		link := repo.Links{
			ID:     uuid.New().String(),
			TripID: tripID,
			Link:   url,
			Title:  title,
		}

		if err := linksRepo.CreateLink(&link); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	log.Print("Links registados.")
	c.JSON(http.StatusOK, gin.H{"message": "Links registrados."})
}

func FindLinks(c *gin.Context) {
	tripID := c.Param("id")

	if _, err := uuid.Parse(tripID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viagem inválido"})
		return
	}

	links, err := linksRepo.FindLinksFromTrip(tripID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Print("Participantes encontrado.")
	c.JSON(http.StatusOK, links)
}
