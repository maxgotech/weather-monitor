package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID  `db:"id" json:"id"`
	Email    string     `db:"email" json:"email"`
	Password string     `db:"password" json:"-"`
	CityID   *uuid.UUID `db:"city_id" json:"-"`
	*City    `json:"city,omitempty"`
}
