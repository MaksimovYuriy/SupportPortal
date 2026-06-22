package app

import (
	"log"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	"github.com/MaksimovYuriy/SupportPortal/internal/infra/postgres"
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

	router := httptransport.NewRouter()
	addr := ":8080"

	log.Printf("SupportPortal API started at %s", addr)
	return http.ListenAndServe(addr, router)
}
