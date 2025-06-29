
CREATE TYPE approval_template_state AS ENUM ('draf','active', 'inactive','archived');
CREATE TYPE approval_state AS ENUM ('pending', 'approved', 'rejected');
CREATE TYPE step_type AS ENUM ('manual', 'auto');
CREATE TYPE approver_type AS ENUM ('user', 'role');
CREATE TYPE sla_unit AS ENUM ('minute', 'hour', 'day');