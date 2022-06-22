# CveGetCveDetailResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | Pointer to [**[]AdvisoryResponseBody**](AdvisoryResponseBody.md) | advisory of cve | [optional] 
**CreatedAt** | **time.Time** | created time | 
**CveID** | **string** | Cve ID string of cve | 
**Cvss** | ***os.File** | cvss of cve | 
**Cwes** | [**[]CweResponseBody**](CweResponseBody.md) | cwes of cve | 
**EnvMetricV2s** | [**[]EnvMetricV2ResponseBody**](EnvMetricV2ResponseBody.md) | envMetricV2 of cve | 
**EnvMetricV3s** | [**[]EnvMetricV3ResponseBody**](EnvMetricV3ResponseBody.md) | envMetricV3 of cve | 
**References** | **map[string]string** | references of cve | 
**SecMetrics** | [**[]SecMetricResponseBody**](SecMetricResponseBody.md) | secMetric of cve | 
**ServerOsFamilies** | **[]string** | serverOsFamilies of cve | 
**TmpMetricV2** | [**TmpMetricResponseBody**](TmpMetricResponseBody.md) |  | 
**TmpMetricV3** | [**TmpMetricResponseBody**](TmpMetricResponseBody.md) |  | 
**UpdatedAt** | **time.Time** | updated time | 

## Methods

### NewCveGetCveDetailResponseBody

`func NewCveGetCveDetailResponseBody(createdAt time.Time, cveID string, cvss *os.File, cwes []CweResponseBody, envMetricV2s []EnvMetricV2ResponseBody, envMetricV3s []EnvMetricV3ResponseBody, references map[string]string, secMetrics []SecMetricResponseBody, serverOsFamilies []string, tmpMetricV2 TmpMetricResponseBody, tmpMetricV3 TmpMetricResponseBody, updatedAt time.Time, ) *CveGetCveDetailResponseBody`

NewCveGetCveDetailResponseBody instantiates a new CveGetCveDetailResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveGetCveDetailResponseBodyWithDefaults

`func NewCveGetCveDetailResponseBodyWithDefaults() *CveGetCveDetailResponseBody`

NewCveGetCveDetailResponseBodyWithDefaults instantiates a new CveGetCveDetailResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisories

`func (o *CveGetCveDetailResponseBody) GetAdvisories() []AdvisoryResponseBody`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *CveGetCveDetailResponseBody) GetAdvisoriesOk() (*[]AdvisoryResponseBody, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *CveGetCveDetailResponseBody) SetAdvisories(v []AdvisoryResponseBody)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *CveGetCveDetailResponseBody) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CveGetCveDetailResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CveGetCveDetailResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CveGetCveDetailResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *CveGetCveDetailResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *CveGetCveDetailResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *CveGetCveDetailResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetCvss

`func (o *CveGetCveDetailResponseBody) GetCvss() *os.File`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *CveGetCveDetailResponseBody) GetCvssOk() (**os.File, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *CveGetCveDetailResponseBody) SetCvss(v *os.File)`

SetCvss sets Cvss field to given value.


### GetCwes

`func (o *CveGetCveDetailResponseBody) GetCwes() []CweResponseBody`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *CveGetCveDetailResponseBody) GetCwesOk() (*[]CweResponseBody, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *CveGetCveDetailResponseBody) SetCwes(v []CweResponseBody)`

SetCwes sets Cwes field to given value.


### GetEnvMetricV2s

`func (o *CveGetCveDetailResponseBody) GetEnvMetricV2s() []EnvMetricV2ResponseBody`

GetEnvMetricV2s returns the EnvMetricV2s field if non-nil, zero value otherwise.

### GetEnvMetricV2sOk

`func (o *CveGetCveDetailResponseBody) GetEnvMetricV2sOk() (*[]EnvMetricV2ResponseBody, bool)`

GetEnvMetricV2sOk returns a tuple with the EnvMetricV2s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvMetricV2s

`func (o *CveGetCveDetailResponseBody) SetEnvMetricV2s(v []EnvMetricV2ResponseBody)`

SetEnvMetricV2s sets EnvMetricV2s field to given value.


### GetEnvMetricV3s

`func (o *CveGetCveDetailResponseBody) GetEnvMetricV3s() []EnvMetricV3ResponseBody`

GetEnvMetricV3s returns the EnvMetricV3s field if non-nil, zero value otherwise.

### GetEnvMetricV3sOk

`func (o *CveGetCveDetailResponseBody) GetEnvMetricV3sOk() (*[]EnvMetricV3ResponseBody, bool)`

GetEnvMetricV3sOk returns a tuple with the EnvMetricV3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvMetricV3s

`func (o *CveGetCveDetailResponseBody) SetEnvMetricV3s(v []EnvMetricV3ResponseBody)`

SetEnvMetricV3s sets EnvMetricV3s field to given value.


### GetReferences

`func (o *CveGetCveDetailResponseBody) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CveGetCveDetailResponseBody) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CveGetCveDetailResponseBody) SetReferences(v map[string]string)`

SetReferences sets References field to given value.


### GetSecMetrics

`func (o *CveGetCveDetailResponseBody) GetSecMetrics() []SecMetricResponseBody`

GetSecMetrics returns the SecMetrics field if non-nil, zero value otherwise.

### GetSecMetricsOk

`func (o *CveGetCveDetailResponseBody) GetSecMetricsOk() (*[]SecMetricResponseBody, bool)`

GetSecMetricsOk returns a tuple with the SecMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecMetrics

`func (o *CveGetCveDetailResponseBody) SetSecMetrics(v []SecMetricResponseBody)`

SetSecMetrics sets SecMetrics field to given value.


### GetServerOsFamilies

`func (o *CveGetCveDetailResponseBody) GetServerOsFamilies() []string`

GetServerOsFamilies returns the ServerOsFamilies field if non-nil, zero value otherwise.

### GetServerOsFamiliesOk

`func (o *CveGetCveDetailResponseBody) GetServerOsFamiliesOk() (*[]string, bool)`

GetServerOsFamiliesOk returns a tuple with the ServerOsFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOsFamilies

`func (o *CveGetCveDetailResponseBody) SetServerOsFamilies(v []string)`

SetServerOsFamilies sets ServerOsFamilies field to given value.


### GetTmpMetricV2

`func (o *CveGetCveDetailResponseBody) GetTmpMetricV2() TmpMetricResponseBody`

GetTmpMetricV2 returns the TmpMetricV2 field if non-nil, zero value otherwise.

### GetTmpMetricV2Ok

`func (o *CveGetCveDetailResponseBody) GetTmpMetricV2Ok() (*TmpMetricResponseBody, bool)`

GetTmpMetricV2Ok returns a tuple with the TmpMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpMetricV2

`func (o *CveGetCveDetailResponseBody) SetTmpMetricV2(v TmpMetricResponseBody)`

SetTmpMetricV2 sets TmpMetricV2 field to given value.


### GetTmpMetricV3

`func (o *CveGetCveDetailResponseBody) GetTmpMetricV3() TmpMetricResponseBody`

GetTmpMetricV3 returns the TmpMetricV3 field if non-nil, zero value otherwise.

### GetTmpMetricV3Ok

`func (o *CveGetCveDetailResponseBody) GetTmpMetricV3Ok() (*TmpMetricResponseBody, bool)`

GetTmpMetricV3Ok returns a tuple with the TmpMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpMetricV3

`func (o *CveGetCveDetailResponseBody) SetTmpMetricV3(v TmpMetricResponseBody)`

SetTmpMetricV3 sets TmpMetricV3 field to given value.


### GetUpdatedAt

`func (o *CveGetCveDetailResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CveGetCveDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CveGetCveDetailResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


