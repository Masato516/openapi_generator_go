# ServerCreatePkgPasteServerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KernelRelease** | **string** | Kernel Release | 
**KernelVersion** | Pointer to **string** | Kernel Version | [optional] 
**OsFamily** | **string** | Server OS Family | 
**OsVersion** | **string** | Server OS Version | 
**PkgPasteText** | **string** | pkg paste text | 
**ServerName** | **string** | Server Name | 

## Methods

### NewServerCreatePkgPasteServerRequestBody

`func NewServerCreatePkgPasteServerRequestBody(kernelRelease string, osFamily string, osVersion string, pkgPasteText string, serverName string, ) *ServerCreatePkgPasteServerRequestBody`

NewServerCreatePkgPasteServerRequestBody instantiates a new ServerCreatePkgPasteServerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCreatePkgPasteServerRequestBodyWithDefaults

`func NewServerCreatePkgPasteServerRequestBodyWithDefaults() *ServerCreatePkgPasteServerRequestBody`

NewServerCreatePkgPasteServerRequestBodyWithDefaults instantiates a new ServerCreatePkgPasteServerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKernelRelease

`func (o *ServerCreatePkgPasteServerRequestBody) GetKernelRelease() string`

GetKernelRelease returns the KernelRelease field if non-nil, zero value otherwise.

### GetKernelReleaseOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetKernelReleaseOk() (*string, bool)`

GetKernelReleaseOk returns a tuple with the KernelRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelRelease

`func (o *ServerCreatePkgPasteServerRequestBody) SetKernelRelease(v string)`

SetKernelRelease sets KernelRelease field to given value.


### GetKernelVersion

`func (o *ServerCreatePkgPasteServerRequestBody) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ServerCreatePkgPasteServerRequestBody) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *ServerCreatePkgPasteServerRequestBody) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetOsFamily

`func (o *ServerCreatePkgPasteServerRequestBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerCreatePkgPasteServerRequestBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *ServerCreatePkgPasteServerRequestBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerCreatePkgPasteServerRequestBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetPkgPasteText

`func (o *ServerCreatePkgPasteServerRequestBody) GetPkgPasteText() string`

GetPkgPasteText returns the PkgPasteText field if non-nil, zero value otherwise.

### GetPkgPasteTextOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetPkgPasteTextOk() (*string, bool)`

GetPkgPasteTextOk returns a tuple with the PkgPasteText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgPasteText

`func (o *ServerCreatePkgPasteServerRequestBody) SetPkgPasteText(v string)`

SetPkgPasteText sets PkgPasteText field to given value.


### GetServerName

`func (o *ServerCreatePkgPasteServerRequestBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerCreatePkgPasteServerRequestBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerCreatePkgPasteServerRequestBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


