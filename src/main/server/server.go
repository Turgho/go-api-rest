package server

import (
	"log"

	"github.com/Turgho/Go-Api-Rest/src/controllers"
	"github.com/Turgho/Go-Api-Rest/src/main/routes"
	repo "github.com/Turgho/Go-Api-Rest/src/models/repositories"
	"github.com/Turgho/Go-Api-Rest/src/models/settings"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	// Testa a conexão com o banco de dados
	dbHandler, err := settings.DBConnect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbHandler.Close()

	// Configura o repositório
	tripRepo := repo.NewTripRepository(dbHandler.DB)
	emailRepo := repo.NewEmailsRepository(dbHandler.DB)
	participantRepo := repo.NewParticipantsRepository(dbHandler.DB)
	linkRepo := repo.NewLinksRepository(dbHandler.DB)
	activityRepo := repo.NewActivitiesRepository(dbHandler.DB)

	// Inicializa os controladores dos repostórios
	controllers.InitializeTripRepo(tripRepo)
	controllers.InitializePartiRepo(participantRepo, emailRepo)
	controllers.InitializeLinksRepo(linkRepo)
	controllers.InitializeActivitiesRepo(activityRepo)

	// Inicializa o servidor e as rotas
	r := gin.Default()
	routes.SetupRoutes(r)

	// Inicialização do servidor na porta 5050
	if err := r.Run(":5050"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
