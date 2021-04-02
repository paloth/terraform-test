terraform {
  required_providers {
    aws = {
      version = ">= 3"
      source  = "hashicorp/aws"
    }
  }
  required_version = ">= 0.15"
}
