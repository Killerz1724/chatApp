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

type ReqUserFriendDetailBody struct {
	CurrentUserEmail string
	Id               int
}

type ResUserFriend struct {
	Id       int
	Username string
	Tag      string
	Img      string
}
