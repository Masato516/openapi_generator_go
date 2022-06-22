# ServerChildResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | crated time of server | 
**DefaultUserId** | Pointer to **int64** | default user ID of server | [optional] 
**DefaultUserName** | Pointer to **string** | default user name of server | [optional] 
**HostUuid** | **string** | UUID of server | 
**Id** | **int64** | ID of server | 
**LastScannedAt** | Pointer to **string** | last scanned time of server | [optional] 
**LastUploadedAt** | Pointer to **string** | last uploaded time of server | [optional] 
**NeedKernelRestart** | **bool** | Whether server needs kernel restart | 
**OsFamily** | **string** | OS Name of server | 
**OsVersion** | **string** | OS Version of server | 
**ServerName** | **string** | Name of server | 
**ServerUuid** | **string** | UUID of server | 
**ServerroleId** | **int64** | ID of server role | 
**ServerroleName** | **string** | Name of server role | 
**Tags** | Pointer to **[]string** | tags is list of server tag | [optional] 
**UpdatedAt** | **string** | updated time of server | 

## Methods

### NewServerChildResponseBody

`func NewServerChildResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string, ) *ServerChildResponseBody`

NewServerChildResponseBody instantiates a new ServerChildResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerChildResponseBodyWithDefaults

`func NewServerChildResponseBodyWithDefaults() *ServerChildResponseBody`

NewServerChildResponseBodyWithDefaults instantiates a new ServerChildResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServerChildResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerChildResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerChildResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultUserId

`func (o *ServerChildResponseBody) GetDefaultUserId() int64`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *ServerChildResponseBody) GetDefaultUserIdOk() (*int64, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *ServerChildResponseBody) SetDefaultUserId(v int64)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *ServerChildResponseBody) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDefaultUserName

`func (o *ServerChildResponseBody) GetDefaultUserName() string`

GetDefaultUserName returns the DefaultUserName field if non-nil, zero value otherwise.

### GetDefaultUserNameOk

`func (o *ServerChildResponseBody) GetDefaultUserNameOk() (*string, bool)`

GetDefaultUserNameOk returns a tuple with the DefaultUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserName

`func (o *ServerChildResponseBody) SetDefaultUserName(v string)`

SetDefaultUserName sets DefaultUserName field to given value.

### HasDefaultUserName

`func (o *ServerChildResponseBody) HasDefaultUserName() bool`

HasDefaultUserName returns a boolean if a field has been set.

### GetHostUuid

`func (o *ServerChildResponseBody) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *ServerChildResponseBody) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *ServerChildResponseBody) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.


### GetId

`func (o *ServerChildResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerChildResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerChildResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLastScannedAt

`func (o *ServerChildResponseBody) GetLastScannedAt() string`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ServerChildResponseBody) GetLastScannedAtOk() (*string, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ServerChildResponseBody) SetLastScannedAt(v string)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ServerChildResponseBody) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetLastUploadedAt

`func (o *ServerChildResponseBody) GetLastUploadedAt() string`

GetLastUploadedAt returns the LastUploadedAt field if non-nil, zero value otherwise.

### GetLastUploadedAtOk

`func (o *ServerChildResponseBody) GetLastUploadedAtOk() (*string, bool)`

GetLastUploadedAtOk returns a tuple with the LastUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUploadedAt

`func (o *ServerChildResponseBody) SetLastUploadedAt(v string)`

SetLastUploadedAt sets LastUploadedAt field to given value.

### HasLastUploadedAt

`func (o *ServerChildResponseBody) HasLastUploadedAt() bool`

HasLastUploadedAt returns a boolean if a field has been set.

### GetNeedKernelRestart

`func (o *ServerChildResponseBody) GetNeedKernelRestart() bool`

GetNeedKernelRestart returns the NeedKernelRestart field if non-nil, zero value otherwise.

### GetNeedKernelRestartOk

`func (o *ServerChildResponseBody) GetNeedKernelRestartOk() (*bool, bool)`

GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedKernelRestart

`func (o *ServerChildResponseBody) SetNeedKernelRestart(v bool)`

SetNeedKernelRestart sets NeedKernelRestart field to given value.


### GetOsFamily

`func (o *ServerChildResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerChildResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerChildResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *ServerChildResponseBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerChildResponseBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerChildResponseBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetServerName

`func (o *ServerChildResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerChildResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerChildResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerUuid

`func (o *ServerChildResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *ServerChildResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *ServerChildResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetServerroleId

`func (o *ServerChildResponseBody) GetServerroleId() int64`

GetServerroleId returns the ServerroleId field if non-nil, zero value otherwise.

### GetServerroleIdOk

`func (o *ServerChildResponseBody) GetServerroleIdOk() (*int64, bool)`

GetServerroleIdOk returns a tuple with the ServerroleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleId

`func (o *ServerChildResponseBody) SetServerroleId(v int64)`

SetServerroleId sets ServerroleId field to given value.


### GetServerroleName

`func (o *ServerChildResponseBody) GetServerroleName() string`

GetServerroleName returns the ServerroleName field if non-nil, zero value otherwise.

### GetServerroleNameOk

`func (o *ServerChildResponseBody) GetServerroleNameOk() (*string, bool)`

GetServerroleNameOk returns a tuple with the ServerroleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleName

`func (o *ServerChildResponseBody) SetServerroleName(v string)`

SetServerroleName sets ServerroleName field to given value.


### GetTags

`func (o *ServerChildResponseBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerChildResponseBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerChildResponseBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerChildResponseBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServerChildResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerChildResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerChildResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


