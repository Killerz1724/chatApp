package dto

import "time"

type LastMessage struct {
	Message string `json:"message"`
	Time    time.Time `json:"last_message_time"`
}

type User struct {
	Id       int `json:"id"` 
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Img      string `json:"img"`
}

type ResListConvBody struct {
	ConversationId int `json:"conversation_id"` 
	IsGroup           bool `json:"is_group"`
	User           User `json:"user_data"`
	Message        LastMessage `json:"last_message"`
}