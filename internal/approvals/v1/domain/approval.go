package domain

import (
	"time"

	"github.com/google/uuid"
)

type ApprovalState string

const (
	Pending    ApprovalState = "pending"
	InProgress ApprovalState = "in_progress"
	Approved   ApprovalState = "approved"
	Completed  ApprovalState = "completed"
	Rejected   ApprovalState = "rejected"
	Expired    ApprovalState = "expired"
	Cancelled  ApprovalState = "cancelled"
)

func (ap ApprovalState) String() string {
	switch ap {
	case Pending, InProgress, Approved, Completed, Rejected, Expired, Cancelled:
		return string(ap)
	default:
		return ""
	}
}

func (ap ApprovalState) IsValid() bool {
	switch ap {
	case Pending, InProgress, Completed, Rejected, Expired, Cancelled:
		return true
	default:
		return false
	}
}

type Approval struct {
	ID                 string           `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	ApprovalTemplateID string           `gorm:"column:approval_template_id"`
	ApprovalTemplate   ApprovalTemplate `gorm:"->;foreignKey:ApprovalTemplateID"`
	ResourceID         string           `gorm:"column:resource_id"`
	ResourceType       string           `gorm:"column:resource_type"`
	RequestedBy        string           `gorm:"column:requested_by"`
	CurrentStepID      string           `gorm:"column:current_step_id"`
	CompletedAt        *time.Time       `gorm:"column:completed_at"`
	Status             ApprovalState    `gorm:"column:status"`
	CreatedAt          time.Time        `gorm:"column:created_at"`
	UpdatedAt          time.Time        `gorm:"column:updated_at"`
}

type ApprovalParams struct {
	ApprovalTemplateID string
	ResourceID         string
	ResourceType       string
	RequestedBy        string
}

func NewApproval(p ApprovalParams) *Approval {
	return &Approval{
		ID:                 uuid.NewString(),
		ApprovalTemplateID: p.ApprovalTemplateID,
		ResourceID:         p.ResourceID,
		ResourceType:       p.ResourceType,
		RequestedBy:        p.RequestedBy,
		Status:             Pending,
	}
}

func (m *Approval) SetCurrentStepID(stepID string) {
	m.CurrentStepID = stepID
}

func (m *Approval) SetCompleted() {
	t := time.Now()
	m.CompletedAt = &t
	m.Status = Completed
}

func (m *Approval) SetCancelled() {
	t := time.Now()
	m.CompletedAt = &t
	m.Status = Cancelled
}
