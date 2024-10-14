package entities

type Role struct {
	Id    int64  `gorm:"primaryKey;autoIncrement:false"`
	Name  string `gorm:"size:100;not null;unique"`
	Users []User `gorm:"many2many:user_roles"`
}
