# RoleUpdateRoleResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTaskCount** | Pointer to **int64** | AllTaskCount of server role | [optional] 
**CreatedAt** | **time.Time** | created time of server role | 
**EnvMetricV2s** | Pointer to [**[]EnvMetricV2ResponseBody**](EnvMetricV2ResponseBody.md) | envMetricV2s of server role | [optional] 
**EnvMetricV3s** | Pointer to [**[]EnvMetricV3ResponseBody**](EnvMetricV3ResponseBody.md) | envMetricV3s of server role | [optional] 
**Id** | **int64** | ID of server role | 
**IsDefault** | Pointer to **bool** | isDefault of server role | [optional] 
**Name** | **string** | Name of server role | 
**NewTaskCount** | Pointer to **int64** | NewTaskCount of server role | [optional] 
**SecMetric** | Pointer to [**SecMetricResponseBody**](SecMetricResponseBody.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerChildResponseBody**](ServerChildResponseBody.md) | Servers of server role | [optional] 
**UpdatedAt** | **time.Time** | updated time of server role | 

## Methods

### NewRoleUpdateRoleResponseBody

`func NewRoleUpdateRoleResponseBody(createdAt time.Time, id int64, name string, updatedAt time.Time, ) *RoleUpdateRoleResponseBody`

NewRoleUpdateRoleResponseBody instantiates a new RoleUpdateRoleResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateRoleResponseBodyWithDefaults

`func NewRoleUpdateRoleResponseBodyWithDefaults() *RoleUpdateRoleResponseBody`

NewRoleUpdateRoleResponseBodyWithDefaults instantiates a new RoleUpdateRoleResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTaskCount

`func (o *RoleUpdateRoleResponseBody) GetAllTaskCount() int64`

GetAllTaskCount returns the AllTaskCount field if non-nil, zero value otherwise.

### GetAllTaskCountOk

`func (o *RoleUpdateRoleResponseBody) GetAllTaskCountOk() (*int64, bool)`

GetAllTaskCountOk returns a tuple with the AllTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTaskCount

`func (o *RoleUpdateRoleResponseBody) SetAllTaskCount(v int64)`

SetAllTaskCount sets AllTaskCount field to given value.

### HasAllTaskCount

`func (o *RoleUpdateRoleResponseBody) HasAllTaskCount() bool`

HasAllTaskCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleUpdateRoleResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleUpdateRoleResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleUpdateRoleResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvMetricV2s

`func (o *RoleUpdateRoleResponseBody) GetEnvMetricV2s() []EnvMetricV2ResponseBody`

GetEnvMetricV2s returns the EnvMetricV2s field if non-nil, zero value otherwise.

### GetEnvMetricV2sOk

`func (o *RoleUpdateRoleResponseBody) GetEnvMetricV2sOk() (*[]EnvMetricV2ResponseBody, bool)`

GetEnvMetricV2sOk returns a tuple with the EnvMetricV2s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvMetricV2s

`func (o *RoleUpdateRoleResponseBody) SetEnvMetricV2s(v []EnvMetricV2ResponseBody)`

SetEnvMetricV2s sets EnvMetricV2s field to given value.

### HasEnvMetricV2s

`func (o *RoleUpdateRoleResponseBody) HasEnvMetricV2s() bool`

HasEnvMetricV2s returns a boolean if a field has been set.

### GetEnvMetricV3s

`func (o *RoleUpdateRoleResponseBody) GetEnvMetricV3s() []EnvMetricV3ResponseBody`

GetEnvMetricV3s returns the EnvMetricV3s field if non-nil, zero value otherwise.

### GetEnvMetricV3sOk

`func (o *RoleUpdateRoleResponseBody) GetEnvMetricV3sOk() (*[]EnvMetricV3ResponseBody, bool)`

GetEnvMetricV3sOk returns a tuple with the EnvMetricV3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvMetricV3s

`func (o *RoleUpdateRoleResponseBody) SetEnvMetricV3s(v []EnvMetricV3ResponseBody)`

SetEnvMetricV3s sets EnvMetricV3s field to given value.

### HasEnvMetricV3s

`func (o *RoleUpdateRoleResponseBody) HasEnvMetricV3s() bool`

HasEnvMetricV3s returns a boolean if a field has been set.

### GetId

`func (o *RoleUpdateRoleResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleUpdateRoleResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleUpdateRoleResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *RoleUpdateRoleResponseBody) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RoleUpdateRoleResponseBody) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RoleUpdateRoleResponseBody) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RoleUpdateRoleResponseBody) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *RoleUpdateRoleResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleUpdateRoleResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleUpdateRoleResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetNewTaskCount

`func (o *RoleUpdateRoleResponseBody) GetNewTaskCount() int64`

GetNewTaskCount returns the NewTaskCount field if non-nil, zero value otherwise.

### GetNewTaskCountOk

`func (o *RoleUpdateRoleResponseBody) GetNewTaskCountOk() (*int64, bool)`

GetNewTaskCountOk returns a tuple with the NewTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskCount

`func (o *RoleUpdateRoleResponseBody) SetNewTaskCount(v int64)`

SetNewTaskCount sets NewTaskCount field to given value.

### HasNewTaskCount

`func (o *RoleUpdateRoleResponseBody) HasNewTaskCount() bool`

HasNewTaskCount returns a boolean if a field has been set.

### GetSecMetric

`func (o *RoleUpdateRoleResponseBody) GetSecMetric() SecMetricResponseBody`

GetSecMetric returns the SecMetric field if non-nil, zero value otherwise.

### GetSecMetricOk

`func (o *RoleUpdateRoleResponseBody) GetSecMetricOk() (*SecMetricResponseBody, bool)`

GetSecMetricOk returns a tuple with the SecMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecMetric

`func (o *RoleUpdateRoleResponseBody) SetSecMetric(v SecMetricResponseBody)`

SetSecMetric sets SecMetric field to given value.

### HasSecMetric

`func (o *RoleUpdateRoleResponseBody) HasSecMetric() bool`

HasSecMetric returns a boolean if a field has been set.

### GetServers

`func (o *RoleUpdateRoleResponseBody) GetServers() []ServerChildResponseBody`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RoleUpdateRoleResponseBody) GetServersOk() (*[]ServerChildResponseBody, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RoleUpdateRoleResponseBody) SetServers(v []ServerChildResponseBody)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RoleUpdateRoleResponseBody) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoleUpdateRoleResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoleUpdateRoleResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoleUpdateRoleResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


