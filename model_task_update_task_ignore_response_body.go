/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TaskUpdateTaskIgnoreResponseBody struct for TaskUpdateTaskIgnoreResponseBody
type TaskUpdateTaskIgnoreResponseBody struct {
	// advisoryIDs of cve
	AdvisoryIDs []string `json:"advisoryIDs,omitempty"`
	// ApplyingPatchOn of task
	ApplyingPatchOn *string `json:"applyingPatchOn,omitempty"`
	// Comment of task
	Comments []TaskCommentResponseBody `json:"comments,omitempty"`
	// created time of task
	CreatedAt time.Time `json:"createdAt"`
	// CVE ID of task
	CveID string `json:"cveID"`
	// Key Value of CveID and Cvss of task
	Cvss *map[string]*os.File `json:"cvss,omitempty"`
	// DetectionMethod of task
	DetectionMethods []DetectionMethodResponseBody `json:"detectionMethods,omitempty"`
	// DetectionTools of task
	DetectionTools []DetectionToolResponseBody `json:"detectionTools,omitempty"`
	// ID of task
	Id int64 `json:"id"`
	// Ignore of task
	Ignore bool `json:"ignore"`
	// Ignore until of task
	IgnoreUntil *string `json:"ignoreUntil,omitempty"`
	// MainUserID of task
	MainUserID *int64 `json:"mainUserID,omitempty"`
	// MainUserName of task
	MainUserName *string `json:"mainUserName,omitempty"`
	// packageStatus of task
	PackageStatuses *map[string]string `json:"packageStatuses,omitempty"`
	// Pcakge And Cpe list of task
	PkgCpes []PkgCpeChildResponseBody `json:"pkgCpes,omitempty"`
	// Priority of task
	Priority string `json:"priority"`
	// ServerRoleID of task
	RoleID int64 `json:"roleID"`
	// ServerRoleName of task
	RoleName string `json:"roleName"`
	Server ServerChildResponseBody `json:"server"`
	// ServerID of task
	ServerID int64 `json:"serverID"`
	// Status of task
	Status string `json:"status"`
	// SubUserID of task
	SubUserID *int64 `json:"subUserID,omitempty"`
	// SubUserName of task
	SubUserName *string `json:"subUserName,omitempty"`
	// updated time of task
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewTaskUpdateTaskIgnoreResponseBody instantiates a new TaskUpdateTaskIgnoreResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskUpdateTaskIgnoreResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, priority string, roleID int64, roleName string, server ServerChildResponseBody, serverID int64, status string, updatedAt time.Time) *TaskUpdateTaskIgnoreResponseBody {
	this := TaskUpdateTaskIgnoreResponseBody{}
	this.CreatedAt = createdAt
	this.CveID = cveID
	this.Id = id
	this.Ignore = ignore
	this.Priority = priority
	this.RoleID = roleID
	this.RoleName = roleName
	this.Server = server
	this.ServerID = serverID
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewTaskUpdateTaskIgnoreResponseBodyWithDefaults instantiates a new TaskUpdateTaskIgnoreResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskUpdateTaskIgnoreResponseBodyWithDefaults() *TaskUpdateTaskIgnoreResponseBody {
	this := TaskUpdateTaskIgnoreResponseBody{}
	return &this
}

// GetAdvisoryIDs returns the AdvisoryIDs field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetAdvisoryIDs() []string {
	if o == nil || o.AdvisoryIDs == nil {
		var ret []string
		return ret
	}
	return o.AdvisoryIDs
}

// GetAdvisoryIDsOk returns a tuple with the AdvisoryIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetAdvisoryIDsOk() ([]string, bool) {
	if o == nil || o.AdvisoryIDs == nil {
		return nil, false
	}
	return o.AdvisoryIDs, true
}

// HasAdvisoryIDs returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasAdvisoryIDs() bool {
	if o != nil && o.AdvisoryIDs != nil {
		return true
	}

	return false
}

// SetAdvisoryIDs gets a reference to the given []string and assigns it to the AdvisoryIDs field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetAdvisoryIDs(v []string) {
	o.AdvisoryIDs = v
}

// GetApplyingPatchOn returns the ApplyingPatchOn field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetApplyingPatchOn() string {
	if o == nil || o.ApplyingPatchOn == nil {
		var ret string
		return ret
	}
	return *o.ApplyingPatchOn
}

// GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetApplyingPatchOnOk() (*string, bool) {
	if o == nil || o.ApplyingPatchOn == nil {
		return nil, false
	}
	return o.ApplyingPatchOn, true
}

// HasApplyingPatchOn returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasApplyingPatchOn() bool {
	if o != nil && o.ApplyingPatchOn != nil {
		return true
	}

	return false
}

// SetApplyingPatchOn gets a reference to the given string and assigns it to the ApplyingPatchOn field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetApplyingPatchOn(v string) {
	o.ApplyingPatchOn = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetComments() []TaskCommentResponseBody {
	if o == nil || o.Comments == nil {
		var ret []TaskCommentResponseBody
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetCommentsOk() ([]TaskCommentResponseBody, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []TaskCommentResponseBody and assigns it to the Comments field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetComments(v []TaskCommentResponseBody) {
	o.Comments = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCveID returns the CveID field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetCveID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CveID
}

// GetCveIDOk returns a tuple with the CveID field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetCveIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveID, true
}

// SetCveID sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetCveID(v string) {
	o.CveID = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetCvss() map[string]*os.File {
	if o == nil || o.Cvss == nil {
		var ret map[string]*os.File
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetCvssOk() (*map[string]*os.File, bool) {
	if o == nil || o.Cvss == nil {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasCvss() bool {
	if o != nil && o.Cvss != nil {
		return true
	}

	return false
}

// SetCvss gets a reference to the given map[string]*os.File and assigns it to the Cvss field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetCvss(v map[string]*os.File) {
	o.Cvss = &v
}

// GetDetectionMethods returns the DetectionMethods field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetDetectionMethods() []DetectionMethodResponseBody {
	if o == nil || o.DetectionMethods == nil {
		var ret []DetectionMethodResponseBody
		return ret
	}
	return o.DetectionMethods
}

// GetDetectionMethodsOk returns a tuple with the DetectionMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetDetectionMethodsOk() ([]DetectionMethodResponseBody, bool) {
	if o == nil || o.DetectionMethods == nil {
		return nil, false
	}
	return o.DetectionMethods, true
}

// HasDetectionMethods returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasDetectionMethods() bool {
	if o != nil && o.DetectionMethods != nil {
		return true
	}

	return false
}

// SetDetectionMethods gets a reference to the given []DetectionMethodResponseBody and assigns it to the DetectionMethods field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetDetectionMethods(v []DetectionMethodResponseBody) {
	o.DetectionMethods = v
}

// GetDetectionTools returns the DetectionTools field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetDetectionTools() []DetectionToolResponseBody {
	if o == nil || o.DetectionTools == nil {
		var ret []DetectionToolResponseBody
		return ret
	}
	return o.DetectionTools
}

// GetDetectionToolsOk returns a tuple with the DetectionTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetDetectionToolsOk() ([]DetectionToolResponseBody, bool) {
	if o == nil || o.DetectionTools == nil {
		return nil, false
	}
	return o.DetectionTools, true
}

// HasDetectionTools returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasDetectionTools() bool {
	if o != nil && o.DetectionTools != nil {
		return true
	}

	return false
}

// SetDetectionTools gets a reference to the given []DetectionToolResponseBody and assigns it to the DetectionTools field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetDetectionTools(v []DetectionToolResponseBody) {
	o.DetectionTools = v
}

// GetId returns the Id field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetId(v int64) {
	o.Id = v
}

// GetIgnore returns the Ignore field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetIgnore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ignore
}

// GetIgnoreOk returns a tuple with the Ignore field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetIgnoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ignore, true
}

// SetIgnore sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetIgnore(v bool) {
	o.Ignore = v
}

// GetIgnoreUntil returns the IgnoreUntil field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetIgnoreUntil() string {
	if o == nil || o.IgnoreUntil == nil {
		var ret string
		return ret
	}
	return *o.IgnoreUntil
}

// GetIgnoreUntilOk returns a tuple with the IgnoreUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetIgnoreUntilOk() (*string, bool) {
	if o == nil || o.IgnoreUntil == nil {
		return nil, false
	}
	return o.IgnoreUntil, true
}

// HasIgnoreUntil returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasIgnoreUntil() bool {
	if o != nil && o.IgnoreUntil != nil {
		return true
	}

	return false
}

// SetIgnoreUntil gets a reference to the given string and assigns it to the IgnoreUntil field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetIgnoreUntil(v string) {
	o.IgnoreUntil = &v
}

// GetMainUserID returns the MainUserID field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetMainUserID() int64 {
	if o == nil || o.MainUserID == nil {
		var ret int64
		return ret
	}
	return *o.MainUserID
}

// GetMainUserIDOk returns a tuple with the MainUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetMainUserIDOk() (*int64, bool) {
	if o == nil || o.MainUserID == nil {
		return nil, false
	}
	return o.MainUserID, true
}

// HasMainUserID returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasMainUserID() bool {
	if o != nil && o.MainUserID != nil {
		return true
	}

	return false
}

// SetMainUserID gets a reference to the given int64 and assigns it to the MainUserID field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetMainUserID(v int64) {
	o.MainUserID = &v
}

// GetMainUserName returns the MainUserName field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetMainUserName() string {
	if o == nil || o.MainUserName == nil {
		var ret string
		return ret
	}
	return *o.MainUserName
}

// GetMainUserNameOk returns a tuple with the MainUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetMainUserNameOk() (*string, bool) {
	if o == nil || o.MainUserName == nil {
		return nil, false
	}
	return o.MainUserName, true
}

// HasMainUserName returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasMainUserName() bool {
	if o != nil && o.MainUserName != nil {
		return true
	}

	return false
}

// SetMainUserName gets a reference to the given string and assigns it to the MainUserName field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetMainUserName(v string) {
	o.MainUserName = &v
}

// GetPackageStatuses returns the PackageStatuses field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetPackageStatuses() map[string]string {
	if o == nil || o.PackageStatuses == nil {
		var ret map[string]string
		return ret
	}
	return *o.PackageStatuses
}

// GetPackageStatusesOk returns a tuple with the PackageStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetPackageStatusesOk() (*map[string]string, bool) {
	if o == nil || o.PackageStatuses == nil {
		return nil, false
	}
	return o.PackageStatuses, true
}

// HasPackageStatuses returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasPackageStatuses() bool {
	if o != nil && o.PackageStatuses != nil {
		return true
	}

	return false
}

// SetPackageStatuses gets a reference to the given map[string]string and assigns it to the PackageStatuses field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetPackageStatuses(v map[string]string) {
	o.PackageStatuses = &v
}

// GetPkgCpes returns the PkgCpes field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetPkgCpes() []PkgCpeChildResponseBody {
	if o == nil || o.PkgCpes == nil {
		var ret []PkgCpeChildResponseBody
		return ret
	}
	return o.PkgCpes
}

// GetPkgCpesOk returns a tuple with the PkgCpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetPkgCpesOk() ([]PkgCpeChildResponseBody, bool) {
	if o == nil || o.PkgCpes == nil {
		return nil, false
	}
	return o.PkgCpes, true
}

// HasPkgCpes returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasPkgCpes() bool {
	if o != nil && o.PkgCpes != nil {
		return true
	}

	return false
}

// SetPkgCpes gets a reference to the given []PkgCpeChildResponseBody and assigns it to the PkgCpes field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetPkgCpes(v []PkgCpeChildResponseBody) {
	o.PkgCpes = v
}

// GetPriority returns the Priority field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetPriority(v string) {
	o.Priority = v
}

// GetRoleID returns the RoleID field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetRoleID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RoleID
}

// GetRoleIDOk returns a tuple with the RoleID field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetRoleIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleID, true
}

// SetRoleID sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetRoleID(v int64) {
	o.RoleID = v
}

// GetRoleName returns the RoleName field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetRoleName(v string) {
	o.RoleName = v
}

// GetServer returns the Server field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetServer() ServerChildResponseBody {
	if o == nil {
		var ret ServerChildResponseBody
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetServerOk() (*ServerChildResponseBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetServer(v ServerChildResponseBody) {
	o.Server = v
}

// GetServerID returns the ServerID field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetServerID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerID
}

// GetServerIDOk returns a tuple with the ServerID field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetServerIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerID, true
}

// SetServerID sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetServerID(v int64) {
	o.ServerID = v
}

// GetStatus returns the Status field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetStatus(v string) {
	o.Status = v
}

// GetSubUserID returns the SubUserID field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetSubUserID() int64 {
	if o == nil || o.SubUserID == nil {
		var ret int64
		return ret
	}
	return *o.SubUserID
}

// GetSubUserIDOk returns a tuple with the SubUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetSubUserIDOk() (*int64, bool) {
	if o == nil || o.SubUserID == nil {
		return nil, false
	}
	return o.SubUserID, true
}

// HasSubUserID returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasSubUserID() bool {
	if o != nil && o.SubUserID != nil {
		return true
	}

	return false
}

// SetSubUserID gets a reference to the given int64 and assigns it to the SubUserID field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetSubUserID(v int64) {
	o.SubUserID = &v
}

// GetSubUserName returns the SubUserName field value if set, zero value otherwise.
func (o *TaskUpdateTaskIgnoreResponseBody) GetSubUserName() string {
	if o == nil || o.SubUserName == nil {
		var ret string
		return ret
	}
	return *o.SubUserName
}

// GetSubUserNameOk returns a tuple with the SubUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetSubUserNameOk() (*string, bool) {
	if o == nil || o.SubUserName == nil {
		return nil, false
	}
	return o.SubUserName, true
}

// HasSubUserName returns a boolean if a field has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) HasSubUserName() bool {
	if o != nil && o.SubUserName != nil {
		return true
	}

	return false
}

// SetSubUserName gets a reference to the given string and assigns it to the SubUserName field.
func (o *TaskUpdateTaskIgnoreResponseBody) SetSubUserName(v string) {
	o.SubUserName = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaskUpdateTaskIgnoreResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskUpdateTaskIgnoreResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaskUpdateTaskIgnoreResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o TaskUpdateTaskIgnoreResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdvisoryIDs != nil {
		toSerialize["advisoryIDs"] = o.AdvisoryIDs
	}
	if o.ApplyingPatchOn != nil {
		toSerialize["applyingPatchOn"] = o.ApplyingPatchOn
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["cveID"] = o.CveID
	}
	if o.Cvss != nil {
		toSerialize["cvss"] = o.Cvss
	}
	if o.DetectionMethods != nil {
		toSerialize["detectionMethods"] = o.DetectionMethods
	}
	if o.DetectionTools != nil {
		toSerialize["detectionTools"] = o.DetectionTools
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ignore"] = o.Ignore
	}
	if o.IgnoreUntil != nil {
		toSerialize["ignoreUntil"] = o.IgnoreUntil
	}
	if o.MainUserID != nil {
		toSerialize["mainUserID"] = o.MainUserID
	}
	if o.MainUserName != nil {
		toSerialize["mainUserName"] = o.MainUserName
	}
	if o.PackageStatuses != nil {
		toSerialize["packageStatuses"] = o.PackageStatuses
	}
	if o.PkgCpes != nil {
		toSerialize["pkgCpes"] = o.PkgCpes
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["roleID"] = o.RoleID
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if true {
		toSerialize["serverID"] = o.ServerID
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.SubUserID != nil {
		toSerialize["subUserID"] = o.SubUserID
	}
	if o.SubUserName != nil {
		toSerialize["subUserName"] = o.SubUserName
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTaskUpdateTaskIgnoreResponseBody struct {
	value *TaskUpdateTaskIgnoreResponseBody
	isSet bool
}

func (v NullableTaskUpdateTaskIgnoreResponseBody) Get() *TaskUpdateTaskIgnoreResponseBody {
	return v.value
}

func (v *NullableTaskUpdateTaskIgnoreResponseBody) Set(val *TaskUpdateTaskIgnoreResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskUpdateTaskIgnoreResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskUpdateTaskIgnoreResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskUpdateTaskIgnoreResponseBody(val *TaskUpdateTaskIgnoreResponseBody) *NullableTaskUpdateTaskIgnoreResponseBody {
	return &NullableTaskUpdateTaskIgnoreResponseBody{value: val, isSet: true}
}

func (v NullableTaskUpdateTaskIgnoreResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskUpdateTaskIgnoreResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


