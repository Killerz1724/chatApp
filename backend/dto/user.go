package dto

type ResUserProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Img      string `json:"img"`
}

type ResUserFriendProf struct {
	Id       int    `json:"id" omitempty:"true"`
	Username string `json:"username"  omitempty:"true"`
	Tag      string `json:"tag" omitempty:"true"`
	Img      string `json:"img" omitempty:"true"`
}

type ResUserFriend struct {
	Friend_list []ResUserFriendProf `json:"friend_list"`
}