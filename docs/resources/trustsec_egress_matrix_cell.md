---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_trustsec_egress_matrix_cell Resource - terraform-provider-ise"
subcategory: "TrustSec"
description: |-
  This resource can manage a TrustSec Egress Matrix Cell.
---

# ise_trustsec_egress_matrix_cell (Resource)

This resource can manage a TrustSec Egress Matrix Cell.

## Example Usage

```terraform
resource "ise_trustsec_egress_matrix_cell" "example" {
  description        = "EgressMatrixCell Description"
  default_rule       = "NONE"
  matrix_cell_status = "ENABLED"
  sgacls             = ["26b76b10-66e6-11ee-9cc1-9eb2a3ecc82a, 9d64dcd0-6384-11ee-9cc1-9eb2a3ecc82a"]
  source_sgt_id      = "93c66ed0-8c01-11e6-996c-525400b48521"
  destination_sgt_id = "93e1bf00-8c01-11e6-996c-525400b48521"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_sgt_id` (String) Destination Trustsec Security Group ID
- `source_sgt_id` (String) Source Trustsec Security Group ID

### Optional

- `default_rule` (String) Can be used only if sgacls not specified.
  - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
  - Default value: `NONE`
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
terraform import ise_trustsec_egress_matrix_cell.example "76d24097-41c4-4558-a4d0-a8c07ac08470"
```
