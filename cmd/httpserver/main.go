package httpserver

import (
	"business-rule-manager/internal/inbound/httpserver"
	"business-rule-manager/internal/pkg/config"
	"business-rule-manager/internal/pkg/repository"
	"log"
	"net/http"
)

func StartHTTPServer() {

	cfg := config.LoadConfig()
	// Initialize the database
	repository.Init(cfg.DatabaseURL)
	defer repository.Close()

	// Setup routes
	router := httpserver.SetupRoutes()

	// Start the HTTP server
	log.Println("Starting server on :9000...")
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
