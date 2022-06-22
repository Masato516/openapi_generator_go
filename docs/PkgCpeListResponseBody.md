# PkgCpeListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeID** | Pointer to **int64** | CpeID of cpe | [optional] 
**CpeURI** | Pointer to **string** | Cpe URI of cpe | [optional] 
**CreatedAt** | **time.Time** | crated time of package or cpe | 
**Name** | **string** | Name of package or cpe | 
**NeedRestartProcs** | Pointer to [**[]NeedRestartProcResponseBody**](NeedRestartProcResponseBody.md) | NeedRestartProcess list of package | [optional] 
**NewRelease** | Pointer to **string** | New Release of package | [optional] 
**NewVersion** | Pointer to **string** | New Version of package | [optional] 
**NotFixedYet** | Pointer to **bool** | Flag of Not fixed yet of package | [optional] 
**PkgID** | Pointer to **int64** | Package ID of package | [optional] 
**Release** | Pointer to **string** | Release of package | [optional] 
**Repository** | Pointer to **string** | Repository of package | [optional] 
**ServerID** | **int64** | ServerID of package or cpe | 
**ServerName** | **string** | ServerName of package or cpe | 
**ServerUuid** | **string** | ServerUUID of package or cpe | 
**UpdatedAt** | **time.Time** | updated time of package or cpe | 
**Version** | **string** | Version of package or cpe | 

## Methods

### NewPkgCpeListResponseBody

`func NewPkgCpeListResponseBody(createdAt time.Time, name string, serverID int64, serverName string, serverUuid string, updatedAt time.Time, version string, ) *PkgCpeListResponseBody`

NewPkgCpeListResponseBody instantiates a new PkgCpeListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeListResponseBodyWithDefaults

`func NewPkgCpeListResponseBodyWithDefaults() *PkgCpeListResponseBody`

NewPkgCpeListResponseBodyWithDefaults instantiates a new PkgCpeListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeID

`func (o *PkgCpeListResponseBody) GetCpeID() int64`

GetCpeID returns the CpeID field if non-nil, zero value otherwise.

### GetCpeIDOk

`func (o *PkgCpeListResponseBody) GetCpeIDOk() (*int64, bool)`

GetCpeIDOk returns a tuple with the CpeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeID

`func (o *PkgCpeListResponseBody) SetCpeID(v int64)`

SetCpeID sets CpeID field to given value.

### HasCpeID

`func (o *PkgCpeListResponseBody) HasCpeID() bool`

HasCpeID returns a boolean if a field has been set.

### GetCpeURI

`func (o *PkgCpeListResponseBody) GetCpeURI() string`

GetCpeURI returns the CpeURI field if non-nil, zero value otherwise.

### GetCpeURIOk

`func (o *PkgCpeListResponseBody) GetCpeURIOk() (*string, bool)`

GetCpeURIOk returns a tuple with the CpeURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeURI

`func (o *PkgCpeListResponseBody) SetCpeURI(v string)`

SetCpeURI sets CpeURI field to given value.

### HasCpeURI

`func (o *PkgCpeListResponseBody) HasCpeURI() bool`

HasCpeURI returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PkgCpeListResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkgCpeListResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkgCpeListResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *PkgCpeListResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgCpeListResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgCpeListResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetNeedRestartProcs

`func (o *PkgCpeListResponseBody) GetNeedRestartProcs() []NeedRestartProcResponseBody`

GetNeedRestartProcs returns the NeedRestartProcs field if non-nil, zero value otherwise.

### GetNeedRestartProcsOk

`func (o *PkgCpeListResponseBody) GetNeedRestartProcsOk() (*[]NeedRestartProcResponseBody, bool)`

GetNeedRestartProcsOk returns a tuple with the NeedRestartProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedRestartProcs

`func (o *PkgCpeListResponseBody) SetNeedRestartProcs(v []NeedRestartProcResponseBody)`

SetNeedRestartProcs sets NeedRestartProcs field to given value.

### HasNeedRestartProcs

`func (o *PkgCpeListResponseBody) HasNeedRestartProcs() bool`

HasNeedRestartProcs returns a boolean if a field has been set.

### GetNewRelease

`func (o *PkgCpeListResponseBody) GetNewRelease() string`

GetNewRelease returns the NewRelease field if non-nil, zero value otherwise.

### GetNewReleaseOk

`func (o *PkgCpeListResponseBody) GetNewReleaseOk() (*string, bool)`

GetNewReleaseOk returns a tuple with the NewRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelease

`func (o *PkgCpeListResponseBody) SetNewRelease(v string)`

SetNewRelease sets NewRelease field to given value.

### HasNewRelease

`func (o *PkgCpeListResponseBody) HasNewRelease() bool`

HasNewRelease returns a boolean if a field has been set.

### GetNewVersion

`func (o *PkgCpeListResponseBody) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *PkgCpeListResponseBody) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *PkgCpeListResponseBody) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *PkgCpeListResponseBody) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetNotFixedYet

`func (o *PkgCpeListResponseBody) GetNotFixedYet() bool`

GetNotFixedYet returns the NotFixedYet field if non-nil, zero value otherwise.

### GetNotFixedYetOk

`func (o *PkgCpeListResponseBody) GetNotFixedYetOk() (*bool, bool)`

GetNotFixedYetOk returns a tuple with the NotFixedYet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFixedYet

`func (o *PkgCpeListResponseBody) SetNotFixedYet(v bool)`

SetNotFixedYet sets NotFixedYet field to given value.

### HasNotFixedYet

`func (o *PkgCpeListResponseBody) HasNotFixedYet() bool`

HasNotFixedYet returns a boolean if a field has been set.

### GetPkgID

`func (o *PkgCpeListResponseBody) GetPkgID() int64`

GetPkgID returns the PkgID field if non-nil, zero value otherwise.

### GetPkgIDOk

`func (o *PkgCpeListResponseBody) GetPkgIDOk() (*int64, bool)`

GetPkgIDOk returns a tuple with the PkgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgID

`func (o *PkgCpeListResponseBody) SetPkgID(v int64)`

SetPkgID sets PkgID field to given value.

### HasPkgID

`func (o *PkgCpeListResponseBody) HasPkgID() bool`

HasPkgID returns a boolean if a field has been set.

### GetRelease

`func (o *PkgCpeListResponseBody) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgCpeListResponseBody) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgCpeListResponseBody) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PkgCpeListResponseBody) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepository

`func (o *PkgCpeListResponseBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PkgCpeListResponseBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PkgCpeListResponseBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PkgCpeListResponseBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetServerID

`func (o *PkgCpeListResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *PkgCpeListResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *PkgCpeListResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetServerName

`func (o *PkgCpeListResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *PkgCpeListResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *PkgCpeListResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerUuid

`func (o *PkgCpeListResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *PkgCpeListResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *PkgCpeListResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetUpdatedAt

`func (o *PkgCpeListResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkgCpeListResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkgCpeListResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *PkgCpeListResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgCpeListResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgCpeListResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


