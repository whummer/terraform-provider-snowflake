---
page_title: "snowflake_account Resource - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  The account resource allows you to create and manage Snowflake accounts. For more information, check account documentation https://docs.snowflake.com/en/user-guide/organizations-manage-accounts.
---

# snowflake_account (Resource)

The account resource allows you to create and manage Snowflake accounts. For more information, check [account documentation](https://docs.snowflake.com/en/user-guide/organizations-manage-accounts).

~> **Note** To use this resource you have to use an account with a privilege to use the ORGADMIN role.

~> **Note** Changes for the following fields won't be detected: `admin_name`, `admin_password`, `admin_rsa_public_key`, `admin_user_type`, `first_name`, `last_name`, `email`, `must_change_password`. This is because these fields only supply initial values for creating the admin user. Once the account is created, the admin user becomes an independent entity. Modifying users from the account resource is challenging since it requires logging into that account. This would require the account resource logging into the account it created to read or alter admin user properties, which is impractical, because any external change to the admin user would disrupt the change detection anyway.

~> **Note** During the import, when Terraform detects changes on a field with `ForceNew`, it will try to recreate the resource. Due to Terraform limitations, `grace_period_in_days` is not set at that moment. This means that Terraform will try to drop the account with the empty grace period which is required, and fail.
Before importing, ensure if the resource configuration matches the actual state.
See more in our [Resource Migration guide](../guides/resource_migration#312-terraform-import) and [issue #3390](https://github.com/snowflakedb/terraform-provider-snowflake/issues/3390).

## Example Usage

```terraform
## Minimal
resource "snowflake_account" "minimal" {
  name                 = "ACCOUNT_NAME"
  admin_name           = var.admin_name
  admin_password       = var.admin_password
  email                = var.email
  edition              = "STANDARD"
  grace_period_in_days = 3
}

## Complete (with SERVICE user type)
resource "snowflake_account" "complete" {
  name                 = "ACCOUNT_NAME"
  admin_name           = var.admin_name
  admin_rsa_public_key = "<public_key>"
  admin_user_type      = "SERVICE"
  email                = var.email
  edition              = "STANDARD"
  region_group         = "PUBLIC"
  region               = "AWS_US_WEST_2"
  comment              = "some comment"
  is_org_admin         = "true"
  grace_period_in_days = 3
}

## Complete (with PERSON user type)
resource "snowflake_account" "complete" {
  name                 = "ACCOUNT_NAME"
  admin_name           = var.admin_name
  admin_password       = var.admin_password
  admin_user_type      = "PERSON"
  first_name           = var.first_name
  last_name            = var.last_name
  email                = var.email
  must_change_password = "false"
  edition              = "STANDARD"
  region_group         = "PUBLIC"
  region               = "AWS_US_WEST_2"
  comment              = "some comment"
  is_org_admin         = "true"
  grace_period_in_days = 3
}

variable "admin_name" {
  type      = string
  sensitive = true
}

variable "email" {
  type      = string
  sensitive = true
}

variable "admin_password" {
  type      = string
  sensitive = true
}

variable "first_name" {
  type      = string
  sensitive = true
}

variable "last_name" {
  type      = string
  sensitive = true
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_name` (String, Sensitive) Login name of the initial administrative user of the account. A new user is created in the new account with this name and password and granted the ACCOUNTADMIN role in the account. A login name can be any string consisting of letters, numbers, and underscores. Login names are always case-insensitive. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `edition` (String) Snowflake Edition of the account. See more about Snowflake Editions in the [official documentation](https://docs.snowflake.com/en/user-guide/intro-editions). Valid options are: `STANDARD` | `ENTERPRISE` | `BUSINESS_CRITICAL`
- `email` (String, Sensitive) Email address of the initial administrative user of the account. This email address is used to send any notifications about the account. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `grace_period_in_days` (Number) Specifies the number of days during which the account can be restored (“undropped”). The minimum is 3 days and the maximum is 90 days.
- `name` (String) Specifies the identifier (i.e. name) for the account. It must be unique within an organization, regardless of which Snowflake Region the account is in and must start with an alphabetic character and cannot contain spaces or special characters except for underscores (_). Note that if the account name includes underscores, features that do not accept account names with underscores (e.g. Okta SSO or SCIM) can reference a version of the account name that substitutes hyphens (-) for the underscores.

### Optional

- `admin_password` (String, Sensitive) Password for the initial administrative user of the account. Either admin_password or admin_rsa_public_key has to be specified. This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `admin_rsa_public_key` (String) Assigns a public key to the initial administrative user of the account. Either admin_password or admin_rsa_public_key has to be specified. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `admin_user_type` (String) Used for setting the type of the first user that is assigned the ACCOUNTADMIN role during account creation. Valid options are: `PERSON` | `SERVICE` | `LEGACY_SERVICE` External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `comment` (String) Specifies a comment for the account.
- `consumption_billing_entity` (String) Determines which billing entity is responsible for the account's consumption-based billing.
- `first_name` (String, Sensitive) First name of the initial administrative user of the account. This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `is_org_admin` (String) (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Sets an account property that determines whether the ORGADMIN role is enabled in the account. Only an organization administrator (i.e. user with the ORGADMIN role) can set the property.
- `last_name` (String, Sensitive) Last name of the initial administrative user of the account. This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `must_change_password` (String) (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether the new user created to administer the account is forced to change their password upon first login into the account. This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
- `region` (String) [Snowflake Region ID](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#label-snowflake-region-ids) of the region where the account is created. If no value is provided, Snowflake creates the account in the same Snowflake Region as the current account (i.e. the account in which the CREATE ACCOUNT statement is executed.)
- `region_group` (String) ID of the region group where the account is created. To retrieve the region group ID for existing accounts in your organization, execute the [SHOW REGIONS](https://docs.snowflake.com/en/sql-reference/sql/show-regions) command. For information about when you might need to specify region group, see [Region groups](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#label-region-groups).
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.
- `show_output` (List of Object) Outputs the result of `SHOW ACCOUNTS` for the given account. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `account_locator` (String)
- `account_locator_url` (String)
- `account_name` (String)
- `account_old_url_last_used` (String)
- `account_old_url_saved_on` (String)
- `account_url` (String)
- `comment` (String)
- `consumption_billing_entity_name` (String)
- `created_on` (String)
- `dropped_on` (String)
- `edition` (String)
- `is_events_account` (Boolean)
- `is_org_admin` (Boolean)
- `is_organization_account` (Boolean)
- `managed_accounts` (Number)
- `marketplace_consumer_billing_entity_name` (String)
- `marketplace_provider_billing_entity_name` (String)
- `moved_on` (String)
- `moved_to_organization` (String)
- `old_account_url` (String)
- `organization_name` (String)
- `organization_old_url` (String)
- `organization_old_url_last_used` (String)
- `organization_old_url_saved_on` (String)
- `organization_url_expiration_on` (String)
- `region_group` (String)
- `restored_on` (String)
- `scheduled_deletion_time` (String)
- `snowflake_region` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_account.example '"<organization_name>"."<account_name>"'
```
