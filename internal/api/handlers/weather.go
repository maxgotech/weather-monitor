package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"weather-monitor/internal/models"
)

func (h *Handlers) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := strings.ToLower(chi.URLParam(r, "city"))

	coords, ok := models.AvailableCities[city]
	if !ok {
		RespondJSON(w, http.StatusNotFound, map[string]string{
			"error": "city not found",
		})

		return
	}

	weather, err := h.openMeteo.GetTestWeather(coords.Latitude, coords.Longitude)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch weather",
		})

		return
	}

	RespondJSON(w, http.StatusOK, weather)
}
