resource "docker_image" "postgres" {
  name = var.image
}

resource "docker_container" "postgres" {
  name  = var.name
  image = docker_image.postgres.name

  ports {
    internal = 5432
    external = var.port
  }

  env = [
    "POSTGRES_DB=${var.database}",
    "POSTGRES_USER=${var.username}",
    "POSTGRES_PASSWORD=${var.password}"
  ]

  must_run = true
  restart  = "unless-stopped"
}
