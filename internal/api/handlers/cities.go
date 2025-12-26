package handlers

import (
	"encoding/json"
	"net/http"

	"weather-monitor/internal/models"
)

func (h *Handlers) ListCities(w http.ResponseWriter, r *http.Request) {
	cities := make([]models.City, 0, len(models.AvailableCities))

	for _, city := range models.AvailableCities {
		cities = append(cities, city)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func (h *Handlers) SaveUserCity(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var req struct {
		City string `json:"city"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	users[userID].City = req.City
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) GetUserCity(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	json.NewEncoder(w).Encode(map[string]string{
		"city": users[userID].City,
	})
}
