package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	WithTrx(tx *gorm.DB) Repository[T]
	Find(ctx context.Context, query *T, opts ...QueryOption) ([]*T, error)
	FindOne(ctx context.Context, query *T) (*T, error)
	Create(ctx context.Context, resource *T) error
	Update(ctx context.Context, resourceID string, resource *T) error
	Delete(ctx context.Context, resourceID string) error
	BatchCreate(ctx context.Context, resources []*T) error
	BatchUpdate(ctx context.Context, query *T, resources *T) error
	Count(ctx context.Context, query *T) (int64, error)
}
