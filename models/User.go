package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Rollno      string `json:"rollno"`
	Branch      string `json:"branch" `
	Course      string `json:"course" `
	PassoutYear string `json:"passoutYear"`
	DOB         string `json:"dob"`
}

type UserLogin struct {
	gorm.Model
	RefreshToken      string  `json:"refresh_token"`
	Verified          bool    `gorm:"default:false" json:"verified"`
	VerificationToken string  `json:"verification_token"`
	Name              string  `json:"name" gorm:"not null"`
	Email             string  `json:"email" gorm:"not null;unique"`
	Role              string  `json:"role" gorm:"not null default:'user'"` // user, admin
	College           College `json:"college"`
	CollegeID         uint    `json:"college_id" gorm:"not null"`
	Password          string
	UserInfo          uint
}

type College struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique"`
	EmailFormat string `gorm:"not null"` // do not expose the college email format while marshalling
	Verified    bool   `json:"verified" gorm:"default:false"`
}
