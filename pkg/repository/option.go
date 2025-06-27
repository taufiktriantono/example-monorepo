package repository

import (
	"time"

	"go.uber.org/zap"
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

func WithPagination(p Pagination) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {

		if p.Limit > 0 {
			db = db.Limit(p.Limit + 1)
		}

		if p.Cursor != "" {
			cursor, err := DecodeCursor(p.Cursor)
			if err == nil {
				// saya tidak tahu ini harus bagaimana
				db = db.Where("id < ?", cursor.ID).Where("created_at < ?", cursor.CreatedAt)
			} else {
				zap.L().Warn("Failed to decode cursor", zap.String("cursor", p.Cursor), zap.Error(err))
			}
		}

		return db
	})
}

type QueryStartAndEndDate struct {
	StartDate time.Time `form:"start_date"`
	EndDate   time.Time `form:"end_date"`
}

func WithStartAndEndDate(p QueryStartAndEndDate) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		if !p.StartDate.IsZero() {
			db = db.Where("created_at >= ?", p.StartDate.Format(time.RFC3339))
		}

		if !p.EndDate.IsZero() {
			db = db.Where("created_at <= ?", p.EndDate.Format(time.RFC3339))
		}
		return db
	})
}

type QueryRange struct {
	Ranges map[string][2]int64 `form:"ranges"`
	Allow  map[string]bool     `form:"-"`
}

func WithRange(p QueryRange) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		for field, val := range p.Ranges {
			if _, ok := p.Allow[field]; ok {
				db = db.Where(field+" BETWEEN ? AND ?", val[0], val[1])
			}
		}

		return db
	})
}

type QuerySortBy struct {
	SortBy  string          `form:"sort_by"`
	OrderBy string          `form:"order_by"`
	Allow   map[string]bool `form:"-"`
}

func WithSortBy(p QuerySortBy) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		sort := "created_at"
		if _, ok := p.Allow[p.SortBy]; ok {
			sort = p.SortBy
		}

		order := "ASC"
		if p.OrderBy == "desc" {
			order = "DESC"
		}

		return db.Order(sort + " " + order)
	})
}

func WithCustomField(field string, types []string) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		if len(types) > 0 {
			return db.Where("? IN ?", field, types)
		}
		return db
	})
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

func WithPreload(values []string) QueryOption {
	return applyQueryOption(func(db *gorm.DB) *gorm.DB {
		for _, v := range values {
			db = db.Preload(v)
		}
		return db
	})
}
