package logfields

// General request-level fields
const (
	RequestID = "request.id"
	TraceID   = "trace.id"
	UserID    = "user.id"
)

// Template-specific fields
const (
	TemplateID   = "template.id"
	TemplateSlug = "template.slug"
	DisplayName  = "template.display_name"
	ResourceType = "template.resource_type"
	Status       = "template.status"
)

// Approval Step fields
const (
	StepID        = "step.id"
	StepOrder     = "step.order"
	StepType      = "step.type"
	ApproverType  = "step.approver_type"
	ApproverValue = "step.approver_value"
	SLAUnit       = "step.sla_unit"
	SLAValue      = "step.sla_value"
)
