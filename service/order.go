package service

import (
	"github.com/rzqmhb/top-up-center/models"
	repo "github.com/rzqmhb/top-up-center/repository"
)

type OrderService interface {
	GetAllUserOrder(userId int) (*[]models.JoinedOrderData, error)
}

type orderService struct {
	orderRepository repo.OrderRepository
}

func NewOrderService(orderRepository repo.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (o *orderService) GetAllUserOrder(userId int) (*[]models.JoinedOrderData, error) {
	return o.orderRepository.GetJoinedOrderByUserID(userId)
}