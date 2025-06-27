package dto

import (
	"time"

	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
)

type ApprovalTemplateResponse struct {
	ID           string                          `json:"id"`
	Slug         string                          `json:"slug"`
	DisplayName  string                          `json:"display_name"`
	ResourceType string                          `json:"resource_type"`
	Status       string                          `json:"status"`
	CreatedAt    time.Time                       `json:"created_at"`
	UpdatedAt    time.Time                       `json:"updated_at"`
	Steps        []*ApprovalTemplateStepResponse `json:"steps"`
}

type ListTemplateResponse struct {
	PageInfo repository.PageInfo         `json:"page_info"`
	Data     []*ApprovalTemplateResponse `json:"data"`
}

type ApprovalTemplateStepResponse struct {
	ID                 string    `json:"id"`
	ApprovalTemplateID string    `json:"approval_template_id"`
	StepOrder          int       `json:"step_order"`
	StepType           string    `json:"step_type"`
	ApproverType       string    `json:"approver_type"`
	ApproverValue      string    `json:"approver_value"`
	ConditionExpr      string    `json:"condition_expr"`
	SLAUnit            string    `json:"sla_unit"`
	SLAValue           int       `json:"sla_value"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
