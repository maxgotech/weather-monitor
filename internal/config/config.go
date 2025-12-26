package config

type Config struct {
	OpenMeteoApiURL string   `env:"OPEN_METEO_URL,required"`
	DB              Database `envPrefix:"DB_"`
}

type Database struct {
	Host     string `env:"HOST,required"`
	Port     string `env:"PORT,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Name     string `env:"NAME,required"`
	SSL      string `env:"SSLMODE"`
}
