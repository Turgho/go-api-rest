package controllers

import (
	"fmt"
	"net/http"

	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var emailsRepo *repo.EmailsRepository

// Função para inicializar o repositório de emails
func IniciatilizeEmailsRepo(repo *repo.EmailsRepository) {
	emailsRepo = repo
}

func InviteEmails(c *gin.Context) {
	var requestBody struct {
		Emails []string `json:"emails"`
	}

	// Obtém o ID da viagem da URL
	tripID := c.Param("id")

	// Decodifica o corpo da requisição para a estrutura definida
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for _, emailStr := range requestBody.Emails {
		email := repo.EmailToInvite{
			ID:     uuid.New().String(), // Cria um UUID único para cada email
			TripId: tripID,              // Atribui o ID da viagem
			Email:  emailStr,            // Define o email
		}

		fmt.Print(email)

		if err := emailsRepo.EmailsToInvite(&email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Emails invited successfully"})
}
