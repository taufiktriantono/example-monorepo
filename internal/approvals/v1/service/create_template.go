package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *approvalservice) CreateTemplate(ctx context.Context, req *dto.CreateTemplateRequest) (*dto.ApprovalTemplateResponse, error) {

	newTemplate := domain.NewTemplate(domain.ApprovalTemplateParams{
		DisplayName:  req.DisplayName,
		ResourceType: req.ResourceType,
		Status:       domain.ApprovalTemplateState(req.Status),
	})

	if err := s.db.Transaction(func(tx *gorm.DB) error {

		if err := s.processBuildTemplate(ctx, tx, newTemplate); err != nil {
			return err
		}

		if err := s.processBuildTemplateStep(ctx, tx, newTemplate.ID, req.Steps); err != nil {
			zap.L().Error("Failed to create approval template steps", zap.Error(err))
			return err
		}

		return nil
	}); err != nil {
		zap.L().Error("Failed to create approval template and steps", zap.Error(err))
		return nil, err
	}

	return s.getTemplateByID(ctx, newTemplate.ID)

}

func (s *approvalservice) processBuildTemplateStep(ctx context.Context, tx *gorm.DB, templateID string, steps []dto.ApprovalTemplateStepRequest) error {
	if len(steps) == 0 {
		zap.L().Warn("Approval template steps is empty")
		return nil
	}

	newsteps := make([]*domain.ApprovalTemplateStep, 0)
	for _, step := range steps {
		newsteps = append(newsteps, domain.NewTemplateStep(domain.ApprovalTemplateStepParams{
			ApprovalTemplateID: templateID,
			StepOrder:          step.StepOrder,
			StepType:           domain.StepType(step.StepType),
			ApproverType:       domain.ApproverType(step.ApproverType),
			ApproverVAlue:      step.ApproverValue,
			ConditionExpr:      step.ConditionExpr,
			SLAUnit:            domain.SLAUnit(step.SLAUnit),
			SLAValue:           step.SLAValue,
		}))
	}

	if err := s.approvaltemplatestep.WithTrx(tx).BatchCreate(ctx, newsteps); err != nil {
		return err
	}

	return nil
}

func (s *approvalservice) processBuildTemplate(ctx context.Context, tx *gorm.DB, template *domain.ApprovalTemplate) error {
	templaterepo := s.approvaltemplate.WithTrx(tx)
	exist, err := templaterepo.FindOne(ctx, &domain.ApprovalTemplate{
		Slug: template.Slug,
	})
	if err != nil {
		return err
	}

	if exist != nil {
		return domain.ErrTemplateAlreadyExists
	}

	if err := templaterepo.Create(ctx, template); err != nil {
		zap.L().Error("Failed to create approval template", zap.Error(err))
		return err
	}

	return nil
}
