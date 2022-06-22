# TaskListResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | Pointer to **[]string** | advisoryIDs of cve | [optional] 
**ApplyingPatchOn** | Pointer to **string** | ApplyingPatchOn of task | [optional] 
**CreatedAt** | **time.Time** | created time of task | 
**CveID** | **string** | CVE ID of task | 
**DetectionTools** | Pointer to [**[]DetectionToolResponseBody**](DetectionToolResponseBody.md) | DetectionTools of task | [optional] 
**HasExploit** | Pointer to **bool** | hasExploit of cve | [optional] 
**HasMitigation** | Pointer to **bool** | hasMitigation of cve | [optional] 
**HasWorkaround** | Pointer to **bool** | hasWorkaroundof cve | [optional] 
**Id** | **int64** | ID of task | 
**Ignore** | **bool** | Ignore of task | 
**IgnoreUntil** | Pointer to **string** | Ignore until of task | [optional] 
**MainUserID** | Pointer to **int64** | MainUserID of task | [optional] 
**MainUserName** | Pointer to **string** | MainUserName of task | [optional] 
**OsFamily** | **string** | OS Name of server | 
**OsVersion** | **string** | OS Version of server | 
**PkgCpeNames** | Pointer to **[]string** | Package And CPE Names of task | [optional] 
**PkgNotFixedYet** | Pointer to **bool** | Flag of Not Fixed Yet of task | [optional] 
**Priority** | **string** | Priority of task | 
**RoleID** | **int64** | ServerRoleID of task | 
**RoleName** | **string** | ServerRoleName of task | 
**ServerID** | **int64** | ServerID of task | 
**ServerName** | **string** | ServerName of task | 
**ServerTags** | Pointer to **[]string** | ServerTags of task | [optional] 
**ServerUuid** | **string** | ServerUUID of task | 
**Status** | **string** | Status of task | 
**SubUserID** | Pointer to **int64** | SubUserID of task | [optional] 
**SubUserName** | Pointer to **string** | SubUserName of task | [optional] 
**UpdatedAt** | **time.Time** | updated time of task | 

## Methods

### NewTaskListResponseBody

`func NewTaskListResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, osFamily string, osVersion string, priority string, roleID int64, roleName string, serverID int64, serverName string, serverUuid string, status string, updatedAt time.Time, ) *TaskListResponseBody`

NewTaskListResponseBody instantiates a new TaskListResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListResponseBodyWithDefaults

`func NewTaskListResponseBodyWithDefaults() *TaskListResponseBody`

NewTaskListResponseBodyWithDefaults instantiates a new TaskListResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryIDs

`func (o *TaskListResponseBody) GetAdvisoryIDs() []string`

GetAdvisoryIDs returns the AdvisoryIDs field if non-nil, zero value otherwise.

### GetAdvisoryIDsOk

`func (o *TaskListResponseBody) GetAdvisoryIDsOk() (*[]string, bool)`

GetAdvisoryIDsOk returns a tuple with the AdvisoryIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryIDs

`func (o *TaskListResponseBody) SetAdvisoryIDs(v []string)`

SetAdvisoryIDs sets AdvisoryIDs field to given value.

### HasAdvisoryIDs

`func (o *TaskListResponseBody) HasAdvisoryIDs() bool`

HasAdvisoryIDs returns a boolean if a field has been set.

### GetApplyingPatchOn

`func (o *TaskListResponseBody) GetApplyingPatchOn() string`

GetApplyingPatchOn returns the ApplyingPatchOn field if non-nil, zero value otherwise.

### GetApplyingPatchOnOk

`func (o *TaskListResponseBody) GetApplyingPatchOnOk() (*string, bool)`

GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyingPatchOn

`func (o *TaskListResponseBody) SetApplyingPatchOn(v string)`

SetApplyingPatchOn sets ApplyingPatchOn field to given value.

### HasApplyingPatchOn

`func (o *TaskListResponseBody) HasApplyingPatchOn() bool`

HasApplyingPatchOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskListResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskListResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskListResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *TaskListResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *TaskListResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *TaskListResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetDetectionTools

`func (o *TaskListResponseBody) GetDetectionTools() []DetectionToolResponseBody`

GetDetectionTools returns the DetectionTools field if non-nil, zero value otherwise.

### GetDetectionToolsOk

`func (o *TaskListResponseBody) GetDetectionToolsOk() (*[]DetectionToolResponseBody, bool)`

GetDetectionToolsOk returns a tuple with the DetectionTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionTools

`func (o *TaskListResponseBody) SetDetectionTools(v []DetectionToolResponseBody)`

SetDetectionTools sets DetectionTools field to given value.

### HasDetectionTools

`func (o *TaskListResponseBody) HasDetectionTools() bool`

HasDetectionTools returns a boolean if a field has been set.

### GetHasExploit

`func (o *TaskListResponseBody) GetHasExploit() bool`

GetHasExploit returns the HasExploit field if non-nil, zero value otherwise.

### GetHasExploitOk

`func (o *TaskListResponseBody) GetHasExploitOk() (*bool, bool)`

GetHasExploitOk returns a tuple with the HasExploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExploit

`func (o *TaskListResponseBody) SetHasExploit(v bool)`

SetHasExploit sets HasExploit field to given value.

### HasHasExploit

`func (o *TaskListResponseBody) HasHasExploit() bool`

HasHasExploit returns a boolean if a field has been set.

### GetHasMitigation

`func (o *TaskListResponseBody) GetHasMitigation() bool`

GetHasMitigation returns the HasMitigation field if non-nil, zero value otherwise.

### GetHasMitigationOk

`func (o *TaskListResponseBody) GetHasMitigationOk() (*bool, bool)`

GetHasMitigationOk returns a tuple with the HasMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMitigation

`func (o *TaskListResponseBody) SetHasMitigation(v bool)`

SetHasMitigation sets HasMitigation field to given value.

### HasHasMitigation

`func (o *TaskListResponseBody) HasHasMitigation() bool`

HasHasMitigation returns a boolean if a field has been set.

### GetHasWorkaround

`func (o *TaskListResponseBody) GetHasWorkaround() bool`

GetHasWorkaround returns the HasWorkaround field if non-nil, zero value otherwise.

### GetHasWorkaroundOk

`func (o *TaskListResponseBody) GetHasWorkaroundOk() (*bool, bool)`

GetHasWorkaroundOk returns a tuple with the HasWorkaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWorkaround

`func (o *TaskListResponseBody) SetHasWorkaround(v bool)`

SetHasWorkaround sets HasWorkaround field to given value.

### HasHasWorkaround

`func (o *TaskListResponseBody) HasHasWorkaround() bool`

HasHasWorkaround returns a boolean if a field has been set.

### GetId

`func (o *TaskListResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskListResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskListResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIgnore

`func (o *TaskListResponseBody) GetIgnore() bool`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *TaskListResponseBody) GetIgnoreOk() (*bool, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *TaskListResponseBody) SetIgnore(v bool)`

SetIgnore sets Ignore field to given value.


### GetIgnoreUntil

`func (o *TaskListResponseBody) GetIgnoreUntil() string`

GetIgnoreUntil returns the IgnoreUntil field if non-nil, zero value otherwise.

### GetIgnoreUntilOk

`func (o *TaskListResponseBody) GetIgnoreUntilOk() (*string, bool)`

GetIgnoreUntilOk returns a tuple with the IgnoreUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUntil

`func (o *TaskListResponseBody) SetIgnoreUntil(v string)`

SetIgnoreUntil sets IgnoreUntil field to given value.

### HasIgnoreUntil

`func (o *TaskListResponseBody) HasIgnoreUntil() bool`

HasIgnoreUntil returns a boolean if a field has been set.

### GetMainUserID

`func (o *TaskListResponseBody) GetMainUserID() int64`

GetMainUserID returns the MainUserID field if non-nil, zero value otherwise.

### GetMainUserIDOk

`func (o *TaskListResponseBody) GetMainUserIDOk() (*int64, bool)`

GetMainUserIDOk returns a tuple with the MainUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserID

`func (o *TaskListResponseBody) SetMainUserID(v int64)`

SetMainUserID sets MainUserID field to given value.

### HasMainUserID

`func (o *TaskListResponseBody) HasMainUserID() bool`

HasMainUserID returns a boolean if a field has been set.

### GetMainUserName

`func (o *TaskListResponseBody) GetMainUserName() string`

GetMainUserName returns the MainUserName field if non-nil, zero value otherwise.

### GetMainUserNameOk

`func (o *TaskListResponseBody) GetMainUserNameOk() (*string, bool)`

GetMainUserNameOk returns a tuple with the MainUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserName

`func (o *TaskListResponseBody) SetMainUserName(v string)`

SetMainUserName sets MainUserName field to given value.

### HasMainUserName

`func (o *TaskListResponseBody) HasMainUserName() bool`

HasMainUserName returns a boolean if a field has been set.

### GetOsFamily

`func (o *TaskListResponseBody) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *TaskListResponseBody) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *TaskListResponseBody) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.


### GetOsVersion

`func (o *TaskListResponseBody) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *TaskListResponseBody) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *TaskListResponseBody) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetPkgCpeNames

`func (o *TaskListResponseBody) GetPkgCpeNames() []string`

GetPkgCpeNames returns the PkgCpeNames field if non-nil, zero value otherwise.

### GetPkgCpeNamesOk

`func (o *TaskListResponseBody) GetPkgCpeNamesOk() (*[]string, bool)`

GetPkgCpeNamesOk returns a tuple with the PkgCpeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgCpeNames

`func (o *TaskListResponseBody) SetPkgCpeNames(v []string)`

SetPkgCpeNames sets PkgCpeNames field to given value.

### HasPkgCpeNames

`func (o *TaskListResponseBody) HasPkgCpeNames() bool`

HasPkgCpeNames returns a boolean if a field has been set.

### GetPkgNotFixedYet

`func (o *TaskListResponseBody) GetPkgNotFixedYet() bool`

GetPkgNotFixedYet returns the PkgNotFixedYet field if non-nil, zero value otherwise.

### GetPkgNotFixedYetOk

`func (o *TaskListResponseBody) GetPkgNotFixedYetOk() (*bool, bool)`

GetPkgNotFixedYetOk returns a tuple with the PkgNotFixedYet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgNotFixedYet

`func (o *TaskListResponseBody) SetPkgNotFixedYet(v bool)`

SetPkgNotFixedYet sets PkgNotFixedYet field to given value.

### HasPkgNotFixedYet

`func (o *TaskListResponseBody) HasPkgNotFixedYet() bool`

HasPkgNotFixedYet returns a boolean if a field has been set.

### GetPriority

`func (o *TaskListResponseBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskListResponseBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskListResponseBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetRoleID

`func (o *TaskListResponseBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *TaskListResponseBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *TaskListResponseBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.


### GetRoleName

`func (o *TaskListResponseBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TaskListResponseBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TaskListResponseBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetServerID

`func (o *TaskListResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *TaskListResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *TaskListResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetServerName

`func (o *TaskListResponseBody) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TaskListResponseBody) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TaskListResponseBody) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerTags

`func (o *TaskListResponseBody) GetServerTags() []string`

GetServerTags returns the ServerTags field if non-nil, zero value otherwise.

### GetServerTagsOk

`func (o *TaskListResponseBody) GetServerTagsOk() (*[]string, bool)`

GetServerTagsOk returns a tuple with the ServerTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTags

`func (o *TaskListResponseBody) SetServerTags(v []string)`

SetServerTags sets ServerTags field to given value.

### HasServerTags

`func (o *TaskListResponseBody) HasServerTags() bool`

HasServerTags returns a boolean if a field has been set.

### GetServerUuid

`func (o *TaskListResponseBody) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *TaskListResponseBody) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *TaskListResponseBody) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetStatus

`func (o *TaskListResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskListResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskListResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubUserID

`func (o *TaskListResponseBody) GetSubUserID() int64`

GetSubUserID returns the SubUserID field if non-nil, zero value otherwise.

### GetSubUserIDOk

`func (o *TaskListResponseBody) GetSubUserIDOk() (*int64, bool)`

GetSubUserIDOk returns a tuple with the SubUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserID

`func (o *TaskListResponseBody) SetSubUserID(v int64)`

SetSubUserID sets SubUserID field to given value.

### HasSubUserID

`func (o *TaskListResponseBody) HasSubUserID() bool`

HasSubUserID returns a boolean if a field has been set.

### GetSubUserName

`func (o *TaskListResponseBody) GetSubUserName() string`

GetSubUserName returns the SubUserName field if non-nil, zero value otherwise.

### GetSubUserNameOk

`func (o *TaskListResponseBody) GetSubUserNameOk() (*string, bool)`

GetSubUserNameOk returns a tuple with the SubUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserName

`func (o *TaskListResponseBody) SetSubUserName(v string)`

SetSubUserName sets SubUserName field to given value.

### HasSubUserName

`func (o *TaskListResponseBody) HasSubUserName() bool`

HasSubUserName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskListResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskListResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskListResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


