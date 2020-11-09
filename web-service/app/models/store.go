package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	PhoneNumber string `gorm:"type:varchar(10);" json:"phone_number"` // assume use in thailand only
	Address     string `gorm:"type:text;" json:"address"`

	StoreProducts []StoreProduct `gorm:"foreignKey:StoreID;" json:"store_products,omitempty"`
}

// SearchStore - Search Store
type SearchStore struct {
	Text *int `json:"text" form:"text"`
	Page *int `json:"page" form:"page"`
	Size *int `json:"size" form:"size"`
}

// ListStore - model to use as []ListStore
type ListStore struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*ListStore) TableName() string {
	return "stores"
}

// StoreInfo - Product infomation
type StoreInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`

	Products []ListProduct `gorm:"-" json:"products"`

	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*StoreInfo) TableName() string {
	return "stores"
}

type CreateStore struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

func (*CreateStore) TableName() string {
	return "stores"
}

type UpdateStore struct {
	ID          uint    `json:"id" binding:"required"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	PhoneNumber *string `json:"phone_number"`
	Address     *string `json:"address"`
}

func (*UpdateStore) TableName() string {
	return "stores"
}
