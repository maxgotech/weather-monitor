package services

type WeatherResponse struct {
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
}
