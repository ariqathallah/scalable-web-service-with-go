package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID          string `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     string `json:"order_id"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.New().String()
	return nil
}
