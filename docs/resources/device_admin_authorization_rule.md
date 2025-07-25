---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_device_admin_authorization_rule Resource - terraform-provider-ise"
subcategory: "Device Administration"
description: |-
  This resource can manage a Device Admin Authorization Rule.
---

# ise_device_admin_authorization_rule (Resource)

This resource can manage a Device Admin Authorization Rule.

## Example Usage

```terraform
resource "ise_device_admin_authorization_rule" "example" {
  policy_set_id             = "d82952cb-b901-4b09-b363-5ebf39bdbaf9"
  name                      = "Rule1"
  default                   = false
  rank                      = 0
  state                     = "enabled"
  condition_type            = "ConditionAttributes"
  condition_is_negate       = false
  condition_attribute_name  = "Location"
  condition_attribute_value = "All Locations"
  condition_dictionary_name = "DEVICE"
  condition_operator        = "equals"
  command_sets              = ["DenyAllCommands"]
  profile                   = "Default Shell Profile"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
- `policy_set_id` (String) Policy set ID

### Optional

- `children` (Attributes Set) List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`. (see [below for nested schema](#nestedatt--children))
- `command_sets` (Set of String) Command sets enforce the specified list of commands that can be executed by a device administrator
- `condition_attribute_name` (String) Dictionary attribute name
- `condition_attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `condition_dictionary_name` (String) Dictionary name
- `condition_dictionary_value` (String) Dictionary value
- `condition_id` (String) UUID for condition
- `condition_is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `condition_operator` (String) Equality operator
  - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
  - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
- `default` (Boolean) Indicates if this rule is the default one
- `profile` (String) Device admin profiles control the initial login session of the device administrator
- `rank` (Number) The rank (priority) in relation to other rules. Lower rank is higher priority.
- `state` (String) The state that the rule is in. A disabled rule cannot be matched.
  - Choices: `disabled`, `enabled`, `monitor`

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--children"></a>
### Nested Schema for `children`

Required:

- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
  - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`

Optional:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `children` (Attributes Set) List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`. (see [below for nested schema](#nestedatt--children--children))
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `id` (String) UUID for condition
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `operator` (String) Equality operator
  - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`

<a id="nestedatt--children--children"></a>
### Nested Schema for `children.children`

Required:

- `condition_type` (String) Condition type.
  - Choices: `ConditionAttributes`, `ConditionReference`

Optional:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `id` (String) UUID for condition
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `operator` (String) Equality operator
  - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import ise_device_admin_authorization_rule.example "76d24097-41c4-4558-a4d0-a8c07ac08470,76d24097-41c4-4558-a4d0-a8c07ac08470"
```
