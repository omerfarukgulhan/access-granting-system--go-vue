package entities

import (
	"time"
)

type User struct {
	Id                 int64     `gorm:"primaryKey;autoIncrement"`
	Username           string    `gorm:"size:255;not null;unique"`
	Email              string    `gorm:"size:255;not null;unique"`
	Password           string    `gorm:"size:255;not null"`
	ProfileImage       string    `gorm:"size:255;default:'default.png'"`
	IsActive           bool      `gorm:"default:false"`
	ActivationToken    string    `gorm:"size:255"`
	PasswordResetToken string    `gorm:"size:255"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
	Roles              []Role    `gorm:"many2many:user_roles"`
}
