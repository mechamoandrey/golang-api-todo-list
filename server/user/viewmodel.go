package server

type CreateUserRequest struct {
	Name     string
	Login    string
	Password string
}

type LoginRequest struct {
	Login    string
	Password string
}
