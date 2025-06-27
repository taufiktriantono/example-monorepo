package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *approvalservice) syncTemplateSteps(ctx context.Context, tx *gorm.DB, templateID string, steps []dto.UpdateTemplateStep) error {
	templatesteprepo := s.approvaltemplatestep.WithTrx(tx)
	existingSteps, err := templatesteprepo.Find(ctx, &domain.ApprovalTemplateStep{
		ApprovalTemplateID: templateID,
	})
	if err != nil {
		return err
	}

	existingMap := make(map[string]*domain.ApprovalTemplateStep)
	for _, step := range existingSteps {
		existingMap[step.ID] = step
	}

	processed := make(map[string]bool)

	for _, step := range steps {
		if err := s.processTemplateStep(ctx, templatesteprepo, templateID, step, existingMap, processed); err != nil {
			return err
		}
	}

	return s.deleteUnprocessedSteps(ctx, templatesteprepo, existingSteps, processed)
}

func (s *approvalservice) processTemplateStep(
	ctx context.Context,
	repo repository.ApprovalTemplateStepRepository,
	templateID string,
	step dto.UpdateTemplateStep,
	existingMap map[string]*domain.ApprovalTemplateStep,
	processed map[string]bool,
) error {
	params := domain.ApprovalTemplateStepParams{
		ApprovalTemplateID: templateID,
		StepOrder:          step.StepOrder,
		StepType:           domain.StepType(step.StepType),
		ApproverType:       domain.ApproverType(step.ApproverType),
		ApproverVAlue:      step.ApproverValue,
		ConditionExpr:      step.ConditionExpr,
		SLAUnit:            domain.SLAUnit(step.SLAUnit),
		SLAValue:           step.SLAValue,
	}

	if step.ID == "" {
		return repo.Create(ctx, domain.NewTemplateStep(params))
	}

	existing, ok := existingMap[step.ID]
	if !ok {
		zap.L().Warn("Step not found for update", zap.String("id", step.ID))
		return nil
	}

	processed[step.ID] = true

	return repo.Update(ctx, existing.ID, &domain.ApprovalTemplateStep{
		StepOrder:     params.StepOrder,
		StepType:      params.StepType,
		ApproverType:  params.ApproverType,
		ApproverValue: params.ApproverVAlue,
		ConditionExpr: params.ConditionExpr,
		SLAUnit:       params.SLAUnit,
		SLAValue:      params.SLAValue,
	})
}

func (s *approvalservice) deleteUnprocessedSteps(
	ctx context.Context,
	repo repository.ApprovalTemplateStepRepository,
	steps []*domain.ApprovalTemplateStep,
	processed map[string]bool,
) error {
	for _, step := range steps {
		if !processed[step.ID] {
			if err := repo.Delete(ctx, step.ID); err != nil {
				return err
			}
		}
	}
	return nil
}
