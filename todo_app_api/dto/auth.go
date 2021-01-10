package dto

// LoginDTO defines the parameters for the login endpoint
type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterDTO defines the parameters for the register endpoint
type RegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
