package repository

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Pagination struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit,default=10" validate:"gte=1,lte=250"` // Min 1, Max 250
}

type Cursor struct {
	CreatedAt string `json:"created_at,omitempty"`
	ID        string `json:"id,omitempty"`
}

type PageInfo struct {
	NextCursor     string `json:"next_cursor"`
	PreviousCursor string `json:"previous_cursor"`
	HasMore        bool   `json:"has_more"`
}

func EncodeCursor(data Cursor) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func DecodeCursor(data string) (*Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var cursor Cursor
	if err := json.Unmarshal(b, &cursor); err != nil {
		return nil, err
	}

	return &cursor, nil
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

func BuildCursorPageInfo[T any](data []*T, limit int, extractCursor func(*T) string) (pageInfo PageInfo) {
	if len(data) > limit {
		pageInfo.HasMore = true
		pageInfo.NextCursor = extractCursor(data[limit-1])
	}
	return
}
