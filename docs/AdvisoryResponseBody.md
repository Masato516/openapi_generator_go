# AdvisoryResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryID** | **string** | AdvisoryID of advisory | 
**OsFamily** | **string** | osFamily of advisory | 

## Methods

### NewAdvisoryResponseBody

`func NewAdvisoryResponseBody(advisoryID string, osFamily string, ) *AdvisoryResponseBody`

NewAdvisoryResponseBody instantiates a new AdvisoryResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryResponseBodyWithDefaults

`func NewAdvisoryResponseBodyWithDefaults() *AdvisoryResponseBody`

NewAdvisoryResponseBodyWithDefaults instantiates a new AdvisoryResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryID

`func (o *AdvisoryResponseBody) GetAdvisoryID() string`

GetAdvisoryID returns the AdvisoryID field if non-nil, zero value otherwise.

### GetAdvisoryIDOk

`func (o *AdvisoryResponseBody) GetAdvisoryIDOk() (*string, bool)`

GetAdvisoryIDOk returns a tuple with the AdvisoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryID

`func (o *AdvisoryResponseBody) SetAdvisoryID(v string)`

SetAdvisoryID sets AdvisoryID field to given value.


### GetOsFamily

`func (o *AdvisoryResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *AdvisoryResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *AdvisoryResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


