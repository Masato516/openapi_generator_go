# CweResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CweID** | **string** | cweID of cwe | 
**English** | **string** | english summary of cwe | 
**Japanese** | **string** | japanese summary of cwe | 
**OwaspTopTen2017** | Pointer to **string** | owaspTopTen2017 of cwe  | [optional] 
**SourceType** | **string** | sourceType of cwe | 

## Methods

### NewCweResponseBody

`func NewCweResponseBody(cweID string, english string, japanese string, sourceType string, ) *CweResponseBody`

NewCweResponseBody instantiates a new CweResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCweResponseBodyWithDefaults

`func NewCweResponseBodyWithDefaults() *CweResponseBody`

NewCweResponseBodyWithDefaults instantiates a new CweResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCweID

`func (o *CweResponseBody) GetCweID() string`

GetCweID returns the CweID field if non-nil, zero value otherwise.

### GetCweIDOk

`func (o *CweResponseBody) GetCweIDOk() (*string, bool)`

GetCweIDOk returns a tuple with the CweID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCweID

`func (o *CweResponseBody) SetCweID(v string)`

SetCweID sets CweID field to given value.


### GetEnglish

`func (o *CweResponseBody) GetEnglish() string`

GetEnglish returns the English field if non-nil, zero value otherwise.

### GetEnglishOk

`func (o *CweResponseBody) GetEnglishOk() (*string, bool)`

GetEnglishOk returns a tuple with the English field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglish

`func (o *CweResponseBody) SetEnglish(v string)`

SetEnglish sets English field to given value.


### GetJapanese

`func (o *CweResponseBody) GetJapanese() string`

GetJapanese returns the Japanese field if non-nil, zero value otherwise.

### GetJapaneseOk

`func (o *CweResponseBody) GetJapaneseOk() (*string, bool)`

GetJapaneseOk returns a tuple with the Japanese field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJapanese

`func (o *CweResponseBody) SetJapanese(v string)`

SetJapanese sets Japanese field to given value.


### GetOwaspTopTen2017

`func (o *CweResponseBody) GetOwaspTopTen2017() string`

GetOwaspTopTen2017 returns the OwaspTopTen2017 field if non-nil, zero value otherwise.

### GetOwaspTopTen2017Ok

`func (o *CweResponseBody) GetOwaspTopTen2017Ok() (*string, bool)`

GetOwaspTopTen2017Ok returns a tuple with the OwaspTopTen2017 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwaspTopTen2017

`func (o *CweResponseBody) SetOwaspTopTen2017(v string)`

SetOwaspTopTen2017 sets OwaspTopTen2017 field to given value.

### HasOwaspTopTen2017

`func (o *CweResponseBody) HasOwaspTopTen2017() bool`

HasOwaspTopTen2017 returns a boolean if a field has been set.

### GetSourceType

`func (o *CweResponseBody) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CweResponseBody) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CweResponseBody) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


