package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);" json:"name"`
	Description string `gorm:"type:text;" json:"description"`

	Products []Product `gorm:"foreignKey:CategoryID;" json:"store_products,omitempty"`
}

type SearchProductCategory struct {
	Text *int `json:"text" form:"text"`
	Page *int `json:"page" form:"page"`
	Size *int `json:"size" form:"size"`
}

type ListProductCategory struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*ListProductCategory) TableName() string {
	return "product_categories"
}

type CreateProductCategory struct {
	gorm.Model

	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (*CreateProductCategory) TableName() string {
	return "product_categories"
}

type UpdateProductCategory struct {
	ID          uint    `json:"id" binding:"required"`
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

func (*UpdateProductCategory) TableName() string {
	return "product_categories"
}
