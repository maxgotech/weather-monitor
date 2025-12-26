package models

import "github.com/google/uuid"

type City struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
}
