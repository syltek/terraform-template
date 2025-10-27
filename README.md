# terraform-{provider}-{service}

Brief description of what the module creates.

## Important Notes

Key considerations, limitations, or special requirements for this module.

## Inputs
<!--
- `metadata`: The metadata information for the module.
  - `name`: The name of the resource
  - `description`: The description of the resource
  - `tags`: Additional tags to apply to the resources

For account-specific resources, use this structure instead:
- `metadata`: The metadata information for the module.
  - `name`: The name of the resource
  - `description`: (Optional) The description of the resource
  - `account`: The account identifier
  - `environment`: The environment name
  - `tags`: Additional tags to apply to the resources
-->

<!-- Add your config variables here:
- `config`: (Optional) Configuration for the resource. Defaults to `{}`.
  - `field_name`: (Optional) Description of the field. Defaults to `default_value`
-->

## Outputs

<!-- Add your outputs here:
- `arn`: The Amazon Resource Name (ARN) of the resource.
- `id`: The identifier of the resource.
-->

## Usage

### Example for account-agnostic resources

```hcl
module "example" {
  source  = "app.terraform.io/Playtomic/{service}/{provider}"
  version = "~> 0.1"

  metadata = {
    name        = "ExampleResource"
    description = "Example resource description"
    tags = {
      team = "infrastructure"
    }
  }

  # Add your configuration here
  # config = {
  #   field_name = "value"
  # }
}
```

### Example for account-specific resources

```hcl
module "example" {
  source  = "app.terraform.io/Playtomic/{service}/{provider}"
  version = "~> 0.1"

  metadata = {
    name        = "example"
    description = "Example resource description"
    account     = "example-account"
    environment = "develop"
    tags = {
      team = "infrastructure"
    }
  }

  # Add your configuration here
  # config = {
  #   field_name = "value"
  # }
}
```
