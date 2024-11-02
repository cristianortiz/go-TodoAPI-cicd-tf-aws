module "vpc" {
    source = "terraform-aws-modules/vpc/aws"
    name ="go-todo-api-vpc"
    cidr = "10.0.0.0/16"

    azs = ["us-east-1a","us-east-1b","us-east-1c"]
    private_subnets = ["10.0.1.0/24","10.0.2.0/24","10.0.3.0/24"]
    public_subnets = ["10.0.101.0/24","10.0.102.0/24","10.0.103.0/24"]
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