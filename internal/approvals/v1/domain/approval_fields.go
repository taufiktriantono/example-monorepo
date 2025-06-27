package domain

import (
	"time"

	"github.com/google/uuid"
)

type ApprovalField struct {
	ID            string    `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	ApprovalID    string    `gorm:"column:approval_id"`
	Approval      Approval  `gorm:"foreignKey:ApprovalID"`
	Field         string    `gorm:"column:field"`
	PreviousValue string    `gorm:"column:previous_value"`
	NewValue      string    `gorm:"column:new_value"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type ApprovalFieldParams struct {
	ApprovalID    string
	Field         string
	PreviousValue string
	NewValue      string
}

func NewApprovalField(p ApprovalFieldParams) *ApprovalField {
	return &ApprovalField{
		ID:            uuid.NewString(),
		ApprovalID:    p.ApprovalID,
		Field:         p.Field,
		PreviousValue: p.PreviousValue,
		NewValue:      p.NewValue,
	}
}
