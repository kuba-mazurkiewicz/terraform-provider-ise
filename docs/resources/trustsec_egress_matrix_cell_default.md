---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_trustsec_egress_matrix_cell_default Resource - terraform-provider-ise"
subcategory: "TrustSec"
description: |-
  Allows modifications to the default egress policy matrix rule
---

# ise_trustsec_egress_matrix_cell_default (Resource)

Allows modifications to the default egress policy matrix rule

## Example Usage

```terraform
resource "ise_trustsec_egress_matrix_cell_default" "example" {
  description        = "Default egress rule"
  default_rule       = "PERMIT_IP"
  matrix_cell_status = "ENABLED"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `default_rule` (String) Can be used only if sgacls not specified. Final Catch All Rule
  - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
- `description` (String) Description
- `matrix_cell_status` (String) Matrix Cell Status
  - Choices: `DISABLED`, `ENABLED`, `MONITOR`
  - Default value: `DISABLED`
- `sgacls` (Set of String) List of TrustSec Security Groups ACLs

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import ise_trustsec_egress_matrix_cell_default.example "92c1a900-8c01-11e6-996c-525400b48521"
```
