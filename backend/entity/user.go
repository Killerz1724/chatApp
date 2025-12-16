package entity

type ReqUserProfileBody struct {
	Email string
}

type ResUserProfile struct {
	Email    string
	Username string
	Tag      string
	Img      string
}