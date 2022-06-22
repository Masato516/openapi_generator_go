# RoleGetRoleListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingResponseBody**](PagingResponseBody.md) |  | [optional] 
**Roles** | Pointer to [**[]RoleListResponseBody**](RoleListResponseBody.md) | ServerRole list | [optional] 

## Methods

### NewRoleGetRoleListResponseBody

`func NewRoleGetRoleListResponseBody() *RoleGetRoleListResponseBody`

NewRoleGetRoleListResponseBody instantiates a new RoleGetRoleListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleGetRoleListResponseBodyWithDefaults

`func NewRoleGetRoleListResponseBodyWithDefaults() *RoleGetRoleListResponseBody`

NewRoleGetRoleListResponseBodyWithDefaults instantiates a new RoleGetRoleListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *RoleGetRoleListResponseBody) GetPaging() PagingResponseBody`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *RoleGetRoleListResponseBody) GetPagingOk() (*PagingResponseBody, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *RoleGetRoleListResponseBody) SetPaging(v PagingResponseBody)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *RoleGetRoleListResponseBody) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetRoles

`func (o *RoleGetRoleListResponseBody) GetRoles() []RoleListResponseBody`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleGetRoleListResponseBody) GetRolesOk() (*[]RoleListResponseBody, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleGetRoleListResponseBody) SetRoles(v []RoleListResponseBody)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RoleGetRoleListResponseBody) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


