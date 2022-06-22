# LockfileAddLockfileResponseBody

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

### NewLockfileAddLockfileResponseBody

`func NewLockfileAddLockfileResponseBody(createdAt string, fileContent string, id int64, path string, updatedAt string, ) *LockfileAddLockfileResponseBody`

NewLockfileAddLockfileResponseBody instantiates a new LockfileAddLockfileResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockfileAddLockfileResponseBodyWithDefaults

`func NewLockfileAddLockfileResponseBodyWithDefaults() *LockfileAddLockfileResponseBody`

NewLockfileAddLockfileResponseBodyWithDefaults instantiates a new LockfileAddLockfileResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LockfileAddLockfileResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LockfileAddLockfileResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LockfileAddLockfileResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileContent

`func (o *LockfileAddLockfileResponseBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *LockfileAddLockfileResponseBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *LockfileAddLockfileResponseBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetId

`func (o *LockfileAddLockfileResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockfileAddLockfileResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockfileAddLockfileResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLibraryPkgs

`func (o *LockfileAddLockfileResponseBody) GetLibraryPkgs() []LibraryPkgChildResponseBody`

GetLibraryPkgs returns the LibraryPkgs field if non-nil, zero value otherwise.

### GetLibraryPkgsOk

`func (o *LockfileAddLockfileResponseBody) GetLibraryPkgsOk() (*[]LibraryPkgChildResponseBody, bool)`

GetLibraryPkgsOk returns a tuple with the LibraryPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryPkgs

`func (o *LockfileAddLockfileResponseBody) SetLibraryPkgs(v []LibraryPkgChildResponseBody)`

SetLibraryPkgs sets LibraryPkgs field to given value.

### HasLibraryPkgs

`func (o *LockfileAddLockfileResponseBody) HasLibraryPkgs() bool`

HasLibraryPkgs returns a boolean if a field has been set.

### GetPath

`func (o *LockfileAddLockfileResponseBody) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LockfileAddLockfileResponseBody) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LockfileAddLockfileResponseBody) SetPath(v string)`

SetPath sets Path field to given value.


### GetServer

`func (o *LockfileAddLockfileResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LockfileAddLockfileResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LockfileAddLockfileResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.

### HasServer

`func (o *LockfileAddLockfileResponseBody) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LockfileAddLockfileResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LockfileAddLockfileResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LockfileAddLockfileResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


