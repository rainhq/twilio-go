# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialAws**](DefaultApi.md#CreateCredentialAws) | **Post** /v1/Credentials/AWS | 
[**CreateCredentialPublicKey**](DefaultApi.md#CreateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys | 
[**CreateSecondaryAuthToken**](DefaultApi.md#CreateSecondaryAuthToken) | **Post** /v1/AuthTokens/Secondary | 
[**DeleteCredentialAws**](DefaultApi.md#DeleteCredentialAws) | **Delete** /v1/Credentials/AWS/{Sid} | 
[**DeleteCredentialPublicKey**](DefaultApi.md#DeleteCredentialPublicKey) | **Delete** /v1/Credentials/PublicKeys/{Sid} | 
[**DeleteSecondaryAuthToken**](DefaultApi.md#DeleteSecondaryAuthToken) | **Delete** /v1/AuthTokens/Secondary | 
[**FetchCredentialAws**](DefaultApi.md#FetchCredentialAws) | **Get** /v1/Credentials/AWS/{Sid} | 
[**FetchCredentialPublicKey**](DefaultApi.md#FetchCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys/{Sid} | 
[**ListCredentialAws**](DefaultApi.md#ListCredentialAws) | **Get** /v1/Credentials/AWS | 
[**ListCredentialPublicKey**](DefaultApi.md#ListCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys | 
[**UpdateAuthTokenPromotion**](DefaultApi.md#UpdateAuthTokenPromotion) | **Post** /v1/AuthTokens/Promote | 
[**UpdateCredentialAws**](DefaultApi.md#UpdateCredentialAws) | **Post** /v1/Credentials/AWS/{Sid} | 
[**UpdateCredentialPublicKey**](DefaultApi.md#UpdateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys/{Sid} | 



## CreateCredentialAws

> AccountsV1CredentialCredentialAws CreateCredentialAws(ctx, optional)



Create a new AWS Credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCredentialAwsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialAwsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AccountSid** | **String**| The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request. | 
**Credentials** | **String**| A string that contains the AWS access credentials in the format &#x60;&lt;AWS_ACCESS_KEY_ID&gt;:&lt;AWS_SECRET_ACCESS_KEY&gt;&#x60;. For example, &#x60;AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY&#x60; | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**AccountsV1CredentialCredentialAws**](accounts.v1.credential.credential_aws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey CreateCredentialPublicKey(ctx, optional)



Create a new Public Key Credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCredentialPublicKeyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialPublicKeyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AccountSid** | **String**| The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**PublicKey** | **String**| A URL encoded representation of the public key. For example, &#x60;-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----&#x60; | 

### Return type

[**AccountsV1CredentialCredentialPublicKey**](accounts.v1.credential.credential_public_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecondaryAuthToken

> AccountsV1SecondaryAuthToken CreateSecondaryAuthToken(ctx, )



Create a new secondary Auth Token

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AccountsV1SecondaryAuthToken**](accounts.v1.secondary_auth_token.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialAws

> DeleteCredentialAws(ctx, Sid)



Delete a Credential from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the AWS resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialPublicKey

> DeleteCredentialPublicKey(ctx, Sid)



Delete a Credential from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the PublicKey resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecondaryAuthToken

> DeleteSecondaryAuthToken(ctx, )



Delete the secondary Auth Token from your account

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredentialAws

> AccountsV1CredentialCredentialAws FetchCredentialAws(ctx, Sid)



Fetch the AWS credentials specified by the provided Credential Sid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the AWS resource to fetch. | 

### Return type

[**AccountsV1CredentialCredentialAws**](accounts.v1.credential.credential_aws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey FetchCredentialPublicKey(ctx, Sid)



Fetch the public key specified by the provided Credential Sid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the PublicKey resource to fetch. | 

### Return type

[**AccountsV1CredentialCredentialPublicKey**](accounts.v1.credential.credential_public_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialAws

> ListCredentialAwsResponse ListCredentialAws(ctx, optional)



Retrieves a collection of AWS Credentials belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCredentialAwsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialAwsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCredentialAwsResponse**](ListCredentialAwsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialPublicKey

> ListCredentialPublicKeyResponse ListCredentialPublicKey(ctx, optional)



Retrieves a collection of Public Key Credentials belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCredentialPublicKeyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialPublicKeyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCredentialPublicKeyResponse**](ListCredentialPublicKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthTokenPromotion

> AccountsV1AuthTokenPromotion UpdateAuthTokenPromotion(ctx, )



Promote the secondary Auth Token to primary. After promoting the new token, all requests to Twilio using your old primary Auth Token will result in an error.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AccountsV1AuthTokenPromotion**](accounts.v1.auth_token_promotion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialAws

> AccountsV1CredentialCredentialAws UpdateCredentialAws(ctx, Sid, optional)



Modify the properties of a given Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the AWS resource to update. | 
 **optional** | ***UpdateCredentialAwsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialAwsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**AccountsV1CredentialCredentialAws**](accounts.v1.credential.credential_aws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey UpdateCredentialPublicKey(ctx, Sid, optional)



Modify the properties of a given Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the PublicKey resource to update. | 
 **optional** | ***UpdateCredentialPublicKeyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialPublicKeyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**AccountsV1CredentialCredentialPublicKey**](accounts.v1.credential.credential_public_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
