package models

import (
	"github.com/google/uuid"
	"time"
)

type File struct {
	ID               string    `gorm:"primaryKey"`
	AdministrationID uuid.UUID `gorm:"index"`
	Filename         string
	Type             string
	Size             int
	Fingerprint      string
	SelfDestruct     bool
	Downloads        int64
	Status           FileStatus `gorm:"index"`
	CreatedAt        time.Time
	DestroyOn        time.Time `gorm:"index"`
}
