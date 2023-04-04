package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id              uuid.UUID     `gorm:"primaryKey;column:id;type:varchar(36)"`
	TotalPrice      int64         `gorm:"column:total_price"`
	OrderDetails    []OrderDetail `gorm:"ForeignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TransactionDate time.Time     `gorm:"column:transaction_date"`
	Status          string        `gorm:"column:status;default:pending"`
	UserId          uuid.UUID
	User            User
}

func (Order) TableName() string {
	return "order"
}
