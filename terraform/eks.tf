module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 20.0"

  cluster_name           = var.cluster_name
  cluster_version        = "1.31"
  vpc_id                 = module.vpc.vpc_id
  subnet_ids             = module.vpc.private_subnets
 
  cluster_endpoint_public_access  = true #use this in false in production
  cluster_endpoint_private_access  = true

  cluster_addons = {
    coredns = {
        resolve_conflict = "OVERWRITE"
    }
    csi = {
        resolve_conflict = "OVERWRITE"
    }
    kube-proxy = {
        resolve_conflict = "OVERWRITE"
    }
    vpc-cni = {
        resolve_conflict = "OVERWRITE"
    }
  }

  eks_managed_node_groups = {
    node_group ={
      desired_capacity = 1
      max_capacity     = 2
      min_capacity     = 1
      instance_type    = ["t3.medium"]
      tags ={
        Environment = "develop"
      }
        
    }
  }
 # Cluster access entry
  # To add the current caller identity as an administrator
  enable_cluster_creator_admin_permissions = true

  tags = {
    Terraform ="true"
    Environment = "develop"
  }

}