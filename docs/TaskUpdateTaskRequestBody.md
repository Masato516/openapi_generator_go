# TaskUpdateTaskRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyingPatchOn** | Pointer to **string** | applyingPatchOn (YYYY-MM-DD) UTC | [optional] 
**MainUserID** | Pointer to **int64** | mainUserID of task | [optional] 
**Priority** | Pointer to **string** | Priority of task | [optional] 
**Status** | Pointer to **string** | Status of task | [optional] 
**SubUserID** | Pointer to **int64** | subUserID of task | [optional] 

## Methods

### NewTaskUpdateTaskRequestBody

`func NewTaskUpdateTaskRequestBody() *TaskUpdateTaskRequestBody`

NewTaskUpdateTaskRequestBody instantiates a new TaskUpdateTaskRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskUpdateTaskRequestBodyWithDefaults

`func NewTaskUpdateTaskRequestBodyWithDefaults() *TaskUpdateTaskRequestBody`

NewTaskUpdateTaskRequestBodyWithDefaults instantiates a new TaskUpdateTaskRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyingPatchOn

`func (o *TaskUpdateTaskRequestBody) GetApplyingPatchOn() string`

GetApplyingPatchOn returns the ApplyingPatchOn field if non-nil, zero value otherwise.

### GetApplyingPatchOnOk

`func (o *TaskUpdateTaskRequestBody) GetApplyingPatchOnOk() (*string, bool)`

GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyingPatchOn

`func (o *TaskUpdateTaskRequestBody) SetApplyingPatchOn(v string)`

SetApplyingPatchOn sets ApplyingPatchOn field to given value.

### HasApplyingPatchOn

`func (o *TaskUpdateTaskRequestBody) HasApplyingPatchOn() bool`

HasApplyingPatchOn returns a boolean if a field has been set.

### GetMainUserID

`func (o *TaskUpdateTaskRequestBody) GetMainUserID() int64`

GetMainUserID returns the MainUserID field if non-nil, zero value otherwise.

### GetMainUserIDOk

`func (o *TaskUpdateTaskRequestBody) GetMainUserIDOk() (*int64, bool)`

GetMainUserIDOk returns a tuple with the MainUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserID

`func (o *TaskUpdateTaskRequestBody) SetMainUserID(v int64)`

SetMainUserID sets MainUserID field to given value.

### HasMainUserID

`func (o *TaskUpdateTaskRequestBody) HasMainUserID() bool`

HasMainUserID returns a boolean if a field has been set.

### GetPriority

`func (o *TaskUpdateTaskRequestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskUpdateTaskRequestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskUpdateTaskRequestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskUpdateTaskRequestBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *TaskUpdateTaskRequestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskUpdateTaskRequestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskUpdateTaskRequestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskUpdateTaskRequestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubUserID

`func (o *TaskUpdateTaskRequestBody) GetSubUserID() int64`

GetSubUserID returns the SubUserID field if non-nil, zero value otherwise.

### GetSubUserIDOk

`func (o *TaskUpdateTaskRequestBody) GetSubUserIDOk() (*int64, bool)`

GetSubUserIDOk returns a tuple with the SubUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserID

`func (o *TaskUpdateTaskRequestBody) SetSubUserID(v int64)`

SetSubUserID sets SubUserID field to given value.

### HasSubUserID

`func (o *TaskUpdateTaskRequestBody) HasSubUserID() bool`

HasSubUserID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


