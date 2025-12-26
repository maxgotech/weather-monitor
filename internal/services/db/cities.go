package db

import (
	"context"

	"github.com/jackc/pgx/v5"

	"weather-monitor/internal/models"
)

func (d *Database) GetCityByName(
	ctx context.Context,
	name string,
) (*models.City, error) {
	query := `
		SELECT id, name, latitude, longitude
		FROM cities
		WHERE name = $1
	`

	var city models.City
	err := d.pool.QueryRow(ctx, query, name).Scan(
		&city.ID,
		&city.Name,
		&city.Latitude,
		&city.Longitude,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &city, err
}

func (d *Database) ListCities(ctx context.Context) ([]models.City, error) {
	query := `
		SELECT id, name, latitude, longitude
		FROM cities
		ORDER BY name
	`

	rows, err := d.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := []models.City{}

	for rows.Next() {
		var c models.City
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Latitude,
			&c.Longitude,
		); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}

	return cities, rows.Err()
}
