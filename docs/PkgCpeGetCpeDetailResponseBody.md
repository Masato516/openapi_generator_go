# PkgCpeGetCpeDetailResponseBody

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

### NewPkgCpeGetCpeDetailResponseBody

`func NewPkgCpeGetCpeDetailResponseBody(createdAt time.Time, id int64, name string, server ServerChildResponseBody, updatedAt time.Time, version string, ) *PkgCpeGetCpeDetailResponseBody`

NewPkgCpeGetCpeDetailResponseBody instantiates a new PkgCpeGetCpeDetailResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeGetCpeDetailResponseBodyWithDefaults

`func NewPkgCpeGetCpeDetailResponseBodyWithDefaults() *PkgCpeGetCpeDetailResponseBody`

NewPkgCpeGetCpeDetailResponseBodyWithDefaults instantiates a new PkgCpeGetCpeDetailResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeURI

`func (o *PkgCpeGetCpeDetailResponseBody) GetCpeURI() string`

GetCpeURI returns the CpeURI field if non-nil, zero value otherwise.

### GetCpeURIOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetCpeURIOk() (*string, bool)`

GetCpeURIOk returns a tuple with the CpeURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeURI

`func (o *PkgCpeGetCpeDetailResponseBody) SetCpeURI(v string)`

SetCpeURI sets CpeURI field to given value.

### HasCpeURI

`func (o *PkgCpeGetCpeDetailResponseBody) HasCpeURI() bool`

HasCpeURI returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PkgCpeGetCpeDetailResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkgCpeGetCpeDetailResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PkgCpeGetCpeDetailResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkgCpeGetCpeDetailResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PkgCpeGetCpeDetailResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgCpeGetCpeDetailResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *PkgCpeGetCpeDetailResponseBody) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgCpeGetCpeDetailResponseBody) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PkgCpeGetCpeDetailResponseBody) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServer

`func (o *PkgCpeGetCpeDetailResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PkgCpeGetCpeDetailResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.


### GetTasks

`func (o *PkgCpeGetCpeDetailResponseBody) GetTasks() []ChildTaskResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetTasksOk() (*[]ChildTaskResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PkgCpeGetCpeDetailResponseBody) SetTasks(v []ChildTaskResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PkgCpeGetCpeDetailResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PkgCpeGetCpeDetailResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkgCpeGetCpeDetailResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *PkgCpeGetCpeDetailResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgCpeGetCpeDetailResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgCpeGetCpeDetailResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


