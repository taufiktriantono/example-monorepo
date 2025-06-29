package repository

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -source=approval_step_repository.go -destination=mock_approval_step_repository.go -package=repository
type ApprovalStepRepository interface {
	WithTrx(tx *gorm.DB) ApprovalStepRepository
	Find(ctx context.Context, f *domain.ApprovalStep, opts ...option.QueryOption) ([]*domain.ApprovalStep, error)
	FindOne(ctx context.Context, f *domain.ApprovalStep) (*domain.ApprovalStep, error)
	Create(ctx context.Context, resource *domain.ApprovalStep) error
	Update(ctx context.Context, resourceID string, resource *domain.ApprovalStep) error
	Delete(ctx context.Context, resourceID string) error
	Count(ctx context.Context, f *domain.ApprovalStep) (int64, error)
}

type approvalstep struct {
	db   *gorm.DB
	repo repository.Repository[domain.ApprovalStep]
}

func NewApprovalStepRepository(
	db *gorm.DB,
	repo repository.Repository[domain.ApprovalStep],
) ApprovalStepRepository {
	return &approvalstep{
		db:   db,
		repo: repo,
	}
}

func (r *approvalstep) WithTrx(tx *gorm.DB) ApprovalStepRepository {
	return &approvalstep{
		db:   tx,
		repo: repository.ProvideStore[domain.ApprovalStep](tx),
	}
}

func (r *approvalstep) Find(ctx context.Context, f *domain.ApprovalStep, opts ...option.QueryOption) ([]*domain.ApprovalStep, error) {
	return r.repo.Find(ctx, f, opts...)
}

func (r *approvalstep) FindOne(ctx context.Context, f *domain.ApprovalStep) (*domain.ApprovalStep, error) {
	return r.repo.FindOne(ctx, f)
}

func (r *approvalstep) Create(ctx context.Context, resource *domain.ApprovalStep) error {
	return r.repo.Create(ctx, resource)
}

func (r *approvalstep) Update(ctx context.Context, resourceID string, resource *domain.ApprovalStep) error {
	return r.repo.Update(ctx, resourceID, resource)
}

func (r *approvalstep) Delete(ctx context.Context, resourceID string) error {
	return r.repo.Delete(ctx, resourceID)
}

func (r *approvalstep) Count(ctx context.Context, f *domain.ApprovalStep) (int64, error) {
	return r.repo.Count(ctx, f)
}
