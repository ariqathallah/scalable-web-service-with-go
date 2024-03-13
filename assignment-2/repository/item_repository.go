package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *model.Item) error
	GetItemByID(id uint) (model.Item, error)
	GetItemsByOrderID(orderID uint) ([]model.Item, error)
	UpdateItem(item *model.Item) error
	DeleteItem(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) CreateItem(item *model.Item) error {
	return r.db.Create(&item).Error
}

func (r *itemRepository) GetItemByID(id uint) (model.Item, error) {
	var item model.Item
	err := r.db.Where("item_id = ?", id).First(&item).Error
	return item, err
}

func (r *itemRepository) GetItemsByOrderID(orderID uint) ([]model.Item, error) {
	var items []model.Item
	err := r.db.Where("order_id = ?", orderID).Find(&items).Error
	return items, err
}

func (r *itemRepository) UpdateItem(item *model.Item) error {
	return r.db.Save(&item).Error
}

func (r *itemRepository) DeleteItem(id uint) error {
	return r.db.Delete(&model.Item{}, id).Error
}
