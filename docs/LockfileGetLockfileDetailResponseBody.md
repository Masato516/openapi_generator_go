# LockfileGetLockfileDetailResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | created time of lockfile | 
**FileContent** | **string** | FileContent of lockfile | 
**Id** | **int64** | ID of lockfile | 
**LibraryPkgs** | Pointer to [**[]LibraryPkgChildResponseBody**](LibraryPkgChildResponseBody.md) | LibraryPkgs of lockfile | [optional] 
**Path** | **string** | Path of lockfile | 
**Server** | Pointer to [**ServerChildResponseBody**](ServerChildResponseBody.md) |  | [optional] 
**UpdatedAt** | **string** | updated time of lockfile | 

## Methods

### NewLockfileGetLockfileDetailResponseBody

`func NewLockfileGetLockfileDetailResponseBody(createdAt string, fileContent string, id int64, path string, updatedAt string, ) *LockfileGetLockfileDetailResponseBody`

NewLockfileGetLockfileDetailResponseBody instantiates a new LockfileGetLockfileDetailResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockfileGetLockfileDetailResponseBodyWithDefaults

`func NewLockfileGetLockfileDetailResponseBodyWithDefaults() *LockfileGetLockfileDetailResponseBody`

NewLockfileGetLockfileDetailResponseBodyWithDefaults instantiates a new LockfileGetLockfileDetailResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LockfileGetLockfileDetailResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LockfileGetLockfileDetailResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LockfileGetLockfileDetailResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileContent

`func (o *LockfileGetLockfileDetailResponseBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *LockfileGetLockfileDetailResponseBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *LockfileGetLockfileDetailResponseBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetId

`func (o *LockfileGetLockfileDetailResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockfileGetLockfileDetailResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockfileGetLockfileDetailResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLibraryPkgs

`func (o *LockfileGetLockfileDetailResponseBody) GetLibraryPkgs() []LibraryPkgChildResponseBody`

GetLibraryPkgs returns the LibraryPkgs field if non-nil, zero value otherwise.

### GetLibraryPkgsOk

`func (o *LockfileGetLockfileDetailResponseBody) GetLibraryPkgsOk() (*[]LibraryPkgChildResponseBody, bool)`

GetLibraryPkgsOk returns a tuple with the LibraryPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryPkgs

`func (o *LockfileGetLockfileDetailResponseBody) SetLibraryPkgs(v []LibraryPkgChildResponseBody)`

SetLibraryPkgs sets LibraryPkgs field to given value.

### HasLibraryPkgs

`func (o *LockfileGetLockfileDetailResponseBody) HasLibraryPkgs() bool`

HasLibraryPkgs returns a boolean if a field has been set.

### GetPath

`func (o *LockfileGetLockfileDetailResponseBody) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LockfileGetLockfileDetailResponseBody) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LockfileGetLockfileDetailResponseBody) SetPath(v string)`

SetPath sets Path field to given value.


### GetServer

`func (o *LockfileGetLockfileDetailResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LockfileGetLockfileDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LockfileGetLockfileDetailResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.

### HasServer

`func (o *LockfileGetLockfileDetailResponseBody) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LockfileGetLockfileDetailResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LockfileGetLockfileDetailResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LockfileGetLockfileDetailResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


