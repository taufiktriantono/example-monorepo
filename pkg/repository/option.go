package repository

import (
	"gorm.io/gorm"
)

//go:generate mockgen -source=option.go -destination=mock_option.go -package=repository
type QueryOption interface {
	Apply(*gorm.DB) *gorm.DB
}

type applyQueryOption func(*gorm.DB) *gorm.DB

func (f applyQueryOption) Apply(db *gorm.DB) *gorm.DB {
	return f(db)
}

func WithResourceTypes(types []string) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		if len(types) > 0 {
			return db.Where("resource_type IN ?", types)
		}
		return db
	})
}

func WithResourceIDs(ids []string) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where("resource_id IN ?", ids)
		}
		return db
	})
}
