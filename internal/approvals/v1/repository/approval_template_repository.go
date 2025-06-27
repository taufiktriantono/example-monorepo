package repository

import (
	"context"
	"errors"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -source=approval_template_repository.go -destination=mock_approval_template_repository.go -package=repository
type ApprovalTemplateRepository interface {
	WithTrx(tx *gorm.DB) ApprovalTemplateRepository
	Find(ctx context.Context, f *domain.ApprovalTemplate, opts ...repository.QueryOption) ([]*domain.ApprovalTemplate, error)
	FindOne(ctx context.Context, f *domain.ApprovalTemplate) (*domain.ApprovalTemplate, error)
	Create(ctx context.Context, resource *domain.ApprovalTemplate) error
	Update(ctx context.Context, resourceID string, resource *domain.ApprovalTemplate) error
	Delete(ctx context.Context, resourceID string) error
	Count(ctx context.Context, f *domain.ApprovalTemplate) (int64, error)
}

type approvaltemplate struct {
	db   *gorm.DB
	repo repository.Repository[domain.ApprovalTemplate]
}

func NewApprovalTemplateRepository(
	db *gorm.DB,
	repo repository.Repository[domain.ApprovalTemplate],
) ApprovalTemplateRepository {
	return &approvaltemplate{
		db:   db,
		repo: repo,
	}
}

func (r *approvaltemplate) WithTrx(tx *gorm.DB) ApprovalTemplateRepository {
	return &approvaltemplate{
		db:   tx,
		repo: repository.ProvideStore[domain.ApprovalTemplate](tx),
	}
}

func (r *approvaltemplate) Find(ctx context.Context, f *domain.ApprovalTemplate, opts ...repository.QueryOption) ([]*domain.ApprovalTemplate, error) {
	return r.repo.Find(ctx, f, opts...)
}

func (r *approvaltemplate) FindOne(ctx context.Context, f *domain.ApprovalTemplate) (*domain.ApprovalTemplate, error) {
	return r.repo.FindOne(ctx, f)
}

func (r *approvaltemplate) Create(ctx context.Context, resource *domain.ApprovalTemplate) error {
	return r.repo.Create(ctx, resource)
}

func (r *approvaltemplate) Update(ctx context.Context, resourceID string, resource *domain.ApprovalTemplate) error {
	return r.repo.Update(ctx, resourceID, resource)
}

func (r *approvaltemplate) Delete(ctx context.Context, resourceID string) error {
	return errors.New("Unimplement")
}

func (r *approvaltemplate) Count(ctx context.Context, f *domain.ApprovalTemplate) (int64, error) {
	return r.repo.Count(ctx, f)
}
