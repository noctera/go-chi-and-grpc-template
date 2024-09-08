package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"my-app/internal/inbound/httpserver/handlers"
	"net/http"
)

// SetupRoutes initializes the HTTP routes and applies middleware
func SetupRoutes() http.Handler {
	router := chi.NewRouter()

	// Adding some default middleware
	router.Use(middleware.Logger)    // Logs the incoming HTTP requests
	router.Use(middleware.Recoverer) // Recovers from panics in the application

	// CORS middleware configuration
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://192.168.178.58:3000"}, // Allow only this origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not to send preflight requests again within this time
	})

	router.Use(corsMiddleware.Handler)

	routerV1 := chi.NewRouter()

	// Define your routes

	routerV1.Post("/test", handlers.CreateTestHandler)
	routerV1.Delete("/test/{id}", handlers.DeleteTestByIdHandler)

	router.Mount("/api/v1", routerV1)

	return router
}
