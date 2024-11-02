terraform {
  backend "s3" {
    bucket = "todoapi-tfbackend-bkt"
    key = "vpc/terraform.tfstate"
    region = "us-east-1"
    
  }
}