package organization

import "time"

type OrganizationStatus string

var (
	OrganizationStatusActive    OrganizationStatus = "active"
	OrganizationStatusSuspended OrganizationStatus = "suspended"
)

func (s OrganizationStatus) String() string {
	switch s {
	case OrganizationStatusActive, OrganizationStatusSuspended:
		return string(s)
	default:
		return ""
	}
}

func (s OrganizationStatus) IsValid() bool {
	switch s {
	case OrganizationStatusActive, OrganizationStatusSuspended:
		return true
	default:
		return false
	}
}

type Organization struct {
	ID          string             `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	GroupID     *string            `gorm:"column:group_id"`
	Slug        string             `gorm:"column:slug"`
	DisplayName string             `gorm:"column:display_name"`
	Status      OrganizationStatus `gorm:"column:status"`
	CreatedAt   time.Time          `gorm:"column:created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at"`
}

func (m *Organization) IsValid() bool {
	if !m.Status.IsValid() {
		return m.Status.IsValid()
	}

	if m.Slug == "" {
		return false
	}

	if m.DisplayName == "" {
		return false
	}

	return true
}
