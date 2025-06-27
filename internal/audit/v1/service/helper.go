package service

import (
	"time"

	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"go.uber.org/zap"
)

func AuditLogCursorExtractor(t *domain.AuditLog) string {
	nextcursor, err := repository.EncodeCursor(repository.Cursor{
		ID:        t.ID,
		CreatedAt: t.CreatedAt.Format(time.RFC3339Nano),
	})
	if err != nil {
		zap.L().Error("Failed to encoded cursor", zap.Error(err))
		return ""
	}

	return nextcursor
}
