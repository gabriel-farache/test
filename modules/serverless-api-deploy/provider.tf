terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
    google-beta = {
      source = "hashicorp/google-beta"
    }
    postgresql = {
      source  = "cyrilgdn/postgresql"
      version = "1.16.0"
    }
  }
}


provider "postgresql" {
  scheme   = "gcppostgres"
  host     = module.pgsql-singlezone.instance_connection_name
  username = module.pgsql-singlezone.user_name
  password = module.pgsql-singlezone.user_password
}