package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Role     int
}

var a int = 10
