package dto

import "github.com/taufiktriantono/api-first-monorepo/pkg/repository"

type ListAuditLogRequest struct {
	repository.Pagination
	repository.QuerySortBy
	repository.QueryStartAndEndDate
	Types          []string `form:"types[]"`
	UserID         string   `form:"user_id"`
	OrganizationID string   `form:"organization_id"`
	ResouceID      string   `form:"resource_id"`
	Action         string   `form:"action"`
}
