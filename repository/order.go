package repository

import (
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type OrderRepository interface {
	Store(order *models.Order) error
	GetAll() (*[]models.Order, error)
	GetByID(id int) (*models.Order, error)
	GetByUserID(userId int) (*models.Order, error)
	Update(id int, order *models.Order) error
	Delete(id int) error
}

type orderRepository struct {
	postgresDB *database.PostgresDB
}

func NewOrderRepository(postgresDB *database.PostgresDB) OrderRepository {
	return &orderRepository{postgresDB: postgresDB}
}

func (o *orderRepository) Store(order *models.Order) error {
	return o.postgresDB.StoreOrder(order)
}

func (o *orderRepository) GetAll() (*[]models.Order, error) {
	return o.postgresDB.FetchOrders()
}

func (o *orderRepository) GetByID(id int) (*models.Order, error) {
	return o.postgresDB.FetchOrderByID(id)
}

func (o *orderRepository) GetByUserID(userId int) (*models.Order, error) {
	return o.postgresDB.FetchOrderByUserID(userId)
}

func (o *orderRepository) Update(id int, order *models.Order) error {
	return o.postgresDB.UpdateOrder(id, order)
}

func (o *orderRepository) Delete(id int) error {
	return o.postgresDB.DeleteOrder(id)
}