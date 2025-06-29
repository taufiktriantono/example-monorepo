package option

import (
	"fmt"
	"time"

	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockgen -source=option.go -destination=mock_option.go -package=option
type QueryOption interface {
	Apply(*gorm.DB) *gorm.DB
}

type applyQuery func(*gorm.DB) *gorm.DB

func (f applyQuery) Apply(db *gorm.DB) *gorm.DB {
	return f(db)
}

type Condition struct {
	Field    string
	Operator Operator
	Value    interface{}
}

// Operator representation the ope
type Operator string

const (
	// Comparison
	EQUAL     Operator = "="
	NOTEQUAL  Operator = "!="
	NOTEQUAL2 Operator = "<>"
	GT        Operator = ">"
	GTE       Operator = ">="
	LT        Operator = "<"
	LTE       Operator = "<="
	BETWEEN   Operator = "BETWEEN"
	IN        Operator = "IN"
	LIKE      Operator = "LIKE"
	NOTLIKE   Operator = "NOT LIKE"
	ILIKE     Operator = "ILIKE"

	// Logical
	AND Operator = "AND"
	OR  Operator = "OR"
	NOT Operator = "NOT"

	// Null checks
	ISNULL    Operator = "IS NULL"
	ISNOTNULL Operator = "IS NOT NULL"
	EXISTS    Operator = "EXISTS"
	NOTEXISTS Operator = "NOT EXISTS"

	// Arithmetic
	ADD Operator = "+"
	SUB Operator = "-"
	MUL Operator = "*"
	DIV Operator = "/"
	MOD Operator = "%"
)

func (o Operator) Valid() bool {
	switch o {
	case GT, GTE, LT, LTE, LIKE, EQUAL, NOTEQUAL, BETWEEN:
		return true
	default:
		return false
	}
}

func ApplyOperator(cond Condition) QueryOption {
	return applyQuery(func(db *gorm.DB) *gorm.DB {
		switch cond.Operator {
		case EQUAL, NOTEQUAL, GT, GTE, LT, LTE, LIKE, NOTLIKE, ILIKE:
			return db.Where(fmt.Sprintf("%s %s ?", cond.Field, cond.Operator), cond.Value)

		case IN:
			// pastikan value slice
			return db.Where(fmt.Sprintf("%s IN ?", cond.Field), cond.Value)

		case BETWEEN:
			vals, ok := cond.Value.([2]interface{})
			if !ok {
				zap.L().Warn("BETWEEN expects [2]interface{}", zap.Any("given", cond.Value))
				return db
			}
			return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", cond.Field), vals[0], vals[1])

		case ISNULL:
			return db.Where(fmt.Sprintf("%s IS NULL", cond.Field))

		case ISNOTNULL:
			return db.Where(fmt.Sprintf("%s IS NOT NULL", cond.Field))

		case EXISTS:
			subQuery, ok := cond.Value.(*gorm.DB)
			if !ok {
				zap.L().Warn("EXISTS expects *gorm.DB")
				return db
			}
			return db.Where("EXISTS (?)", subQuery)

		case NOTEXISTS:
			subQuery, ok := cond.Value.(*gorm.DB)
			if !ok {
				zap.L().Warn("NOT EXISTS expects *gorm.DB")
				return db
			}
			return db.Where("NOT EXISTS (?)", subQuery)

		default:
			zap.L().Warn("Unsupported operator", zap.String("operator", string(cond.Operator)))
			return db
		}
	})
}

func ApplyPagination(p pagination.Pagination) QueryOption {
	return applyQuery(func(db *gorm.DB) *gorm.DB {

		if p.Limit > 0 {
			db = db.Limit(p.Limit + 1)
		}

		if p.Cursor != "" {
			cursor, err := pagination.DecodeCursor(p.Cursor)
			if err == nil {
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
	return applyQuery(func(db *gorm.DB) *gorm.DB {
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
	return applyQuery(func(db *gorm.DB) *gorm.DB {
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
	return applyQuery(func(db *gorm.DB) *gorm.DB {
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

func WithSelect(selects []string) QueryOption {
	return applyQuery(func(db *gorm.DB) *gorm.DB {
		if len(selects) > 0 {
			return db.Select(selects)
		}
		return db
	})
}

func WithPreloads(values ...string) QueryOption {
	return applyQuery(func(db *gorm.DB) *gorm.DB {
		for _, v := range values {
			db = db.Preload(v)
		}
		return db
	})
}

func WithQuerySortBy(sortBy, orderBy string, allows map[string]bool) QuerySortBy {
	return QuerySortBy{
		SortBy:  sortBy,
		OrderBy: orderBy,
		Allow:   allows,
	}
}
