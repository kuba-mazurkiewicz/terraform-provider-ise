
---
layout: ""
page_title: "Provider: ISE"
description: |-
  The ISE provider provides resources to interact with a Cisco ISE (Identity Service Engine) instance.
---

# ISE Provider

The ISE provider provides resources to interact with a Cisco ISE (Identity Service Engine) instance. It communicates with ISE via the REST API.

This provider uses both, the ERS and Open API. Instructions on how to enable API access can be found here: https://developer.cisco.com/docs/identity-services-engine/latest/#!cisco-ise-api-framework

All resources and data sources have been tested with the following releases.

| Platform | Version       |
| -------- | ------------- |
| ISE      | 3.2.0 Patch 4 |
| ISE      | 3.3.0         |
| ISE      | 3.4.0         |

## Getting Started

The following guides with examples exist to demonstrate the use of the provider:

- [Getting Started](https://registry.terraform.io/providers/CiscoDevNet/ise/latest/docs/guides/getting_started)
- [Authentication Rules](https://registry.terraform.io/providers/CiscoDevNet/ise/latest/docs/guides/authentication_rules)

## Example Usage

```terraform
provider "ise" {
  username = "admin"
  password = "password"
  url      = "https://10.1.1.1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `insecure` (Boolean) Allow insecure HTTPS client. This can also be set as the ISE_INSECURE environment variable. Defaults to `true`.
- `password` (String, Sensitive) Password for the ISE instance. This can also be set as the ISE_PASSWORD environment variable.
- `retries` (Number) Number of retries for REST API calls. This can also be set as the ISE_RETRIES environment variable. Defaults to `3`.
- `url` (String) URL of the Cisco ISE instance. This can also be set as the ISE_URL environment variable.
- `username` (String) Username for the ISE instance. This can also be set as the ISE_USERNAME environment variable.
