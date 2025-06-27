package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *approvalservice) UpdateTemplate(ctx context.Context, req *dto.UpdateTemplateRequest) (*dto.ApprovalTemplateResponse, error) {
	if err := s.db.Transaction(func(tx *gorm.DB) error {

		templaterepo := s.approvaltemplate.WithTrx(tx)
		existing, err := templaterepo.FindOne(ctx, &domain.ApprovalTemplate{
			ID: req.ID,
		})
		if err != nil {
			return err
		}
		if existing == nil {
			return domain.ErrTemplateNotFound
		}

		if len(req.Steps) > 0 {
			if err := s.syncTemplateSteps(ctx, tx, req.ID, req.Steps); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		zap.L().Error("Failed to update approval template", zap.Error(err))
		return nil, err
	}

	return s.getTemplateByID(ctx, req.ID)
}
