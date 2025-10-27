terraform {
  required_version = ">= 1.9.0"
  required_providers {
    # Add your required providers here
    # Example for AWS:
    # aws = {
    #   source  = "hashicorp/aws"
    #   version = "~> 5.92"
    # }
    #
    # Example for Datadog:
    # datadog = {
    #   source  = "datadog/datadog"
    #   version = "~> 3.0"
    # }
  }
}

# Add provider configurations only when necessary (e.g., database providers)
# Example:
# provider "mysql" {
#   endpoint   = aws_rds_cluster.cluster.endpoint
#   username   = var.cluster.master_username
#   password   = aws_rds_cluster.cluster.master_password
#   private_ip = true
# }
