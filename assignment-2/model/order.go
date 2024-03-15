package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID           string         `json:"id"`
	CustomerName string         `json:"customer_name"`
	Items        []Item         `json:"items"`
	OrderedAt    time.Time      `json:"ordered_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	o.OrderedAt = time.Now()
	return nil
}

func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	o.UpdatedAt = time.Now()
	return nil
}
