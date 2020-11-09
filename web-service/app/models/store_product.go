package models

import (
	"time"

	"gorm.io/gorm"
)

type StoreProduct struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	StoreID         uint `json:"store_id"`
	ProductID       uint `json:"product_id,omitempty"`
	ProductQuantity uint `json:"product_quantity"`
}

// SearchStoreProduct - Search StoreProduct
type SearchStoreProduct struct {
	Text *int `json:"text" form:"text"`
	Page *int `json:"page" form:"page"`
	Size *int `json:"size" form:"size"`
}

// ListStoreProduct - model to use as []ListStoreProduct
type ListStoreProduct struct {
	StoreID         uint `json:"store_id"`
	ProductID       uint `json:"product_id,omitempty"`
	ProductQuantity uint `json:"product_quantity"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*ListStoreProduct) TableName() string {
	return "store_products"
}

// StoreProductInfo - StoreProduct infomation
type StoreProductInfo struct {
	StoreID         uint `json:"store_id"`
	ProductID       uint `json:"product_id,omitempty"`
	ProductQuantity uint `json:"product_quantity"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*StoreProductInfo) TableName() string {
	return "store_products"
}

type CreateStoreProduct struct {
	gorm.Model
	StoreID         uint `json:"store_id" binding:"required"`
	ProductID       uint `json:"product_id,omitempty" binding:"required"`
	ProductQuantity uint `json:"product_quantity" binding:"required"`
}

func (*CreateStoreProduct) TableName() string {
	return "store_products"
}

type UpdateStoreProduct struct {
	ID              uint  `json:"id" binding:"required"`
	StoreID         *uint `json:"store_id"`
	ProductID       *uint `json:"product_id,omitempty"`
	ProductQuantity *uint `json:"product_quantity"`
}

func (*UpdateStoreProduct) TableName() string {
	return "store_products"
}
