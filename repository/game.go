package repository

import (
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type GameRepository interface {
	Store(game *models.Game) error
	GetAll() (*[]models.Game, error)
	GetByID(id int) (*models.Game, error)
	Update(id int, game *models.Game) error
	Delete(id int) error
}

type gameRepository struct {
	postgresDB *database.PostgresDB
}

func NewGameRepository(postgresDB *database.PostgresDB) GameRepository {
	return &gameRepository{postgresDB: postgresDB}
}

func (g *gameRepository) Store(game *models.Game) error {
	return g.postgresDB.StoreGame(game)
}

func (g *gameRepository) GetAll() (*[]models.Game, error) {
	return g.postgresDB.FetchGames()
}

func (g *gameRepository) GetByID(id int) (*models.Game, error) {
	return g.postgresDB.FetchGameByID(id)
}

func (g *gameRepository) Update(id int, game *models.Game) error {
	return g.postgresDB.UpdateGame(id, game)
}

func (g *gameRepository) Delete(id int) error {
	return g.postgresDB.DeleteGame(id)
}