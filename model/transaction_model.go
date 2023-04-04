package model

import "github.com/google/uuid"

type OrderModel struct {
	Id           string             `json:"id"`
	TotalPrice   int64              `json:"total_price"`
	OrderDetails []OrderDetailModel `json:"transaction_details"`
}

type OrderCreateUpdateModel struct {
	Id           string                         `json:"id"`
	TotalPrice   *int64                         `json:"total_price"`
	OrderDetails []OrderDetailCreateUpdateModel `json:"transaction_details"`
}

type OrderCreateModel struct {
	Product []string `json:"product"`
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
