package database

import (
	"errors"
	"os"

	"github.com/rzqmhb/top-up-center/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

// get DSN for postgres database from loaded .env file
var dsn = os.Getenv("DSN")

// connecting to postgres database using gorm with the provided credentials
func InitDB() (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDB{DB: db}, nil
}

//------------------------------------------
// DML and DQL operations to 'users' table
//------------------------------------------

func (postgres *PostgresDB) StoreUser(user models.User) error {
	result := postgres.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchUsers() ([]models.User, error) {
	var users []models.User = []models.User{}
	result := postgres.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	if result.RowsAffected == 0 {
		return users, errors.New("no data found")
	}
	return users, nil
}

func (postgres *PostgresDB) FetchUserByUsername(username string) (models.User, error) {
	var user models.User = models.User{}
	result := postgres.DB.Raw("SELECT * FROM users WHERE name = ?", username).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("no data found")
	}
	return user, nil
}

func (postgres *PostgresDB) FetchUserByEmail(email string) (models.User, error) {
	var user models.User = models.User{}
	result := postgres.DB.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("no data found")
	}
	return user, nil
}

func (postgres *PostgresDB) UpdateUser(id int, user models.User) error {
	result := postgres.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (postgres *PostgresDB) DeleteUser(id int) error {
	result := postgres.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

//---------------------------------------------
// DML and DQL operations to 'sessions' table
//---------------------------------------------

func (postgres *PostgresDB) StoreSession(session models.Session) error {
	result := postgres.DB.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchSessions() ([]models.Session, error) {
	var sessions []models.Session = []models.Session{}
	result := postgres.DB.Find(&sessions)
	if result.Error != nil {
		return sessions, result.Error
	}
	if result.RowsAffected == 0 {
		return sessions, errors.New("no data found")
	}
	return sessions, nil
}

func (postgres *PostgresDB) FetchSessionByToken(token string) (models.Session, error) {
	var session models.Session = models.Session{}
	result := postgres.DB.Raw("SELECT * FROM sessions WHERE token = ?", token).Scan(&session)
	if result.Error != nil {
		return session, result.Error
	}
	if result.RowsAffected == 0 {
		return session, errors.New("no data found")
	}
	return session, nil
}

func (postgres *PostgresDB) FetchSessionByUsername(username string) (models.Session, error) {
	var session models.Session = models.Session{}
	result := postgres.DB.Raw("SELECT * FROM sessions WHERE user_name = ?", username).Scan(&session)
	if result.Error != nil {
		return session, result.Error
	}
	if result.RowsAffected == 0 {
		return session, errors.New("no data found")
	}
	return session, nil
}

func(postgres *PostgresDB) UpdateSession(username string, session models.Session) error {
	result := postgres.DB.Model(&models.User{}).Where("username = ?", username).Updates(&session)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func(postgres *PostgresDB) DeleteSession(token string) error {
	result := postgres.DB.Delete(&models.Session{}, "token = ?", token)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

//-------------------------------------------
// DML and DQL operations to 'games' table
//-------------------------------------------

func (postgres *PostgresDB) StoreGame(game models.Game) error {
	result := postgres.DB.Create(&game)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//-------------------------------------------
// DML and DQL operations to 'items' table
//-------------------------------------------

func (postgres *PostgresDB) StoreItem(item models.Item) error {
	result := postgres.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//-------------------------------------------
// DML and DQL operations to 'orders' table
//-------------------------------------------

func (postgres *PostgresDB) StoreOrder(order models.Order) error {
	result := postgres.DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//------------------------------------------------
// DML and DQL operations to 'order_items' table
//------------------------------------------------

func (postgres *PostgresDB) StoreOrderItem(orderItem models.OrderItem) error {
	result := postgres.DB.Create(&orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}