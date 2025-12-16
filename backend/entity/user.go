package entity

type ReqUserProfileBody struct {
	Email   string
	SearchP string
}

type ResUserProfile struct {
	Email    string
	Username string
	Tag      string
	Img      string
}

type ResUserFriend struct {
	Username string
	Tag      string
	Img      string
}
