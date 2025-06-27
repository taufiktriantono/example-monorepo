package dto

import (
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
)

type ApprovalTemplateRequest struct {
	DisplayName  string `json:"display_name" validate:"required"`
	ResourceType string `json:"resource_type" validate:"required"`
	Status       string `json:"status" validate:"required"`
}

type ApprovalTemplateStepRequest struct {
	StepOrder     int    `json:"step_order" validate:"required,gte=1"`
	StepType      string `json:"step_type" validate:"required,oneof=manual auto notification"`
	ApproverType  string `json:"approver_type" validate:"required,oneof=user role"`
	ApproverValue string `json:"approver_value" validate:"required"`
	ConditionExpr string `json:"condition_expr"`
	SLAUnit       string `json:"sla_unit" validate:"required,oneof=minutes hours days"`
	SLAValue      int    `json:"sla_value" validate:"required,gte=1"`
}

type CreateTemplateRequest struct {
	ApprovalTemplateRequest
	Steps []ApprovalTemplateStepRequest `json:"steps" validate:"required,dive"`
}

type UpdateTemplateStep struct {
	ID string `json:"id" validate:"required"`
	ApprovalTemplateStepRequest
}

type UpdateTemplateRequest struct {
	ID string `json:"id" validate:"required"`
	ApprovalTemplateStepRequest
	Steps []UpdateTemplateStep `json:"steps" validate:"required,dive"`
}

type ListTemplatRequest struct {
	repository.Pagination
	repository.QueryStartAndEndDate
	repository.QuerySortBy
	repository.QueryRange
	ResourceTypes []string `form:"resource_types[]"` // Multiple resource types via query param
	Status        string   `form:"status"`
}

type ListApprovalRequest struct {
	repository.Pagination
	repository.QueryStartAndEndDate
	repository.QuerySortBy
	repository.QueryRange
	ResourceIDs   []string `form:"resource_ids[]"`
	ResourceTypes []string `form:"resource_types[]"`
	RequestedBy   string   `form:"requested_by"`
	Status        string   `form:"status"`
}
