# ServerGetServerDetailByUUIDResponseBody

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
**PlatformInstanceId** | **string** | platformInstanceId of server | 
**PlatformName** | **string** | platformName of server | 
**ServerIpv4** | **string** | IPv4 of server | 
**ServerName** | **string** | Name of server | 
**ServerUuid** | **string** | UUID of server | 
**ServerroleId** | **int64** | ID of server role | 
**ServerroleName** | **string** | Name of server role | 
**Tags** | Pointer to [**[]ServerTagResponseBody**](ServerTagResponseBody.md) | tags is list of server tag | [optional] 
**Tasks** | Pointer to [**[]ChildTaskResponseBody**](ChildTaskResponseBody.md) | tasks of server | [optional] 
**UpdatedAt** | **string** | updated time of server | 

## Methods

### NewServerGetServerDetailByUUIDResponseBody

`func NewServerGetServerDetailByUUIDResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, platformInstanceId string, platformName string, serverIpv4 string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string, ) *ServerGetServerDetailByUUIDResponseBody`

NewServerGetServerDetailByUUIDResponseBody instantiates a new ServerGetServerDetailByUUIDResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGetServerDetailByUUIDResponseBodyWithDefaults

`func NewServerGetServerDetailByUUIDResponseBodyWithDefaults() *ServerGetServerDetailByUUIDResponseBody`

NewServerGetServerDetailByUUIDResponseBodyWithDefaults instantiates a new ServerGetServerDetailByUUIDResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultUserId

`func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserId() int64`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserIdOk() (*int64, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *ServerGetServerDetailByUUIDResponseBody) SetDefaultUserId(v int64)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *ServerGetServerDetailByUUIDResponseBody) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDefaultUserName

`func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserName() string`

GetDefaultUserName returns the DefaultUserName field if non-nil, zero value otherwise.

### GetDefaultUserNameOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserNameOk() (*string, bool)`

GetDefaultUserNameOk returns a tuple with the DefaultUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserName

`func (o *ServerGetServerDetailByUUIDResponseBody) SetDefaultUserName(v string)`

SetDefaultUserName sets DefaultUserName field to given value.

### HasDefaultUserName

`func (o *ServerGetServerDetailByUUIDResponseBody) HasDefaultUserName() bool`

HasDefaultUserName returns a boolean if a field has been set.

### GetHostUuid

`func (o *ServerGetServerDetailByUUIDResponseBody) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *ServerGetServerDetailByUUIDResponseBody) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.


### GetId

`func (o *ServerGetServerDetailByUUIDResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerGetServerDetailByUUIDResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLastScannedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) GetLastScannedAt() string`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetLastScannedAtOk() (*string, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) SetLastScannedAt(v string)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetLastUploadedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) GetLastUploadedAt() string`

GetLastUploadedAt returns the LastUploadedAt field if non-nil, zero value otherwise.

### GetLastUploadedAtOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetLastUploadedAtOk() (*string, bool)`

GetLastUploadedAtOk returns a tuple with the LastUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUploadedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) SetLastUploadedAt(v string)`

SetLastUploadedAt sets LastUploadedAt field to given value.

### HasLastUploadedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) HasLastUploadedAt() bool`

HasLastUploadedAt returns a boolean if a field has been set.

### GetNeedKernelRestart

`func (o *ServerGetServerDetailByUUIDResponseBody) GetNeedKernelRestart() bool`

GetNeedKernelRestart returns the NeedKernelRestart field if non-nil, zero value otherwise.

### GetNeedKernelRestartOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetNeedKernelRestartOk() (*bool, bool)`

GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedKernelRestart

`func (o *ServerGetServerDetailByUUIDResponseBody) SetNeedKernelRestart(v bool)`

SetNeedKernelRestart sets NeedKernelRestart field to given value.


### GetOsFamily

`func (o *ServerGetServerDetailByUUIDResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerGetServerDetailByUUIDResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *ServerGetServerDetailByUUIDResponseBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerGetServerDetailByUUIDResponseBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetPlatformInstanceId

`func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformInstanceId() string`

GetPlatformInstanceId returns the PlatformInstanceId field if non-nil, zero value otherwise.

### GetPlatformInstanceIdOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformInstanceIdOk() (*string, bool)`

GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInstanceId

`func (o *ServerGetServerDetailByUUIDResponseBody) SetPlatformInstanceId(v string)`

SetPlatformInstanceId sets PlatformInstanceId field to given value.


### GetPlatformName

`func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *ServerGetServerDetailByUUIDResponseBody) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetServerIpv4

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerIpv4() string`

GetServerIpv4 returns the ServerIpv4 field if non-nil, zero value otherwise.

### GetServerIpv4Ok

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerIpv4Ok() (*string, bool)`

GetServerIpv4Ok returns a tuple with the ServerIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpv4

`func (o *ServerGetServerDetailByUUIDResponseBody) SetServerIpv4(v string)`

SetServerIpv4 sets ServerIpv4 field to given value.


### GetServerName

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerGetServerDetailByUUIDResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerUuid

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *ServerGetServerDetailByUUIDResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetServerroleId

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleId() int64`

GetServerroleId returns the ServerroleId field if non-nil, zero value otherwise.

### GetServerroleIdOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleIdOk() (*int64, bool)`

GetServerroleIdOk returns a tuple with the ServerroleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleId

`func (o *ServerGetServerDetailByUUIDResponseBody) SetServerroleId(v int64)`

SetServerroleId sets ServerroleId field to given value.


### GetServerroleName

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleName() string`

GetServerroleName returns the ServerroleName field if non-nil, zero value otherwise.

### GetServerroleNameOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleNameOk() (*string, bool)`

GetServerroleNameOk returns a tuple with the ServerroleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleName

`func (o *ServerGetServerDetailByUUIDResponseBody) SetServerroleName(v string)`

SetServerroleName sets ServerroleName field to given value.


### GetTags

`func (o *ServerGetServerDetailByUUIDResponseBody) GetTags() []ServerTagResponseBody`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetTagsOk() (*[]ServerTagResponseBody, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerGetServerDetailByUUIDResponseBody) SetTags(v []ServerTagResponseBody)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerGetServerDetailByUUIDResponseBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *ServerGetServerDetailByUUIDResponseBody) GetTasks() []ChildTaskResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetTasksOk() (*[]ChildTaskResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ServerGetServerDetailByUUIDResponseBody) SetTasks(v []ChildTaskResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ServerGetServerDetailByUUIDResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerGetServerDetailByUUIDResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerGetServerDetailByUUIDResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


