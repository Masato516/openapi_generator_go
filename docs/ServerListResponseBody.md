# ServerListResponseBody

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
**SuccessScanCount** | **int64** | successScanCount of server | 
**Tags** | Pointer to [**[]ServerTagResponseBody**](ServerTagResponseBody.md) | tags is list of server tag | [optional] 
**UpdatedAt** | **string** | updated time of server | 

## Methods

### NewServerListResponseBody

`func NewServerListResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, platformInstanceId string, platformName string, serverIpv4 string, serverName string, serverUuid string, serverroleId int64, serverroleName string, successScanCount int64, updatedAt string, ) *ServerListResponseBody`

NewServerListResponseBody instantiates a new ServerListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerListResponseBodyWithDefaults

`func NewServerListResponseBodyWithDefaults() *ServerListResponseBody`

NewServerListResponseBodyWithDefaults instantiates a new ServerListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServerListResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerListResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerListResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultUserId

`func (o *ServerListResponseBody) GetDefaultUserId() int64`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *ServerListResponseBody) GetDefaultUserIdOk() (*int64, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *ServerListResponseBody) SetDefaultUserId(v int64)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *ServerListResponseBody) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDefaultUserName

`func (o *ServerListResponseBody) GetDefaultUserName() string`

GetDefaultUserName returns the DefaultUserName field if non-nil, zero value otherwise.

### GetDefaultUserNameOk

`func (o *ServerListResponseBody) GetDefaultUserNameOk() (*string, bool)`

GetDefaultUserNameOk returns a tuple with the DefaultUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserName

`func (o *ServerListResponseBody) SetDefaultUserName(v string)`

SetDefaultUserName sets DefaultUserName field to given value.

### HasDefaultUserName

`func (o *ServerListResponseBody) HasDefaultUserName() bool`

HasDefaultUserName returns a boolean if a field has been set.

### GetHostUuid

`func (o *ServerListResponseBody) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *ServerListResponseBody) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *ServerListResponseBody) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.


### GetId

`func (o *ServerListResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerListResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerListResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLastScannedAt

`func (o *ServerListResponseBody) GetLastScannedAt() string`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ServerListResponseBody) GetLastScannedAtOk() (*string, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ServerListResponseBody) SetLastScannedAt(v string)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ServerListResponseBody) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetLastUploadedAt

`func (o *ServerListResponseBody) GetLastUploadedAt() string`

GetLastUploadedAt returns the LastUploadedAt field if non-nil, zero value otherwise.

### GetLastUploadedAtOk

`func (o *ServerListResponseBody) GetLastUploadedAtOk() (*string, bool)`

GetLastUploadedAtOk returns a tuple with the LastUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUploadedAt

`func (o *ServerListResponseBody) SetLastUploadedAt(v string)`

SetLastUploadedAt sets LastUploadedAt field to given value.

### HasLastUploadedAt

`func (o *ServerListResponseBody) HasLastUploadedAt() bool`

HasLastUploadedAt returns a boolean if a field has been set.

### GetNeedKernelRestart

`func (o *ServerListResponseBody) GetNeedKernelRestart() bool`

GetNeedKernelRestart returns the NeedKernelRestart field if non-nil, zero value otherwise.

### GetNeedKernelRestartOk

`func (o *ServerListResponseBody) GetNeedKernelRestartOk() (*bool, bool)`

GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedKernelRestart

`func (o *ServerListResponseBody) SetNeedKernelRestart(v bool)`

SetNeedKernelRestart sets NeedKernelRestart field to given value.


### GetOsFamily

`func (o *ServerListResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerListResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerListResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *ServerListResponseBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerListResponseBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerListResponseBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetPlatformInstanceId

`func (o *ServerListResponseBody) GetPlatformInstanceId() string`

GetPlatformInstanceId returns the PlatformInstanceId field if non-nil, zero value otherwise.

### GetPlatformInstanceIdOk

`func (o *ServerListResponseBody) GetPlatformInstanceIdOk() (*string, bool)`

GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInstanceId

`func (o *ServerListResponseBody) SetPlatformInstanceId(v string)`

SetPlatformInstanceId sets PlatformInstanceId field to given value.


### GetPlatformName

`func (o *ServerListResponseBody) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *ServerListResponseBody) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *ServerListResponseBody) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetServerIpv4

`func (o *ServerListResponseBody) GetServerIpv4() string`

GetServerIpv4 returns the ServerIpv4 field if non-nil, zero value otherwise.

### GetServerIpv4Ok

`func (o *ServerListResponseBody) GetServerIpv4Ok() (*string, bool)`

GetServerIpv4Ok returns a tuple with the ServerIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpv4

`func (o *ServerListResponseBody) SetServerIpv4(v string)`

SetServerIpv4 sets ServerIpv4 field to given value.


### GetServerName

`func (o *ServerListResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerListResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerListResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerUuid

`func (o *ServerListResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *ServerListResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *ServerListResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetServerroleId

`func (o *ServerListResponseBody) GetServerroleId() int64`

GetServerroleId returns the ServerroleId field if non-nil, zero value otherwise.

### GetServerroleIdOk

`func (o *ServerListResponseBody) GetServerroleIdOk() (*int64, bool)`

GetServerroleIdOk returns a tuple with the ServerroleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleId

`func (o *ServerListResponseBody) SetServerroleId(v int64)`

SetServerroleId sets ServerroleId field to given value.


### GetServerroleName

`func (o *ServerListResponseBody) GetServerroleName() string`

GetServerroleName returns the ServerroleName field if non-nil, zero value otherwise.

### GetServerroleNameOk

`func (o *ServerListResponseBody) GetServerroleNameOk() (*string, bool)`

GetServerroleNameOk returns a tuple with the ServerroleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerroleName

`func (o *ServerListResponseBody) SetServerroleName(v string)`

SetServerroleName sets ServerroleName field to given value.


### GetSuccessScanCount

`func (o *ServerListResponseBody) GetSuccessScanCount() int64`

GetSuccessScanCount returns the SuccessScanCount field if non-nil, zero value otherwise.

### GetSuccessScanCountOk

`func (o *ServerListResponseBody) GetSuccessScanCountOk() (*int64, bool)`

GetSuccessScanCountOk returns a tuple with the SuccessScanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessScanCount

`func (o *ServerListResponseBody) SetSuccessScanCount(v int64)`

SetSuccessScanCount sets SuccessScanCount field to given value.


### GetTags

`func (o *ServerListResponseBody) GetTags() []ServerTagResponseBody`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerListResponseBody) GetTagsOk() (*[]ServerTagResponseBody, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerListResponseBody) SetTags(v []ServerTagResponseBody)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerListResponseBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServerListResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerListResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerListResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


