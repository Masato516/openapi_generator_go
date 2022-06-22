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

// RoleGetRoleDetailResponseBody struct for RoleGetRoleDetailResponseBody
type RoleGetRoleDetailResponseBody struct {
	// AllTaskCount of server role
	AllTaskCount *int64 `json:"allTaskCount,omitempty"`
	// created time of server role
	CreatedAt time.Time `json:"createdAt"`
	// envMetricV2s of server role
	EnvMetricV2s []EnvMetricV2ResponseBody `json:"envMetricV2s,omitempty"`
	// envMetricV3s of server role
	EnvMetricV3s []EnvMetricV3ResponseBody `json:"envMetricV3s,omitempty"`
	// ID of server role
	Id int64 `json:"id"`
	// isDefault of server role
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name of server role
	Name string `json:"name"`
	// NewTaskCount of server role
	NewTaskCount *int64 `json:"newTaskCount,omitempty"`
	SecMetric *SecMetricResponseBody `json:"secMetric,omitempty"`
	// Servers of server role
	Servers []ServerChildResponseBody `json:"servers,omitempty"`
	// updated time of server role
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewRoleGetRoleDetailResponseBody instantiates a new RoleGetRoleDetailResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGetRoleDetailResponseBody(createdAt time.Time, id int64, name string, updatedAt time.Time) *RoleGetRoleDetailResponseBody {
	this := RoleGetRoleDetailResponseBody{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewRoleGetRoleDetailResponseBodyWithDefaults instantiates a new RoleGetRoleDetailResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGetRoleDetailResponseBodyWithDefaults() *RoleGetRoleDetailResponseBody {
	this := RoleGetRoleDetailResponseBody{}
	return &this
}

// GetAllTaskCount returns the AllTaskCount field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetAllTaskCount() int64 {
	if o == nil || o.AllTaskCount == nil {
		var ret int64
		return ret
	}
	return *o.AllTaskCount
}

// GetAllTaskCountOk returns a tuple with the AllTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetAllTaskCountOk() (*int64, bool) {
	if o == nil || o.AllTaskCount == nil {
		return nil, false
	}
	return o.AllTaskCount, true
}

// HasAllTaskCount returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasAllTaskCount() bool {
	if o != nil && o.AllTaskCount != nil {
		return true
	}

	return false
}

// SetAllTaskCount gets a reference to the given int64 and assigns it to the AllTaskCount field.
func (o *RoleGetRoleDetailResponseBody) SetAllTaskCount(v int64) {
	o.AllTaskCount = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RoleGetRoleDetailResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RoleGetRoleDetailResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvMetricV2s returns the EnvMetricV2s field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetEnvMetricV2s() []EnvMetricV2ResponseBody {
	if o == nil || o.EnvMetricV2s == nil {
		var ret []EnvMetricV2ResponseBody
		return ret
	}
	return o.EnvMetricV2s
}

// GetEnvMetricV2sOk returns a tuple with the EnvMetricV2s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetEnvMetricV2sOk() ([]EnvMetricV2ResponseBody, bool) {
	if o == nil || o.EnvMetricV2s == nil {
		return nil, false
	}
	return o.EnvMetricV2s, true
}

// HasEnvMetricV2s returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasEnvMetricV2s() bool {
	if o != nil && o.EnvMetricV2s != nil {
		return true
	}

	return false
}

// SetEnvMetricV2s gets a reference to the given []EnvMetricV2ResponseBody and assigns it to the EnvMetricV2s field.
func (o *RoleGetRoleDetailResponseBody) SetEnvMetricV2s(v []EnvMetricV2ResponseBody) {
	o.EnvMetricV2s = v
}

// GetEnvMetricV3s returns the EnvMetricV3s field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetEnvMetricV3s() []EnvMetricV3ResponseBody {
	if o == nil || o.EnvMetricV3s == nil {
		var ret []EnvMetricV3ResponseBody
		return ret
	}
	return o.EnvMetricV3s
}

// GetEnvMetricV3sOk returns a tuple with the EnvMetricV3s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetEnvMetricV3sOk() ([]EnvMetricV3ResponseBody, bool) {
	if o == nil || o.EnvMetricV3s == nil {
		return nil, false
	}
	return o.EnvMetricV3s, true
}

// HasEnvMetricV3s returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasEnvMetricV3s() bool {
	if o != nil && o.EnvMetricV3s != nil {
		return true
	}

	return false
}

// SetEnvMetricV3s gets a reference to the given []EnvMetricV3ResponseBody and assigns it to the EnvMetricV3s field.
func (o *RoleGetRoleDetailResponseBody) SetEnvMetricV3s(v []EnvMetricV3ResponseBody) {
	o.EnvMetricV3s = v
}

// GetId returns the Id field value
func (o *RoleGetRoleDetailResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleGetRoleDetailResponseBody) SetId(v int64) {
	o.Id = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RoleGetRoleDetailResponseBody) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value
func (o *RoleGetRoleDetailResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleGetRoleDetailResponseBody) SetName(v string) {
	o.Name = v
}

// GetNewTaskCount returns the NewTaskCount field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetNewTaskCount() int64 {
	if o == nil || o.NewTaskCount == nil {
		var ret int64
		return ret
	}
	return *o.NewTaskCount
}

// GetNewTaskCountOk returns a tuple with the NewTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetNewTaskCountOk() (*int64, bool) {
	if o == nil || o.NewTaskCount == nil {
		return nil, false
	}
	return o.NewTaskCount, true
}

// HasNewTaskCount returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasNewTaskCount() bool {
	if o != nil && o.NewTaskCount != nil {
		return true
	}

	return false
}

// SetNewTaskCount gets a reference to the given int64 and assigns it to the NewTaskCount field.
func (o *RoleGetRoleDetailResponseBody) SetNewTaskCount(v int64) {
	o.NewTaskCount = &v
}

// GetSecMetric returns the SecMetric field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetSecMetric() SecMetricResponseBody {
	if o == nil || o.SecMetric == nil {
		var ret SecMetricResponseBody
		return ret
	}
	return *o.SecMetric
}

// GetSecMetricOk returns a tuple with the SecMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetSecMetricOk() (*SecMetricResponseBody, bool) {
	if o == nil || o.SecMetric == nil {
		return nil, false
	}
	return o.SecMetric, true
}

// HasSecMetric returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasSecMetric() bool {
	if o != nil && o.SecMetric != nil {
		return true
	}

	return false
}

// SetSecMetric gets a reference to the given SecMetricResponseBody and assigns it to the SecMetric field.
func (o *RoleGetRoleDetailResponseBody) SetSecMetric(v SecMetricResponseBody) {
	o.SecMetric = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *RoleGetRoleDetailResponseBody) GetServers() []ServerChildResponseBody {
	if o == nil || o.Servers == nil {
		var ret []ServerChildResponseBody
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetServersOk() ([]ServerChildResponseBody, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *RoleGetRoleDetailResponseBody) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerChildResponseBody and assigns it to the Servers field.
func (o *RoleGetRoleDetailResponseBody) SetServers(v []ServerChildResponseBody) {
	o.Servers = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RoleGetRoleDetailResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleGetRoleDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RoleGetRoleDetailResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RoleGetRoleDetailResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllTaskCount != nil {
		toSerialize["allTaskCount"] = o.AllTaskCount
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.EnvMetricV2s != nil {
		toSerialize["envMetricV2s"] = o.EnvMetricV2s
	}
	if o.EnvMetricV3s != nil {
		toSerialize["envMetricV3s"] = o.EnvMetricV3s
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewTaskCount != nil {
		toSerialize["newTaskCount"] = o.NewTaskCount
	}
	if o.SecMetric != nil {
		toSerialize["secMetric"] = o.SecMetric
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRoleGetRoleDetailResponseBody struct {
	value *RoleGetRoleDetailResponseBody
	isSet bool
}

func (v NullableRoleGetRoleDetailResponseBody) Get() *RoleGetRoleDetailResponseBody {
	return v.value
}

func (v *NullableRoleGetRoleDetailResponseBody) Set(val *RoleGetRoleDetailResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGetRoleDetailResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGetRoleDetailResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGetRoleDetailResponseBody(val *RoleGetRoleDetailResponseBody) *NullableRoleGetRoleDetailResponseBody {
	return &NullableRoleGetRoleDetailResponseBody{value: val, isSet: true}
}

func (v NullableRoleGetRoleDetailResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGetRoleDetailResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

