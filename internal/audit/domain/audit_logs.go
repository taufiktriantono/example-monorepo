package domain

import "time"

type AuditLog struct {
	ID             string               `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         string               `gorm:"column:user_id;type:uuid"`
	OrganizationID string               `gorm:"column:organization_id;type:uuid"`
	ResourceID     string               `gorm:"column:resource_id;type:uuid"`
	ResourceName   string               `gorm:"column:resource_name"`
	Action         string               `gorm:"column:action"`
	Fields         []AuditLogFieldValue `gorm:"foreignKey:AuditLogID"`
	CreatedAt      time.Time            `gorm:"column:created_at"`
	UpdatedAt      time.Time            `gorm:"column:updated_at"`
}
