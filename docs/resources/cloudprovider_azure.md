---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_cloudprovider_azure Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for creating an Azure cloud provider
---

# harness_cloudprovider_azure (Resource)

Resource for creating an Azure cloud provider



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **client_id** (String) The client id for the Azure application
- **key** (String) The Name of the Harness secret containing the key for the Azure application
- **name** (String) The name of the cloud provider.
- **tenant_id** (String) The tenant id for the Azure application

### Optional

- **environment_type** (String) The type of environment. Valid options are [AZURE AZURE_US_GOVERNMENT]
- **usage_scope** (Block Set) Usage scopes (see [below for nested schema](#nestedblock--usage_scope))

### Read-Only

- **id** (String) The id of the cloud provider.

<a id="nestedblock--usage_scope"></a>
### Nested Schema for `usage_scope`

Optional:

- **application_filter_type** (String) Type of application filter applied. ALL if not application id supplied, otherwise NULL
- **application_id** (String) Id of the application scoping
- **environment_filter_type** (String) Type of environment filter applied. ALL if not filter applied
- **environment_id** (String) Id of the environment scoping

