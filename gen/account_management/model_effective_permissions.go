/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmanagement

import (
	"encoding/json"
)

// checks if the EffectivePermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectivePermissions{}

// EffectivePermissions struct for EffectivePermissions
type EffectivePermissions struct {
	// List of effective permissions.
	EffectivePermissions []EffectivePermission `json:"effectivePermissions"`
}

// NewEffectivePermissions instantiates a new EffectivePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectivePermissions(effectivePermissions []EffectivePermission) *EffectivePermissions {
	this := EffectivePermissions{}
	this.EffectivePermissions = effectivePermissions
	return &this
}

// NewEffectivePermissionsWithDefaults instantiates a new EffectivePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectivePermissionsWithDefaults() *EffectivePermissions {
	this := EffectivePermissions{}
	return &this
}

// GetEffectivePermissions returns the EffectivePermissions field value
func (o *EffectivePermissions) GetEffectivePermissions() []EffectivePermission {
	if o == nil {
		var ret []EffectivePermission
		return ret
	}

	return o.EffectivePermissions
}

// GetEffectivePermissionsOk returns a tuple with the EffectivePermissions field value
// and a boolean to check if the value has been set.
func (o *EffectivePermissions) GetEffectivePermissionsOk() ([]EffectivePermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectivePermissions, true
}

// SetEffectivePermissions sets field value
func (o *EffectivePermissions) SetEffectivePermissions(v []EffectivePermission) {
	o.EffectivePermissions = v
}

func (o EffectivePermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectivePermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["effectivePermissions"] = o.EffectivePermissions
	return toSerialize, nil
}

type NullableEffectivePermissions struct {
	value *EffectivePermissions
	isSet bool
}

func (v NullableEffectivePermissions) Get() *EffectivePermissions {
	return v.value
}

func (v *NullableEffectivePermissions) Set(val *EffectivePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectivePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectivePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectivePermissions(val *EffectivePermissions) *NullableEffectivePermissions {
	return &NullableEffectivePermissions{value: val, isSet: true}
}

func (v NullableEffectivePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectivePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
