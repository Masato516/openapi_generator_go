# PkgCpeGetPkgCpeListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 
**PkgCpes** | Pointer to [**[]PkgCpeListResponseBody**](PkgCpeListResponseBody.md) | PkgCpes list | [optional] 

## Methods

### NewPkgCpeGetPkgCpeListResponseBody

`func NewPkgCpeGetPkgCpeListResponseBody() *PkgCpeGetPkgCpeListResponseBody`

NewPkgCpeGetPkgCpeListResponseBody instantiates a new PkgCpeGetPkgCpeListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeGetPkgCpeListResponseBodyWithDefaults

`func NewPkgCpeGetPkgCpeListResponseBodyWithDefaults() *PkgCpeGetPkgCpeListResponseBody`

NewPkgCpeGetPkgCpeListResponseBodyWithDefaults instantiates a new PkgCpeGetPkgCpeListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *PkgCpeGetPkgCpeListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PkgCpeGetPkgCpeListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PkgCpeGetPkgCpeListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PkgCpeGetPkgCpeListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetPkgCpes

`func (o *PkgCpeGetPkgCpeListResponseBody) GetPkgCpes() []PkgCpeListResponseBody`

GetPkgCpes returns the PkgCpes field if non-nil, zero value otherwise.

### GetPkgCpesOk

`func (o *PkgCpeGetPkgCpeListResponseBody) GetPkgCpesOk() (*[]PkgCpeListResponseBody, bool)`

GetPkgCpesOk returns a tuple with the PkgCpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgCpes

`func (o *PkgCpeGetPkgCpeListResponseBody) SetPkgCpes(v []PkgCpeListResponseBody)`

SetPkgCpes sets PkgCpes field to given value.

### HasPkgCpes

`func (o *PkgCpeGetPkgCpeListResponseBody) HasPkgCpes() bool`

HasPkgCpes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


