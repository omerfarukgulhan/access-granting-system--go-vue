package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                 int64          `gorm:"primaryKey;autoIncrement:false"`
	Username           string         `gorm:"size:255;not null;unique"`
	Email              string         `gorm:"size:255;not null;unique"`
	Password           string         `gorm:"size:255;not null"`
	ProfileImage       string         `gorm:"size:255;default:'default.png'"`
	IsActive           bool           `gorm:"default:false"`
	ActivationToken    string         `gorm:"size:255"`
	PasswordResetToken string         `gorm:"size:255"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Roles              []Role         `gorm:"many2many:user_roles"`
}
