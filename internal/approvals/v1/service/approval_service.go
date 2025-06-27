package service

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ApprovalService interface {
	CreateTemplate(context.Context, *dto.CreateTemplateRequest) (*dto.ApprovalTemplateResponse, error)
	ListTemplate(context.Context, *dto.ListTemplatRequest) (*dto.ListTemplateResponse, error)
	GetTemplateByID(context.Context, string) (*dto.ApprovalTemplateResponse, error)
}

type approvalservice struct {
	db                   *gorm.DB
	approvaltemplate     repository.ApprovalTemplateRepository
	approvaltemplatestep repository.ApprovalTemplateStepRepository
	approval             repository.ApprovalRepository
	approvalstep         repository.ApprovalStepRepository
}

type ApprovalParams struct {
	fx.In
	DB                             *gorm.DB
	ApprovalTemplateRepository     repository.ApprovalTemplateRepository
	ApprovalTemplateStepRepository repository.ApprovalTemplateStepRepository
	ApprovalRepository             repository.ApprovalRepository
	ApprovalStepRepository         repository.ApprovalStepRepository
}

func ProvideService(p ApprovalParams) ApprovalService {
	return &approvalservice{
		db:                   p.DB,
		approvaltemplate:     p.ApprovalTemplateRepository,
		approvaltemplatestep: p.ApprovalTemplateStepRepository,
		approval:             p.ApprovalRepository,
		approvalstep:         p.ApprovalStepRepository,
	}
}
