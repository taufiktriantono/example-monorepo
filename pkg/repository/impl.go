package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type store[T any] struct {
	db *gorm.DB
}

func ProvideStore[T any](db *gorm.DB) Repository[T] {
	return &store[T]{db: db}
}

func (r *store[T]) WithTrx(tx *gorm.DB) Repository[T] {
	return &store[T]{db: tx}
}

func (r *store[T]) Find(ctx context.Context, query *T, opts ...QueryOption) ([]*T, error) {
	var result []*T
	stmt := r.buildQuery(ctx, query, opts...)
	err := stmt.Find(&result).Error
	return result, err
}

func (r *store[T]) FindOne(ctx context.Context, query *T) (*T, error) {
	var result T
	err := r.db.WithContext(ctx).Where(query).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, err
}

func (r *store[T]) Create(ctx context.Context, resource *T) error {
	return r.db.WithContext(ctx).Create(resource).Error
}

func (r *store[T]) Update(ctx context.Context, resourceID string, resource *T) error {
	return r.db.WithContext(ctx).Model(resource).Where("id = ?", resourceID).Updates(resource).Error
}

func (r *store[T]) Delete(ctx context.Context, resourceID string) error {
	var dummy T
	return r.db.WithContext(ctx).Where("id = ?", resourceID).Delete(&dummy).Error
}

func (r *store[T]) Count(ctx context.Context, query *T) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(query).Where(query).Count(&count).Error
	return count, err
}

func (r *store[T]) BatchCreate(ctx context.Context, resources []*T) error {
	return r.db.WithContext(ctx).Create(resources).Error
}

func (r *store[T]) BatchUpdate(ctx context.Context, query *T, resource *T) error {
	return r.db.WithContext(ctx).Where(query).Updates(resource).Error
}

func (s *store[T]) buildQuery(ctx context.Context, filter *T, opts ...QueryOption) *gorm.DB {
	db := s.db.WithContext(ctx).Where(filter)

	for _, opt := range opts {
		db = opt.Apply(db)
	}

	return db
}
