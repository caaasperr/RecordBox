package handler

import (
	"github.com/go-playground/validator"
)

type CreateVinylRequest struct {
	GenreID      uint
	ShelfslotID  uint
	Name         string `validate:"required,max=100"`
	Artist       string `validate:"required,max=100"`
	Detail       string `validate:"max=455"`
	Price        uint   `validate:"min=0"`
	ImageURL     string `validate:"omitempty,url"`
	Format       uint
	Sleeve       uint
	Media        uint
	ReleasedDate string `validate:"omitempty"`
}

type UpdateVinylRequest struct {
	GenreID      uint
	ShelfslotID  uint
	Name         string `validate:"max=100"`
	Artist       string `validate:"max=100"`
	Detail       string `validate:"max=455"`
	Price        uint   `validate:"min=0"`
	ImageURL     string `validate:"omitempty,url"`
	Format       uint
	Sleeve       uint
	Media        uint
	ReleasedDate string `validate:"omitempty"`
}

type UpdateVinylSlotRequest struct {
	ShelfslotID uint
}

type GetAlbumCoversRequest struct {
	Name   string
	Artist string
}

/*
type Genre struct {
	ID   uint   `gorm:"primaryKey;<-:create"`
	Name string `validate:"required,max=50"`
}
*/

var validate = validator.New()

func (v *CreateVinylRequest) Validate() error {
	return validate.Struct(v)
}
func (v *UpdateVinylRequest) Validate() error {
	return validate.Struct(v)
}
