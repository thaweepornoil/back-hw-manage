package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
}
