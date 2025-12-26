package handlers

import (
	"weather-monitor/internal/services/db"
	"weather-monitor/internal/services/openmeteo"
)

type Handlers struct {
	openMeteo *openmeteo.OpenMeteo
	db        *db.Database
}

func NewHandlers(openMeteo *openmeteo.OpenMeteo, pool *db.Database) Handlers {
	return Handlers{
		openMeteo: openMeteo,
		db:        pool,
	}
}
