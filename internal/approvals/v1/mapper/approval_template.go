package mapper

import (
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
)

func ToDtoApprovalTemplate(template *domain.ApprovalTemplate) *dto.ApprovalTemplateResponse {
	if template == nil {
		return nil
	}

	return &dto.ApprovalTemplateResponse{
		ID:           template.ID,
		Slug:         template.Slug,
		DisplayName:  template.DisplayName,
		ResourceType: template.ResourceType,
		Status:       template.Status.String(),
		CreatedAt:    template.CreatedAt,
		UpdatedAt:    template.UpdatedAt,
	}
}

func ToDtoApprovalTemplateStepList(steps []*domain.ApprovalTemplateStep) []*dto.ApprovalTemplateStepResponse {
	result := make([]*dto.ApprovalTemplateStepResponse, 0, len(steps))
	for _, step := range steps {
		result = append(result, ToDtoApprovalTemplateSteps(step))
	}
	return result
}

func ToDtoApprovalTemplateSteps(step *domain.ApprovalTemplateStep) *dto.ApprovalTemplateStepResponse {
	if step == nil {
		return nil
	}

	return &dto.ApprovalTemplateStepResponse{
		ID:                 step.ID,
		ApprovalTemplateID: step.ApprovalTemplateID,
		StepOrder:          step.StepOrder,
		StepType:           step.StepType.String(),
		ApproverType:       step.ApproverType.String(),
		ApproverValue:      step.ApproverValue,
		ConditionExpr:      step.ConditionExpr,
		SLAUnit:            step.SLAUnit.String(),
		SLAValue:           step.SLAValue,
		CreatedAt:          step.CreatedAt,
		UpdatedAt:          step.UpdatedAt,
	}
}
