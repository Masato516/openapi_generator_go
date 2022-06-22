# TaskGetTaskListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 
**Tasks** | Pointer to [**[]TaskListResponseBody**](TaskListResponseBody.md) | Task list | [optional] 

## Methods

### NewTaskGetTaskListResponseBody

`func NewTaskGetTaskListResponseBody() *TaskGetTaskListResponseBody`

NewTaskGetTaskListResponseBody instantiates a new TaskGetTaskListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGetTaskListResponseBodyWithDefaults

`func NewTaskGetTaskListResponseBodyWithDefaults() *TaskGetTaskListResponseBody`

NewTaskGetTaskListResponseBodyWithDefaults instantiates a new TaskGetTaskListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *TaskGetTaskListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *TaskGetTaskListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *TaskGetTaskListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *TaskGetTaskListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetTasks

`func (o *TaskGetTaskListResponseBody) GetTasks() []TaskListResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskGetTaskListResponseBody) GetTasksOk() (*[]TaskListResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskGetTaskListResponseBody) SetTasks(v []TaskListResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TaskGetTaskListResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


