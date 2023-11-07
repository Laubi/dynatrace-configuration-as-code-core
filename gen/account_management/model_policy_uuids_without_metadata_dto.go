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

// checks if the PolicyUuidsWithoutMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyUuidsWithoutMetadataDto{}

// PolicyUuidsWithoutMetadataDto struct for PolicyUuidsWithoutMetadataDto
type PolicyUuidsWithoutMetadataDto struct {
	// A list of policies bound to the user group.
	PolicyUuids []string `json:"policyUuids"`
}

// NewPolicyUuidsWithoutMetadataDto instantiates a new PolicyUuidsWithoutMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyUuidsWithoutMetadataDto(policyUuids []string) *PolicyUuidsWithoutMetadataDto {
	this := PolicyUuidsWithoutMetadataDto{}
	this.PolicyUuids = policyUuids
	return &this
}

// NewPolicyUuidsWithoutMetadataDtoWithDefaults instantiates a new PolicyUuidsWithoutMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyUuidsWithoutMetadataDtoWithDefaults() *PolicyUuidsWithoutMetadataDto {
	this := PolicyUuidsWithoutMetadataDto{}
	return &this
}

// GetPolicyUuids returns the PolicyUuids field value
func (o *PolicyUuidsWithoutMetadataDto) GetPolicyUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PolicyUuids
}

// GetPolicyUuidsOk returns a tuple with the PolicyUuids field value
// and a boolean to check if the value has been set.
func (o *PolicyUuidsWithoutMetadataDto) GetPolicyUuidsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyUuids, true
}

// SetPolicyUuids sets field value
func (o *PolicyUuidsWithoutMetadataDto) SetPolicyUuids(v []string) {
	o.PolicyUuids = v
}

func (o PolicyUuidsWithoutMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyUuidsWithoutMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policyUuids"] = o.PolicyUuids
	return toSerialize, nil
}

type NullablePolicyUuidsWithoutMetadataDto struct {
	value *PolicyUuidsWithoutMetadataDto
	isSet bool
}

func (v NullablePolicyUuidsWithoutMetadataDto) Get() *PolicyUuidsWithoutMetadataDto {
	return v.value
}

func (v *NullablePolicyUuidsWithoutMetadataDto) Set(val *PolicyUuidsWithoutMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyUuidsWithoutMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyUuidsWithoutMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyUuidsWithoutMetadataDto(val *PolicyUuidsWithoutMetadataDto) *NullablePolicyUuidsWithoutMetadataDto {
	return &NullablePolicyUuidsWithoutMetadataDto{value: val, isSet: true}
}

func (v NullablePolicyUuidsWithoutMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyUuidsWithoutMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
