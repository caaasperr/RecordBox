package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:255;unique" validate:"required,min=3,max=255"`
	Password  string    `validate:"required,min=8,max=100"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	IsAdmin   bool      `gorm:"default:false;"`
}
