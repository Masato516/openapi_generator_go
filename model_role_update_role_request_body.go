/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RoleUpdateRoleRequestBody struct for RoleUpdateRoleRequestBody
type RoleUpdateRoleRequestBody struct {
	// RoleName of role
	RoleName *string `json:"roleName,omitempty"`
}

// NewRoleUpdateRoleRequestBody instantiates a new RoleUpdateRoleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateRoleRequestBody() *RoleUpdateRoleRequestBody {
	this := RoleUpdateRoleRequestBody{}
	return &this
}

// NewRoleUpdateRoleRequestBodyWithDefaults instantiates a new RoleUpdateRoleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateRoleRequestBodyWithDefaults() *RoleUpdateRoleRequestBody {
	this := RoleUpdateRoleRequestBody{}
	return &this
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *RoleUpdateRoleRequestBody) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdateRoleRequestBody) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *RoleUpdateRoleRequestBody) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *RoleUpdateRoleRequestBody) SetRoleName(v string) {
	o.RoleName = &v
}

func (o RoleUpdateRoleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleName != nil {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableRoleUpdateRoleRequestBody struct {
	value *RoleUpdateRoleRequestBody
	isSet bool
}

func (v NullableRoleUpdateRoleRequestBody) Get() *RoleUpdateRoleRequestBody {
	return v.value
}

func (v *NullableRoleUpdateRoleRequestBody) Set(val *RoleUpdateRoleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateRoleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateRoleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateRoleRequestBody(val *RoleUpdateRoleRequestBody) *NullableRoleUpdateRoleRequestBody {
	return &NullableRoleUpdateRoleRequestBody{value: val, isSet: true}
}

func (v NullableRoleUpdateRoleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateRoleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

