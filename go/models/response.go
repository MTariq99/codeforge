package models

type ErrorResponse struct {
	Error string
}

type Response struct {
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type LoginUserRes struct {
	UserId        int64   `json:"id" db:"id"`
	UserName      string  `json:"user_name" db:"user_name"`
	User_email    *string `json:"user_email" db:"user_email"`
	User_password string  `json:"user_password" db:"user_password"`
}
