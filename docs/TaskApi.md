# \TaskApi

All URIs are relative to *http://rest.vuls.biz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskAddTaskComment**](TaskApi.md#TaskAddTaskComment) | **Post** /v1/task/{taskID}/comment | addTaskComment task
[**TaskGetTaskDetail**](TaskApi.md#TaskGetTaskDetail) | **Get** /v1/task/{taskID} | getTaskDetail task
[**TaskGetTaskList**](TaskApi.md#TaskGetTaskList) | **Get** /v1/tasks | getTaskList task
[**TaskUpdateTask**](TaskApi.md#TaskUpdateTask) | **Put** /v1/task/{taskID} | updateTask task
[**TaskUpdateTaskIgnore**](TaskApi.md#TaskUpdateTaskIgnore) | **Put** /v1/task/{taskID}/ignore | updateTaskIgnore task



## TaskAddTaskComment

> TaskAddTaskCommentResponseBody TaskAddTaskComment(ctx, taskID).AddTaskCommentRequestBody(addTaskCommentRequestBody).Authorization(authorization).Execute()

addTaskComment task



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
    taskID := int64(789) // int64 | Task ID
    addTaskCommentRequestBody := *openapiclient.NewTaskAddTaskCommentRequestBody("comment") // TaskAddTaskCommentRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.TaskAddTaskComment(context.Background(), taskID).AddTaskCommentRequestBody(addTaskCommentRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskAddTaskComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskAddTaskComment`: TaskAddTaskCommentResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskAddTaskComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int64** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskAddTaskCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTaskCommentRequestBody** | [**TaskAddTaskCommentRequestBody**](TaskAddTaskCommentRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**TaskAddTaskCommentResponseBody**](TaskAddTaskCommentResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGetTaskDetail

> TaskGetTaskDetailResponseBody TaskGetTaskDetail(ctx, taskID).Authorization(authorization).Execute()

getTaskDetail task



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
    taskID := int64(789) // int64 | Task ID
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.TaskGetTaskDetail(context.Background(), taskID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskGetTaskDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskGetTaskDetail`: TaskGetTaskDetailResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskGetTaskDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int64** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGetTaskDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | api key auth | 

### Return type

[**TaskGetTaskDetailResponseBody**](TaskGetTaskDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGetTaskList

> TaskGetTaskListResponseBody TaskGetTaskList(ctx).Page(page).Limit(limit).Offset(offset).FilterStatus(filterStatus).FilterPriority(filterPriority).FilterIgnore(filterIgnore).FilterMainUserIDs(filterMainUserIDs).FilterSubUserIDs(filterSubUserIDs).FilterCveID(filterCveID).FilterServerID(filterServerID).FilterRoleID(filterRoleID).FilterPkgID(filterPkgID).FilterCpeID(filterCpeID).Authorization(authorization).Execute()

getTaskList task



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
    filterStatus := []string{"FilterStatus_example"} // []string | Status filter (optional) (default to ["new","investigating","ongoing"])
    filterPriority := []string{"FilterPriority_example"} // []string | Priority filter (optional)
    filterIgnore := true // bool | Ignore filter(trueの場合は、非表示のものを取得しない。falseの場合は全件取得) (optional)
    filterMainUserIDs := []int32{int32(123)} // []int32 | MainUserIDs filter (optional)
    filterSubUserIDs := []int32{int32(123)} // []int32 | SubUserIDs filter (optional)
    filterCveID := "filterCveID_example" // string | CveID filter (optional)
    filterServerID := int32(56) // int32 | ServerID filter (optional)
    filterRoleID := int32(56) // int32 | ServerRoleID filter (optional)
    filterPkgID := int32(56) // int32 | PackageID filter (optional)
    filterCpeID := int32(56) // int32 | CpeID filter (optional)
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.TaskGetTaskList(context.Background()).Page(page).Limit(limit).Offset(offset).FilterStatus(filterStatus).FilterPriority(filterPriority).FilterIgnore(filterIgnore).FilterMainUserIDs(filterMainUserIDs).FilterSubUserIDs(filterSubUserIDs).FilterCveID(filterCveID).FilterServerID(filterServerID).FilterRoleID(filterRoleID).FilterPkgID(filterPkgID).FilterCpeID(filterCpeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskGetTaskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskGetTaskList`: TaskGetTaskListResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskGetTaskList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskGetTaskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number | [default to 1]
 **limit** | **int32** | Limit | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **filterStatus** | **[]string** | Status filter | [default to [&quot;new&quot;,&quot;investigating&quot;,&quot;ongoing&quot;]]
 **filterPriority** | **[]string** | Priority filter | 
 **filterIgnore** | **bool** | Ignore filter(trueの場合は、非表示のものを取得しない。falseの場合は全件取得) | 
 **filterMainUserIDs** | **[]int32** | MainUserIDs filter | 
 **filterSubUserIDs** | **[]int32** | SubUserIDs filter | 
 **filterCveID** | **string** | CveID filter | 
 **filterServerID** | **int32** | ServerID filter | 
 **filterRoleID** | **int32** | ServerRoleID filter | 
 **filterPkgID** | **int32** | PackageID filter | 
 **filterCpeID** | **int32** | CpeID filter | 
 **authorization** | **string** | api key auth | 

### Return type

[**TaskGetTaskListResponseBody**](TaskGetTaskListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskUpdateTask

> TaskUpdateTaskResponseBody TaskUpdateTask(ctx, taskID).UpdateTaskRequestBody(updateTaskRequestBody).Authorization(authorization).Execute()

updateTask task



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
    taskID := int64(789) // int64 | Task ID
    updateTaskRequestBody := *openapiclient.NewTaskUpdateTaskRequestBody() // TaskUpdateTaskRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.TaskUpdateTask(context.Background(), taskID).UpdateTaskRequestBody(updateTaskRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskUpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskUpdateTask`: TaskUpdateTaskResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskUpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int64** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskRequestBody** | [**TaskUpdateTaskRequestBody**](TaskUpdateTaskRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**TaskUpdateTaskResponseBody**](TaskUpdateTaskResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskUpdateTaskIgnore

> TaskUpdateTaskIgnoreResponseBody TaskUpdateTaskIgnore(ctx, taskID).UpdateTaskIgnoreRequestBody(updateTaskIgnoreRequestBody).Authorization(authorization).Execute()

updateTaskIgnore task



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
    taskID := int64(789) // int64 | Task ID
    updateTaskIgnoreRequestBody := *openapiclient.NewTaskUpdateTaskIgnoreRequestBody("forever") // TaskUpdateTaskIgnoreRequestBody | 
    authorization := "authorization_example" // string | api key auth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.TaskUpdateTaskIgnore(context.Background(), taskID).UpdateTaskIgnoreRequestBody(updateTaskIgnoreRequestBody).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskUpdateTaskIgnore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskUpdateTaskIgnore`: TaskUpdateTaskIgnoreResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskUpdateTaskIgnore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int64** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskUpdateTaskIgnoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskIgnoreRequestBody** | [**TaskUpdateTaskIgnoreRequestBody**](TaskUpdateTaskIgnoreRequestBody.md) |  | 
 **authorization** | **string** | api key auth | 

### Return type

[**TaskUpdateTaskIgnoreResponseBody**](TaskUpdateTaskIgnoreResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/gob
- **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

