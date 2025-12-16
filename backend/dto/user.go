package dto

type ResUserProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Img      string `json:"img"`
}

type ResUserFriendProf struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Img      string `json:"img"`
}

type ResUserFriend struct {
	Friend_list []ResUserFriendProf `json:"friend_list"`
}