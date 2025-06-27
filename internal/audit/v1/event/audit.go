package event

import "time"

type AuditEvent struct {
	TraceID        string    `json:"trace_id"`
	Timestamp      time.Time `json:"timestamp"`
	Method         string    `json:"method"`
	Path           string    `json:"path"`
	StatusCode     int       `json:"status_code"`
	UserID         string    `json:"user_id"`
	OrganizationID string    `json:"organization_id"`
	RemoteIP       string    `json:"remote_ip"`
	UserAgent      string    `json:"user_agent"`
	ServiceName    string    `json:"service_name"`
	ResponseTimeMs int64     `json:"response_time_ms"`
}
