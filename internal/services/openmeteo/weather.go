package openmeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (om *OpenMeteo) GetWeather(lat, lon float64) (*WeatherResponse, error) {
	url := fmt.Sprintf(
		"%s/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		om.cfg.OpenMeteoApiURL, lat, lon,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
