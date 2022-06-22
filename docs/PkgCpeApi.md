# \PkgCpeApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PkgCpeAddCpe**](PkgCpeApi.md#PkgCpeAddCpe) | **Post** /v1/pkgCpe/cpe | addCpe pkgCpe
[**PkgCpeDeleteCpe**](PkgCpeApi.md#PkgCpeDeleteCpe) | **Delete** /v1/pkgCpe/cpe/{cpeID} | deleteCpe pkgCpe
[**PkgCpeDeleteCpeDeprecated**](PkgCpeApi.md#PkgCpeDeleteCpeDeprecated) | **Delete** /v1/pkgCpe/cpe | deleteCpe_deprecated pkgCpe
[**PkgCpeGetCpeDetail**](PkgCpeApi.md#PkgCpeGetCpeDetail) | **Get** /v1/pkgCpe/cpe/{cpeID} | getCpeDetail pkgCpe
[**PkgCpeGetPkgCpeList**](PkgCpeApi.md#PkgCpeGetPkgCpeList) | **Get** /v1/pkgCpes | getPkgCpeList pkgCpe
[**PkgCpeGetPkgDetail**](PkgCpeApi.md#PkgCpeGetPkgDetail) | **Get** /v1/pkgCpe/pkg/{pkgID} | getPkgDetail pkgCpe



## PkgCpeAddCpe

> PkgCpeAddCpeResponseBody PkgCpeAddCpe(ctx).AddCpeRequestBody(addCpeRequestBody).Authorization(authorization).Execute()

addCpe pkgCpe



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
    addCpeRequestBody := *openapiclient.NewPkgCpeAddCpeRequestBody("cpe:/a:berlios:discussion_forum_2k:3.3", int64(1)) // PkgCpeAddCpeRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeAddCpe(context.Background()).AddCpeRequestBody(addCpeRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeAddCpe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PkgCpeAddCpe`: PkgCpeAddCpeResponseBody
    fmt.Fprintf(os.Stdout, "Response from `PkgCpeApi.PkgCpeAddCpe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeAddCpeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCpeRequestBody** | [**PkgCpeAddCpeRequestBody**](PkgCpeAddCpeRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**PkgCpeAddCpeResponseBody**](PkgCpeAddCpeResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkgCpeDeleteCpe

> PkgCpeDeleteCpe(ctx, cpeID).Authorization(authorization).Execute()

deleteCpe pkgCpe



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
    cpeID := int64(789) // int64 | cpe ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeDeleteCpe(context.Background(), cpeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeDeleteCpe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cpeID** | **int64** | cpe ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeDeleteCpeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkgCpeDeleteCpeDeprecated

> PkgCpeDeleteCpeDeprecated(ctx).DeleteCpeDeprecatedRequestBody(deleteCpeDeprecatedRequestBody).Authorization(authorization).Execute()

deleteCpe_deprecated pkgCpe



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
    deleteCpeDeprecatedRequestBody := *openapiclient.NewPkgCpeDeleteCpeDeprecatedRequestBody(int64(4046142736569201742)) // PkgCpeDeleteCpeDeprecatedRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeDeleteCpeDeprecated(context.Background()).DeleteCpeDeprecatedRequestBody(deleteCpeDeprecatedRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeDeleteCpeDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeDeleteCpeDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCpeDeprecatedRequestBody** | [**PkgCpeDeleteCpeDeprecatedRequestBody**](PkgCpeDeleteCpeDeprecatedRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkgCpeGetCpeDetail

> PkgCpeGetCpeDetailResponseBody PkgCpeGetCpeDetail(ctx, cpeID).Authorization(authorization).Execute()

getCpeDetail pkgCpe



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
    cpeID := int64(789) // int64 | cpe ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeGetCpeDetail(context.Background(), cpeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeGetCpeDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PkgCpeGetCpeDetail`: PkgCpeGetCpeDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `PkgCpeApi.PkgCpeGetCpeDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cpeID** | **int64** | cpe ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeGetCpeDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**PkgCpeGetCpeDetailResponseBody**](PkgCpeGetCpeDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkgCpeGetPkgCpeList

> PkgCpeGetPkgCpeListResponseBody PkgCpeGetPkgCpeList(ctx).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).FilterTaskID(filterTaskID).FilterServerID(filterServerID).FilterRoleID(filterRoleID).Authorization(authorization).Execute()

getPkgCpeList pkgCpe



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
    filterCveID := "filterCveID_example" // string | CveID filter (optional)
    filterTaskID := int32(56) // int32 | TaskID filter (optional)
    filterServerID := int32(56) // int32 | ServerID filter (optional)
    filterRoleID := int32(56) // int32 | ServerRoleID filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeGetPkgCpeList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).FilterTaskID(filterTaskID).FilterServerID(filterServerID).FilterRoleID(filterRoleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeGetPkgCpeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PkgCpeGetPkgCpeList`: PkgCpeGetPkgCpeListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `PkgCpeApi.PkgCpeGetPkgCpeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeGetPkgCpeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number | [default to 1]
 **limit** | **int32** | Limit | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterCveID** | **string** | CveID filter | 
 **filterTaskID** | **int32** | TaskID filter | 
 **filterServerID** | **int32** | ServerID filter | 
 **filterRoleID** | **int32** | ServerRoleID filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**PkgCpeGetPkgCpeListResponseBody**](PkgCpeGetPkgCpeListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkgCpeGetPkgDetail

> PkgCpeGetPkgDetailResponseBody PkgCpeGetPkgDetail(ctx, pkgID).Authorization(authorization).Execute()

getPkgDetail pkgCpe



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
    pkgID := int64(789) // int64 | PackageID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PkgCpeApi.PkgCpeGetPkgDetail(context.Background(), pkgID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PkgCpeApi.PkgCpeGetPkgDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PkgCpeGetPkgDetail`: PkgCpeGetPkgDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `PkgCpeApi.PkgCpeGetPkgDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkgID** | **int64** | PackageID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPkgCpeGetPkgDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**PkgCpeGetPkgDetailResponseBody**](PkgCpeGetPkgDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

