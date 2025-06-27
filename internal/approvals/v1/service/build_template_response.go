package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/mapper"
	"github.com/taufiktriantono/api-first-monorepo/pkg/logfields"
	"go.uber.org/zap"
)

func (s *approvalservice) getTemplateByID(ctx context.Context, templateID string) (*dto.ApprovalTemplateResponse, error) {
	templateIDField := zap.String(logfields.TemplateID, templateID)

	template, err := s.approvaltemplate.FindOne(ctx, &domain.ApprovalTemplate{ID: templateID})
	if err != nil {
		zap.L().Error("Failed to fetch created approval template", templateIDField, zap.Error(err))
		return nil, err
	}

	templateDto := mapper.ToDtoApprovalTemplate(template)

	steps, err := s.approvaltemplatestep.Find(ctx, &domain.ApprovalTemplateStep{ApprovalTemplateID: templateID})
	if err != nil {
		zap.L().Error("Failed to fetch created approval template steps", templateIDField, zap.Error(err))
		return nil, err
	}

	templateDto.Steps = mapper.ToDtoApprovalTemplateStepList(steps)

	return templateDto, nil
}
