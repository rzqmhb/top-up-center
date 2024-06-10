package repository

import (
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type UserRepository interface {
	Store(user *models.User) error
	GetAll() (*[]models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(id int, user *models.User) error
	Delete(id int) error
}

type userRepository struct {
	postgresDB *database.PostgresDB
}

func NewUserRepository(postgresDB *database.PostgresDB) UserRepository {
	return &userRepository{postgresDB: postgresDB}
}

func (u *userRepository) Store(user *models.User) error {
	return u.postgresDB.StoreUser(user)
}

func (u *userRepository) GetAll() (*[]models.User, error) {
	return u.postgresDB.FetchUsers()
}

func (u *userRepository) GetByUsername(username string) (*models.User, error) {
	return u.postgresDB.FetchUserByUsername(username)
}

func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	return u.postgresDB.FetchUserByEmail(email)
}

func (u *userRepository) Update(id int, user *models.User) error {
	return u.postgresDB.UpdateUser(id, user)
}

func (u *userRepository) Delete(id int) error {
	return u.postgresDB.DeleteUser(id)
}