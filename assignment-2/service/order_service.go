package service

import (
	"assignment-2/model"
	"assignment-2/repository"
)

type OrderService interface {
	CreateOrder(order model.OrderRequest) (model.Order, error)
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id uint) (model.Order, error)
	UpdateOrder(order model.OrderRequest) (model.Order, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepo repository.OrderRepository
	itemRepo  repository.ItemRepository
}

func NewOrderService(orderRepo repository.OrderRepository, itemRepo repository.ItemRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		itemRepo:  itemRepo,
	}
}

func (s *orderService) CreateOrder(order model.OrderRequest) (model.Order, error)

func (s *orderService) GetAllOrders() ([]model.Order, error)

func (s *orderService) GetOrderByID(id uint) (model.Order, error)

func (s *orderService) UpdateOrder(order model.OrderRequest) (model.Order, error)

func (s *orderService) DeleteOrder(id uint) error
