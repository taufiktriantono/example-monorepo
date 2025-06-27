package domain

import (
	"time"

	"github.com/google/uuid"
)

type ApprovalStep struct {
	ID            string        `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	ApprovalID    string        `gorm:"column:approval_id"`
	StepOrder     int           `gorm:"column:step_order"`
	StepType      string        `gorm:"column:step_type"`   // e.g., "manual", "auto"
	AssigneeID    string        `gorm:"column:assignee_id"` // User ID of the assignee
	ConditionExpr string        `gorm:"column:condition_expr"`
	SLAUnit       string        `gorm:"column:sla_unit"`
	SLAValue      int           `gorm:"column:sla_value"`
	Comment       string        `gorm:"column:comment"`
	StartedAt     time.Time     `gorm:"column:started_at"`
	CompletedAt   *time.Time    `gorm:"column:completed_at"`
	SLAMet        bool          `gorm:"column:sla_met"`
	Status        ApprovalState `gorm:"column:status"`
	CreatedAt     time.Time     `gorm:"column:created_at"`
	UpdatedAt     time.Time     `gorm:"column:updated_at"`
}

type ApprovalStepParams struct {
	TemplateStepID string
	ApprovalID     string
	StepOrder      int
	StepType       string
	AssigneeID     string
	ConditionExpr  string
	SLAUnit        string
	SLAValue       int
}

func NewApprovalStep(p ApprovalStepParams) *ApprovalStep {
	return &ApprovalStep{
		ID:            uuid.NewString(),
		ApprovalID:    p.ApprovalID,
		StepOrder:     p.StepOrder,
		StepType:      p.StepType,
		AssigneeID:    p.AssigneeID,
		ConditionExpr: p.ConditionExpr,
		SLAUnit:       p.SLAUnit,
		SLAValue:      p.SLAValue,
		StartedAt:     time.Now(),
		Status:        Pending,
	}
}

func (s *ApprovalStep) DueAt() time.Time {
	switch s.SLAUnit {
	case "minute":
		return s.StartedAt.Add(time.Duration(s.SLAValue) * time.Minute)
	case "hours":
		return s.StartedAt.Add(time.Duration(s.SLAValue) * time.Hour)
	default:
		return s.StartedAt.AddDate(0, 0, s.SLAValue)
	}
}

func (s *ApprovalStep) EvaluateSLAMet(now time.Time) bool {
	if s.CompletedAt == nil {
		return now.Before(s.DueAt())
	}
	return s.CompletedAt.Before(s.DueAt())
}
