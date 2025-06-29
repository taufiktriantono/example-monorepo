
output "connection_url" {
  value = "postgres://${var.username}:${var.password}@localhost:${var.port}/${var.database}?sslmode=disable"
}