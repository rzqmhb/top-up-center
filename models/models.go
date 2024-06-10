package models

import "time"

type User struct {
	ID        int       `json:"user_id" gorm:"column:id"` 
	Name      string    `json:"user_name" gorm:"column:name"`
	Email     string    `json:"user_email" gorm:"column:email"`
	Password  string    `json:"user_password" gorm:"column:password"`
	CreatedAt time.Time `json:"user_created_at" gorm:"column:created_at"`
	UpdatedAT time.Time `json:"user_updated_at" gorm:"column:updated_at"`
}

type Session struct {
	ID       int       `json:"session_id" gorm:"column:id"`
	UserName string    `json:"session_user_name" gorm:"column:user_name"`
	Token    string    `json:"session_token" gorm:"column:token"`
	Expiry   time.Time `json:"session_expiry" gorm:"column:expiry"`
}

type Game struct {
	ID        int       `json:"game_id" gorm:"column:id"`
	Name      string    `json:"game_name" gorm:"column:name"`
	CreatedAt time.Time `json:"game_created_at" gorm:"column:created_at"`
	UpdatedAT time.Time `json:"game_updated_at" gorm:"column:updated_at"`
}

type Item struct {
	ID           int       `json:"item_id" gorm:"column:id"`
	GameID       int       `json:"item_game_id" gorm:"column:game_id"`
	VendorItemID string    `json:"item_vendor_item_id" gorm:"column:vendor_item_id"`
	Name         string    `json:"item_name" gorm:"column:name"`
	Price        float64   `json:"item_price" gorm:"column:price"`
	CreatedAt    time.Time `json:"item_created_at" gorm:"column:created_at"`
	UpdatedAT    time.Time `json:"item_updated_at" gorm:"column:updated_at"`
}

type Order struct {
	ID               int       `json:"order_id" gorm:"column:id"`
	ItemID           int       `json:"order_item_id" gorm:"column:item_id"`
	UserID           int       `json:"order_user_id" gorm:"column:user_id"`
	ItemCurrentPrice float64   `json:"order_item_current_price" gorm:"column:item_current_price"`
	InGameUserID     string    `json:"order_in_game_user_id" gorm:"column:in_game_user_id"`
	Status           string    `json:"order_status" gorm:"column:status"`
	Date             time.Time `json:"order_date" gorm:"column:date"`
}

type JoinedOrderData struct {
	ItemName     string    `json:"joined_order_item_name" gorm:"column:item_name"`
	ItemPrice    float64   `json:"joined_order_item_price" gorm:"column:item_price"`
	OrderDate    time.Time `json:"joined_order_date" gorm:"column:date"`
	OrderStatus  string    `json:"joined_order_status" gorm:"column:status"`
	InGameUserID string    `json:"joined_order_in_game_user_id" gorm:"column:in_game_user_id"`
}
