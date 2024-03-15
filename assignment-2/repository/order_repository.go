package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id string) (model.Order, error)
	UpdateOrder(id string, order *model.Order) (model.Order, error)
	DeleteOrder(id string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderByID(id string) (model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").First(&order, "id = ?", id).Error
	return order, err
}

func (r *orderRepository) UpdateOrder(id string, updatedOrder *model.Order) (model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		return model.Order{}, err
	}

	err = r.db.Model(&order).Updates(&updatedOrder).Error
	if err != nil {
		return model.Order{}, err
	}

	if err := r.db.Model(&order).Association("Items").Replace(updatedOrder.Items); err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) DeleteOrder(id string) error {
	return r.db.Delete(&model.Order{}, "id = ?", id).Error
}
