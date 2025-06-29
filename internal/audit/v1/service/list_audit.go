package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/mapper"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
	"go.uber.org/zap"
)

func (s *auditservice) List(ctx context.Context, req *dto.ListAuditLogRequest) (*dto.ListAuditLogResponse, error) {
	opts := []option.QueryOption{
		option.ApplyPagination(req.Pagination),
		option.WithSortBy(req.QuerySortBy),
		option.ApplyOperator(option.Condition{
			Field:    "type",
			Operator: option.IN,
			Value:    req.Types,
		}),
		option.WithPreloads("Fields"),
	}

	audits, err := s.audit.Find(ctx, &domain.AuditLog{
		UserID:         req.UserID,
		OrganizationID: req.OrganizationID,
		ResourceID:     req.ResouceID,
		Action:         req.Action,
	}, opts...)
	if err != nil {
		zap.L().Error("Failed to fetch list audit log", zap.Error(err))
		return nil, err
	}

	return &dto.ListAuditLogResponse{
		PageInfo: pagination.BuildCursorPageInfo(audits, req.Pagination.Limit, AuditLogCursorExtractor),
		Data:     mapper.ToDtoAuditList(audits),
	}, nil
}
