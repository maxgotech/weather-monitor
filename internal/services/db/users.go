package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"weather-monitor/internal/models"
)

func (d *Database) CreateUser(
	ctx context.Context,
	user *models.User,
) error {
	query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
	`

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	_, err := d.pool.Exec(ctx, query,
		user.ID,
		user.Email,
		user.Password,
	)

	return err
}

func (d *Database) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	query := `
		SELECT id, email, password, city_id
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := d.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CityID,
	)

	if err == pgx.ErrNoRows {
		return nil, err
	}

	return &user, err
}

func (d *Database) UpdateUserCity(
	ctx context.Context,
	email string,
	cityName string,
) error {
	query := `
		UPDATE users
		SET city_id = c.id
		FROM cities c
		WHERE users.email = $1
		  AND c.name = $2
	`

	cmd, err := d.pool.Exec(ctx, query, email, cityName)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (d *Database) GetUserWithCity(
	ctx context.Context,
	userEmail string,
) (*models.User, error) {
	query := `
		SELECT
			u.email,
			c.id,
			c.name,
			c.latitude,
			c.longitude
		FROM users u
		LEFT JOIN cities c ON u.city_id = c.id
		WHERE u.email = $1
	`

	row := d.pool.QueryRow(ctx, query, userEmail)

	var (
		user     models.User
		cityID   *uuid.UUID
		cityName *string
		lat      *float64
		lng      *float64
	)

	err := row.Scan(
		&user.Email,
		&cityID,
		&cityName,
		&lat,
		&lng,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if cityID != nil {
		user.City = &models.City{
			ID:        *cityID,
			Name:      *cityName,
			Latitude:  *lat,
			Longitude: *lng,
		}
	}

	return &user, nil
}
