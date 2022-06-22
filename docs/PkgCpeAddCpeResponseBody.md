# PkgCpeAddCpeResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeURI** | Pointer to **string** | Cpe URI of cpe | [optional] 
**CreatedAt** | **time.Time** | crated time of package or cpe | 
**Id** | **int64** | ID of package or cpe | 
**Name** | **string** | Name of package or cpe | 
**Release** | Pointer to **string** | Release of package | [optional] 
**Server** | [**ServerChildResponseBody**](ServerChildResponseBody.md) |  | 
**Tasks** | Pointer to [**[]ChildTaskResponseBody**](ChildTaskResponseBody.md) | updated time of server | [optional] 
**UpdatedAt** | **time.Time** | updated time of package or cpe | 
**Version** | **string** | Version of package or cpe | 

## Methods

### NewPkgCpeAddCpeResponseBody

`func NewPkgCpeAddCpeResponseBody(createdAt time.Time, id int64, name string, server ServerChildResponseBody, updatedAt time.Time, version string, ) *PkgCpeAddCpeResponseBody`

NewPkgCpeAddCpeResponseBody instantiates a new PkgCpeAddCpeResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeAddCpeResponseBodyWithDefaults

`func NewPkgCpeAddCpeResponseBodyWithDefaults() *PkgCpeAddCpeResponseBody`

NewPkgCpeAddCpeResponseBodyWithDefaults instantiates a new PkgCpeAddCpeResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeURI

`func (o *PkgCpeAddCpeResponseBody) GetCpeURI() string`

GetCpeURI returns the CpeURI field if non-nil, zero value otherwise.

### GetCpeURIOk

`func (o *PkgCpeAddCpeResponseBody) GetCpeURIOk() (*string, bool)`

GetCpeURIOk returns a tuple with the CpeURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeURI

`func (o *PkgCpeAddCpeResponseBody) SetCpeURI(v string)`

SetCpeURI sets CpeURI field to given value.

### HasCpeURI

`func (o *PkgCpeAddCpeResponseBody) HasCpeURI() bool`

HasCpeURI returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PkgCpeAddCpeResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkgCpeAddCpeResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkgCpeAddCpeResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PkgCpeAddCpeResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkgCpeAddCpeResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkgCpeAddCpeResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PkgCpeAddCpeResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgCpeAddCpeResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgCpeAddCpeResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *PkgCpeAddCpeResponseBody) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgCpeAddCpeResponseBody) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgCpeAddCpeResponseBody) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PkgCpeAddCpeResponseBody) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServer

`func (o *PkgCpeAddCpeResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PkgCpeAddCpeResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PkgCpeAddCpeResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.


### GetTasks

`func (o *PkgCpeAddCpeResponseBody) GetTasks() []ChildTaskResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PkgCpeAddCpeResponseBody) GetTasksOk() (*[]ChildTaskResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PkgCpeAddCpeResponseBody) SetTasks(v []ChildTaskResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PkgCpeAddCpeResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PkgCpeAddCpeResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkgCpeAddCpeResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkgCpeAddCpeResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *PkgCpeAddCpeResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgCpeAddCpeResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgCpeAddCpeResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


