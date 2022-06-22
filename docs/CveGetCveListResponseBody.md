# CveGetCveListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to [**[]CveListResponseBody**](CveListResponseBody.md) | Cves list | [optional] 
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 

## Methods

### NewCveGetCveListResponseBody

`func NewCveGetCveListResponseBody() *CveGetCveListResponseBody`

NewCveGetCveListResponseBody instantiates a new CveGetCveListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveGetCveListResponseBodyWithDefaults

`func NewCveGetCveListResponseBodyWithDefaults() *CveGetCveListResponseBody`

NewCveGetCveListResponseBodyWithDefaults instantiates a new CveGetCveListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *CveGetCveListResponseBody) GetCves() []CveListResponseBody`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *CveGetCveListResponseBody) GetCvesOk() (*[]CveListResponseBody, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *CveGetCveListResponseBody) SetCves(v []CveListResponseBody)`

SetCves sets Cves field to given value.

### HasCves

`func (o *CveGetCveListResponseBody) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetPaging

`func (o *CveGetCveListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CveGetCveListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CveGetCveListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CveGetCveListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


