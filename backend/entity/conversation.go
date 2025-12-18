package entity

import "time"

type ReqListConvBody struct {
	Email string
}

type LastMessage struct {
	Message string
	Time    time.Time
}

type User struct {
	Id       int
	Username string
	Tag      string
	Img      string
}

type ResListConvBody struct {
	ConversationId    int
	IsGroup           bool
	User 	User
	Message  LastMessage
}
