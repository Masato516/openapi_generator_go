# ChildTaskResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyingPatchOn** | Pointer to **string** | ApplyingPatchOn of task | [optional] 
**CreatedAt** | **time.Time** | created time of task | 
**CveID** | **string** | CVE ID of task | 
**Id** | **int64** | ID of task | 
**Ignore** | **bool** | Ignore of task | 
**IgnoreUntil** | Pointer to **string** | Ignore until of task | [optional] 
**MainUserID** | Pointer to **int64** | MainUserID of task | [optional] 
**MainUserName** | Pointer to **string** | MainUserName of task | [optional] 
**Priority** | **string** | Priority of task | 
**ServerID** | **int64** | ServerID of task | 
**Status** | **string** | Status of task | 
**SubUserID** | Pointer to **int64** | SubUserID of task | [optional] 
**SubUserName** | Pointer to **string** | SubUserName of task | [optional] 
**UpdatedAt** | **time.Time** | updated time of task | 

## Methods

### NewChildTaskResponseBody

`func NewChildTaskResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, priority string, serverID int64, status string, updatedAt time.Time, ) *ChildTaskResponseBody`

NewChildTaskResponseBody instantiates a new ChildTaskResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildTaskResponseBodyWithDefaults

`func NewChildTaskResponseBodyWithDefaults() *ChildTaskResponseBody`

NewChildTaskResponseBodyWithDefaults instantiates a new ChildTaskResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyingPatchOn

`func (o *ChildTaskResponseBody) GetApplyingPatchOn() string`

GetApplyingPatchOn returns the ApplyingPatchOn field if non-nil, zero value otherwise.

### GetApplyingPatchOnOk

`func (o *ChildTaskResponseBody) GetApplyingPatchOnOk() (*string, bool)`

GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyingPatchOn

`func (o *ChildTaskResponseBody) SetApplyingPatchOn(v string)`

SetApplyingPatchOn sets ApplyingPatchOn field to given value.

### HasApplyingPatchOn

`func (o *ChildTaskResponseBody) HasApplyingPatchOn() bool`

HasApplyingPatchOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChildTaskResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChildTaskResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChildTaskResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *ChildTaskResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *ChildTaskResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *ChildTaskResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetId

`func (o *ChildTaskResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChildTaskResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChildTaskResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIgnore

`func (o *ChildTaskResponseBody) GetIgnore() bool`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *ChildTaskResponseBody) GetIgnoreOk() (*bool, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *ChildTaskResponseBody) SetIgnore(v bool)`

SetIgnore sets Ignore field to given value.


### GetIgnoreUntil

`func (o *ChildTaskResponseBody) GetIgnoreUntil() string`

GetIgnoreUntil returns the IgnoreUntil field if non-nil, zero value otherwise.

### GetIgnoreUntilOk

`func (o *ChildTaskResponseBody) GetIgnoreUntilOk() (*string, bool)`

GetIgnoreUntilOk returns a tuple with the IgnoreUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUntil

`func (o *ChildTaskResponseBody) SetIgnoreUntil(v string)`

SetIgnoreUntil sets IgnoreUntil field to given value.

### HasIgnoreUntil

`func (o *ChildTaskResponseBody) HasIgnoreUntil() bool`

HasIgnoreUntil returns a boolean if a field has been set.

### GetMainUserID

`func (o *ChildTaskResponseBody) GetMainUserID() int64`

GetMainUserID returns the MainUserID field if non-nil, zero value otherwise.

### GetMainUserIDOk

`func (o *ChildTaskResponseBody) GetMainUserIDOk() (*int64, bool)`

GetMainUserIDOk returns a tuple with the MainUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserID

`func (o *ChildTaskResponseBody) SetMainUserID(v int64)`

SetMainUserID sets MainUserID field to given value.

### HasMainUserID

`func (o *ChildTaskResponseBody) HasMainUserID() bool`

HasMainUserID returns a boolean if a field has been set.

### GetMainUserName

`func (o *ChildTaskResponseBody) GetMainUserName() string`

GetMainUserName returns the MainUserName field if non-nil, zero value otherwise.

### GetMainUserNameOk

`func (o *ChildTaskResponseBody) GetMainUserNameOk() (*string, bool)`

GetMainUserNameOk returns a tuple with the MainUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserName

`func (o *ChildTaskResponseBody) SetMainUserName(v string)`

SetMainUserName sets MainUserName field to given value.

### HasMainUserName

`func (o *ChildTaskResponseBody) HasMainUserName() bool`

HasMainUserName returns a boolean if a field has been set.

### GetPriority

`func (o *ChildTaskResponseBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ChildTaskResponseBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ChildTaskResponseBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetServerID

`func (o *ChildTaskResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *ChildTaskResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *ChildTaskResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetStatus

`func (o *ChildTaskResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildTaskResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildTaskResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubUserID

`func (o *ChildTaskResponseBody) GetSubUserID() int64`

GetSubUserID returns the SubUserID field if non-nil, zero value otherwise.

### GetSubUserIDOk

`func (o *ChildTaskResponseBody) GetSubUserIDOk() (*int64, bool)`

GetSubUserIDOk returns a tuple with the SubUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserID

`func (o *ChildTaskResponseBody) SetSubUserID(v int64)`

SetSubUserID sets SubUserID field to given value.

### HasSubUserID

`func (o *ChildTaskResponseBody) HasSubUserID() bool`

HasSubUserID returns a boolean if a field has been set.

### GetSubUserName

`func (o *ChildTaskResponseBody) GetSubUserName() string`

GetSubUserName returns the SubUserName field if non-nil, zero value otherwise.

### GetSubUserNameOk

`func (o *ChildTaskResponseBody) GetSubUserNameOk() (*string, bool)`

GetSubUserNameOk returns a tuple with the SubUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserName

`func (o *ChildTaskResponseBody) SetSubUserName(v string)`

SetSubUserName sets SubUserName field to given value.

### HasSubUserName

`func (o *ChildTaskResponseBody) HasSubUserName() bool`

HasSubUserName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChildTaskResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChildTaskResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChildTaskResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


