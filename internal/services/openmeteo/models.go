package openmeteo

type WeatherResponse struct {
	CurrentWeather CurrentWeather `json:"current"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature_2m"`
	Windspeed   float64 `json:"wind_speed_10m"`
}
