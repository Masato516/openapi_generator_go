# \ServerApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServerCreatePkgPasteServer**](ServerApi.md#ServerCreatePkgPasteServer) | **Post** /v1/server/paste | createPkgPasteServer server
[**ServerDeleteServer**](ServerApi.md#ServerDeleteServer) | **Delete** /v1/server/{serverID} | deleteServer server
[**ServerGetServerDetail**](ServerApi.md#ServerGetServerDetail) | **Get** /v1/server/{serverID} | getServerDetail server
[**ServerGetServerDetailByUUID**](ServerApi.md#ServerGetServerDetailByUUID) | **Get** /v1/server/uuid/{serverUuid} | getServerDetailByUUID server
[**ServerGetServerList**](ServerApi.md#ServerGetServerList) | **Get** /v1/servers | getServerList server
[**ServerUpdatePkgPasteServer**](ServerApi.md#ServerUpdatePkgPasteServer) | **Put** /v1/server/paste/{serverID} | updatePkgPasteServer server
[**ServerUpdateServer**](ServerApi.md#ServerUpdateServer) | **Put** /v1/server/{serverID} | updateServer server



## ServerCreatePkgPasteServer

> ServerCreatePkgPasteServerResponseBody ServerCreatePkgPasteServer(ctx).CreatePkgPasteServerRequestBody(createPkgPasteServerRequestBody).Authorization(authorization).Execute()

createPkgPasteServer server



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
    createPkgPasteServerRequestBody := *openapiclient.NewServerCreatePkgPasteServerRequestBody("kernel release", "20", "20", "PkgPasteText_example", "Server Name") // ServerCreatePkgPasteServerRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerCreatePkgPasteServer(context.Background()).CreatePkgPasteServerRequestBody(createPkgPasteServerRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerCreatePkgPasteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerCreatePkgPasteServer`: ServerCreatePkgPasteServerResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ServerCreatePkgPasteServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerCreatePkgPasteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPkgPasteServerRequestBody** | [**ServerCreatePkgPasteServerRequestBody**](ServerCreatePkgPasteServerRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**ServerCreatePkgPasteServerResponseBody**](ServerCreatePkgPasteServerResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerDeleteServer

> ServerDeleteServer(ctx, serverID).Authorization(authorization).Execute()

deleteServer server



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
    serverID := int64(789) // int64 | Server ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerDeleteServer(context.Background(), serverID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerDeleteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverID** | **int64** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerDeleteServerRequest struct via the builder pattern


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


## ServerGetServerDetail

> ServerGetServerDetailResponseBody ServerGetServerDetail(ctx, serverID).Authorization(authorization).Execute()

getServerDetail server



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
    serverID := int64(789) // int64 | Server ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerGetServerDetail(context.Background(), serverID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerGetServerDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGetServerDetail`: ServerGetServerDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ServerGetServerDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverID** | **int64** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerGetServerDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**ServerGetServerDetailResponseBody**](ServerGetServerDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGetServerDetailByUUID

> ServerGetServerDetailByUUIDResponseBody ServerGetServerDetailByUUID(ctx, serverUuid).Authorization(authorization).Execute()

getServerDetailByUUID server



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
    serverUuid := "serverUuid_example" // string | Server UUID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerGetServerDetailByUUID(context.Background(), serverUuid).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerGetServerDetailByUUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGetServerDetailByUUID`: ServerGetServerDetailByUUIDResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ServerGetServerDetailByUUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverUuid** | **string** | Server UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerGetServerDetailByUUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**ServerGetServerDetailByUUIDResponseBody**](ServerGetServerDetailByUUIDResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGetServerList

> ServerGetServerListResponseBody ServerGetServerList(ctx).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).FilterRoleID(filterRoleID).FilterTagName(filterTagName).Authorization(authorization).Execute()

getServerList server



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
    filterRoleID := int32(56) // int32 | ServerRoleID filter (optional)
    filterTagName := "filterTagName_example" // string | ServerTagName filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerGetServerList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterCveID(filterCveID).FilterRoleID(filterRoleID).FilterTagName(filterTagName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerGetServerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGetServerList`: ServerGetServerListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ServerGetServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGetServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number | [default to 1]
 **limit** | **int32** | Limit | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterCveID** | **string** | CveID filter | 
 **filterRoleID** | **int32** | ServerRoleID filter | 
 **filterTagName** | **string** | ServerTagName filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**ServerGetServerListResponseBody**](ServerGetServerListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerUpdatePkgPasteServer

> ServerUpdatePkgPasteServer(ctx, serverID).UpdatePkgPasteServerRequestBody(updatePkgPasteServerRequestBody).Authorization(authorization).Execute()

updatePkgPasteServer server



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
    serverID := int64(789) // int64 | Server ID
    updatePkgPasteServerRequestBody := *openapiclient.NewServerUpdatePkgPasteServerRequestBody("PkgPasteText_example") // ServerUpdatePkgPasteServerRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerUpdatePkgPasteServer(context.Background(), serverID).UpdatePkgPasteServerRequestBody(updatePkgPasteServerRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerUpdatePkgPasteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverID** | **int64** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerUpdatePkgPasteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePkgPasteServerRequestBody** | [**ServerUpdatePkgPasteServerRequestBody**](ServerUpdatePkgPasteServerRequestBody.md) |  | 
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


## ServerUpdateServer

> ServerUpdateServerResponseBody ServerUpdateServer(ctx, serverID).UpdateServerRequestBody(updateServerRequestBody).Authorization(authorization).Execute()

updateServer server



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
    serverID := int64(789) // int64 | Server ID
    updateServerRequestBody := *openapiclient.NewServerUpdateServerRequestBody() // ServerUpdateServerRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ServerUpdateServer(context.Background(), serverID).UpdateServerRequestBody(updateServerRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ServerUpdateServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerUpdateServer`: ServerUpdateServerResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ServerUpdateServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverID** | **int64** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerUpdateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerRequestBody** | [**ServerUpdateServerRequestBody**](ServerUpdateServerRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**ServerUpdateServerResponseBody**](ServerUpdateServerResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

