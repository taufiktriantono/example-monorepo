package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
)

func (s *approvalservice) GetTemplateByID(ctx context.Context, resourceID string) (*dto.ApprovalTemplateResponse, error) {
	filter := &domain.ApprovalTemplate{
		ID: resourceID,
	}

	if _, err := uuid.Parse(resourceID); err != nil {
		filter.Slug = resourceID
	}

	return s.getTemplateByID(ctx, resourceID)
}
