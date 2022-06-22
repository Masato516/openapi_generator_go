# \CveApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CveGetCveDetail**](CveApi.md#CveGetCveDetail) | **Get** /v1/cve/{cveID} | getCveDetail cve
[**CveGetCveList**](CveApi.md#CveGetCveList) | **Get** /v1/cves | getCveList cve



## CveGetCveDetail

> CveGetCveDetailResponseBody CveGetCveDetail(ctx, cveID).Authorization(authorization).Execute()

getCveDetail cve



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cveID := "cveID_example" // string | Cve ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CveApi.CveGetCveDetail(context.Background(), cveID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CveApi.CveGetCveDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CveGetCveDetail`: CveGetCveDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CveApi.CveGetCveDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveID** | **string** | Cve ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCveGetCveDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**CveGetCveDetailResponseBody**](CveGetCveDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CveGetCveList

> CveGetCveListResponseBody CveGetCveList(ctx).Page(page).Limit(limit).Offset(offset).FilterServerID(filterServerID).FilterRoleID(filterRoleID).FilterPkgID(filterPkgID).FilterCpeID(filterCpeID).Authorization(authorization).Execute()

getCveList cve



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page Number (optional) (default to 1)
    limit := int32(56) // int32 | Limit (optional) (default to 20)
    offset := int32(56) // int32 | Offset (optional) (default to 0)
    filterServerID := int32(56) // int32 | ServerID filter (optional)
    filterRoleID := int32(56) // int32 | ServerRoleID filter (optional)
    filterPkgID := int32(56) // int32 | PackageID filter (optional)
    filterCpeID := int32(56) // int32 | CpeID filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CveApi.CveGetCveList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterServerID(filterServerID).FilterRoleID(filterRoleID).FilterPkgID(filterPkgID).FilterCpeID(filterCpeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CveApi.CveGetCveList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CveGetCveList`: CveGetCveListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CveApi.CveGetCveList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCveGetCveListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number | [default to 1]
 **limit** | **int32** | Limit | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterServerID** | **int32** | ServerID filter | 
 **filterRoleID** | **int32** | ServerRoleID filter | 
 **filterPkgID** | **int32** | PackageID filter | 
 **filterCpeID** | **int32** | CpeID filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**CveGetCveListResponseBody**](CveGetCveListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

