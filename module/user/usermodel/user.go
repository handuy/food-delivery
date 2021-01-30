package usermodel

type User struct {
	Id       string
	Email     string
	Password string
}

type NewUser struct {
	Email     string
	Password string
}