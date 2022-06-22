# PkgCpeAddCpeRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeName** | **string** | Cpe Name(cpe uri or cpe format string) | 
**IsURI** | Pointer to **bool** | isURI specifies whether cpeName is in URI format or FormatString format. | [optional] [default to true]
**ServerID** | **int64** | ServerID | 

## Methods

### NewPkgCpeAddCpeRequestBody

`func NewPkgCpeAddCpeRequestBody(cpeName string, serverID int64, ) *PkgCpeAddCpeRequestBody`

NewPkgCpeAddCpeRequestBody instantiates a new PkgCpeAddCpeRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgCpeAddCpeRequestBodyWithDefaults

`func NewPkgCpeAddCpeRequestBodyWithDefaults() *PkgCpeAddCpeRequestBody`

NewPkgCpeAddCpeRequestBodyWithDefaults instantiates a new PkgCpeAddCpeRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeName

`func (o *PkgCpeAddCpeRequestBody) GetCpeName() string`

GetCpeName returns the CpeName field if non-nil, zero value otherwise.

### GetCpeNameOk

`func (o *PkgCpeAddCpeRequestBody) GetCpeNameOk() (*string, bool)`

GetCpeNameOk returns a tuple with the CpeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeName

`func (o *PkgCpeAddCpeRequestBody) SetCpeName(v string)`

SetCpeName sets CpeName field to given value.


### GetIsURI

`func (o *PkgCpeAddCpeRequestBody) GetIsURI() bool`

GetIsURI returns the IsURI field if non-nil, zero value otherwise.

### GetIsURIOk

`func (o *PkgCpeAddCpeRequestBody) GetIsURIOk() (*bool, bool)`

GetIsURIOk returns a tuple with the IsURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsURI

`func (o *PkgCpeAddCpeRequestBody) SetIsURI(v bool)`

SetIsURI sets IsURI field to given value.

### HasIsURI

`func (o *PkgCpeAddCpeRequestBody) HasIsURI() bool`

HasIsURI returns a boolean if a field has been set.

### GetServerID

`func (o *PkgCpeAddCpeRequestBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *PkgCpeAddCpeRequestBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *PkgCpeAddCpeRequestBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


