---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_device_admin_policy_set Data Source - terraform-provider-ise"
subcategory: "Policy"
description: |-
  This data source can read the Device Admin Policy Set.
---

# ise_device_admin_policy_set (Data Source)

This data source can read the Device Admin Policy Set.

## Example Usage

```terraform
data "ise_device_admin_policy_set" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the object
- `name` (String) Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]

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
- `description` (String) The description of the policy set
- `is_proxy` (Boolean) Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
- `rank` (Number) The rank (priority) in relation to other policy sets. Lower rank is higher priority.
- `service_name` (String) Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
- `state` (String) The state that the policy set is in. A disabled policy set cannot be matched.

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