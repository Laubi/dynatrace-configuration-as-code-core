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

// checks if the PolicyOverviewDtoList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyOverviewDtoList{}

// PolicyOverviewDtoList struct for PolicyOverviewDtoList
type PolicyOverviewDtoList struct {
	// A list of policies.
	PolicyOverviewList []PolicyOverview `json:"policyOverviewList"`
}

// NewPolicyOverviewDtoList instantiates a new PolicyOverviewDtoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyOverviewDtoList(policyOverviewList []PolicyOverview) *PolicyOverviewDtoList {
	this := PolicyOverviewDtoList{}
	this.PolicyOverviewList = policyOverviewList
	return &this
}

// NewPolicyOverviewDtoListWithDefaults instantiates a new PolicyOverviewDtoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyOverviewDtoListWithDefaults() *PolicyOverviewDtoList {
	this := PolicyOverviewDtoList{}
	return &this
}

// GetPolicyOverviewList returns the PolicyOverviewList field value
func (o *PolicyOverviewDtoList) GetPolicyOverviewList() []PolicyOverview {
	if o == nil {
		var ret []PolicyOverview
		return ret
	}

	return o.PolicyOverviewList
}

// GetPolicyOverviewListOk returns a tuple with the PolicyOverviewList field value
// and a boolean to check if the value has been set.
func (o *PolicyOverviewDtoList) GetPolicyOverviewListOk() ([]PolicyOverview, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyOverviewList, true
}

// SetPolicyOverviewList sets field value
func (o *PolicyOverviewDtoList) SetPolicyOverviewList(v []PolicyOverview) {
	o.PolicyOverviewList = v
}

func (o PolicyOverviewDtoList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyOverviewDtoList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policyOverviewList"] = o.PolicyOverviewList
	return toSerialize, nil
}

type NullablePolicyOverviewDtoList struct {
	value *PolicyOverviewDtoList
	isSet bool
}

func (v NullablePolicyOverviewDtoList) Get() *PolicyOverviewDtoList {
	return v.value
}

func (v *NullablePolicyOverviewDtoList) Set(val *PolicyOverviewDtoList) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyOverviewDtoList) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyOverviewDtoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyOverviewDtoList(val *PolicyOverviewDtoList) *NullablePolicyOverviewDtoList {
	return &NullablePolicyOverviewDtoList{value: val, isSet: true}
}

func (v NullablePolicyOverviewDtoList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyOverviewDtoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
