package entities

import "time"

type UserRole struct {
	UserId    int64     `gorm:"primaryKey"`
	RoleId    int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
