variable "aws_region"{
    description = "AWS region"
    type = string
    default = "us-east-1"
}

variable "vpc_name" {
    description = "VPC name"
    type = string
    default = "go-todo-api-vpc"
}
variable "vpc_cidr" {
    description = "CIDR block"
    type = string
    default = "10.0.0.0/16"
}
variable "azs" {
    description = "AZs for subnets"
  type = list(string)
  
}