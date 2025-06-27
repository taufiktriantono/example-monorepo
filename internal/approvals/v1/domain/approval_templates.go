package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/taufiktriantono/api-first-monorepo/pkg/errutil"
)

var (
	ErrTemplateAlreadyExists = errutil.BadRequest("TEMPLATE_ALREADY_EXISTS", nil)
	ErrTemplateNotFound      = errutil.BadRequest("TEMPLATE_NOT_FOUND", nil)
)

type ApprovalTemplateState string

var (
	Draft    ApprovalTemplateState = "draft"
	Active   ApprovalTemplateState = "active"
	Archived ApprovalTemplateState = "archived"
)

func (ats ApprovalTemplateState) String() string {
	switch ats {
	case Draft, Active, Archived:
		return string(ats)
	default:
		return ""
	}
}

func (ats ApprovalTemplateState) IsValid() bool {
	switch ats {
	case Draft, Active, Archived:
		return true
	default:
		return false
	}
}

type ApprovalTemplate struct {
	ID           string                `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Slug         string                `gorm:"column:slug"`
	DisplayName  string                `gorm:"column:display_name"`
	ResourceType string                `gorm:"column:resource_type"`
	Status       ApprovalTemplateState `gorm:"column:status"`
	CreatedAt    time.Time             `gorm:"column:created_at"`
	UpdatedAt    time.Time             `gorm:"column:updated_at"`
}

type ApprovalTemplateParams struct {
	DisplayName  string
	ResourceType string
	Steps        []ApprovalTemplateStep
	Status       ApprovalTemplateState
}

func NewTemplate(p ApprovalTemplateParams) *ApprovalTemplate {
	newslug := slug.Make(p.DisplayName)
	return &ApprovalTemplate{
		ID:           uuid.NewString(),
		Slug:         newslug,
		DisplayName:  p.DisplayName,
		ResourceType: p.ResourceType,
		Status:       p.Status,
	}
}
