package repository

import (
	"context"
	"errors"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -source=approval_repository.go -destination=mock_approval_repository.go -package=repository
type ApprovalRepository interface {
	WithTrx(tx *gorm.DB) ApprovalRepository
	Find(ctx context.Context, f *domain.Approval, opts ...option.QueryOption) ([]*domain.Approval, error)
	FindOne(ctx context.Context, f *domain.Approval) (*domain.Approval, error)
	Create(ctx context.Context, resource *domain.Approval) error
	Update(ctx context.Context, resourceID string, resource *domain.Approval) error
	Delete(ctx context.Context, resourceID string) error
	Count(ctx context.Context, f *domain.Approval) (int64, error)
}

type approval struct {
	db   *gorm.DB
	repo repository.Repository[domain.Approval]
}

func NewApprovalRepository(
	db *gorm.DB,
	repo repository.Repository[domain.Approval],
) ApprovalRepository {
	return &approval{
		db,
		repo,
	}
}

func (r *approval) WithTrx(tx *gorm.DB) ApprovalRepository {
	return &approval{
		db:   tx,
		repo: repository.ProvideStore[domain.Approval](tx),
	}
}

func (r *approval) Find(ctx context.Context, f *domain.Approval, opts ...option.QueryOption) ([]*domain.Approval, error) {
	return r.repo.Find(ctx, f, opts...)
}

func (r *approval) FindOne(ctx context.Context, f *domain.Approval) (*domain.Approval, error) {
	return r.repo.FindOne(ctx, f)
}

func (r *approval) Create(ctx context.Context, resource *domain.Approval) error {
	return r.repo.Create(ctx, resource)
}

func (r *approval) Update(ctx context.Context, resourceID string, resource *domain.Approval) error {
	return r.repo.Update(ctx, resourceID, resource)
}

func (r *approval) Delete(ctx context.Context, resourceID string) error {
	return errors.New("Unimplement")
}

func (r *approval) Count(ctx context.Context, f *domain.Approval) (int64, error) {
	return r.repo.Count(ctx, f)
}
