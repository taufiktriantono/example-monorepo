package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/mapper"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
	"github.com/taufiktriantono/api-first-monorepo/pkg/errutil"
	"go.uber.org/zap"
)

func (s *approvalservice) ListTemplate(ctx context.Context, req *dto.ListTemplatRequest) (*dto.ListTemplateResponse, error) {
	f := &domain.ApprovalTemplate{
		Status: domain.ApprovalTemplateState(req.Status),
	}

	opts := []option.QueryOption{
		option.ApplyPagination(req.Pagination),
		option.WithRange(req.QueryRange),
		option.WithSortBy(NewTemplateQuerySortBy(req.SortBy, req.OrderBy)),
	}

	if len(req.ResourceTypes) > 0 {
		opts = append(opts, option.ApplyOperator(option.Condition{
			Field:    "resource_type",
			Operator: option.IN,
			Value:    req.ResourceTypes,
		}))
	}

	templates, err := s.approvaltemplate.Find(ctx, f, opts...)
	if err != nil {
		zap.L().Error("Failed to fetch list template", zap.Any("query", req), zap.Error(err))
		return nil, errutil.Internal(err.Error(), nil)
	}

	newtemplates := make([]*dto.ApprovalTemplateResponse, 0)
	for _, template := range templates {
		newtemplate := mapper.ToDtoApprovalTemplate(template)
		steps, err := s.approvaltemplatestep.Find(ctx, &domain.ApprovalTemplateStep{
			ApprovalTemplateID: template.ID,
		})
		if err != nil {
			zap.L().Error("Failed to fetch approval template steps", zap.Error(err))
			return nil, err
		}

		newtemplate.Steps = mapper.ToDtoApprovalTemplateStepList(steps)
		newtemplates = append(newtemplates, newtemplate)
	}

	return &dto.ListTemplateResponse{
		PageInfo: pagination.BuildCursorPageInfo(templates, req.Pagination.Limit, ApprovalTemplateCursorExtractor),
		Data:     newtemplates,
	}, nil
}
