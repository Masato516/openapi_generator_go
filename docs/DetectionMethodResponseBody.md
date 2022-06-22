# DetectionMethodResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Detection Method Name | 
**ReliabilityScore** | **int64** | ReliabilityScore | 

## Methods

### NewDetectionMethodResponseBody

`func NewDetectionMethodResponseBody(name string, reliabilityScore int64, ) *DetectionMethodResponseBody`

NewDetectionMethodResponseBody instantiates a new DetectionMethodResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionMethodResponseBodyWithDefaults

`func NewDetectionMethodResponseBodyWithDefaults() *DetectionMethodResponseBody`

NewDetectionMethodResponseBodyWithDefaults instantiates a new DetectionMethodResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DetectionMethodResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetectionMethodResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetectionMethodResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetReliabilityScore

`func (o *DetectionMethodResponseBody) GetReliabilityScore() int64`

GetReliabilityScore returns the ReliabilityScore field if non-nil, zero value otherwise.

### GetReliabilityScoreOk

`func (o *DetectionMethodResponseBody) GetReliabilityScoreOk() (*int64, bool)`

GetReliabilityScoreOk returns a tuple with the ReliabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliabilityScore

`func (o *DetectionMethodResponseBody) SetReliabilityScore(v int64)`

SetReliabilityScore sets ReliabilityScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


