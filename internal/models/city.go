package models

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var AvailableCities = map[string]City{
	"london": {
		Name:      "London",
		Latitude:  51.5072,
		Longitude: -0.1276,
	},
	"paris": {
		Name:      "Paris",
		Latitude:  48.8566,
		Longitude: 2.3522,
	},
	"berlin": {
		Name:      "Berlin",
		Latitude:  52.5200,
		Longitude: 13.4050,
	},
	"madrid": {
		Name:      "Madrid",
		Latitude:  40.4168,
		Longitude: -3.7038,
	},
}
