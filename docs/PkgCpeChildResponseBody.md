# PkgCpeChildResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProcs** | Pointer to [**[]AffectedProcResponseBody**](AffectedProcResponseBody.md) | AffectedProcess list of package | [optional] 
**CpeID** | Pointer to **int64** | CpeID of cpe | [optional] 
**CpeURI** | Pointer to **string** | Cpe URI of cpe | [optional] 
**CreatedAt** | **time.Time** | crated time of package or cpe | 
**Name** | **string** | Name of package or cpe | 
**NewRelease** | Pointer to **string** | New Release of package | [optional] 
**NewVersion** | Pointer to **string** | New Version of package | [optional] 
**PkgID** | Pointer to **int64** | Package ID of package | [optional] 
**Release** | Pointer to **string** | Release of package | [optional] 
**Repository** | Pointer to **string** | Repository of package | [optional] 
**ServerID** | **int64** | ServerID of package or cpe | 
**UpdatedAt** | **time.Time** | updated time of package or cpe | 
**Version** | **string** | Version of package or cpe | 

## Methods

### NewPkgCpeChildResponseBody

`func NewPkgCpeChildResponseBody(createdAt time.Time, name string, serverID int64, updatedAt time.Time, version string, ) *PkgCpeChildResponseBody`

NewPkgCpeChildResponseBody instantiates a new PkgCpeChildResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeChildResponseBodyWithDefaults

`func NewPkgCpeChildResponseBodyWithDefaults() *PkgCpeChildResponseBody`

NewPkgCpeChildResponseBodyWithDefaults instantiates a new PkgCpeChildResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProcs

`func (o *PkgCpeChildResponseBody) GetAffectedProcs() []AffectedProcResponseBody`

GetAffectedProcs returns the AffectedProcs field if non-nil, zero value otherwise.

### GetAffectedProcsOk

`func (o *PkgCpeChildResponseBody) GetAffectedProcsOk() (*[]AffectedProcResponseBody, bool)`

GetAffectedProcsOk returns a tuple with the AffectedProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcs

`func (o *PkgCpeChildResponseBody) SetAffectedProcs(v []AffectedProcResponseBody)`

SetAffectedProcs sets AffectedProcs field to given value.

### HasAffectedProcs

`func (o *PkgCpeChildResponseBody) HasAffectedProcs() bool`

HasAffectedProcs returns a boolean if a field has been set.

### GetCpeID

`func (o *PkgCpeChildResponseBody) GetCpeID() int64`

GetCpeID returns the CpeID field if non-nil, zero value otherwise.

### GetCpeIDOk

`func (o *PkgCpeChildResponseBody) GetCpeIDOk() (*int64, bool)`

GetCpeIDOk returns a tuple with the CpeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeID

`func (o *PkgCpeChildResponseBody) SetCpeID(v int64)`

SetCpeID sets CpeID field to given value.

### HasCpeID

`func (o *PkgCpeChildResponseBody) HasCpeID() bool`

HasCpeID returns a boolean if a field has been set.

### GetCpeURI

`func (o *PkgCpeChildResponseBody) GetCpeURI() string`

GetCpeURI returns the CpeURI field if non-nil, zero value otherwise.

### GetCpeURIOk

`func (o *PkgCpeChildResponseBody) GetCpeURIOk() (*string, bool)`

GetCpeURIOk returns a tuple with the CpeURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeURI

`func (o *PkgCpeChildResponseBody) SetCpeURI(v string)`

SetCpeURI sets CpeURI field to given value.

### HasCpeURI

`func (o *PkgCpeChildResponseBody) HasCpeURI() bool`

HasCpeURI returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PkgCpeChildResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkgCpeChildResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkgCpeChildResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *PkgCpeChildResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgCpeChildResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgCpeChildResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetNewRelease

`func (o *PkgCpeChildResponseBody) GetNewRelease() string`

GetNewRelease returns the NewRelease field if non-nil, zero value otherwise.

### GetNewReleaseOk

`func (o *PkgCpeChildResponseBody) GetNewReleaseOk() (*string, bool)`

GetNewReleaseOk returns a tuple with the NewRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelease

`func (o *PkgCpeChildResponseBody) SetNewRelease(v string)`

SetNewRelease sets NewRelease field to given value.

### HasNewRelease

`func (o *PkgCpeChildResponseBody) HasNewRelease() bool`

HasNewRelease returns a boolean if a field has been set.

### GetNewVersion

`func (o *PkgCpeChildResponseBody) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *PkgCpeChildResponseBody) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *PkgCpeChildResponseBody) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *PkgCpeChildResponseBody) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPkgID

`func (o *PkgCpeChildResponseBody) GetPkgID() int64`

GetPkgID returns the PkgID field if non-nil, zero value otherwise.

### GetPkgIDOk

`func (o *PkgCpeChildResponseBody) GetPkgIDOk() (*int64, bool)`

GetPkgIDOk returns a tuple with the PkgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgID

`func (o *PkgCpeChildResponseBody) SetPkgID(v int64)`

SetPkgID sets PkgID field to given value.

### HasPkgID

`func (o *PkgCpeChildResponseBody) HasPkgID() bool`

HasPkgID returns a boolean if a field has been set.

### GetRelease

`func (o *PkgCpeChildResponseBody) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgCpeChildResponseBody) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgCpeChildResponseBody) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PkgCpeChildResponseBody) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepository

`func (o *PkgCpeChildResponseBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PkgCpeChildResponseBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PkgCpeChildResponseBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PkgCpeChildResponseBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetServerID

`func (o *PkgCpeChildResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *PkgCpeChildResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *PkgCpeChildResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetUpdatedAt

`func (o *PkgCpeChildResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkgCpeChildResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkgCpeChildResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *PkgCpeChildResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgCpeChildResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgCpeChildResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


