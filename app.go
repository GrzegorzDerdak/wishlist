package main

import (
	"log"
	"net/http"

	"github.com/grzegorzderdak/saleor-wishlist/config"
	"github.com/grzegorzderdak/saleor-wishlist/healthcheck"
	"github.com/grzegorzderdak/saleor-wishlist/saleor"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/manifest", saleor.ManifestHandler)

	// Manifest URLs
	handler.HandleFunc("/app", saleor.AppHandler)
	handler.HandleFunc("/register", saleor.RegisterHandler)
	handler.HandleFunc("/support", saleor.BaseHandler)
	handler.HandleFunc("/homepage", saleor.BaseHandler)

	handler.HandleFunc("/healthcheck", healthcheck.HealthcheckHandler)

	config := config.Load()

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler,
	}

	log.Println("Starting server on :" + config.Port)
	server.ListenAndServe()
}
