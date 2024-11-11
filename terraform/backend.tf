terraform {
  backend "s3" {
    bucket         = "todoapi-tfbackend-bkt"
    key            = "develop/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "todoapiadmin-tf-statelocking"
    encrypt        = true
  }
}
