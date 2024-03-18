package model

import "gorm.io/gorm"

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
