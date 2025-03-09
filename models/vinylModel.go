package models

import "time"

type Vinyl struct {
	ID           uint `gorm:"primaryKey;<-:create"`
	UserID       uint `gorm:"<-:create"`
	GenreID      uint
	ShelfslotID  *uint `gorm:"default:null"`
	Name         string
	Artist       string
	Detail       string
	Price        uint
	ImageURL     string
	Format       uint
	Sleeve       uint
	Media        uint
	ReleasedDate string
	CreatedAt    time.Time `gorm:"<-:create"`

	Shelfslot Shelfslot `gorm:"foreignKey:ShelfslotID"`
}

type Genre struct {
	ID   uint `gorm:"primaryKey;<-:create"`
	Name string
}
