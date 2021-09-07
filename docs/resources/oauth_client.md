---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingfederate_oauth_client Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Provides configuration for OAuth Clients within PingFederate.
---

# pingfederate_oauth_client (Resource)

Provides configuration for OAuth Clients within PingFederate.

## Example Usage

```terraform
resource "pingfederate_oauth_client" "example" {
  client_id   = "example"
  name        = "example"
  grant_types = ["CLIENT_CREDENTIALS"]

  client_auth {
    secret = "super_top_secret"
    type   = "SECRET"
  }

  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }

  oidc_policy {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout", "https://foo"]
    ping_access_logout_capable          = true
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **client_id** (String) A unique identifier the client provides to the Resource Server to identify itself. This identifier is included with every request the client makes.
- **grant_types** (Set of String) The grant types allowed for this client. The EXTENSION grant type applies to SAML/JWT assertion grants.
- **name** (String) A descriptive name for the client instance. This name appears when the user is prompted for authorization.

### Optional

- **bypass_activation_code_confirmation_override** (Boolean) Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device. This overrides the 'bypassUseCodeConfirmation' value present in Authorization Server Settings.
- **bypass_approval_page** (Boolean) Use this setting, for example, when you want to deploy a trusted application and authenticate end users via an IdP adapter or IdP connection.
- **ciba_delivery_mode** (String) The token delivery mode for the client.  The default value is 'POLL'.
- **ciba_notification_endpoint** (String) The endpoint the OP will call after a successful or failed end-user authentication.
- **ciba_polling_interval** (Number) The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds.
- **ciba_request_object_signing_algorithm** (String) The JSON Web Signature [JWS] algorithm that must be used to sign the CIBA Request Object. All signing algorithms are allowed if value is not present.
- **ciba_require_signed_requests** (Boolean) Determines whether CIBA signed requests are required for this client.
- **ciba_user_code_supported** (Boolean) Determines whether CIBA user code is supported for this client.
- **client_auth** (Block List, Max: 1) Client authentication settings.  If this model is null, it indicates that no client authentication will be used. (see [below for nested schema](#nestedblock--client_auth))
- **default_access_token_manager_ref** (Block List, Max: 1) The default access token manager for this client. (see [below for nested schema](#nestedblock--default_access_token_manager_ref))
- **description** (String) A description of what the client application does. This description appears when the user is prompted for authorization.
- **device_flow_setting_type** (String) Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT.
- **device_polling_interval_override** (Number) The amount of time client should wait between polling requests, in seconds. This overrides the 'devicePollingInterval' value present in Authorization Server Settings.
- **enabled** (Boolean) Specifies whether the client is enabled. The default value is true.
- **exclusive_scopes** (Set of String) The exclusive scopes available for this client.
- **extended_properties** (Block Set) OAuth Client Metadata can be extended to use custom Client Metadata Parameters. The names of these custom parameters should be defined in /extendedProperties. (see [below for nested schema](#nestedblock--extended_properties))
- **jwks_settings** (Block List, Max: 1) JSON Web Key Set Settings of the OAuth client. Required if private key JWT client authentication or signed requests is enabled. (see [below for nested schema](#nestedblock--jwks_settings))
- **logo_url** (String) The location of the logo used on user-facing OAuth grant authorization and revocation pages.
- **oidc_policy** (Block List, Max: 1) Open ID Connect Policy settings.  This is included in the message only when OIDC is enabled. (see [below for nested schema](#nestedblock--oidc_policy))
- **pending_authorization_timeout_override** (Number) The 'device_code' and 'user_code' timeout, in seconds. This overrides the 'pendingAuthorizationTimeout' value present in Authorization Server Settings.
- **persistent_grant_expiration_time** (Number) The persistent grant expiration time. -1 indicates an indefinite amount of time.
- **persistent_grant_expiration_time_unit** (String) The persistent grant expiration time unit.
- **persistent_grant_expiration_type** (String) Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT.
- **persistent_grant_idle_timeout** (Number) The persistent grant idle timeout.
- **persistent_grant_idle_timeout_time_unit** (String) The persistent grant idle timeout time unit.
- **persistent_grant_idle_timeout_type** (String) Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT.
- **redirect_uris** (Set of String) URIs to which the OAuth AS may redirect the resource owner's user agent after authorization is obtained. A redirection URI is used with the Authorization Code and Implicit grant types. Wildcards are allowed. However, for security reasons, make the URL as restrictive as possible.For example: https://*.company.com/* Important: If more than one URI is added or if a single URI uses wildcards, then Authorization Code grant and token requests must contain a specific matching redirect uri parameter.
- **refresh_rolling** (String) Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT.
- **request_object_signing_algorithm** (String) The JSON Web Signature [JWS] algorithm that must be used to sign the Request Object. All signing algorithms are allowed if value is not present.
- **request_policy_ref** (Block List, Max: 1) The CIBA request policy. (see [below for nested schema](#nestedblock--request_policy_ref))
- **require_proof_key_for_code_exchange** (Boolean) Determines whether Proof Key for Code Exchange (PKCE) is required for this client.
- **require_pushed_authorization_requests** (Boolean) Determines whether pushed authorization requests are required when initiating an authorization request. The default is false.
- **require_signed_requests** (Boolean) Determines whether signed requests are required for this client
- **restrict_scopes** (Boolean) Restricts this client's access to specific scopes.
- **restrict_to_default_access_token_manager** (Boolean) Determines whether the client is restricted to using only its default access token manager. The default is false.
- **restricted_response_types** (Set of String) The response types allowed for this client. If omitted all response types are available to the client.
- **restricted_scopes** (Set of String) The scopes available for this client.
- **token_exchange_processor_policy_ref** (Block List, Max: 1) The Token Exchange Processor policy. (see [below for nested schema](#nestedblock--token_exchange_processor_policy_ref))
- **user_authorization_url_override** (String) The URL used as 'verification_url' and 'verification_url_complete' values in a Device Authorization request. This property overrides the 'userAuthorizationUrl' value present in Authorization Server Settings.
- **validate_using_all_eligible_atms** (Boolean) Validates token using all eligible access token managers for the client. This setting is ignored if 'restrict_to_default_access_token_manager' is set to true.

### Read-Only

- **id** (String) The ID of this resource.

<a id="nestedblock--client_auth"></a>
### Nested Schema for `client_auth`

Required:

- **type** (String) Client authentication type.

Optional:

- **client_cert_issuer_dn** (String) Client TLS Certificate Issuer DN.
- **client_cert_subject_dn** (String) Client TLS Certificate Subject DN.
- **enforce_replay_prevention** (Boolean) Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication.
- **secret** (String, Sensitive) Client secret for Basic Authentication. To update the client secret, specify the plaintext value in this field. This field will not be populated for GET requests.
- **token_endpoint_auth_signing_algorithm** (String) The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present.


<a id="nestedblock--default_access_token_manager_ref"></a>
### Nested Schema for `default_access_token_manager_ref`

Required:

- **id** (String) The ID of the resource.

Read-Only:

- **location** (String) A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.


<a id="nestedblock--extended_properties"></a>
### Nested Schema for `extended_properties`

Required:

- **key_name** (String) The extended property name.

Optional:

- **values** (Set of String) A List of parameter values.


<a id="nestedblock--jwks_settings"></a>
### Nested Schema for `jwks_settings`

Optional:

- **jwks** (String) JSON Web Key Set (JWKS) document of the OAuth client. Either 'jwks' or 'jwks_url' must be provided if private key JWT client authentication or signed requests is enabled. If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.
- **jwks_url** (String) JSON Web Key Set (JWKS) URL of the OAuth client. Either 'jwks' or 'jwks_url' must be provided if private key JWT client authentication or signed requests is enabled. If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.


<a id="nestedblock--oidc_policy"></a>
### Nested Schema for `oidc_policy`

Optional:

- **grant_access_session_revocation_api** (Boolean) Determines whether this client is allowed to access the Session Revocation API.
- **id_token_content_encryption_algorithm** (String) The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.
- **id_token_encryption_algorithm** (String) The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.
- **id_token_signing_algorithm** (String) The JSON Web Signature [JWS] algorithm required for the ID Token.
- **logout_uris** (List of String) A list of client logout URI's which will be invoked when a user logs out through one of PingFederate's SLO endpoints.
- **pairwise_identifier_user_type** (Boolean) Determines whether the subject identifier type is pairwise.
- **ping_access_logout_capable** (Boolean) Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention.
- **policy_group** (Block List, Max: 1) The Open ID Connect policy. A null value will represent the default policy group. (see [below for nested schema](#nestedblock--oidc_policy--policy_group))
- **sector_identifier_uri** (String) The URI references a file with a single JSON array of Redirect URI and JWKS URL values.

<a id="nestedblock--oidc_policy--policy_group"></a>
### Nested Schema for `oidc_policy.policy_group`

Required:

- **id** (String) The ID of the resource.

Read-Only:

- **location** (String) A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.



<a id="nestedblock--request_policy_ref"></a>
### Nested Schema for `request_policy_ref`

Required:

- **id** (String) The ID of the resource.

Read-Only:

- **location** (String) A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.


<a id="nestedblock--token_exchange_processor_policy_ref"></a>
### Nested Schema for `token_exchange_processor_policy_ref`

Required:

- **id** (String) The ID of the resource.

Read-Only:

- **location** (String) A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.

## Import

Import is supported using the following syntax:

```shell
terraform import pingfederate_oauth_client.example 123
```