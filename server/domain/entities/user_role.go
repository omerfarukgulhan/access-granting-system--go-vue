package entities

import "time"

type UserRole struct {
	UserId    int64     `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	RoleId    int64     `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
