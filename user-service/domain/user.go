package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
