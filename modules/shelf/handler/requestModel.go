package handler

import "github.com/go-playground/validator"

type CreateShelfReq struct {
	Name    string  `json:"name" validate:"required,min=1,max=100"`
	Detail  string  `json:"detail" validate:"max=255"`
	Columns float64 `json:"columns" validate:"required,gt=0"`
	Rows    float64 `json:"rows" validate:"required,gt=0"`
}

type UpdateShelfReq struct {
	Name   string `json:"name" validate:"min=1,max=100"`
	Detail string `json:"detail" validate:"max=255"`
}

var validate = validator.New()

func ValidateCreateShelf(req CreateShelfReq) error {
	return validate.Struct(req)
}

func ValidateUpdateShelf(req UpdateShelfReq) error {
	return validate.Struct(req)
}
