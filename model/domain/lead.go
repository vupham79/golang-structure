package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	ID               int    `gorm:"AUTO_INCREMENT;unique;not null"`
	Firstname        string `gorm:"not null"`
	Middlename       string
	Lastname         string `gorm:"not null"`
	Birthdate        *time.Time
	Phone            string `gorm:"not null"`
	SourceID         int    `gorm:"not null"`
	Address          string
	Role             string `gorm:"not null"`
	OriginalSourceID string
	Email            string
}
