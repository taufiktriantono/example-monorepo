package dto

import (
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
)

type ListAuditLogRequest struct {
	pagination.Pagination
	option.QuerySortBy
	Types          []string `form:"types[]"`
	UserID         string   `form:"user_id"`
	OrganizationID string   `form:"organization_id"`
	ResouceID      string   `form:"resource_id"`
	Action         string   `form:"action"`
}
