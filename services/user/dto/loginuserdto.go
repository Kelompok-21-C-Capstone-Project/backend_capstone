package dto

type LoginUserDTO struct {
	Identifier string `validate:"required,min=5"`
	Password   string `validate:"required,min=8"`
}
