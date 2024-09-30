# Golang TODO API

Example project to study deployment of Go apps usign terraform as IaC tool, Github Actions as CI-CD tool and AWS
Example project to study deployment of Go apps usign terraform as IaC tool, Github Actions as CI-CD tool and AWS

## Project description

This API is an example project, the basic idea is an app to manage projects, define a project owner and team members, to assign roles of any members of a project team, and finnaly to assign and manage tasks for every team member inside the project


## AWS Infrastructure Deployment API with Terraform

We are goin to use Terraform to create and deploy the infra in AWS. The architecture utilizes API Gateway, a VPC with a subnet, Kubernetes for deployment, and CloudWatch for monitoring and logging.********

## Prerequisites

* AWS account with administrator privileges, also create an specific user for this projecto with all the needed permissions (policies)****
* Terraform installed and configured with your AWS credentials.


## Steps for declaring and organize IaC with terraform
Before geeting dirty is helpfull have a detailed plan about wich components will be created with  terraform and how to organize those componentes just like an app project, for example here is the terraform folder structure:
```ultree
terraform
├── backendSetup
│   └── main.tf
├── security_groups
│   └── main.tf
├── vpc
│   └── main.tf
├── nat_gateway
│   └── main.tf
├── kubernetes
│   └── main.tf
├── iam
│   └── main.tf
├── api_gateway
│   └── main.tf
├── cloudwatch_logs
│   └── main.tf
└── main.tf

```
**1. Configure Terraform Remote Backend:**

This step is about define a remote backend for terraform, is not included in the repo because security reasons, besides you can go to the cloud provider console an created it manually in this case i just prefeer to define de remote backend also in terraform, before anything else, then its deployed from my local machine with the addional bonus that i can remove it, or edit from the project it self if i need to. In this case the backend definition is in `terraform/backendSetup/main.tf` is all about defining the S3 bucket, versioning, encryption, DynamoDB table for state locking, and the AWS region for the remote backend, a little trick S3 buckets name must be unique and written in lowercase only
  
```hcl
#TF remote backend component definition in aws
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"

    }
  }
}
provider "aws" {
  region = "us-east-1"
}

#define a S3 bucket to store the TF state file
resource "aws_s3_bucket" "globalbucketname" {
  #this must be unique globally chkeck carefully, DO NOT use upper case
  bucket = "bucket-name-internal"
  #enables to destroy the bucket including the content
  force_destroy = true
}

#versioning config for the s3 bucket
resource "aws_s3_bucket_versioning" "globalbucketname-versioning" {
  #get id from bucket
  bucket = aws_s3_bucket.globalbucketname
  .id
  versioning_configuration {
    status = "Enabled"
  }
}

#server side encryption for the created bucket
resource "aws_s3_bucket_server_side_encryption_configuration" "globalbucketname-encrypt-conf" {
  bucket = aws_s3_bucket.globalbucketname
  .bucket
  #encryption ruls
  rule {
    #default encryption applied and sse algo to apply
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}
#dynamoDB table to enable terraform state lock (avoid state overwriting)
resource "aws_dynamodb_table" "terraform-locks" {
  name         = "tf-statelocking"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}

```
  

**2. Create VPC and Subnet:**

* **Define VPC:** Create a VPC (Virtual Private Cloud) in `terraform/vpc/main.tf`. Specify the IP range (CIDR) and configure routing options.
* **Create Subnet:** Define a subnet within the VPC in `terraform/vpc/main.tf`. Specify the IP range (CIDR) and the availability zone (AZ) where the subnet will be deployed.

**3. Define Security Groups:**

* **API Gateway Security Group:** Create a security group in `terraform/security_groups/main.tf` that allows access to your API Gateway from the outside (e.g., port 443 for HTTPS).
* **Kubernetes Security Group:** Create another security group in `terraform/security_groups/main.tf` for your Kubernetes cluster. Allow access to the application pods from the VPC's private network.

**4. Provide Internet Access:**

* **NAT Gateway:** If your Kubernetes pods need internet access, define a NAT Gateway in `terraform/nat_gateway/main.tf`. This will allow your pods to communicate with the internet without exposing them to the public network.

**5. Create Kubernetes Cluster:**

* **Define Cluster:** Create a Kubernetes cluster in `terraform/kubernetes/main.tf`. Define the number of nodes, node type, Kubernetes version, and cluster configuration options.

**6. Install CloudWatch Agent for Monitoring:**

* **Create IAM Role:** Define an IAM role in `terraform/iam/main.tf` that allows the CloudWatch Agent to collect metrics and logs from your Kubernetes cluster.
* **Install CloudWatch Agent:** Within the definition of your Kubernetes cluster in `terraform/kubernetes/main.tf`, configure the CloudWatch Agent to install on each node.

**7. Configure API Gateway:**

* **Define API Gateway:** Create a REST API in `terraform/api_gateway/main.tf`. Specify the API type (REST) and configuration options.
* **Define Resources:** Define API resources (e.g., `/users`, `/tasks`, etc.) in `terraform/api_gateway/main.tf`. Each resource will represent a route of your API.
* **Define Methods:** For each resource, define the allowed HTTP methods (GET, POST, PUT, DELETE, etc.).
* **Configure Integrations:**
    * **Kubernetes Integration:** In `terraform/api_gateway/main.tf`, define an integration that connects your API Gateway to your Kubernetes service. The integration should specify the URL of the Kubernetes service exposing your Go API.
    * **Configure Route Mapping:** Configure the mapping between API Gateway routes and the routes of your application within Kubernetes.

**8. Deploy the Go Application to Kubernetes:**

* **Define Deployment:** Create a deployment in `terraform/kubernetes/main.tf` responsible for deploying your Go API as a pod within the Kubernetes cluster.
* **Define Service:** Create a service in `terraform/kubernetes/main.tf` that exposes your application as a service within the Kubernetes cluster. Ensure the service is configured to be exposed on the private network of your VPC.

**9. Configure Ingress for API Gateway:**

* **Create Ingress:** Define an ingress in `terraform/kubernetes/main.tf`. The Ingress will redirect requests from the API Gateway to your Kubernetes service.
* **Configure Custom Domain:** If you want to use a custom domain for your API, configure a custom DNS in AWS Route 53 and configure the Ingress to use that domain.

**10. Configure CloudWatch Logs:**

* **Create Logs Group:** Define a logs group in `terraform/cloudwatch_logs/main.tf` to store your API logs.
* **Configure Filter:** Define a filter in `terraform/cloudwatch_logs/main.tf` to log only the logs of your API.

**Important Notes:**

*  Consider using a separate Security Group for your API Gateway to ensure that your application is not exposed to the public internet.
* Implement best security practices for your API and infrastructure.