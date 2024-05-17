package dto

type RegisterNewUserRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterNewUserRepoStruct struct {
	Username string
	Email    string
	Password string
}

type RegisterNewUserUseCaseStruct struct {
	Username string
	Email    string
	Password string
}
type LoginParamUseCaseStruct struct {
	Username string
	Password string
}
