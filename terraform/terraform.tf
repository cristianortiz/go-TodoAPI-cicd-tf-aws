terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
  }
  backend "s3" {
    bucket = "todoapi-tfbackend-bkt"
    key = "develop/terraform.tfstate"
    region = "us-east-1"
    dynamodb_table="tf-statelocking"
    encrypt = true
    
  }
}