package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderModel struct {
	Id           string             `json:"id"`
	TotalPrice   int64              `json:"total_price"`
	OrderDetails []OrderDetailModel `json:"transaction_details"`
}

type OrderModelWithUser struct {
	Id              string             `json:"id"`
	TotalPrice      int64              `json:"total_price"`
	TransactionDate time.Time          `json:"transaction_date"`
	Status          string             `json:"status"`
	OrderDetails    []OrderDetailModel `json:"transaction_details"`
	User            UserModel          `json:"customer"`
}

type OrderCreateUpdateModel struct {
	Id           string                         `json:"id"`
	TotalPrice   *int64                         `json:"total_price"`
	OrderDetails []OrderDetailCreateUpdateModel `json:"transaction_details"`
}

type OrderCreateModel struct {
	Product []ProductWithQty `json:"product" validate:"min=1"`
}

type ProductWithQty struct {
	Id  string `json:"productId"`
	Qty int64  `json:"qty"`
}

type OrderDetailModel struct {
	Id            string `json:"id"`
	SubTotalPrice int64  `json:"sub_total_price" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	Quantity      int32  `json:"quantity" validate:"required"`
	Product       ProductModel
}

type OrderDetailCreateUpdateModel struct {
	Id            string    `json:"id"`
	SubTotalPrice int64     `json:"sub_total_price" validate:"required"`
	Price         int64     `json:"price" validate:"required"`
	Quantity      int32     `json:"quantity" validate:"required"`
	ProductId     uuid.UUID `json:"product_id" validate:"required"`
	Product       ProductModel
}
