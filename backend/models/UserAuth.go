package models

import "gorm.io/gorm"

// UserAuth struct
type UserAuth struct {
	gorm.Model
	Username     string `gorm:"unique;not null" json:"username"`
	Password     string `json:"password"`
	Phone        string `gorm:"" json:"phone"`
	Email        string `gorm:"unique;not null" json:"email"`
	RefreshToken string `gorm:"" json:"refresh_token"`
	Verified     bool   `gorm:"default:false" json:"verified"`
	DOB          string `json:"dob"`
	OTP          uint   `json:"otp"`
}
