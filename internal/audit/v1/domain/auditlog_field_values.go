package domain

import "time"

type AuditLogFieldValue struct {
	ID            string    `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	AuditLogID    string    `gorm:"column:audit_log_id;type:uuid"`
	Field         string    `gorm:"column:field"`
	PreviousValue string    `gorm:"column:previous_value"`
	NewValue      string    `gorm:"column:new_value"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
