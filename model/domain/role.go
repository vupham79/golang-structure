package domain

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	ID   int    `gorm:"AUTO_INCREMENT;unique;not null"`
	Name string `gorm:"not null"`
}
