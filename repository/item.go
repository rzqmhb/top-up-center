package repository

import (
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type ItemRepository interface {
	Store(item *models.Item) error
	GetAll() (*[]models.Item, error)
	GetByID(id int) (*models.Item, error)
	GetByGameID(gameId int) (*models.Item, error)
	Update(id int, item *models.Item) error
	Delete(id int) error
}

type itemRepository struct {
	postgresDB *database.PostgresDB
}

func NewItemRepository(postgresDB *database.PostgresDB) ItemRepository {
	return &itemRepository{postgresDB: postgresDB}
}

func (i *itemRepository) Store(item *models.Item) error {
	return i.postgresDB.StoreItem(item)
}

func (i *itemRepository) GetAll() (*[]models.Item, error) {
	return i.postgresDB.FetchItems()
}

func (i *itemRepository) GetByID(id int) (*models.Item, error) {
	return i.postgresDB.FetchItemByID(id)
}

func (i *itemRepository) GetByGameID(gameId int) (*models.Item, error) {
	return i.postgresDB.FetchItemByGameID(gameId)
}

func (i *itemRepository) Update(id int, item *models.Item) error {
	return i.postgresDB.UpdateItem(id, item)
}

func (i *itemRepository) Delete(id int) error {
	return i.postgresDB.DeleteItem(id)
}