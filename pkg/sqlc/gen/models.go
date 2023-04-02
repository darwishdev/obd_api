// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	BranchID  int64        `json:"branch_id"`
	Name      string       `json:"name"`
	Location  string       `json:"location"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Order struct {
	OrderID     int64     `json:"order_id"`
	OrderNumber string    `json:"order_number"`
	TotalAmount float32   `json:"total_amount"`
	UserID      int64     `json:"user_id"`
	BranchID    int64     `json:"branch_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderProduct struct {
	OrderProductsID int64        `json:"order_products_id"`
	OrderID         int64        `json:"order_id"`
	ProductID       int64        `json:"product_id"`
	Qty             float32      `json:"qty"`
	UnitPrice       float32      `json:"unit_price"`
	TotalAmount     float32      `json:"total_amount"`
	DeletedAt       sql.NullTime `json:"deleted_at"`
}

type Product struct {
	ProductID   int64          `json:"product_id"`
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	Price       float32        `json:"price"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	UserID            int64        `json:"user_id"`
	Name              string       `json:"name"`
	Phone             string       `json:"phone"`
	Email             string       `json:"email"`
	Password          string       `json:"password"`
	PasswordChangedAt time.Time    `json:"password_changed_at"`
	CreatedAt         time.Time    `json:"created_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
}
