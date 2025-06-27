package mapper

import (
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/dto"
)

func ToDtoAudit(audit *domain.AuditLog) *dto.AuditLogResponse {
	if audit == nil {
		return nil
	}

	this := &dto.AuditLogResponse{
		ID:             audit.ID,
		OrganizationID: audit.OrganizationID,
		UserID:         audit.UserID,
		Type:           string(audit.Type),
		ResourceID:     audit.ResourceID,
		ResourceName:   audit.ResourceName,
		Action:         audit.Action,
	}

	if len(audit.Fields) > 0 {
		this.Fields = ToDtoAuditFieldList(audit.Fields)
	}

	return this
}

func ToDtoAuditList(audits []*domain.AuditLog) []*dto.AuditLogResponse {
	result := make([]*dto.AuditLogResponse, 0, len(audits))
	for _, audit := range audits {
		result = append(result, ToDtoAudit(audit))
	}
	return result
}

func ToDtoAuditFieldList(fields []*domain.AuditLogFieldValue) []*dto.AuditLogFieldResponse {
	result := make([]*dto.AuditLogFieldResponse, 0, len(fields))
	for _, field := range fields {
		result = append(result, ToDtoAuditField(field))
	}
	return result
}

func ToDtoAuditField(field *domain.AuditLogFieldValue) *dto.AuditLogFieldResponse {
	if field == nil {
		return nil
	}

	return &dto.AuditLogFieldResponse{
		ID:            field.ID,
		AuditLogID:    field.AuditLogID,
		Field:         field.Field,
		PreviousValue: field.PreviousValue,
		NewValue:      field.NewValue,
	}
}
