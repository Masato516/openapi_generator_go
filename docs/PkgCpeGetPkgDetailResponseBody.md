# PkgCpeGetPkgDetailResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProcs** | Pointer to [**[]AffectedProcResponseBody**](AffectedProcResponseBody.md) | AffectedProcess list of package | [optional] 
**CreatedAt** | **time.Time** | crated time of package or cpe | 
**Id** | Pointer to **int64** | ID of package | [optional] 
**Name** | **string** | Name of package or cpe | 
**NeedRestartProcs** | Pointer to [**[]NeedRestartProcResponseBody**](NeedRestartProcResponseBody.md) | NeedRestartProcess list of package | [optional] 
**NewRelease** | Pointer to **string** | New Release of package | [optional] 
**NewVersion** | Pointer to **string** | New Version of package | [optional] 
**PackageStatuses** | Pointer to **map[string]string** | package status of package | [optional] 
**Release** | **string** | Release of package | 
**Repository** | Pointer to **string** | Repository of package | [optional] 
**Server** | [**ServerChildResponseBody**](ServerChildResponseBody.md) |  | 
**Tasks** | Pointer to [**[]ChildTaskResponseBody**](ChildTaskResponseBody.md) | updated time of server | [optional] 
**UpdatedAt** | **time.Time** | updated time of package or cpe | 
**Version** | **string** | Version of package or cpe | 

## Methods

### NewPkgCpeGetPkgDetailResponseBody

`func NewPkgCpeGetPkgDetailResponseBody(createdAt time.Time, name string, release string, server ServerChildResponseBody, updatedAt time.Time, version string, ) *PkgCpeGetPkgDetailResponseBody`

NewPkgCpeGetPkgDetailResponseBody instantiates a new PkgCpeGetPkgDetailResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeGetPkgDetailResponseBodyWithDefaults

`func NewPkgCpeGetPkgDetailResponseBodyWithDefaults() *PkgCpeGetPkgDetailResponseBody`

NewPkgCpeGetPkgDetailResponseBodyWithDefaults instantiates a new PkgCpeGetPkgDetailResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProcs

`func (o *PkgCpeGetPkgDetailResponseBody) GetAffectedProcs() []AffectedProcResponseBody`

GetAffectedProcs returns the AffectedProcs field if non-nil, zero value otherwise.

### GetAffectedProcsOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetAffectedProcsOk() (*[]AffectedProcResponseBody, bool)`

GetAffectedProcsOk returns a tuple with the AffectedProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcs

`func (o *PkgCpeGetPkgDetailResponseBody) SetAffectedProcs(v []AffectedProcResponseBody)`

SetAffectedProcs sets AffectedProcs field to given value.

### HasAffectedProcs

`func (o *PkgCpeGetPkgDetailResponseBody) HasAffectedProcs() bool`

HasAffectedProcs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PkgCpeGetPkgDetailResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkgCpeGetPkgDetailResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PkgCpeGetPkgDetailResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkgCpeGetPkgDetailResponseBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PkgCpeGetPkgDetailResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PkgCpeGetPkgDetailResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgCpeGetPkgDetailResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetNeedRestartProcs

`func (o *PkgCpeGetPkgDetailResponseBody) GetNeedRestartProcs() []NeedRestartProcResponseBody`

GetNeedRestartProcs returns the NeedRestartProcs field if non-nil, zero value otherwise.

### GetNeedRestartProcsOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetNeedRestartProcsOk() (*[]NeedRestartProcResponseBody, bool)`

GetNeedRestartProcsOk returns a tuple with the NeedRestartProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedRestartProcs

`func (o *PkgCpeGetPkgDetailResponseBody) SetNeedRestartProcs(v []NeedRestartProcResponseBody)`

SetNeedRestartProcs sets NeedRestartProcs field to given value.

### HasNeedRestartProcs

`func (o *PkgCpeGetPkgDetailResponseBody) HasNeedRestartProcs() bool`

HasNeedRestartProcs returns a boolean if a field has been set.

### GetNewRelease

`func (o *PkgCpeGetPkgDetailResponseBody) GetNewRelease() string`

GetNewRelease returns the NewRelease field if non-nil, zero value otherwise.

### GetNewReleaseOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetNewReleaseOk() (*string, bool)`

GetNewReleaseOk returns a tuple with the NewRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelease

`func (o *PkgCpeGetPkgDetailResponseBody) SetNewRelease(v string)`

SetNewRelease sets NewRelease field to given value.

### HasNewRelease

`func (o *PkgCpeGetPkgDetailResponseBody) HasNewRelease() bool`

HasNewRelease returns a boolean if a field has been set.

### GetNewVersion

`func (o *PkgCpeGetPkgDetailResponseBody) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *PkgCpeGetPkgDetailResponseBody) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *PkgCpeGetPkgDetailResponseBody) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPackageStatuses

`func (o *PkgCpeGetPkgDetailResponseBody) GetPackageStatuses() map[string]string`

GetPackageStatuses returns the PackageStatuses field if non-nil, zero value otherwise.

### GetPackageStatusesOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetPackageStatusesOk() (*map[string]string, bool)`

GetPackageStatusesOk returns a tuple with the PackageStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageStatuses

`func (o *PkgCpeGetPkgDetailResponseBody) SetPackageStatuses(v map[string]string)`

SetPackageStatuses sets PackageStatuses field to given value.

### HasPackageStatuses

`func (o *PkgCpeGetPkgDetailResponseBody) HasPackageStatuses() bool`

HasPackageStatuses returns a boolean if a field has been set.

### GetRelease

`func (o *PkgCpeGetPkgDetailResponseBody) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgCpeGetPkgDetailResponseBody) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetRepository

`func (o *PkgCpeGetPkgDetailResponseBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PkgCpeGetPkgDetailResponseBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PkgCpeGetPkgDetailResponseBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetServer

`func (o *PkgCpeGetPkgDetailResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PkgCpeGetPkgDetailResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.


### GetTasks

`func (o *PkgCpeGetPkgDetailResponseBody) GetTasks() []ChildTaskResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetTasksOk() (*[]ChildTaskResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PkgCpeGetPkgDetailResponseBody) SetTasks(v []ChildTaskResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PkgCpeGetPkgDetailResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PkgCpeGetPkgDetailResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkgCpeGetPkgDetailResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *PkgCpeGetPkgDetailResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgCpeGetPkgDetailResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgCpeGetPkgDetailResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


