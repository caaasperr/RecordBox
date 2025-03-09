package handler

import "github.com/go-playground/validator"

type LoginForm struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type CreateUserRequest struct {
	Username string `gorm:"size:255;unique" validate:"required,min=3,max=255"`
	Password string `validate:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
	Username string `gorm:"size:255;unique" validate:"min=3,max=255"`
	Password string `validate:"min=8,max=100"`
}

var validate = validator.New()

func (v *CreateUserRequest) Validate() error {
	return validate.Struct(v)
}

func (v *UpdateUserRequest) Validate() error {
	return validate.Struct(v)
}

func (v *LoginForm) Validate() error {
	return validate.Struct(v)
}
