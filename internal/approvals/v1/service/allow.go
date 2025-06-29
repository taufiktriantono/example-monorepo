package service

import "github.com/taufiktriantono/api-first-monorepo/pkg/db/option"

// templateSortAllowFields specifies the allowed fields for sorting templates.
var templateSortAllowFields = map[string]bool{
	"created_at":   true,
	"display_name": true,
	"status":       true,
}

// templateStepRangeAllowFields specifies the allowed fields for range filtering on template steps.
var templateStepRangeAllowFields = map[string]bool{
	"sla_value": true,
}

// approvalAllowFields specifies the allowed fields for sorting approvals.
var approvalAllowFields = map[string]bool{
	"created_at": true,
	"status":     true,
}

// approvalStepRangeAllowFields specifies the allowed fields for range filtering on approval steps.
var approvalStepRangeAllowFields = map[string]bool{
	"sla_value":  true,
	"step_order": true,
}

func NewTemplateQuerySortBy(sortBy, orderBy string) option.QuerySortBy {
	return option.QuerySortBy{
		SortBy:  sortBy,
		OrderBy: orderBy,
		Allow:   templateSortAllowFields,
	}
}
