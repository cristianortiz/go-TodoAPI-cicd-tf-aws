module "vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "~> 5.15"
    name = var.vpc_name
    cidr = var.vpc_cidr

    azs = var.azs
    private_subnets = var.private_subnets
    public_subnets = var.public_subnets
    #internet outbound anly access for private subnets
    enable_nat_gateway = true
    enable_vpn_gateway = false
    tags = {
        Terraform ="true"
        Enviroment ="develop"
        #optional, to add repo url for infra resource tagging
        #repo_url=""

    }
    #optionals custom tags for deployment of private and public LBs for subnets
    public_subnet_tags = {
        "kubernetes.io/role/elb"="1"
        #deploy public ELB in this shared vpc
        "kubernetes.io/cluster/go-todo-api-vpc"="shared"
    }
    private_subnet_tags = {
        "kubernetes.io/role/internal/elb"="1"
        #deploy private ELB in this shared vpc
        "kubernetes.io/cluster/go-todo-api-vpc"="shared"
    }

}