package service

import (
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type AuditService interface{}

type auditservice struct {
	db    *gorm.DB
	audit repository.AuditRepository
}

type AuditParams struct {
	fx.In
	DB    *gorm.DB
	audit repository.AuditRepository
}

func ProvideService(p AuditParams) AuditService {
	return &auditservice{
		db:    p.DB,
		audit: p.audit,
	}
}
