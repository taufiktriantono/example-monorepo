package repository

import (
	"context"
	"errors"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -source=approval_template_step_repository.go -destination=mock_approval_template_step_repository.go -package=repository
type ApprovalTemplateStepRepository interface {
	WithTrx(tx *gorm.DB) ApprovalTemplateStepRepository
	Find(ctx context.Context, f *domain.ApprovalTemplateStep, opts ...option.QueryOption) ([]*domain.ApprovalTemplateStep, error)
	FindOne(ctx context.Context, f *domain.ApprovalTemplateStep) (*domain.ApprovalTemplateStep, error)
	Create(ctx context.Context, resource *domain.ApprovalTemplateStep) error
	Update(ctx context.Context, resourceID string, resource *domain.ApprovalTemplateStep) error
	Delete(ctx context.Context, resourceID string) error
	Count(ctx context.Context, f *domain.ApprovalTemplateStep) (int64, error)
	BatchCreate(ctx context.Context, resources []*domain.ApprovalTemplateStep) error
}

type approvaltemplatestep struct {
	db   *gorm.DB
	repo repository.Repository[domain.ApprovalTemplateStep]
}

func NewApprovalTemplateStepRepository(
	db *gorm.DB,
	repo repository.Repository[domain.ApprovalTemplateStep],
) ApprovalTemplateStepRepository {
	return &approvaltemplatestep{
		db:   db,
		repo: repo,
	}
}

func (r *approvaltemplatestep) WithTrx(tx *gorm.DB) ApprovalTemplateStepRepository {
	return &approvaltemplatestep{
		db:   tx,
		repo: repository.ProvideStore[domain.ApprovalTemplateStep](tx),
	}
}

func (r *approvaltemplatestep) Find(ctx context.Context, f *domain.ApprovalTemplateStep, opts ...option.QueryOption) ([]*domain.ApprovalTemplateStep, error) {
	return r.repo.Find(ctx, f, opts...)
}

func (r *approvaltemplatestep) FindOne(ctx context.Context, f *domain.ApprovalTemplateStep) (*domain.ApprovalTemplateStep, error) {
	return r.repo.FindOne(ctx, f)
}

func (r *approvaltemplatestep) Create(ctx context.Context, resource *domain.ApprovalTemplateStep) error {
	return r.repo.Create(ctx, resource)
}

func (r *approvaltemplatestep) Update(ctx context.Context, resourceID string, resource *domain.ApprovalTemplateStep) error {
	return r.repo.Update(ctx, resourceID, resource)
}

func (r *approvaltemplatestep) Delete(ctx context.Context, resourceID string) error {
	return errors.New("Unimplement")
}

func (r *approvaltemplatestep) Count(ctx context.Context, f *domain.ApprovalTemplateStep) (int64, error) {
	return r.repo.Count(ctx, f)
}

func (r *approvaltemplatestep) BatchCreate(ctx context.Context, resources []*domain.ApprovalTemplateStep) error {
	return r.repo.BatchCreate(ctx, resources)
}
