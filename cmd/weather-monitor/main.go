package main

import (
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"weather-monitor/internal/api/handlers"
	"weather-monitor/internal/config"
	"weather-monitor/internal/services"
)

func main() {
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	openMeteoService := services.NewOpenMeteo(cfg)
	handlers := handlers.NewHandlers(openMeteoService)

	// Routes
	r.Get("/health", handlers.HealthHandler)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.Register)
		r.Post("/login", handlers.Login)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/cities", handlers.ListCities)
		r.Get("/weather/{city}", handlers.GetWeatherByCity)
		r.Get("/user/city", handlers.GetUserCity)
		r.Post("/user/city", handlers.SaveUserCity)
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
