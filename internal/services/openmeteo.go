package services

import (
	"weather-monitor/internal/config"
)

type OpenMeteo struct {
	cfg config.Config
}

func NewOpenMeteo(cfg config.Config) *OpenMeteo {
	return &OpenMeteo{
		cfg: cfg,
	}
}
