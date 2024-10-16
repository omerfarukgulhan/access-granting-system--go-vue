package entities

import "time"

type UserRole struct {
	UserID    int64     `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	RoleID    int64     `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
