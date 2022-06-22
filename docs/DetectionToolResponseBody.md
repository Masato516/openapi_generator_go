# DetectionToolResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Detection Tool Name | 
**PatchAppliedAt** | Pointer to **time.Time** | PatchAppliedAt | [optional] 

## Methods

### NewDetectionToolResponseBody

`func NewDetectionToolResponseBody(name string, ) *DetectionToolResponseBody`

NewDetectionToolResponseBody instantiates a new DetectionToolResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionToolResponseBodyWithDefaults

`func NewDetectionToolResponseBodyWithDefaults() *DetectionToolResponseBody`

NewDetectionToolResponseBodyWithDefaults instantiates a new DetectionToolResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DetectionToolResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetectionToolResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetectionToolResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetPatchAppliedAt

`func (o *DetectionToolResponseBody) GetPatchAppliedAt() time.Time`

GetPatchAppliedAt returns the PatchAppliedAt field if non-nil, zero value otherwise.

### GetPatchAppliedAtOk

`func (o *DetectionToolResponseBody) GetPatchAppliedAtOk() (*time.Time, bool)`

GetPatchAppliedAtOk returns a tuple with the PatchAppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchAppliedAt

`func (o *DetectionToolResponseBody) SetPatchAppliedAt(v time.Time)`

SetPatchAppliedAt sets PatchAppliedAt field to given value.

### HasPatchAppliedAt

`func (o *DetectionToolResponseBody) HasPatchAppliedAt() bool`

HasPatchAppliedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


