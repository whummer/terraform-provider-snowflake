---
page_title: "snowflake_account_password_policy_attachment Resource - terraform-provider-snowflake"
subcategory: "Preview"
description: |-
  Specifies the password policy to use for the current account. To set the password policy of a different account, use a provider alias.
---

!> **Caution: Preview Feature** This feature is considered a preview feature in the provider, regardless of the state of the resource in Snowflake. We do not guarantee its stability. It will be reworked and marked as a stable feature in future releases. Breaking changes are expected, even without bumping the major version. To use this feature, add the relevant feature name to `preview_features_enabled` field in the [provider configuration](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#schema). Please always refer to the [Getting Help](https://github.com/snowflakedb/terraform-provider-snowflake?tab=readme-ov-file#getting-help) section in our Github repo to best determine how to get help for your questions.

!> **Warning** This resource shouldn't be used with `snowflake_current_account` resource in the same configuration, as it may lead to unexpected behavior.

# snowflake_account_password_policy_attachment (Resource)

Specifies the password policy to use for the current account. To set the password policy of a different account, use a provider alias.

## Example Usage

```terraform
resource "snowflake_password_policy" "default" {
  database = "prod"
  schema   = "security"
  name     = "default_policy"
}

resource "snowflake_account_password_policy_attachment" "attachment" {
  password_policy = snowflake_password_policy.default.fully_qualified_name
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `password_policy` (String) Qualified name (`"db"."schema"."policy_name"`) of the password policy to apply to the current account.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
