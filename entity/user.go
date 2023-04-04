package entity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID  `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name      string     `gorm:"column:name;type:varchar(100)"`
	Password  string     `gorm:"column:password;type:varchar(200)"`
	Email     string     `gorm:"column:email;type:varchar(200)"`
	IsActive  bool       `gorm:"column:is_active;type:boolean"`
	UserRoles []UserRole `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Orders    []Order    `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (User) TableName() string {
	return "user"
}
