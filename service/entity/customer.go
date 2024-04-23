package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        string         `gorm:"primary_key;"`
	CreatedAt time.Time      `form:"created_at"`
	UpdatedAt time.Time      `form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;default:null"`
	Name      string
	Age       int
}
