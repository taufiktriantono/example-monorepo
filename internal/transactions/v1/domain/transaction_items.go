package domain

import "time"

type TransactionItem struct {
	ID        string    `gorm:"column:id"`
	OrderID   string    `gorm:"column:order_id"`
	ItemID    string    `gorm:"column:item_id"`
	ItemName  string    `gorm:"column:item_name"`
	UnitPrice float64   `gorm:"column:unit_price"`
	Quantity  int       `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
