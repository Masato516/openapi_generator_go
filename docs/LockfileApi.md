# \LockfileApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LockfileAddLockfile**](LockfileApi.md#LockfileAddLockfile) | **Post** /v1/lockfile | addLockfile lockfile
[**LockfileDeleteLockfile**](LockfileApi.md#LockfileDeleteLockfile) | **Delete** /v1/lockfile/{lockfileID} | deleteLockfile lockfile
[**LockfileGetLockfileDetail**](LockfileApi.md#LockfileGetLockfileDetail) | **Get** /v1/lockfile/{lockfileID} | getLockfileDetail lockfile
[**LockfileGetLockfileList**](LockfileApi.md#LockfileGetLockfileList) | **Get** /v1/lockfiles | getLockfileList lockfile
[**LockfileUpdateLockfile**](LockfileApi.md#LockfileUpdateLockfile) | **Put** /v1/lockfile/{lockfileID} | updateLockfile lockfile



## LockfileAddLockfile

> LockfileAddLockfileResponseBody LockfileAddLockfile(ctx).AddLockfileRequestBody(addLockfileRequestBody).Authorization(authorization).Execute()

addLockfile lockfile



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
    addLockfileRequestBody := *openapiclient.NewLockfileAddLockfileRequestBody("FileContent_example", "/FutureVuls/package.json", int64(1)) // LockfileAddLockfileRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LockfileApi.LockfileAddLockfile(context.Background()).AddLockfileRequestBody(addLockfileRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockfileApi.LockfileAddLockfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockfileAddLockfile`: LockfileAddLockfileResponseBody
    fmt.Fprintf(os.Stdout, "Response from `LockfileApi.LockfileAddLockfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockfileAddLockfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLockfileRequestBody** | [**LockfileAddLockfileRequestBody**](LockfileAddLockfileRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**LockfileAddLockfileResponseBody**](LockfileAddLockfileResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockfileDeleteLockfile

> LockfileDeleteLockfile(ctx, lockfileID).Authorization(authorization).Execute()

deleteLockfile lockfile



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
    lockfileID := int64(789) // int64 | Lockfile ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LockfileApi.LockfileDeleteLockfile(context.Background(), lockfileID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockfileApi.LockfileDeleteLockfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockfileID** | **int64** | Lockfile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockfileDeleteLockfileRequest struct via the builder pattern


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


## LockfileGetLockfileDetail

> LockfileGetLockfileDetailResponseBody LockfileGetLockfileDetail(ctx, lockfileID).Authorization(authorization).Execute()

getLockfileDetail lockfile



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
    lockfileID := int64(789) // int64 | Lockfile ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LockfileApi.LockfileGetLockfileDetail(context.Background(), lockfileID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockfileApi.LockfileGetLockfileDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockfileGetLockfileDetail`: LockfileGetLockfileDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `LockfileApi.LockfileGetLockfileDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockfileID** | **int64** | Lockfile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockfileGetLockfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**LockfileGetLockfileDetailResponseBody**](LockfileGetLockfileDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockfileGetLockfileList

> LockfileGetLockfileListResponseBody LockfileGetLockfileList(ctx).Page(page).Limit(limit).Offset(offset).FilterServerID(filterServerID).FilterPath(filterPath).Authorization(authorization).Execute()

getLockfileList lockfile



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
    filterServerID := int64(789) // int64 | ServerID filter (optional)
    filterPath := "filterPath_example" // string | Path filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LockfileApi.LockfileGetLockfileList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterServerID(filterServerID).FilterPath(filterPath).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockfileApi.LockfileGetLockfileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockfileGetLockfileList`: LockfileGetLockfileListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `LockfileApi.LockfileGetLockfileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockfileGetLockfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number | [default to 1]
 **limit** | **int32** | Limit | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterServerID** | **int64** | ServerID filter | 
 **filterPath** | **string** | Path filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**LockfileGetLockfileListResponseBody**](LockfileGetLockfileListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockfileUpdateLockfile

> LockfileUpdateLockfileResponseBody LockfileUpdateLockfile(ctx, lockfileID).UpdateLockfileRequestBody(updateLockfileRequestBody).Authorization(authorization).Execute()

updateLockfile lockfile



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
    lockfileID := int64(789) // int64 | Lockfile ID
    updateLockfileRequestBody := *openapiclient.NewLockfileUpdateLockfileRequestBody() // LockfileUpdateLockfileRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LockfileApi.LockfileUpdateLockfile(context.Background(), lockfileID).UpdateLockfileRequestBody(updateLockfileRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockfileApi.LockfileUpdateLockfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockfileUpdateLockfile`: LockfileUpdateLockfileResponseBody
    fmt.Fprintf(os.Stdout, "Response from `LockfileApi.LockfileUpdateLockfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockfileID** | **int64** | Lockfile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockfileUpdateLockfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLockfileRequestBody** | [**LockfileUpdateLockfileRequestBody**](LockfileUpdateLockfileRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**LockfileUpdateLockfileResponseBody**](LockfileUpdateLockfileResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

