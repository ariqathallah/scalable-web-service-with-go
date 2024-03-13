package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id uint) (model.Order, error)
	UpdateOrder(order *model.Order) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(&order).Error
}

func (r *orderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderByID(id uint) (model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").Where("id = ?", id).First(&order).Error
	return order, err
}

func (r *orderRepository) UpdateOrder(order *model.Order) error {
	return r.db.Save(&order).Error
}

func (r *orderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&model.Order{}, id).Error
}
