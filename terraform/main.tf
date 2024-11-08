provider "aws" {
    region = var.aws_region
}

# #VPC input variables, 
module "vpc" {
    source = "./vpc"
   
  
}