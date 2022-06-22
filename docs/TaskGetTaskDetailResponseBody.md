# TaskGetTaskDetailResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | Pointer to **[]string** | advisoryIDs of cve | [optional] 
**ApplyingPatchOn** | Pointer to **string** | ApplyingPatchOn of task | [optional] 
**Comments** | Pointer to [**[]TaskCommentResponseBody**](TaskCommentResponseBody.md) | Comment of task | [optional] 
**CreatedAt** | **time.Time** | created time of task | 
**CveID** | **string** | CVE ID of task | 
**Cvss** | Pointer to **map[string]*os.File** | Key Value of CveID and Cvss of task | [optional] 
**DetectionMethods** | Pointer to [**[]DetectionMethodResponseBody**](DetectionMethodResponseBody.md) | DetectionMethod of task | [optional] 
**DetectionTools** | Pointer to [**[]DetectionToolResponseBody**](DetectionToolResponseBody.md) | DetectionTools of task | [optional] 
**Id** | **int64** | ID of task | 
**Ignore** | **bool** | Ignore of task | 
**IgnoreUntil** | Pointer to **string** | Ignore until of task | [optional] 
**MainUserID** | Pointer to **int64** | MainUserID of task | [optional] 
**MainUserName** | Pointer to **string** | MainUserName of task | [optional] 
**PackageStatuses** | Pointer to **map[string]string** | packageStatus of task | [optional] 
**PkgCpes** | Pointer to [**[]PkgCpeChildResponseBody**](PkgCpeChildResponseBody.md) | Pcakge And Cpe list of task | [optional] 
**Priority** | **string** | Priority of task | 
**RoleID** | **int64** | ServerRoleID of task | 
**RoleName** | **string** | ServerRoleName of task | 
**Server** | [**ServerChildResponseBody**](ServerChildResponseBody.md) |  | 
**ServerID** | **int64** | ServerID of task | 
**Status** | **string** | Status of task | 
**SubUserID** | Pointer to **int64** | SubUserID of task | [optional] 
**SubUserName** | Pointer to **string** | SubUserName of task | [optional] 
**UpdatedAt** | **time.Time** | updated time of task | 

## Methods

### NewTaskGetTaskDetailResponseBody

`func NewTaskGetTaskDetailResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, priority string, roleID int64, roleName string, server ServerChildResponseBody, serverID int64, status string, updatedAt time.Time, ) *TaskGetTaskDetailResponseBody`

NewTaskGetTaskDetailResponseBody instantiates a new TaskGetTaskDetailResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGetTaskDetailResponseBodyWithDefaults

`func NewTaskGetTaskDetailResponseBodyWithDefaults() *TaskGetTaskDetailResponseBody`

NewTaskGetTaskDetailResponseBodyWithDefaults instantiates a new TaskGetTaskDetailResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryIDs

`func (o *TaskGetTaskDetailResponseBody) GetAdvisoryIDs() []string`

GetAdvisoryIDs returns the AdvisoryIDs field if non-nil, zero value otherwise.

### GetAdvisoryIDsOk

`func (o *TaskGetTaskDetailResponseBody) GetAdvisoryIDsOk() (*[]string, bool)`

GetAdvisoryIDsOk returns a tuple with the AdvisoryIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryIDs

`func (o *TaskGetTaskDetailResponseBody) SetAdvisoryIDs(v []string)`

SetAdvisoryIDs sets AdvisoryIDs field to given value.

### HasAdvisoryIDs

`func (o *TaskGetTaskDetailResponseBody) HasAdvisoryIDs() bool`

HasAdvisoryIDs returns a boolean if a field has been set.

### GetApplyingPatchOn

`func (o *TaskGetTaskDetailResponseBody) GetApplyingPatchOn() string`

GetApplyingPatchOn returns the ApplyingPatchOn field if non-nil, zero value otherwise.

### GetApplyingPatchOnOk

`func (o *TaskGetTaskDetailResponseBody) GetApplyingPatchOnOk() (*string, bool)`

GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyingPatchOn

`func (o *TaskGetTaskDetailResponseBody) SetApplyingPatchOn(v string)`

SetApplyingPatchOn sets ApplyingPatchOn field to given value.

### HasApplyingPatchOn

`func (o *TaskGetTaskDetailResponseBody) HasApplyingPatchOn() bool`

HasApplyingPatchOn returns a boolean if a field has been set.

### GetComments

`func (o *TaskGetTaskDetailResponseBody) GetComments() []TaskCommentResponseBody`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TaskGetTaskDetailResponseBody) GetCommentsOk() (*[]TaskCommentResponseBody, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TaskGetTaskDetailResponseBody) SetComments(v []TaskCommentResponseBody)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TaskGetTaskDetailResponseBody) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskGetTaskDetailResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskGetTaskDetailResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskGetTaskDetailResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *TaskGetTaskDetailResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *TaskGetTaskDetailResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *TaskGetTaskDetailResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetCvss

`func (o *TaskGetTaskDetailResponseBody) GetCvss() map[string]*os.File`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *TaskGetTaskDetailResponseBody) GetCvssOk() (*map[string]*os.File, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *TaskGetTaskDetailResponseBody) SetCvss(v map[string]*os.File)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *TaskGetTaskDetailResponseBody) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDetectionMethods

`func (o *TaskGetTaskDetailResponseBody) GetDetectionMethods() []DetectionMethodResponseBody`

GetDetectionMethods returns the DetectionMethods field if non-nil, zero value otherwise.

### GetDetectionMethodsOk

`func (o *TaskGetTaskDetailResponseBody) GetDetectionMethodsOk() (*[]DetectionMethodResponseBody, bool)`

GetDetectionMethodsOk returns a tuple with the DetectionMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethods

`func (o *TaskGetTaskDetailResponseBody) SetDetectionMethods(v []DetectionMethodResponseBody)`

SetDetectionMethods sets DetectionMethods field to given value.

### HasDetectionMethods

`func (o *TaskGetTaskDetailResponseBody) HasDetectionMethods() bool`

HasDetectionMethods returns a boolean if a field has been set.

### GetDetectionTools

`func (o *TaskGetTaskDetailResponseBody) GetDetectionTools() []DetectionToolResponseBody`

GetDetectionTools returns the DetectionTools field if non-nil, zero value otherwise.

### GetDetectionToolsOk

`func (o *TaskGetTaskDetailResponseBody) GetDetectionToolsOk() (*[]DetectionToolResponseBody, bool)`

GetDetectionToolsOk returns a tuple with the DetectionTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionTools

`func (o *TaskGetTaskDetailResponseBody) SetDetectionTools(v []DetectionToolResponseBody)`

SetDetectionTools sets DetectionTools field to given value.

### HasDetectionTools

`func (o *TaskGetTaskDetailResponseBody) HasDetectionTools() bool`

HasDetectionTools returns a boolean if a field has been set.

### GetId

`func (o *TaskGetTaskDetailResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskGetTaskDetailResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskGetTaskDetailResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIgnore

`func (o *TaskGetTaskDetailResponseBody) GetIgnore() bool`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *TaskGetTaskDetailResponseBody) GetIgnoreOk() (*bool, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *TaskGetTaskDetailResponseBody) SetIgnore(v bool)`

SetIgnore sets Ignore field to given value.


### GetIgnoreUntil

`func (o *TaskGetTaskDetailResponseBody) GetIgnoreUntil() string`

GetIgnoreUntil returns the IgnoreUntil field if non-nil, zero value otherwise.

### GetIgnoreUntilOk

`func (o *TaskGetTaskDetailResponseBody) GetIgnoreUntilOk() (*string, bool)`

GetIgnoreUntilOk returns a tuple with the IgnoreUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUntil

`func (o *TaskGetTaskDetailResponseBody) SetIgnoreUntil(v string)`

SetIgnoreUntil sets IgnoreUntil field to given value.

### HasIgnoreUntil

`func (o *TaskGetTaskDetailResponseBody) HasIgnoreUntil() bool`

HasIgnoreUntil returns a boolean if a field has been set.

### GetMainUserID

`func (o *TaskGetTaskDetailResponseBody) GetMainUserID() int64`

GetMainUserID returns the MainUserID field if non-nil, zero value otherwise.

### GetMainUserIDOk

`func (o *TaskGetTaskDetailResponseBody) GetMainUserIDOk() (*int64, bool)`

GetMainUserIDOk returns a tuple with the MainUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserID

`func (o *TaskGetTaskDetailResponseBody) SetMainUserID(v int64)`

SetMainUserID sets MainUserID field to given value.

### HasMainUserID

`func (o *TaskGetTaskDetailResponseBody) HasMainUserID() bool`

HasMainUserID returns a boolean if a field has been set.

### GetMainUserName

`func (o *TaskGetTaskDetailResponseBody) GetMainUserName() string`

GetMainUserName returns the MainUserName field if non-nil, zero value otherwise.

### GetMainUserNameOk

`func (o *TaskGetTaskDetailResponseBody) GetMainUserNameOk() (*string, bool)`

GetMainUserNameOk returns a tuple with the MainUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserName

`func (o *TaskGetTaskDetailResponseBody) SetMainUserName(v string)`

SetMainUserName sets MainUserName field to given value.

### HasMainUserName

`func (o *TaskGetTaskDetailResponseBody) HasMainUserName() bool`

HasMainUserName returns a boolean if a field has been set.

### GetPackageStatuses

`func (o *TaskGetTaskDetailResponseBody) GetPackageStatuses() map[string]string`

GetPackageStatuses returns the PackageStatuses field if non-nil, zero value otherwise.

### GetPackageStatusesOk

`func (o *TaskGetTaskDetailResponseBody) GetPackageStatusesOk() (*map[string]string, bool)`

GetPackageStatusesOk returns a tuple with the PackageStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageStatuses

`func (o *TaskGetTaskDetailResponseBody) SetPackageStatuses(v map[string]string)`

SetPackageStatuses sets PackageStatuses field to given value.

### HasPackageStatuses

`func (o *TaskGetTaskDetailResponseBody) HasPackageStatuses() bool`

HasPackageStatuses returns a boolean if a field has been set.

### GetPkgCpes

`func (o *TaskGetTaskDetailResponseBody) GetPkgCpes() []PkgCpeChildResponseBody`

GetPkgCpes returns the PkgCpes field if non-nil, zero value otherwise.

### GetPkgCpesOk

`func (o *TaskGetTaskDetailResponseBody) GetPkgCpesOk() (*[]PkgCpeChildResponseBody, bool)`

GetPkgCpesOk returns a tuple with the PkgCpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgCpes

`func (o *TaskGetTaskDetailResponseBody) SetPkgCpes(v []PkgCpeChildResponseBody)`

SetPkgCpes sets PkgCpes field to given value.

### HasPkgCpes

`func (o *TaskGetTaskDetailResponseBody) HasPkgCpes() bool`

HasPkgCpes returns a boolean if a field has been set.

### GetPriority

`func (o *TaskGetTaskDetailResponseBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskGetTaskDetailResponseBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskGetTaskDetailResponseBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetRoleID

`func (o *TaskGetTaskDetailResponseBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *TaskGetTaskDetailResponseBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *TaskGetTaskDetailResponseBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.


### GetRoleName

`func (o *TaskGetTaskDetailResponseBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TaskGetTaskDetailResponseBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TaskGetTaskDetailResponseBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetServer

`func (o *TaskGetTaskDetailResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *TaskGetTaskDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *TaskGetTaskDetailResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.


### GetServerID

`func (o *TaskGetTaskDetailResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *TaskGetTaskDetailResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *TaskGetTaskDetailResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetStatus

`func (o *TaskGetTaskDetailResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskGetTaskDetailResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskGetTaskDetailResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubUserID

`func (o *TaskGetTaskDetailResponseBody) GetSubUserID() int64`

GetSubUserID returns the SubUserID field if non-nil, zero value otherwise.

### GetSubUserIDOk

`func (o *TaskGetTaskDetailResponseBody) GetSubUserIDOk() (*int64, bool)`

GetSubUserIDOk returns a tuple with the SubUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserID

`func (o *TaskGetTaskDetailResponseBody) SetSubUserID(v int64)`

SetSubUserID sets SubUserID field to given value.

### HasSubUserID

`func (o *TaskGetTaskDetailResponseBody) HasSubUserID() bool`

HasSubUserID returns a boolean if a field has been set.

### GetSubUserName

`func (o *TaskGetTaskDetailResponseBody) GetSubUserName() string`

GetSubUserName returns the SubUserName field if non-nil, zero value otherwise.

### GetSubUserNameOk

`func (o *TaskGetTaskDetailResponseBody) GetSubUserNameOk() (*string, bool)`

GetSubUserNameOk returns a tuple with the SubUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserName

`func (o *TaskGetTaskDetailResponseBody) SetSubUserName(v string)`

SetSubUserName sets SubUserName field to given value.

### HasSubUserName

`func (o *TaskGetTaskDetailResponseBody) HasSubUserName() bool`

HasSubUserName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskGetTaskDetailResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskGetTaskDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskGetTaskDetailResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


