package app

import (
	"log"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	"github.com/MaksimovYuriy/SupportPortal/internal/database/postgres"
	"github.com/MaksimovYuriy/SupportPortal/internal/database/postgres/repositories"
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
	defer db.Close()

	ticketRepository := repositories.NewPostgresTicketRepository(db)
	queueRepository := repositories.NewPostgresQueueRepository(db)

	ticketService := services.NewTicketService(ticketRepository, queueRepository)
	queueService := services.NewQueueService(queueRepository)

	ticketHandler := handlers.NewTicketHandler(ticketService)
	queueHandler := handlers.NewQueueHandler(queueService)

	handlers := &rest.Handlers{
		TicketHandler: ticketHandler,
		QueueHandler:  queueHandler,
	}

	router := rest.NewRouter(handlers)
	addr := ":8080"

	log.Printf("SupportPortal API started at %s", addr)
	return http.ListenAndServe(addr, router)
}
