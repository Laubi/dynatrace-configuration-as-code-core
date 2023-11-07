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

// checks if the AppendLevelPolicyBindingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppendLevelPolicyBindingsRequestDto{}

// AppendLevelPolicyBindingsRequestDto struct for AppendLevelPolicyBindingsRequestDto
type AppendLevelPolicyBindingsRequestDto struct {
	// A list of user groups (specified by IDs) to which the policy applies.
	Groups []string `json:"groups"`
}

// NewAppendLevelPolicyBindingsRequestDto instantiates a new AppendLevelPolicyBindingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppendLevelPolicyBindingsRequestDto(groups []string) *AppendLevelPolicyBindingsRequestDto {
	this := AppendLevelPolicyBindingsRequestDto{}
	this.Groups = groups
	return &this
}

// NewAppendLevelPolicyBindingsRequestDtoWithDefaults instantiates a new AppendLevelPolicyBindingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppendLevelPolicyBindingsRequestDtoWithDefaults() *AppendLevelPolicyBindingsRequestDto {
	this := AppendLevelPolicyBindingsRequestDto{}
	return &this
}

// GetGroups returns the Groups field value
func (o *AppendLevelPolicyBindingsRequestDto) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *AppendLevelPolicyBindingsRequestDto) GetGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *AppendLevelPolicyBindingsRequestDto) SetGroups(v []string) {
	o.Groups = v
}

func (o AppendLevelPolicyBindingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppendLevelPolicyBindingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

type NullableAppendLevelPolicyBindingsRequestDto struct {
	value *AppendLevelPolicyBindingsRequestDto
	isSet bool
}

func (v NullableAppendLevelPolicyBindingsRequestDto) Get() *AppendLevelPolicyBindingsRequestDto {
	return v.value
}

func (v *NullableAppendLevelPolicyBindingsRequestDto) Set(val *AppendLevelPolicyBindingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppendLevelPolicyBindingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppendLevelPolicyBindingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppendLevelPolicyBindingsRequestDto(val *AppendLevelPolicyBindingsRequestDto) *NullableAppendLevelPolicyBindingsRequestDto {
	return &NullableAppendLevelPolicyBindingsRequestDto{value: val, isSet: true}
}

func (v NullableAppendLevelPolicyBindingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppendLevelPolicyBindingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
