package app

import (
	"log"
	"net/http"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	"github.com/MaksimovYuriy/SupportPortal/internal/database/postgres"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/rest"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	db, err := postgres.NewDB(cfg.DB)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
	}()

	userRepository := postgres.NewUserRepository(db)
	agentRepository := postgres.NewAgentRepository(db)
	queueRepository := postgres.NewQueueRepository(db)

	agentService := services.NewAgentService(agentRepository)
	userService := services.NewUserService(userRepository, agentService)
	queueService := services.NewQueueService(queueRepository)

	userHandler := handlers.NewUserHandler(userService)
	agentHandler := handlers.NewAgentHandler(agentService)
	queueHandler := handlers.NewQueueHandler(queueService)

	handlers := &rest.Handlers{
		UserHandler:  userHandler,
		AgentHandler: agentHandler,
		QueueHandler: queueHandler,
	}

	router := rest.NewRouter(handlers)
	addr := ":8080"
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}

	log.Printf("SupportPortal API started at %s", addr)
	return server.ListenAndServe()
}
