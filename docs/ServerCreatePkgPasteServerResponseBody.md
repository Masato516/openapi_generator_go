# ServerCreatePkgPasteServerResponseBody

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

### NewServerCreatePkgPasteServerResponseBody

`func NewServerCreatePkgPasteServerResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, platformInstanceId string, platformName string, serverIpv4 string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string, ) *ServerCreatePkgPasteServerResponseBody`

NewServerCreatePkgPasteServerResponseBody instantiates a new ServerCreatePkgPasteServerResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCreatePkgPasteServerResponseBodyWithDefaults

`func NewServerCreatePkgPasteServerResponseBodyWithDefaults() *ServerCreatePkgPasteServerResponseBody`

NewServerCreatePkgPasteServerResponseBodyWithDefaults instantiates a new ServerCreatePkgPasteServerResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServerCreatePkgPasteServerResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerCreatePkgPasteServerResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultUserId

`func (o *ServerCreatePkgPasteServerResponseBody) GetDefaultUserId() int64`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetDefaultUserIdOk() (*int64, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *ServerCreatePkgPasteServerResponseBody) SetDefaultUserId(v int64)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *ServerCreatePkgPasteServerResponseBody) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDefaultUserName

`func (o *ServerCreatePkgPasteServerResponseBody) GetDefaultUserName() string`

GetDefaultUserName returns the DefaultUserName field if non-nil, zero value otherwise.

### GetDefaultUserNameOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetDefaultUserNameOk() (*string, bool)`

GetDefaultUserNameOk returns a tuple with the DefaultUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserName

`func (o *ServerCreatePkgPasteServerResponseBody) SetDefaultUserName(v string)`

SetDefaultUserName sets DefaultUserName field to given value.

### HasDefaultUserName

`func (o *ServerCreatePkgPasteServerResponseBody) HasDefaultUserName() bool`

HasDefaultUserName returns a boolean if a field has been set.

### GetHostUuid

`func (o *ServerCreatePkgPasteServerResponseBody) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *ServerCreatePkgPasteServerResponseBody) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.


### GetId

`func (o *ServerCreatePkgPasteServerResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerCreatePkgPasteServerResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLastScannedAt

`func (o *ServerCreatePkgPasteServerResponseBody) GetLastScannedAt() string`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetLastScannedAtOk() (*string, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ServerCreatePkgPasteServerResponseBody) SetLastScannedAt(v string)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ServerCreatePkgPasteServerResponseBody) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetLastUploadedAt

`func (o *ServerCreatePkgPasteServerResponseBody) GetLastUploadedAt() string`

GetLastUploadedAt returns the LastUploadedAt field if non-nil, zero value otherwise.

### GetLastUploadedAtOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetLastUploadedAtOk() (*string, bool)`

GetLastUploadedAtOk returns a tuple with the LastUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUploadedAt

`func (o *ServerCreatePkgPasteServerResponseBody) SetLastUploadedAt(v string)`

SetLastUploadedAt sets LastUploadedAt field to given value.

### HasLastUploadedAt

`func (o *ServerCreatePkgPasteServerResponseBody) HasLastUploadedAt() bool`

HasLastUploadedAt returns a boolean if a field has been set.

### GetNeedKernelRestart

`func (o *ServerCreatePkgPasteServerResponseBody) GetNeedKernelRestart() bool`

GetNeedKernelRestart returns the NeedKernelRestart field if non-nil, zero value otherwise.

### GetNeedKernelRestartOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetNeedKernelRestartOk() (*bool, bool)`

GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedKernelRestart

`func (o *ServerCreatePkgPasteServerResponseBody) SetNeedKernelRestart(v bool)`

SetNeedKernelRestart sets NeedKernelRestart field to given value.


### GetOsFamily

`func (o *ServerCreatePkgPasteServerResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerCreatePkgPasteServerResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *ServerCreatePkgPasteServerResponseBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerCreatePkgPasteServerResponseBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetPlatformInstanceId

`func (o *ServerCreatePkgPasteServerResponseBody) GetPlatformInstanceId() string`

GetPlatformInstanceId returns the PlatformInstanceId field if non-nil, zero value otherwise.

### GetPlatformInstanceIdOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetPlatformInstanceIdOk() (*string, bool)`

GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInstanceId

`func (o *ServerCreatePkgPasteServerResponseBody) SetPlatformInstanceId(v string)`

SetPlatformInstanceId sets PlatformInstanceId field to given value.


### GetPlatformName

`func (o *ServerCreatePkgPasteServerResponseBody) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *ServerCreatePkgPasteServerResponseBody) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetServerIpv4

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerIpv4() string`

GetServerIpv4 returns the ServerIpv4 field if non-nil, zero value otherwise.

### GetServerIpv4Ok

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerIpv4Ok() (*string, bool)`

GetServerIpv4Ok returns a tuple with the ServerIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpv4

`func (o *ServerCreatePkgPasteServerResponseBody) SetServerIpv4(v string)`

SetServerIpv4 sets ServerIpv4 field to given value.


### GetServerName

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerCreatePkgPasteServerResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerUuid

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *ServerCreatePkgPasteServerResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetServerroleId

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerroleId() int64`

GetServerroleId returns the ServerroleId field if non-nil, zero value otherwise.

### GetServerroleIdOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerroleIdOk() (*int64, bool)`

GetServerroleIdOk returns a tuple with the ServerroleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleId

`func (o *ServerCreatePkgPasteServerResponseBody) SetServerroleId(v int64)`

SetServerroleId sets ServerroleId field to given value.


### GetServerroleName

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerroleName() string`

GetServerroleName returns the ServerroleName field if non-nil, zero value otherwise.

### GetServerroleNameOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetServerroleNameOk() (*string, bool)`

GetServerroleNameOk returns a tuple with the ServerroleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleName

`func (o *ServerCreatePkgPasteServerResponseBody) SetServerroleName(v string)`

SetServerroleName sets ServerroleName field to given value.


### GetTags

`func (o *ServerCreatePkgPasteServerResponseBody) GetTags() []ServerTagResponseBody`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetTagsOk() (*[]ServerTagResponseBody, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerCreatePkgPasteServerResponseBody) SetTags(v []ServerTagResponseBody)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerCreatePkgPasteServerResponseBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *ServerCreatePkgPasteServerResponseBody) GetTasks() []ChildTaskResponseBody`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetTasksOk() (*[]ChildTaskResponseBody, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ServerCreatePkgPasteServerResponseBody) SetTasks(v []ChildTaskResponseBody)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ServerCreatePkgPasteServerResponseBody) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServerCreatePkgPasteServerResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerCreatePkgPasteServerResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerCreatePkgPasteServerResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


