terraform {
  required_version = ">= 1.3.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.0.0"
    }
  }
  backend "s3" {
    bucket = "certificate-generator-76534819"
    key    = "terraform/infra.tfstate"
    region = "sa-east-1"
  }
}

provider "aws" {
  region              = var.aws_region
  access_key          = var.access_key
  secret_key          = var.secret_key
  allowed_account_ids = [var.aws_account_number]
}