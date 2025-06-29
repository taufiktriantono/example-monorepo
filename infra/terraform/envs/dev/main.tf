terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "postgres" {
  name = "postgres:15"
}

resource "docker_container" "approval_db" {
  name  = "approval-db"
  image = docker_image.postgres.name

  ports {
    internal = 5432
    external = 5433 # gunakan 5433 agar tidak bentrok
  }

  env = [
    "POSTGRES_DB=approval_db",
    "POSTGRES_USER=postgres",
    "POSTGRES_PASSWORD=postgres"
  ]

  must_run = true
  restart  = "unless-stopped"
}
