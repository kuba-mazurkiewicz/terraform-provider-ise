---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_network_access_time_and_date_condition Resource - terraform-provider-ise"
subcategory: "Network Access"
description: |-
  This resource can manage a Network Access Time And Date Condition.
---

# ise_network_access_time_and_date_condition (Resource)

This resource can manage a Network Access Time And Date Condition.

## Example Usage

```terraform
resource "ise_network_access_time_and_date_condition" "example" {
  name                 = "Cond1"
  description          = "My description"
  is_negate            = false
  week_days            = ["Monday"]
  week_days_exception  = ["Tuesday"]
  start_date           = "2022-05-06"
  end_date             = "2022-05-10"
  exception_start_date = "2022-06-06"
  exception_end_date   = "2022-06-10"
  start_time           = "08:00"
  end_time             = "15:00"
  exception_start_time = "20:00"
  exception_end_time   = "22:00"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Condition name

### Optional

- `description` (String) Condition description
- `end_date` (String) End date
- `end_time` (String) End time
- `exception_end_date` (String) Exception end date
- `exception_end_time` (String) Exception end time
- `exception_start_date` (String) Exception start date
- `exception_start_time` (String) Exception start time
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `start_date` (String) Start date
- `start_time` (String) Start time
- `week_days` (Set of String) Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
- `week_days_exception` (Set of String) Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import ise_network_access_time_and_date_condition.example "76d24097-41c4-4558-a4d0-a8c07ac08470"
```
