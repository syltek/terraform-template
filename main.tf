locals {
  # Common tags used for all resources in the module
  module_tags = merge(
    {
      managed_by       = "terraform"
      terraform_module = "terraform-{provider}-{service}"
    },
    var.metadata.tags
  )
}

# Add data sources here
# Example:
# data "aws_..." "example" {}

# Add resources here
# Example:
# resource "aws_..." "main" {
#   name        = var.metadata.name
#   description = var.metadata.description
#   tags        = local.module_tags
# }
