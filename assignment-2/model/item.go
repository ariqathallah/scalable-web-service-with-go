package model

type Item struct {
	ItemID      uint   `json:"item_id" gorm:"unique;primaryKey;autoIncrement"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type ItemRequest struct {
	ItemCode    string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
