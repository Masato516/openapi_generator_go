# ServerUpdateServerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **string** | Additional information of the server | [optional] 
**DefaultUserID** | Pointer to **int64** | DefaultUserID of server | [optional] 
**RoleID** | Pointer to **int64** | ServerRoleID of server | [optional] 
**ServerName** | Pointer to **string** | ServerName of server | [optional] 

## Methods

### NewServerUpdateServerRequestBody

`func NewServerUpdateServerRequestBody() *ServerUpdateServerRequestBody`

NewServerUpdateServerRequestBody instantiates a new ServerUpdateServerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerUpdateServerRequestBodyWithDefaults

`func NewServerUpdateServerRequestBodyWithDefaults() *ServerUpdateServerRequestBody`

NewServerUpdateServerRequestBodyWithDefaults instantiates a new ServerUpdateServerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ServerUpdateServerRequestBody) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ServerUpdateServerRequestBody) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ServerUpdateServerRequestBody) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ServerUpdateServerRequestBody) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetDefaultUserID

`func (o *ServerUpdateServerRequestBody) GetDefaultUserID() int64`

GetDefaultUserID returns the DefaultUserID field if non-nil, zero value otherwise.

### GetDefaultUserIDOk

`func (o *ServerUpdateServerRequestBody) GetDefaultUserIDOk() (*int64, bool)`

GetDefaultUserIDOk returns a tuple with the DefaultUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserID

`func (o *ServerUpdateServerRequestBody) SetDefaultUserID(v int64)`

SetDefaultUserID sets DefaultUserID field to given value.

### HasDefaultUserID

`func (o *ServerUpdateServerRequestBody) HasDefaultUserID() bool`

HasDefaultUserID returns a boolean if a field has been set.

### GetRoleID

`func (o *ServerUpdateServerRequestBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *ServerUpdateServerRequestBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *ServerUpdateServerRequestBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.

### HasRoleID

`func (o *ServerUpdateServerRequestBody) HasRoleID() bool`

HasRoleID returns a boolean if a field has been set.

### GetServerName

`func (o *ServerUpdateServerRequestBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerUpdateServerRequestBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerUpdateServerRequestBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ServerUpdateServerRequestBody) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


