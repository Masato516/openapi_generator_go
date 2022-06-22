# RoleListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTaskCount** | Pointer to **int64** | AllTaskCount of server role | [optional] 
**CreatedAt** | **time.Time** | created time of server role | 
**Id** | **int64** | ID of server role | 
**IsDefault** | **bool** | isDefault of server role | 
**Name** | **string** | Name of server role | 
**NewTaskCount** | Pointer to **int64** | NewTaskCount of server role | [optional] 
**SecMetric** | Pointer to [**SecMetricResponseBody**](SecMetricResponseBody.md) |  | [optional] 
**ServerCount** | Pointer to **int64** | Server Count of server role | [optional] 
**UpdatedAt** | **time.Time** | updated time of server role | 

## Methods

### NewRoleListResponseBody

`func NewRoleListResponseBody(createdAt time.Time, id int64, isDefault bool, name string, updatedAt time.Time, ) *RoleListResponseBody`

NewRoleListResponseBody instantiates a new RoleListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleListResponseBodyWithDefaults

`func NewRoleListResponseBodyWithDefaults() *RoleListResponseBody`

NewRoleListResponseBodyWithDefaults instantiates a new RoleListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTaskCount

`func (o *RoleListResponseBody) GetAllTaskCount() int64`

GetAllTaskCount returns the AllTaskCount field if non-nil, zero value otherwise.

### GetAllTaskCountOk

`func (o *RoleListResponseBody) GetAllTaskCountOk() (*int64, bool)`

GetAllTaskCountOk returns a tuple with the AllTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTaskCount

`func (o *RoleListResponseBody) SetAllTaskCount(v int64)`

SetAllTaskCount sets AllTaskCount field to given value.

### HasAllTaskCount

`func (o *RoleListResponseBody) HasAllTaskCount() bool`

HasAllTaskCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleListResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleListResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleListResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *RoleListResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleListResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleListResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *RoleListResponseBody) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RoleListResponseBody) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RoleListResponseBody) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetName

`func (o *RoleListResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleListResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleListResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetNewTaskCount

`func (o *RoleListResponseBody) GetNewTaskCount() int64`

GetNewTaskCount returns the NewTaskCount field if non-nil, zero value otherwise.

### GetNewTaskCountOk

`func (o *RoleListResponseBody) GetNewTaskCountOk() (*int64, bool)`

GetNewTaskCountOk returns a tuple with the NewTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskCount

`func (o *RoleListResponseBody) SetNewTaskCount(v int64)`

SetNewTaskCount sets NewTaskCount field to given value.

### HasNewTaskCount

`func (o *RoleListResponseBody) HasNewTaskCount() bool`

HasNewTaskCount returns a boolean if a field has been set.

### GetSecMetric

`func (o *RoleListResponseBody) GetSecMetric() SecMetricResponseBody`

GetSecMetric returns the SecMetric field if non-nil, zero value otherwise.

### GetSecMetricOk

`func (o *RoleListResponseBody) GetSecMetricOk() (*SecMetricResponseBody, bool)`

GetSecMetricOk returns a tuple with the SecMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecMetric

`func (o *RoleListResponseBody) SetSecMetric(v SecMetricResponseBody)`

SetSecMetric sets SecMetric field to given value.

### HasSecMetric

`func (o *RoleListResponseBody) HasSecMetric() bool`

HasSecMetric returns a boolean if a field has been set.

### GetServerCount

`func (o *RoleListResponseBody) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *RoleListResponseBody) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *RoleListResponseBody) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *RoleListResponseBody) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoleListResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoleListResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoleListResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


