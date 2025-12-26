package services

func (om *OpenMeteo) GetTestWeather(lat, lon float64) (*WeatherResponse, error) {
	data := WeatherResponse{
		CurrentWeather: CurrentWeather{
			Temperature: -3,
			Windspeed:   3,
		},
	}

	return &data, nil
}
