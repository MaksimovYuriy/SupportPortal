package main

import (
	"log"

	"github.com/MaksimovYuriy/SupportPortal/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
