---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_network_access_authorization_rule Resource - terraform-provider-ise"
subcategory: "Policy"
description: |-
  This resource can manage a Network Access Authorization Rule.
---

# ise_network_access_authorization_rule (Resource)

This resource can manage a Network Access Authorization Rule.

## Example Usage

```terraform
resource "ise_network_access_authorization_rule" "example" {
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
  profile                   = ["PermitAccess"]
  security_group            = "BYOD"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]

### Optional

- `condition_attribute_name` (String) Dictionary attribute name
- `condition_attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `condition_dictionary_name` (String) Dictionary name
- `condition_dictionary_value` (String) Dictionary value
- `condition_is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `condition_operator` (String) Equality operator
  - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
  - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`, `LibraryConditionAndBlock`, `LibraryConditionAttributes`, `LibraryConditionOrBlock`, `TimeAndDateCondition`
- `default` (Boolean) Indicates if this rule is the default one
- `policy_set_id` (String) Policy set ID
- `profile` (List of String) The authorization profile(s)
- `rank` (Number) The rank (priority) in relation to other rules. Lower rank is higher priority.
- `security_group` (String) Security group used in authorization policies
- `state` (String) The state that the rule is in. A disabled rule cannot be matched.
  - Choices: `disabled`, `enabled`, `monitor`

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import ise_network_access_authorization_rule.example "76d24097-41c4-4558-a4d0-a8c07ac08470"
```