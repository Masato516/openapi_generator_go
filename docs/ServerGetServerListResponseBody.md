# ServerGetServerListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerListResponseBody**](ServerListResponseBody.md) | Servers list | [optional] 

## Methods

### NewServerGetServerListResponseBody

`func NewServerGetServerListResponseBody() *ServerGetServerListResponseBody`

NewServerGetServerListResponseBody instantiates a new ServerGetServerListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGetServerListResponseBodyWithDefaults

`func NewServerGetServerListResponseBodyWithDefaults() *ServerGetServerListResponseBody`

NewServerGetServerListResponseBodyWithDefaults instantiates a new ServerGetServerListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *ServerGetServerListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ServerGetServerListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ServerGetServerListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ServerGetServerListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetServers

`func (o *ServerGetServerListResponseBody) GetServers() []ServerListResponseBody`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ServerGetServerListResponseBody) GetServersOk() (*[]ServerListResponseBody, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ServerGetServerListResponseBody) SetServers(v []ServerListResponseBody)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ServerGetServerListResponseBody) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


