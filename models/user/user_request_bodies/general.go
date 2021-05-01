package user_request_bodies

//NOTE We removed the binding out and use validate only instead for validation phase

type RegisterUserBody struct {
	Username string `json:"username" xml:"username" validate:"required"`
	Password string `json:"password" xml:"password" validate:"required"`
}

type LoginUserBody = RegisterUserBody

