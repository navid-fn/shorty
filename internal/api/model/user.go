package model

type UserRegister struct {
	Username string
	Password string
	Email    string
}

type UserLogin struct {
	Username string
	Password string
}
