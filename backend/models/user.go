package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"not null"`  // Omit password in JSON responses
}
