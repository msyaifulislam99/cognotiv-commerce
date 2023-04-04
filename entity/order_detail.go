package entity

import "github.com/google/uuid"

type OrderDetail struct {
	Id            uuid.UUID `gorm:"primaryKey;column:order_detail_id;type:varchar(36)"`
	SubTotalPrice int64     `gorm:"column:sub_total_price"`
	Price         int64     `gorm:"column:price"`
	Quantity      int32     `gorm:"column:quantity"`
	OrderId       uuid.UUID
	Order         Order
	ProductId     uuid.UUID
	Product       Product
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
