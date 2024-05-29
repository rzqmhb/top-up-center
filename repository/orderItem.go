package repository

import (
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type OrderItemRepository interface {
	Store(orderItem *models.OrderItem) error
	GetAll() (*[]models.OrderItem, error)
	GetByID(id int) (*models.OrderItem, error)
	GetByOrderID(orderId int) (*models.OrderItem, error)
	GetByItemID(itemId int) (*models.OrderItem, error)
	Update(id int, orderItem *models.OrderItem) error
	Delete(id int) error
}

type orderItemRepository struct {
	postgresDB *database.PostgresDB
}

func NewItemOrderRepository(postgresDB *database.PostgresDB) OrderItemRepository {
	return &orderItemRepository{postgresDB: postgresDB}
}

func (o *orderItemRepository) Store(orderItem *models.OrderItem) error {
	return o.postgresDB.StoreOrderItem(orderItem)
}

func (o *orderItemRepository) GetAll() (*[]models.OrderItem, error) {
	return o.postgresDB.FetchOrderItems()
}

func (o *orderItemRepository) GetByID(id int) (*models.OrderItem, error) {
	return o.postgresDB.FetchOrderItemByID(id)
}

func (o *orderItemRepository) GetByOrderID(orderId int) (*models.OrderItem, error) {
	return o.postgresDB.FetchOrderItemByOrderID(orderId)
}

func (o *orderItemRepository) GetByItemID(itemId int) (*models.OrderItem, error) {
	return o.postgresDB.FetchOrderItemByItemID(itemId)
}

func (o *orderItemRepository) Update(id int, orderItem *models.OrderItem) error {
	return o.postgresDB.UpdateOrderItem(id, orderItem)
}

func (o *orderItemRepository) Delete(id int) error {
	return o.postgresDB.DeleteOrderItem(id)
}