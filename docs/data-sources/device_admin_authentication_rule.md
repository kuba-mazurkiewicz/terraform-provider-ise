---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_device_admin_authentication_rule Data Source - terraform-provider-ise"
subcategory: "Policy"
description: |-
  This data source can read the Device Admin Authentication Rule.
---

# ise_device_admin_authentication_rule (Data Source)

This data source can read the Device Admin Authentication Rule.

## Example Usage

```terraform
data "ise_device_admin_authentication_rule" "example" {
  id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  policy_set_id = "d82952cb-b901-4b09-b363-5ebf39bdbaf9"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_set_id` (String) Policy set ID

### Optional

- `id` (String) The id of the object
- `name` (String) Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]

### Read-Only

- `children` (Attributes List) List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `ConditionAttributes` or `ConditionReference`. (see [below for nested schema](#nestedatt--children))
- `condition_attribute_name` (String) Dictionary attribute name
- `condition_attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `condition_dictionary_name` (String) Dictionary name
- `condition_dictionary_value` (String) Dictionary value
- `condition_id` (String) UUID for condition
- `condition_is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `condition_operator` (String) Equality operator
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `default` (Boolean) Indicates if this rule is the default one
- `identity_source_name` (String) Identity source name from the identity stores
- `if_auth_fail` (String) Action to perform when authentication fails such as Bad credentials, disabled user and so on
- `if_process_fail` (String) Action to perform when ISE is unable to access the identity database
- `if_user_not_found` (String) Action to perform when user is not found in any of identity stores
- `rank` (Number) The rank (priority) in relation to other rules. Lower rank is higher priority.
- `state` (String) The state that the rule is in. A disabled rule cannot be matched.

<a id="nestedatt--children"></a>
### Nested Schema for `children`

Read-Only:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `children` (Attributes List) List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `ConditionAttributes` or `ConditionReference`. (see [below for nested schema](#nestedatt--children--children))
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `id` (String) UUID for condition
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `operator` (String) Equality operator

<a id="nestedatt--children--children"></a>
### Nested Schema for `children.children`

Read-Only:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `id` (String) UUID for condition
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `operator` (String) Equality operator