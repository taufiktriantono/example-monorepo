package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/event"
	"github.com/taufiktriantono/api-first-monorepo/pkg/message"
)

func Audit(producer message.Publisher, serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		eventID := uuid.NewString()
		evt := event.AuditEvent{
			TraceID:        c.GetString("trace_id"),
			Timestamp:      start,
			Method:         c.Request.Method,
			Path:           c.FullPath(),
			StatusCode:     c.Writer.Status(),
			UserID:         c.GetString("user_id"),
			OrganizationID: c.GetString("organization_id"),
			UserAgent:      c.Request.UserAgent(),
			ServiceName:    serviceName,
			ResponseTimeMs: time.Since(start).Milliseconds(),
		}
		_ = producer.Publish(c.Request.Context(), "audit-api-call", eventID, evt)
	}
}
