package model

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"ordered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type OrderRequest struct {
	CustomerName string        `json:"customer_name" binding:"required"`
	OrderedAt    time.Time     `json:"ordered_at" binding:"required"`
	Items        []ItemRequest `json:"items" binding:"required"`
}
