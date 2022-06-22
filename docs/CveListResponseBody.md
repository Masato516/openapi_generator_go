# CveListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | Pointer to **[]string** | advisoryIDs of cve | [optional] 
**AllTaskCount** | **int64** | AllTaskCount of cve | 
**CreatedAt** | **time.Time** | created time | 
**CveID** | **string** | Cve ID string of cve | 
**Cwes** | [**[]CweResponseBody**](CweResponseBody.md) | cwes of cve | 
**HasExploit** | Pointer to **bool** | hasExploit of cve | [optional] 
**HasMitigation** | Pointer to **bool** | hasMitigation of cve | [optional] 
**HasWorkaround** | Pointer to **bool** | hasWorkaroundof cve | [optional] 
**IsNotActive** | **bool** | Flag of active cve | 
**IsOwaspTopTen2017** | **bool** | isOwaspTopTen2017 of cve | 
**MaxV2** | **float64** | maxV2 of cve | 
**MaxV3** | **float64** | maxV3 of cve | 
**NewTaskCount** | **int64** | NewTaskCount of cve | 
**ScoreV2s** | **map[string]float64** | cvss v2 scores of cve | 
**ScoreV3s** | **map[string]float64** | cvss v3 scores of cve | 
**Title** | **string** | Title of cve | 
**TopicCount** | **int64** | topicCount of cve | 
**TopicLastUpdatedAt** | **time.Time** | topicLastUpdatedAt of cve | 
**UpdatedAt** | **time.Time** | updated time | 
**VectorV2s** | **map[string]string** | cvss v2 vectors of cve | 
**VectorV3s** | **map[string]string** | cvss v3 vectors of cve | 

## Methods

### NewCveListResponseBody

`func NewCveListResponseBody(allTaskCount int64, createdAt time.Time, cveID string, cwes []CweResponseBody, isNotActive bool, isOwaspTopTen2017 bool, maxV2 float64, maxV3 float64, newTaskCount int64, scoreV2s map[string]float64, scoreV3s map[string]float64, title string, topicCount int64, topicLastUpdatedAt time.Time, updatedAt time.Time, vectorV2s map[string]string, vectorV3s map[string]string, ) *CveListResponseBody`

NewCveListResponseBody instantiates a new CveListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveListResponseBodyWithDefaults

`func NewCveListResponseBodyWithDefaults() *CveListResponseBody`

NewCveListResponseBodyWithDefaults instantiates a new CveListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryIDs

`func (o *CveListResponseBody) GetAdvisoryIDs() []string`

GetAdvisoryIDs returns the AdvisoryIDs field if non-nil, zero value otherwise.

### GetAdvisoryIDsOk

`func (o *CveListResponseBody) GetAdvisoryIDsOk() (*[]string, bool)`

GetAdvisoryIDsOk returns a tuple with the AdvisoryIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryIDs

`func (o *CveListResponseBody) SetAdvisoryIDs(v []string)`

SetAdvisoryIDs sets AdvisoryIDs field to given value.

### HasAdvisoryIDs

`func (o *CveListResponseBody) HasAdvisoryIDs() bool`

HasAdvisoryIDs returns a boolean if a field has been set.

### GetAllTaskCount

`func (o *CveListResponseBody) GetAllTaskCount() int64`

GetAllTaskCount returns the AllTaskCount field if non-nil, zero value otherwise.

### GetAllTaskCountOk

`func (o *CveListResponseBody) GetAllTaskCountOk() (*int64, bool)`

GetAllTaskCountOk returns a tuple with the AllTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTaskCount

`func (o *CveListResponseBody) SetAllTaskCount(v int64)`

SetAllTaskCount sets AllTaskCount field to given value.


### GetCreatedAt

`func (o *CveListResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CveListResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CveListResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *CveListResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *CveListResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *CveListResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetCwes

`func (o *CveListResponseBody) GetCwes() []CweResponseBody`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *CveListResponseBody) GetCwesOk() (*[]CweResponseBody, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *CveListResponseBody) SetCwes(v []CweResponseBody)`

SetCwes sets Cwes field to given value.


### GetHasExploit

`func (o *CveListResponseBody) GetHasExploit() bool`

GetHasExploit returns the HasExploit field if non-nil, zero value otherwise.

### GetHasExploitOk

`func (o *CveListResponseBody) GetHasExploitOk() (*bool, bool)`

GetHasExploitOk returns a tuple with the HasExploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExploit

`func (o *CveListResponseBody) SetHasExploit(v bool)`

SetHasExploit sets HasExploit field to given value.

### HasHasExploit

`func (o *CveListResponseBody) HasHasExploit() bool`

HasHasExploit returns a boolean if a field has been set.

### GetHasMitigation

`func (o *CveListResponseBody) GetHasMitigation() bool`

GetHasMitigation returns the HasMitigation field if non-nil, zero value otherwise.

### GetHasMitigationOk

`func (o *CveListResponseBody) GetHasMitigationOk() (*bool, bool)`

GetHasMitigationOk returns a tuple with the HasMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMitigation

`func (o *CveListResponseBody) SetHasMitigation(v bool)`

SetHasMitigation sets HasMitigation field to given value.

### HasHasMitigation

`func (o *CveListResponseBody) HasHasMitigation() bool`

HasHasMitigation returns a boolean if a field has been set.

### GetHasWorkaround

`func (o *CveListResponseBody) GetHasWorkaround() bool`

GetHasWorkaround returns the HasWorkaround field if non-nil, zero value otherwise.

### GetHasWorkaroundOk

`func (o *CveListResponseBody) GetHasWorkaroundOk() (*bool, bool)`

GetHasWorkaroundOk returns a tuple with the HasWorkaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWorkaround

`func (o *CveListResponseBody) SetHasWorkaround(v bool)`

SetHasWorkaround sets HasWorkaround field to given value.

### HasHasWorkaround

`func (o *CveListResponseBody) HasHasWorkaround() bool`

HasHasWorkaround returns a boolean if a field has been set.

### GetIsNotActive

`func (o *CveListResponseBody) GetIsNotActive() bool`

GetIsNotActive returns the IsNotActive field if non-nil, zero value otherwise.

### GetIsNotActiveOk

`func (o *CveListResponseBody) GetIsNotActiveOk() (*bool, bool)`

GetIsNotActiveOk returns a tuple with the IsNotActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotActive

`func (o *CveListResponseBody) SetIsNotActive(v bool)`

SetIsNotActive sets IsNotActive field to given value.


### GetIsOwaspTopTen2017

`func (o *CveListResponseBody) GetIsOwaspTopTen2017() bool`

GetIsOwaspTopTen2017 returns the IsOwaspTopTen2017 field if non-nil, zero value otherwise.

### GetIsOwaspTopTen2017Ok

`func (o *CveListResponseBody) GetIsOwaspTopTen2017Ok() (*bool, bool)`

GetIsOwaspTopTen2017Ok returns a tuple with the IsOwaspTopTen2017 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwaspTopTen2017

`func (o *CveListResponseBody) SetIsOwaspTopTen2017(v bool)`

SetIsOwaspTopTen2017 sets IsOwaspTopTen2017 field to given value.


### GetMaxV2

`func (o *CveListResponseBody) GetMaxV2() float64`

GetMaxV2 returns the MaxV2 field if non-nil, zero value otherwise.

### GetMaxV2Ok

`func (o *CveListResponseBody) GetMaxV2Ok() (*float64, bool)`

GetMaxV2Ok returns a tuple with the MaxV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxV2

`func (o *CveListResponseBody) SetMaxV2(v float64)`

SetMaxV2 sets MaxV2 field to given value.


### GetMaxV3

`func (o *CveListResponseBody) GetMaxV3() float64`

GetMaxV3 returns the MaxV3 field if non-nil, zero value otherwise.

### GetMaxV3Ok

`func (o *CveListResponseBody) GetMaxV3Ok() (*float64, bool)`

GetMaxV3Ok returns a tuple with the MaxV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxV3

`func (o *CveListResponseBody) SetMaxV3(v float64)`

SetMaxV3 sets MaxV3 field to given value.


### GetNewTaskCount

`func (o *CveListResponseBody) GetNewTaskCount() int64`

GetNewTaskCount returns the NewTaskCount field if non-nil, zero value otherwise.

### GetNewTaskCountOk

`func (o *CveListResponseBody) GetNewTaskCountOk() (*int64, bool)`

GetNewTaskCountOk returns a tuple with the NewTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskCount

`func (o *CveListResponseBody) SetNewTaskCount(v int64)`

SetNewTaskCount sets NewTaskCount field to given value.


### GetScoreV2s

`func (o *CveListResponseBody) GetScoreV2s() map[string]float64`

GetScoreV2s returns the ScoreV2s field if non-nil, zero value otherwise.

### GetScoreV2sOk

`func (o *CveListResponseBody) GetScoreV2sOk() (*map[string]float64, bool)`

GetScoreV2sOk returns a tuple with the ScoreV2s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreV2s

`func (o *CveListResponseBody) SetScoreV2s(v map[string]float64)`

SetScoreV2s sets ScoreV2s field to given value.


### GetScoreV3s

`func (o *CveListResponseBody) GetScoreV3s() map[string]float64`

GetScoreV3s returns the ScoreV3s field if non-nil, zero value otherwise.

### GetScoreV3sOk

`func (o *CveListResponseBody) GetScoreV3sOk() (*map[string]float64, bool)`

GetScoreV3sOk returns a tuple with the ScoreV3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreV3s

`func (o *CveListResponseBody) SetScoreV3s(v map[string]float64)`

SetScoreV3s sets ScoreV3s field to given value.


### GetTitle

`func (o *CveListResponseBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CveListResponseBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CveListResponseBody) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTopicCount

`func (o *CveListResponseBody) GetTopicCount() int64`

GetTopicCount returns the TopicCount field if non-nil, zero value otherwise.

### GetTopicCountOk

`func (o *CveListResponseBody) GetTopicCountOk() (*int64, bool)`

GetTopicCountOk returns a tuple with the TopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCount

`func (o *CveListResponseBody) SetTopicCount(v int64)`

SetTopicCount sets TopicCount field to given value.


### GetTopicLastUpdatedAt

`func (o *CveListResponseBody) GetTopicLastUpdatedAt() time.Time`

GetTopicLastUpdatedAt returns the TopicLastUpdatedAt field if non-nil, zero value otherwise.

### GetTopicLastUpdatedAtOk

`func (o *CveListResponseBody) GetTopicLastUpdatedAtOk() (*time.Time, bool)`

GetTopicLastUpdatedAtOk returns a tuple with the TopicLastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLastUpdatedAt

`func (o *CveListResponseBody) SetTopicLastUpdatedAt(v time.Time)`

SetTopicLastUpdatedAt sets TopicLastUpdatedAt field to given value.


### GetUpdatedAt

`func (o *CveListResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CveListResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CveListResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVectorV2s

`func (o *CveListResponseBody) GetVectorV2s() map[string]string`

GetVectorV2s returns the VectorV2s field if non-nil, zero value otherwise.

### GetVectorV2sOk

`func (o *CveListResponseBody) GetVectorV2sOk() (*map[string]string, bool)`

GetVectorV2sOk returns a tuple with the VectorV2s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorV2s

`func (o *CveListResponseBody) SetVectorV2s(v map[string]string)`

SetVectorV2s sets VectorV2s field to given value.


### GetVectorV3s

`func (o *CveListResponseBody) GetVectorV3s() map[string]string`

GetVectorV3s returns the VectorV3s field if non-nil, zero value otherwise.

### GetVectorV3sOk

`func (o *CveListResponseBody) GetVectorV3sOk() (*map[string]string, bool)`

GetVectorV3sOk returns a tuple with the VectorV3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorV3s

`func (o *CveListResponseBody) SetVectorV3s(v map[string]string)`

SetVectorV3s sets VectorV3s field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


