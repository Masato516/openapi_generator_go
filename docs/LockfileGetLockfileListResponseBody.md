# LockfileGetLockfileListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lockfiles** | Pointer to [**[]LockfileListResponseBody**](LockfileListResponseBody.md) | Lockfiles list | [optional] 
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 

## Methods

### NewLockfileGetLockfileListResponseBody

`func NewLockfileGetLockfileListResponseBody() *LockfileGetLockfileListResponseBody`

NewLockfileGetLockfileListResponseBody instantiates a new LockfileGetLockfileListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockfileGetLockfileListResponseBodyWithDefaults

`func NewLockfileGetLockfileListResponseBodyWithDefaults() *LockfileGetLockfileListResponseBody`

NewLockfileGetLockfileListResponseBodyWithDefaults instantiates a new LockfileGetLockfileListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockfiles

`func (o *LockfileGetLockfileListResponseBody) GetLockfiles() []LockfileListResponseBody`

GetLockfiles returns the Lockfiles field if non-nil, zero value otherwise.

### GetLockfilesOk

`func (o *LockfileGetLockfileListResponseBody) GetLockfilesOk() (*[]LockfileListResponseBody, bool)`

GetLockfilesOk returns a tuple with the Lockfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockfiles

`func (o *LockfileGetLockfileListResponseBody) SetLockfiles(v []LockfileListResponseBody)`

SetLockfiles sets Lockfiles field to given value.

### HasLockfiles

`func (o *LockfileGetLockfileListResponseBody) HasLockfiles() bool`

HasLockfiles returns a boolean if a field has been set.

### GetPaging

`func (o *LockfileGetLockfileListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *LockfileGetLockfileListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *LockfileGetLockfileListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *LockfileGetLockfileListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


