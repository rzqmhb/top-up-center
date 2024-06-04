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

func (postgres *PostgresDB) StoreUser(user *models.User) error {
	result := postgres.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchUsers() (*[]models.User, error) {
	var users []models.User = []models.User{}
	result := postgres.DB.Find(&users)
	if result.Error != nil {
		return &users, result.Error
	}
	if result.RowsAffected == 0 {
		return &users, errors.New("no data found")
	}
	return &users, nil
}

func (postgres *PostgresDB) FetchUserByUsername(username string) (*models.User, error) {
	var user models.User = models.User{}
	result := postgres.DB.Raw("SELECT * FROM users WHERE name = ?", username).Scan(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	if result.RowsAffected == 0 {
		return &user, errors.New("no data found")
	}
	return &user, nil
}

func (postgres *PostgresDB) FetchUserByEmail(email string) (*models.User, error) {
	var user models.User = models.User{}
	result := postgres.DB.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	if result.RowsAffected == 0 {
		return &user, errors.New("no data found")
	}
	return &user, nil
}

func (postgres *PostgresDB) UpdateUser(id int, user *models.User) error {
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

func (postgres *PostgresDB) StoreSession(session *models.Session) error {
	result := postgres.DB.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchSessions() (*[]models.Session, error) {
	var sessions []models.Session = []models.Session{}
	result := postgres.DB.Find(&sessions)
	if result.Error != nil {
		return &sessions, result.Error
	}
	if result.RowsAffected == 0 {
		return &sessions, errors.New("no data found")
	}
	return &sessions, nil
}

func (postgres *PostgresDB) FetchSessionByToken(token string) (*models.Session, error) {
	var session models.Session = models.Session{}
	result := postgres.DB.Raw("SELECT * FROM sessions WHERE token = ?", token).Scan(&session)
	if result.Error != nil {
		return &session, result.Error
	}
	if result.RowsAffected == 0 {
		return &session, errors.New("no data found")
	}
	return &session, nil
}

func (postgres *PostgresDB) FetchSessionByUsername(username string) (*models.Session, error) {
	var session models.Session = models.Session{}
	result := postgres.DB.Raw("SELECT * FROM sessions WHERE user_name = ?", username).Scan(&session)
	if result.Error != nil {
		return &session, result.Error
	}
	if result.RowsAffected == 0 {
		return &session, errors.New("no data found")
	}
	return &session, nil
}

func(postgres *PostgresDB) UpdateSession(username string, session *models.Session) error {
	result := postgres.DB.Exec("UPDATE sessions SET token = ?, expiry = ? WHERE user_name = ?;", session.Token, session.Expiry, username)
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

func (postgres *PostgresDB) StoreGame(game *models.Game) error {
	result := postgres.DB.Create(&game)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchGames() (*[]models.Game, error) {
	var games []models.Game = []models.Game{}
	result := postgres.DB.Find(&games)
	if result.Error != nil {
		return &games, result.Error
	}
	if result.RowsAffected == 0 {
		return &games, errors.New("no data found")
	}
	return &games, nil
}

func (postgres *PostgresDB) FetchGameByID(id int) (*models.Game, error) {
	var game models.Game = models.Game{}
	result := postgres.DB.First(&game, id)
	if result.Error != nil {
		return &game, result.Error
	}
	if result.RowsAffected == 0 {
		return &game, errors.New("no data found")
	}
	return &game, nil
}

func (postgres *PostgresDB) UpdateGame(id int, game *models.Game) error {
	result := postgres.DB.Model(&models.Game{}).Where("id = ?", id).Updates(&game)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (postgres *PostgresDB) DeleteGame(id int) error {
	result := postgres.DB.Delete(&models.Game{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

//-------------------------------------------
// DML and DQL operations to 'items' table
//-------------------------------------------

func (postgres *PostgresDB) StoreItem(item *models.Item) error {
	result := postgres.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchItems() (*[]models.Item, error) {
	var items []models.Item = []models.Item{}
	result := postgres.DB.Find(&items)
	if result.Error != nil {
		return &items, result.Error
	}
	if result.RowsAffected == 0 {
		return &items, errors.New("no data found")
	}
	return &items, nil
}

func (postgres *PostgresDB) FetchItemByID(id int) (*models.Item, error) {
	var item models.Item = models.Item{}
	result := postgres.DB.First(&item, id)
	if result.Error != nil {
		return &item, result.Error
	}
	if result.RowsAffected == 0 {
		return &item, errors.New("no data found")
	}
	return &item, nil
}

func (postgres *PostgresDB) FetchItemByGameID(gameId int) (*models.Item, error) {
	var item models.Item = models.Item{}
	result := postgres.DB.Raw("SELECT * FROM items WHERE game_id = ?", gameId).Scan(&item)
	if result.Error != nil {
		return &item, result.Error
	}
	if result.RowsAffected == 0 {
		return &item, errors.New("no data found")
	}
	return &item, nil
}

func (postgres *PostgresDB) UpdateItem(id int, item *models.Item) error {
	result := postgres.DB.Model(&models.Item{}).Where("id = ?", id).Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (postgres *PostgresDB) DeleteItem(id int) error {
	result := postgres.DB.Delete(&models.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

//-------------------------------------------
// DML and DQL operations to 'orders' table
//-------------------------------------------

func (postgres *PostgresDB) StoreOrder(order *models.Order) error {
	result := postgres.DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchOrders() (*[]models.Order, error) {
	var orders []models.Order = []models.Order{}
	result := postgres.DB.Find(&orders)
	if result.Error != nil {
		return &orders, result.Error
	}
	if result.RowsAffected == 0 {
		return &orders, errors.New("no data found")
	}
	return &orders, nil
}

func (postgres *PostgresDB) FetchOrderByID(id int) (*models.Order, error) {
	var order models.Order = models.Order{}
	result := postgres.DB.First(&order, id)
	if result.Error != nil {
		return &order, result.Error
	}
	if result.RowsAffected == 0 {
		return &order, errors.New("no data found")
	}
	return &order, nil
}

func (postgres *PostgresDB) FetchOrderByUserID(userId int) (*models.Order, error) {
	var order models.Order = models.Order{}
	result := postgres.DB.Raw("SELECT * FROM order WHERE user_id = ?", userId).Scan(&order)
	if result.Error != nil {
		return &order, result.Error
	}
	if result.RowsAffected == 0 {
		return &order, errors.New("no data found")
	}
	return &order, nil
}

func (postgres *PostgresDB) UpdateOrder(id int, order *models.Order) error {
	result := postgres.DB.Model(&models.Order{}).Where("id = ?", id).Updates(&order)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (postgres *PostgresDB) DeleteOrder(id int) error {
	result := postgres.DB.Delete(&models.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

//------------------------------------------------
// DML and DQL operations to 'order_items' table
//------------------------------------------------

func (postgres *PostgresDB) StoreOrderItem(orderItem *models.OrderItem) error {
	result := postgres.DB.Create(&orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postgres *PostgresDB) FetchOrderItems() (*[]models.OrderItem, error) {
	var orderItems []models.OrderItem = []models.OrderItem{}
	result := postgres.DB.Find(&orderItems)
	if result.Error != nil {
		return &orderItems, result.Error
	}
	if result.RowsAffected == 0 {
		return &orderItems, errors.New("no data found")
	}
	return &orderItems, nil
}

func (postgres *PostgresDB) FetchOrderItemByID(id int) (*models.OrderItem, error) {
	var orderItem models.OrderItem = models.OrderItem{}
	result := postgres.DB.First(&orderItem, id)
	if result.Error != nil {
		return &orderItem, result.Error
	}
	if result.RowsAffected == 0 {
		return &orderItem, errors.New("no data found")
	}
	return &orderItem, nil
}

func (postgres *PostgresDB) FetchOrderItemByOrderID(orderId int) (*models.OrderItem, error) {
	var orderItem models.OrderItem = models.OrderItem{}
	result := postgres.DB.Raw("SELECT * FROM order WHERE user_id = ?", orderId).Scan(&orderItem)
	if result.Error != nil {
		return &orderItem, result.Error
	}
	if result.RowsAffected == 0 {
		return &orderItem, errors.New("no data found")
	}
	return &orderItem, nil
}

func (postgres *PostgresDB) FetchOrderItemByItemID(itemId int) (*models.OrderItem, error) {
	var orderItem models.OrderItem = models.OrderItem{}
	result := postgres.DB.Raw("SELECT * FROM order WHERE user_id = ?", itemId).Scan(&orderItem)
	if result.Error != nil {
		return &orderItem, result.Error
	}
	if result.RowsAffected == 0 {
		return &orderItem, errors.New("no data found")
	}
	return &orderItem, nil
}

func (postgres *PostgresDB) UpdateOrderItem(id int, orderItem *models.OrderItem) error {
	result := postgres.DB.Model(&models.OrderItem{}).Where("id = ?", id).Updates(&orderItem)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (postgres *PostgresDB) DeleteOrderItem(id int) error {
	result := postgres.DB.Delete(&models.OrderItem{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}