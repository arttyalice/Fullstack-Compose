package models

import "gorm.io/gorm"

// Product
type Product struct {
	gorm.Model

	Name         string  `gorm:"type:varchar(255);" json:"name"`
	Description  string  `gorm:"type:text;" json:"description"`
	PricePerUnit float64 `json:"price_per_unit"`
	UnitName     string  `gorm:"type:varchar(50)" json:"unit_name"`

	CategoryID uint `json:"category_id,omitempty"`

	Quantity uint `json:"quantity"`

	StoreProducts []StoreProduct `gorm:"foreignKey:ProductID;" json:"store_products,omitempty"`
}

// SearchProduct - Search Product
type SearchProduct struct {
	Text *int `json:"text" form:"text"`
	Page *int `json:"page" form:"page"`
	Size *int `json:"size" form:"size"`
}

// ListProduct - model to use as []ListProduct
type ListProduct struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	PricePerUnit float64 `json:"price_per_unit"`
	UnitName     string  `json:"unit_name"`
	Quantity     uint    `json:"quantity"`
	CategoryID   uint    `json:"category_id"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*ListProduct) TableName() string {
	return "products"
}

// ProductInfo - Product infomation
type ProductInfo struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	PricePerUnit float64 `json:"price_per_unit"`
	UnitName     string  `json:"unit_name"`

	CategoryID uint `gorm:"category_id" json:"-"`

	Quantity uint `json:"quantity"`

	Category ProductCategory `gorm:"foreignKey:ID;reference:CategoryID;" json:"category"`
	Stores   []ListStore     `gorm:"-" json:"stores"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*ProductInfo) TableName() string {
	return "products"
}

type CreateProduct struct {
	gorm.Model
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description"`
	PricePerUnit float64 `json:"price_per_unit" binding:"required"`
	UnitName     string  `json:"unit_name" binding:"required"`
	Quantity     uint    `json:"quantity" binding:"required"`

	CategoryID uint `json:"category_id" binding:"required"`
}

func (*CreateProduct) TableName() string {
	return "products"
}

type UpdateProduct struct {
	ID           uint     `json:"id" binding:"required"`
	Name         *string  `json:"name"`
	Description  *string  `json:"description"`
	PricePerUnit *float64 `json:"price_per_unit"`
	UnitName     *string  `json:"unit_name"`
	Quantity     *uint    `json:"quantity"`

	CategoryID *uint `json:"category_id"`
}

func (*UpdateProduct) TableName() string {
	return "products"
}
