package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model base model for domain type
type Model struct {
	ID        UUID       `sql:",type:uuid" json:"id"`
	CreatedAt time.Time  `sql:"default:now()" json:"created_at"`
	UpdatedAt time.Time  `sql:"default:now()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate prepare data before create data
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeDelete prepare data before delete data
func (m *Model) BeforeDelete(scope *gorm.Scope) error {
	scope.SetColumn("DeletedAt", time.Now())
	return nil
}

// BeforeUpdate prepare data before update data
func (m *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
