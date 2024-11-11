# VPC Outputs 
output "vpc_id" {
  description = "vpc id"
  value       = module.vpc.vpc_id
}

output "private_subnets" {
  description = "private subnets ids"
  value       = module.vpc.private_subnets
}

output "public_subnets" {
  description = "public subnets ids"
  value       = module.vpc.public_subnets
}

# EKS cluster outputs
output "eks_cluster_name" {
  description = "EKS cluster name"
  value       = module.eks.cluster_name
}

output "eks_cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = module.eks.cluster_endpoint
}

output "eks_cluster_security_group" {
  description = "EKS security group"
  value       = module.eks.cluster_security_group_id
}
