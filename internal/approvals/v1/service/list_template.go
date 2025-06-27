package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/mapper"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"go.uber.org/zap"
)

func (s *approvalservice) ListTemplate(ctx context.Context, req *dto.ListTemplatRequest) (*dto.ListTemplateResponse, error) {
	f := &domain.ApprovalTemplate{
		Status: domain.ApprovalTemplateState(req.Status),
	}

	queryopts := []repository.QueryOption{
		repository.WithPagination(req.Pagination),
		repository.WithStartAndEndDate(req.QueryStartAndEndDate),
		repository.WithSortBy(NewTemplateQuerySortBy(req.SortBy, req.OrderBy)),
		repository.WithRange(req.QueryRange),
		repository.WithResourceTypes(req.ResourceTypes),
	}

	templates, err := s.approvaltemplate.Find(ctx, f, queryopts...)
	if err != nil {
		zap.L().Error("Failed to fetch list template", zap.Any("query", req), zap.Error(err))
		return nil, err
	}

	newtemplates := make([]*dto.ApprovalTemplateResponse, 0)
	for _, template := range templates {
		newtemplate := &dto.ApprovalTemplateResponse{
			ID:          template.ID,
			Slug:        template.Slug,
			DisplayName: template.DisplayName,
			Status:      template.Status.String(),
			CreatedAt:   template.CreatedAt,
			UpdatedAt:   template.UpdatedAt,
		}

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
		PageInfo: repository.BuildCursorPageInfo(templates, req.Pagination.Limit, ApprovalTemplateCursorExtractor),
		Data:     newtemplates,
	}, nil
}
