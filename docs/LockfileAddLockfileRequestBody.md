# LockfileAddLockfileRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileContent** | **string** | fileContent of the lockfile | 
**Path** | **string** | Path of lockfile | 
**ServerID** | **int64** | ServerID | 

## Methods

### NewLockfileAddLockfileRequestBody

`func NewLockfileAddLockfileRequestBody(fileContent string, path string, serverID int64, ) *LockfileAddLockfileRequestBody`

NewLockfileAddLockfileRequestBody instantiates a new LockfileAddLockfileRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockfileAddLockfileRequestBodyWithDefaults

`func NewLockfileAddLockfileRequestBodyWithDefaults() *LockfileAddLockfileRequestBody`

NewLockfileAddLockfileRequestBodyWithDefaults instantiates a new LockfileAddLockfileRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileContent

`func (o *LockfileAddLockfileRequestBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *LockfileAddLockfileRequestBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *LockfileAddLockfileRequestBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetPath

`func (o *LockfileAddLockfileRequestBody) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LockfileAddLockfileRequestBody) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LockfileAddLockfileRequestBody) SetPath(v string)`

SetPath sets Path field to given value.


### GetServerID

`func (o *LockfileAddLockfileRequestBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *LockfileAddLockfileRequestBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *LockfileAddLockfileRequestBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


