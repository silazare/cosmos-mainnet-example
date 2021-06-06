terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.0"
    }
    http = {
      source  = "hashicorp/http"
      version = "~> 2.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 2.2.1"
    }
  }
  required_version = ">= 0.15"
}
