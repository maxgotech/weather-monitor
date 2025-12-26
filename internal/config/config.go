package config

type Config struct {
	OpenMeteoApiURL string `env:"OPEN_METEO_URL,required"`
}
