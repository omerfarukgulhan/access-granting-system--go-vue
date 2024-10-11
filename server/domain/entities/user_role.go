package entities

type UserRole struct {
	UserId int64 `gorm:"primaryKey"`
	RoleId int64 `gorm:"primaryKey"`
}
