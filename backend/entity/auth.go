package entity

type LoginBody struct {
	Email    string
	Password string
}

type RegisterBody struct {
	Email    string
	Username string
	Password string
	Tag      string
}

type LoginRes struct {
	Token string
}