# EnvMetricV2ResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdp** | **string** | CDP of envMetricV2 | 
**CreatedAt** | **time.Time** | created time | 
**CveID** | **string** | CveID of envMetricV2 | 
**RoleID** | **int64** | ServerRoleID of envMetricV2 | 
**RoleName** | **string** | ServerRoleName of envMetricV2 | 
**Td** | **string** | TD of envMetricV2 | 
**UpdatedAt** | **time.Time** | updated time | 

## Methods

### NewEnvMetricV2ResponseBody

`func NewEnvMetricV2ResponseBody(cdp string, createdAt time.Time, cveID string, roleID int64, roleName string, td string, updatedAt time.Time, ) *EnvMetricV2ResponseBody`

NewEnvMetricV2ResponseBody instantiates a new EnvMetricV2ResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvMetricV2ResponseBodyWithDefaults

`func NewEnvMetricV2ResponseBodyWithDefaults() *EnvMetricV2ResponseBody`

NewEnvMetricV2ResponseBodyWithDefaults instantiates a new EnvMetricV2ResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdp

`func (o *EnvMetricV2ResponseBody) GetCdp() string`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *EnvMetricV2ResponseBody) GetCdpOk() (*string, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *EnvMetricV2ResponseBody) SetCdp(v string)`

SetCdp sets Cdp field to given value.


### GetCreatedAt

`func (o *EnvMetricV2ResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvMetricV2ResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvMetricV2ResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *EnvMetricV2ResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *EnvMetricV2ResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *EnvMetricV2ResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetRoleID

`func (o *EnvMetricV2ResponseBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *EnvMetricV2ResponseBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *EnvMetricV2ResponseBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.


### GetRoleName

`func (o *EnvMetricV2ResponseBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *EnvMetricV2ResponseBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *EnvMetricV2ResponseBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetTd

`func (o *EnvMetricV2ResponseBody) GetTd() string`

GetTd returns the Td field if non-nil, zero value otherwise.

### GetTdOk

`func (o *EnvMetricV2ResponseBody) GetTdOk() (*string, bool)`

GetTdOk returns a tuple with the Td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTd

`func (o *EnvMetricV2ResponseBody) SetTd(v string)`

SetTd sets Td field to given value.


### GetUpdatedAt

`func (o *EnvMetricV2ResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvMetricV2ResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvMetricV2ResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


