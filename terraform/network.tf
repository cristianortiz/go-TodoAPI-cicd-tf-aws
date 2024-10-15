module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  name   = "go-todoapi-vpc"
  cidr   = "10.0.0.0/16"

  azs             = ["us-east-1", "us-east-1b", "us-east-1d"]
  private_subnets = []

}
