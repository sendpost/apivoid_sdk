# \DefaultApi

All URIs are relative to *https://endpoint.apivoid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainageV1PayAsYouGoGet**](DefaultApi.md#DomainageV1PayAsYouGoGet) | **Get** /domainage/v1/pay-as-you-go/ | 
[**DomainblV1PayAsYouGoGet**](DefaultApi.md#DomainblV1PayAsYouGoGet) | **Get** /domainbl/v1/pay-as-you-go/ | 
[**SslinfoV1PayAsYouGoGet**](DefaultApi.md#SslinfoV1PayAsYouGoGet) | **Get** /sslinfo/v1/pay-as-you-go/ | 


# **DomainageV1PayAsYouGoGet**
> ModelsAge DomainageV1PayAsYouGoGet(ctx, key, host)


Get domain reputation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| string (your key) | 
  **host** | **string**| domain name | 

### Return type

[**ModelsAge**](models.age.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainblV1PayAsYouGoGet**
> ModelsBlacklist DomainblV1PayAsYouGoGet(ctx, key, host)


Get domain reputation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| string (your key) | 
  **host** | **string**| domain name | 

### Return type

[**ModelsBlacklist**](models.blacklist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SslinfoV1PayAsYouGoGet**
> ModelsSsl SslinfoV1PayAsYouGoGet(ctx, key, host)


Get domain ssl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| string (your key) | 
  **host** | **string**| domain name | 

### Return type

[**ModelsSsl**](models.ssl.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

