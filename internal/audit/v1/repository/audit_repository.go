package repository

import (
	"context"

	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -source=audit_repository.go -destination=mock_audit_repository.go -package=repository
type AuditRepository interface {
	WithTrx(tx *gorm.DB) AuditRepository
	Find(context.Context, *domain.AuditLog, ...option.QueryOption) ([]*domain.AuditLog, error)
}

type audit struct {
	db   *gorm.DB
	repo repository.Repository[domain.AuditLog]
}

func NewAuditRepository(
	db *gorm.DB,
	repo repository.Repository[domain.AuditLog],
) AuditRepository {
	return &audit{
		db:   db,
		repo: repo,
	}
}

func (r *audit) WithTrx(tx *gorm.DB) AuditRepository {
	return &audit{
		db:   tx,
		repo: repository.ProvideStore[domain.AuditLog](tx),
	}
}

func (r *audit) Find(ctx context.Context, f *domain.AuditLog, opts ...option.QueryOption) ([]*domain.AuditLog, error) {
	return r.repo.Find(ctx, f, opts...)
}
