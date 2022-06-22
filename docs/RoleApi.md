# \RoleApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDeleteRole**](RoleApi.md#RoleDeleteRole) | **Delete** /v1/role/{roleID} | deleteRole role
[**RoleGetRoleDetail**](RoleApi.md#RoleGetRoleDetail) | **Get** /v1/role/{roleID} | getRoleDetail role
[**RoleGetRoleList**](RoleApi.md#RoleGetRoleList) | **Get** /v1/roles | getRoleList role
[**RoleUpdateRole**](RoleApi.md#RoleUpdateRole) | **Put** /v1/role/{roleID} | updateRole role



## RoleDeleteRole

> RoleDeleteRole(ctx, roleID).Authorization(authorization).Execute()

deleteRole role



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
    roleID := int64(789) // int64 | Role ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleApi.RoleDeleteRole(context.Background(), roleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.RoleDeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **int64** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoleDeleteRoleRequest struct via the builder pattern


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


## RoleGetRoleDetail

> RoleGetRoleDetailResponseBody RoleGetRoleDetail(ctx, roleID).Authorization(authorization).Execute()

getRoleDetail role



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
    roleID := int64(789) // int64 | Role ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleApi.RoleGetRoleDetail(context.Background(), roleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.RoleGetRoleDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoleGetRoleDetail`: RoleGetRoleDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.RoleGetRoleDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **int64** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoleGetRoleDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**RoleGetRoleDetailResponseBody**](RoleGetRoleDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleGetRoleList

> RoleGetRoleListResponseBody RoleGetRoleList(ctx).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).Authorization(authorization).Execute()

getRoleList role



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
    page := int32(56) // int32 | Page Number (default: 1) (optional) (default to 1)
    limit := int32(56) // int32 | Limit (default: 20, max: 100) (optional) (default to 20)
    offset := int32(56) // int32 | Offset (optional) (default to 0)
    filterCveID := "filterCveID_example" // string | CveID filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleApi.RoleGetRoleList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.RoleGetRoleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoleGetRoleList`: RoleGetRoleListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.RoleGetRoleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleGetRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number (default: 1) | [default to 1]
 **limit** | **int32** | Limit (default: 20, max: 100) | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterCveID** | **string** | CveID filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**RoleGetRoleListResponseBody**](RoleGetRoleListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleUpdateRole

> RoleUpdateRoleResponseBody RoleUpdateRole(ctx, roleID).UpdateRoleRequestBody(updateRoleRequestBody).Authorization(authorization).Execute()

updateRole role



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
    roleID := int64(789) // int64 | Role ID
    updateRoleRequestBody := *openapiclient.NewRoleUpdateRoleRequestBody() // RoleUpdateRoleRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleApi.RoleUpdateRole(context.Background(), roleID).UpdateRoleRequestBody(updateRoleRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.RoleUpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoleUpdateRole`: RoleUpdateRoleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.RoleUpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **int64** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoleUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleRequestBody** | [**RoleUpdateRoleRequestBody**](RoleUpdateRoleRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**RoleUpdateRoleResponseBody**](RoleUpdateRoleResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

