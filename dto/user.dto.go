package dto

// * START REGISTER DTO
type RegisterNewUserRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type RegisterNewUserDTO struct {
	Username string
	Email    string
	Password string
}

type RegisterNewUserResponseDTO struct {
	Token string `json:"token"`
}

// * END REGISTER DTO

// * START LOGIN DTO
type LoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginDTO struct {
	Username string
	Password string
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

// * END LOGIN DTO
