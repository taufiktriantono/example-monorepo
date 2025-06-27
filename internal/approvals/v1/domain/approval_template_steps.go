package domain

import (
	"time"

	"github.com/google/uuid"
)

type ApproverType string

var (
	User ApproverType = "user"
	Role ApproverType = "role"
)

func (at ApproverType) String() string {
	switch at {
	case Role, User:
		return at.String()
	default:
		return ""
	}
}

func (at ApproverType) IsValid() bool {
	switch at {
	case Role, User:
		return true
	default:
		return false
	}
}

type StepType string

var (
	Manual       StepType = "manual"
	Notification StepType = "notification"
	Auto         StepType = "auto"
)

func (st StepType) String() string {
	switch st {
	case Manual, Auto, Notification:
		return string(st)
	default:
		return ""
	}
}

func (st StepType) IsValid() bool {
	switch st {
	case Manual, Auto, Notification:
		return true
	default:
		return false
	}
}

type SLAUnit string

var (
	Minute SLAUnit = "minutes"
	Hours  SLAUnit = "hours"
	Days   SLAUnit = "days"
)

func (su SLAUnit) String() string {
	switch su {
	case Minute, Hours, Days:
		return string(su)
	default:
		return ""
	}
}

func (su SLAUnit) IsValid() bool {
	switch su {
	case Minute, Hours, Days:
		return true
	default:
		return false
	}
}

type ApprovalTemplateStep struct {
	ID                 string           `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	ApprovalTemplateID string           `gorm:"column:approval_template_id"`
	ApprovalTemplate   ApprovalTemplate `gorm:"foreignKey:ApprovalTemplateID"`
	StepOrder          int              `gorm:"column:step_order"`
	StepType           StepType         `gorm:"column:step_type"`
	ApproverType       ApproverType     `gorm:"column:approver_type"`
	ApproverValue      string           `gorm:"column:approver_value"`
	ConditionExpr      string           `gorm:"column:condition_expr"`
	SLAUnit            SLAUnit          `gorm:"column:sla_unit"`
	SLAValue           int              `gorm:"column:sla_value"`
	CreatedAt          time.Time        `gorm:"column:created_at"`
	UpdatedAt          time.Time        `gorm:"column:updated_at"`
}

type ApprovalTemplateStepParams struct {
	ApprovalTemplateID string
	StepOrder          int
	StepType           StepType
	ApproverType       ApproverType
	ApproverVAlue      string
	ConditionExpr      string
	SLAUnit            SLAUnit
	SLAValue           int
}

func NewTemplateStep(p ApprovalTemplateStepParams) *ApprovalTemplateStep {
	return &ApprovalTemplateStep{
		ID:                 uuid.NewString(),
		ApprovalTemplateID: p.ApprovalTemplateID,
		StepOrder:          p.StepOrder,
		StepType:           p.StepType,
		ApproverType:       p.ApproverType,
		ApproverValue:      p.ApproverVAlue,
		ConditionExpr:      p.ConditionExpr,
		SLAUnit:            p.SLAUnit,
		SLAValue:           p.SLAValue,
	}
}

func (m *ApprovalTemplateStep) DueAt() time.Time {
	switch m.SLAUnit {
	case "minutes":
		return time.Now().Add(time.Duration(m.SLAValue) * time.Minute)
	case "hours":
		return time.Now().Add(time.Duration(m.SLAValue) * time.Hour)
	default:
		return time.Now().AddDate(0, 0, m.SLAValue)
	}
}

func (m *ApprovalTemplateStep) IsValid() bool {
	if !m.StepType.IsValid() {
		return m.StepType.IsValid()
	}

	if !m.ApproverType.IsValid() {
		return m.ApproverType.IsValid()
	}

	if !m.SLAUnit.IsValid() {
		return m.SLAUnit.IsValid()
	}

	if m.SLAValue < 1 {
		return false
	}

	return true
}
