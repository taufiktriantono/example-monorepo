package service

import (
	"time"

	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db/pagination"
	"go.uber.org/zap"
)

func ApprovalTemplateCursorExtractor(t *domain.ApprovalTemplate) string {
	nextcursor, err := pagination.EncodeCursor(pagination.Cursor{
		ID:        t.ID,
		CreatedAt: t.CreatedAt.Format(time.RFC3339Nano),
	})
	if err != nil {
		zap.L().Error("Failed to encoded cursor", zap.Error(err))
		return ""
	}

	return nextcursor
}
