package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                    string    `gorm:"column:id"`
	OrganizationID        string    `gorm:"column:organization_id"`
	AppID                 string    `gorm:"column:app_id"`
	CustomerID            string    `gorm:"column:customer_id"`
	TransactionID         string    `gorm:"column:transaction_id"`
	ExternalTransactionID string    `gorm:"column:external_transaction_id"`
	Amount                float64   `gorm:"column:amount"`
	TransactionDate       time.Time `gorm:"column:transaction_date"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at"`
}

type TransactionParams struct {
	OrganizationID        string
	AppID                 string
	CustomerID            string
	Amount                float64
	ExternalTransactionID string
	TransactionDate       time.Time
}

func NewTransaction(p TransactionParams) *Transaction {
	return &Transaction{
		ID:                    uuid.NewString(),
		OrganizationID:        p.OrganizationID,
		AppID:                 p.AppID,
		CustomerID:            p.CustomerID,
		ExternalTransactionID: p.ExternalTransactionID,
		Amount:                p.Amount,
		TransactionDate:       p.TransactionDate,
	}
}
