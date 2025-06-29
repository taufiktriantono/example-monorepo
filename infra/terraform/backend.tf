terraform {
  backend "s3" {
    endpoint                    = "http://localhost:9000"         # Ganti dengan host MinIO kamu
    bucket                      = "terraform-state"               # Pastikan sudah dibuat di MinIO
    key                         = "services/approval/dev/terraform.tfstate"
    region                      = "us-east-1"                     # Wajib isi walau dummy
    access_key                  = "minioadmin"
    secret_key                  = "minioadmin"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    force_path_style            = true
  }
}
