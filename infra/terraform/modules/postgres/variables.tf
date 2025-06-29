variable "name" {
  type        = string
  description = "Container name"
}

variable "image" {
  type        = string
  default     = "postgres:15"
  description = "PostgreSQL Docker image"
}

variable "port" {
  type        = number
  default     = 5432
  description = "External port"
}

variable "database" {
  type        = string
  default     = "app_db"
  description = "Database name"
}

variable "username" {
  type        = string
  default     = "postgres"
}

variable "password" {
  type        = string
  default     = "postgres"
}