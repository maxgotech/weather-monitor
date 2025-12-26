package handlers

import "weather-monitor/internal/services"

type Handlers struct {
	openMeteo *services.OpenMeteo
}

func NewHandlers(openMeteo *services.OpenMeteo) Handlers {
	return Handlers{
		openMeteo: openMeteo,
	}
}
