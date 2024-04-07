package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Login Credentials
type LoginCredentials struct {
	Username string `json:"username" form:"username" validation:"required"`
	Password string `json:"password" form:"password" validation:"required"`
}

type Admin struct {
	gorm.Model
	Username string
	Password string
}

type Responce struct {
	Id     int
	Status string
	Token  string
}

type OtpVerify struct {
	Email string
	OTP   string
}

type UserAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type QuizResponse struct {
	Question string `json:"question"`
	A        string `json:"a"`
	B        string `json:"b"`
	C        string `json:"c"`
	D        string `json:"d"`
	Answer   string `json:"answer"`
}
