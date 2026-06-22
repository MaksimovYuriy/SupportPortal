package app

import (
	"log"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	"github.com/MaksimovYuriy/SupportPortal/internal/infra/postgres"
	ticketrepo "github.com/MaksimovYuriy/SupportPortal/internal/repository/postgres"
	"github.com/MaksimovYuriy/SupportPortal/internal/service"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
	httptransport "github.com/MaksimovYuriy/SupportPortal/internal/transport/http"
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

	ticketRepository := ticketrepo.NewPostgresTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handlers.NewTicketHandler(ticketService)

	router := httptransport.NewRouter(ticketHandler)
	addr := ":8080"

	log.Printf("SupportPortal API started at %s", addr)
	return http.ListenAndServe(addr, router)
}
