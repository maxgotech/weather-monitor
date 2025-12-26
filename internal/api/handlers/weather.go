package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather-monitor/internal/models"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := strings.ToLower(chi.URLParam(r, "city"))

	coords, ok := models.AvailableCities[city]
	if !ok {
		http.Error(w, "city not found", http.StatusNotFound)
		return
	}

	weather, err := h.openMeteo.GetTestWeather(coords.Latitude, coords.Longitude)
	if err != nil {
		http.Error(w, "failed to fetch weather", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
