package models

import "time"

type User struct {
	ID        int       `json:"user_id"`
	Name      string    `json:"user_name"`
	Email     string    `json:"user_email"`
	Password  string    `json:"user_password"`
	CreatedAt time.Time `json:"user_created_at"`
	UpdatedAT time.Time `json:"user_updated_at"`
}

type Session struct {
	ID       int       `json:"session_id"`
	UserName string    `json:"session_user_name"`
	Token    string    `json:"session_token"`
	Expiry   time.Time `json:"session_expiry"`
}

type Game struct {
	ID        int       `json:"game_id"`
	Name      string    `json:"game_name"`
	CreatedAt time.Time `json:"game_created_at"`
	UpdatedAT time.Time `json:"game_updated_at"`
}

type Item struct {
	ID           int       `json:"item_id"`
	GameID       int       `json:"item_game_id"`
	VendorItemID int       `json:"item_vendor_item_id"`
	Name         string    `json:"item_name"`
	Price        float64   `json:"item_price"`
	CreatedAt    time.Time `json:"item_created_at"`
	UpdatedAT    time.Time `json:"item_updated_at"`
}

type Order struct {
	ID               int       `json:"order_id"`
	ItemID           int       `json:"order_item_id"`
	UserID           int       `json:"order_user_id"`
	ItemCurrentPrice float64   `json:"order_item_current_price"`
	InGameUserID     int       `json:"order_in_game_user_id"`
	Status           string    `json:"order_status"`
	Date             time.Time `json:"order_date"`
}