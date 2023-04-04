package model

type ProductModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ProductCreateOrUpdateModel struct {
	Name        string `json:"name" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
