package entity

type LoginBody struct {
	Email    string
	Password string
}

type LoginRes struct {
	Token string
}