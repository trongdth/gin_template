package models

import (
	"github.com/jinzhu/gorm"
)

// User : struct
type User struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
	FullName   string
	UserName   string
	Email      string
	Password   string
	RoleID     uint
}

// UserSubscribe : struct
type UserSubscribe struct {
	gorm.Model

	Email   string
	Name    string
	Company string
}
