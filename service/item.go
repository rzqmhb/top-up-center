package service

import (
	"strings"

	"github.com/rzqmhb/top-up-center/models"
	repo "github.com/rzqmhb/top-up-center/repository"
)

type ItemService interface {
	GetAll() (*[]models.Item, error)
	GetByGame(gameId int) (*[]models.Item, error)
	GetByKeywords(keywords string) (*[]models.Item, error)
}

type itemService struct {
	itemRepository repo.ItemRepository
}

func NewItemService(itemRepository repo.ItemRepository) ItemService {
	return &itemService{itemRepository: itemRepository}
}

func (i *itemService) GetAll() (*[]models.Item, error) {
	return i.itemRepository.GetAll()
}

func (i *itemService) GetByGame(gameId int) (*[]models.Item, error) {
	return i.itemRepository.GetByGameID(gameId)
}

func (i *itemService) GetByKeywords(keywords string) (*[]models.Item, error) {
	var items *[]models.Item = &[]models.Item{}
	var keywordsInSlice []string = strings.Split(keywords, " ")
	items, err := i.itemRepository.GetByKeywords(keywordsInSlice)
	if err != nil {
		return items, err
	}
	return items, nil
}