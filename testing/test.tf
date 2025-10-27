# Required test variables
variable "test_author" { type = string }
variable "test_id" { type = string }

# Add provider-specific variables here:
# For AWS modules:
# variable "aws_role_arn" { type = string }
# variable "aws_region" { type = string }
#
# For Datadog modules:
# variable "dd_api_key" { type = string }
# variable "dd_app_key" { type = string }

locals {
  # For account-agnostic resources, use:
  common_tags = {
    test_author = var.test_author
    test_id     = var.test_id
  }

  # For account-specific resources, use:
  # account     = "sre-sandbox"
  # environment = "sandbox"
  # common_tags = {
  #   test_author = var.test_author
  #   test_id     = var.test_id
  # }
}

# Configure provider(s)
# For AWS:
# provider "aws" {
#   region = var.aws_region
#   assume_role {
#     role_arn = var.aws_role_arn
#   }
# }
#
# For Datadog:
# provider "datadog" {
#   api_key = var.dd_api_key
#   app_key = var.dd_app_key
#   api_url = "https://api.datadoghq.eu/"
# }

# Test Case 1 - Minimal Configuration
module "test_1" {
  source = "../"

  metadata = {
    name        = "test-1-${var.test_id}"
    description = "Test case 1 - minimal configuration"
    tags = merge(local.common_tags, {
      test_case = "test_1_minimal"
    })
  }

  # Add minimal configuration here
  # config = {}
}

# Outputs for Test Case 1
# output "test_1_arn" {
#   value       = module.test_1.arn
#   description = "The ARN of test case 1"
# }

# Test Case 2 - Full Configuration
module "test_2" {
  source = "../"

  metadata = {
    name        = "test-2-${var.test_id}"
    description = "Test case 2 - full configuration"
    tags = merge(local.common_tags, {
      test_case = "test_2_full"
    })
  }

  # Add full configuration here
  # config = {
  #   field_name = "value"
  # }
}

# Outputs for Test Case 2
# output "test_2_arn" {
#   value       = module.test_2.arn
#   description = "The ARN of test case 2"
# }

# Test Case 3 - Edge Case
# module "test_3" {
#   source = "../"
#
#   metadata = {
#     name        = "test-3-${var.test_id}"
#     description = "Test case 3 - edge case"
#     tags = merge(local.common_tags, {
#       test_case = "test_3_edge_case"
#     })
#   }
#
#   # Add edge case configuration here
#   # config = {
#   #   field_name = "edge_value"
#   # }
# }
#
# # Outputs for Test Case 3
# # output "test_3_arn" {
# #   value       = module.test_3.arn
# #   description = "The ARN of test case 3"
# # }
