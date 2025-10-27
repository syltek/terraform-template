# Choose the appropriate metadata structure based on your module's context:

# For account-agnostic resources (e.g., IAM, SSO):
# variable "metadata" {
#   type = object({
#     name        = string
#     description = string
#     tags        = map(string)
#   })
#   description = "The metadata information for the module."
# }

# For account-specific resources (e.g., RDS, EC2, VPC), use this instead:
# variable "metadata" {
#   type = object({
#     name        = string
#     description = optional(string)
#     account     = string
#     environment = string
#     tags        = map(string)
#   })
#   description = "The metadata information for the module."
# }

# Add your configuration variables here
# Example:
# variable "config" {
#   type = object({
#     field_name = optional(string, "default_value")
#   })
#   description = <<EOF
# Configuration for the resource.
#   - field_name: (Optional) Description of the field. Defaults to `default_value`
# EOF
#   default     = {}
# }
