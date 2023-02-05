---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "laravelenvoyer_server Resource - terraform-provider-laravelenvoyer"
subcategory: ""
description: |-
  Sample resource in the Terraform laravelenvoyer scaffolding.
---

# laravelenvoyer_server (Resource)

Sample resource in the Terraform laravelenvoyer scaffolding.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connect_as` (String) Connect as user.
- `deployment_path` (String) Path to where project lives on the server.
- `host` (String) Host / IP address.
- `name` (String) Server name
- `php_version` (String) PHP version.
- `project_id` (Number) Project ID.

### Optional

- `composer_path` (String) Path to composer
- `port` (Number) Port.
- `receives_code_deployments` (Boolean) Receives code deployments.
- `should_restart_fpm` (Boolean) Restart PHP after deployment.

### Read-Only

- `connection_status` (String) Connection status.
- `id` (String) The ID of this resource.
- `public_key` (String) Public key.
- `user_id` (String) User ID.

