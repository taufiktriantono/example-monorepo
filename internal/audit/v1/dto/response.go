package dto

import (
	"time"

	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
)

type AuditLogResponse struct {
	ID             string                   `json:"id"`
	UserID         string                   `json:"user_id"`
	OrganizationID string                   `json:"organization_id"`
	Type           string                   `json:"type"`
	ResourceID     string                   `json:"resource_id"`
	ResourceName   string                   `json:"resource_name"`
	Action         string                   `json:"action"`
	Fields         []*AuditLogFieldResponse `json:"fields,omitempty"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
}

type AuditLogFieldResponse struct {
	ID            string    `json:"id"`
	AuditLogID    string    `json:"audit_log_id"`
	Field         string    `json:"field"`
	PreviousValue string    `json:"previous_value"`
	NewValue      string    `json:"new_value"`
	CreatedAt     time.Time `json:"created_at"`
}

type ListAuditLogResponse struct {
	PageInfo pagination.PageInfo `json:"page_info"`
	Data     []*AuditLogResponse `json:"data"`
}
