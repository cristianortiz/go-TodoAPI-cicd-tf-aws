provider "aws" {
    region = var.aws_region
}

#VPC input vars
module "vpc" {
    source = "terraform-aws-modules/vpc/aws"
    name = var.vpc_name
  
}