

-- Table: approval_templates
CREATE TABLE approval.templates (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  slug TEXT NOT NULL UNIQUE,
  display_name TEXT NOT NULL,
  resource_type TEXT NOT NULL,
  status approval_template_state NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table: approval_template_steps
CREATE TABLE approval.template_steps (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  approval_template_id UUID NOT NULL REFERENCES approval.templates(id) ON DELETE CASCADE,
  step_order INT NOT NULL,
  step_type step_type NOT NULL,
  approver_type approver_type NOT NULL,
  approver_value TEXT NOT NULL,
  condition_expr TEXT,
  sla_unit sla_unit NOT NULL,
  sla_value INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table: approvals
CREATE TABLE approval.approvals (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  approval_template_id UUID NOT NULL REFERENCES approval.templates(id) ON DELETE RESTRICT,
  resource_id TEXT NOT NULL,
  resource_type TEXT NOT NULL,
  requested_by TEXT NOT NULL,
  current_step_id UUID,
  completed_at TIMESTAMP,
  status approval_state NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table: approval_steps
CREATE TABLE approval.approval_steps (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  approval_id UUID NOT NULL REFERENCES approval.approvals(id) ON DELETE CASCADE,
  step_order INT NOT NULL,
  step_type step_type NOT NULL,
  assignee_id TEXT NOT NULL,
  condition_expr TEXT,
  sla_unit sla_unit NOT NULL,
  sla_value INT NOT NULL,
  comment TEXT,
  started_at TIMESTAMP NOT NULL,
  completed_at TIMESTAMP,
  sla_met BOOLEAN NOT NULL DEFAULT false,
  status approval_state NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);