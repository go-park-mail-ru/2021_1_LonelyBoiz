package models

import "github.com/google/uuid"

type Image struct {
	Uuid uuid.UUID `json:"photoId" db:"photoUuid"`
}
