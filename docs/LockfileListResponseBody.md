# LockfileListResponseBody

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

### NewLockfileListResponseBody

`func NewLockfileListResponseBody(createdAt string, fileContent string, id int64, path string, updatedAt string, ) *LockfileListResponseBody`

NewLockfileListResponseBody instantiates a new LockfileListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockfileListResponseBodyWithDefaults

`func NewLockfileListResponseBodyWithDefaults() *LockfileListResponseBody`

NewLockfileListResponseBodyWithDefaults instantiates a new LockfileListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LockfileListResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LockfileListResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LockfileListResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileContent

`func (o *LockfileListResponseBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *LockfileListResponseBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *LockfileListResponseBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetId

`func (o *LockfileListResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockfileListResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockfileListResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetLibraryPkgs

`func (o *LockfileListResponseBody) GetLibraryPkgs() []LibraryPkgChildResponseBody`

GetLibraryPkgs returns the LibraryPkgs field if non-nil, zero value otherwise.

### GetLibraryPkgsOk

`func (o *LockfileListResponseBody) GetLibraryPkgsOk() (*[]LibraryPkgChildResponseBody, bool)`

GetLibraryPkgsOk returns a tuple with the LibraryPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryPkgs

`func (o *LockfileListResponseBody) SetLibraryPkgs(v []LibraryPkgChildResponseBody)`

SetLibraryPkgs sets LibraryPkgs field to given value.

### HasLibraryPkgs

`func (o *LockfileListResponseBody) HasLibraryPkgs() bool`

HasLibraryPkgs returns a boolean if a field has been set.

### GetPath

`func (o *LockfileListResponseBody) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LockfileListResponseBody) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LockfileListResponseBody) SetPath(v string)`

SetPath sets Path field to given value.


### GetServer

`func (o *LockfileListResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LockfileListResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LockfileListResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.

### HasServer

`func (o *LockfileListResponseBody) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LockfileListResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LockfileListResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LockfileListResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


