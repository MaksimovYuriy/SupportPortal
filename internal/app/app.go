package app

import (
	"log"
	"net/http"

	httptransport "github.com/MaksimovYuriy/SupportPortal/internal/transport/http"
)

func Run() error {
	router := httptransport.NewRouter()
	addr := ":8080"

	log.Printf("SupportPortal API started at %s", addr)
	return http.ListenAndServe(addr, router)
}
