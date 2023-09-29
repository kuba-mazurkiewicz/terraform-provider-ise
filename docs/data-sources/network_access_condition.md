---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ise_network_access_condition Data Source - terraform-provider-ise"
subcategory: "Network Access"
description: |-
  This data source can read the Network Access Condition.
---

# ise_network_access_condition (Data Source)

This data source can read the Network Access Condition.

## Example Usage

```terraform
data "ise_network_access_condition" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `children` (Attributes List) List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `LibraryConditionAndBlock` or `LibraryConditionOrBlock`. (see [below for nested schema](#nestedatt--children))
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `description` (String) Condition description
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `name` (String) Condition name
- `operator` (String) Equality operator

<a id="nestedatt--children"></a>
### Nested Schema for `children`

Read-Only:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `children` (Attributes List) List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `LibraryConditionAndBlock` or `LibraryConditionOrBlock`. (see [below for nested schema](#nestedatt--children--children))
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `description` (String) Condition description
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `name` (String) Condition name
- `operator` (String) Equality operator

<a id="nestedatt--children--children"></a>
### Nested Schema for `children.children`

Read-Only:

- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) Attribute value for condition. Value type is specified in dictionary object.
- `condition_type` (String) Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
- `description` (String) Condition description
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `is_negate` (Boolean) Indicates whereas this condition is in negate mode
- `name` (String) Condition name
- `operator` (String) Equality operator