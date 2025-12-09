package dto

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type Response struct {
	Success bool           `json:"success"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
}