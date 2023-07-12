package main

import (
	"log"
	handler2 "rest-api-postgres/internal/handler"
	"rest-api-postgres/pkg"
)

func main() {
	srv := new(pkg.Server)
	handler := new(handler2.Handler)
	// TODO pass config
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %v", err)
	}
}
