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

// checks if the ServiceUserUuidDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUserUuidDto{}

// ServiceUserUuidDto struct for ServiceUserUuidDto
type ServiceUserUuidDto struct {
	// The UUID of the new service user
	Uuid string `json:"uuid"`
}

// NewServiceUserUuidDto instantiates a new ServiceUserUuidDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUserUuidDto(uuid string) *ServiceUserUuidDto {
	this := ServiceUserUuidDto{}
	this.Uuid = uuid
	return &this
}

// NewServiceUserUuidDtoWithDefaults instantiates a new ServiceUserUuidDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUserUuidDtoWithDefaults() *ServiceUserUuidDto {
	this := ServiceUserUuidDto{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ServiceUserUuidDto) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ServiceUserUuidDto) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ServiceUserUuidDto) SetUuid(v string) {
	o.Uuid = v
}

func (o ServiceUserUuidDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceUserUuidDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

type NullableServiceUserUuidDto struct {
	value *ServiceUserUuidDto
	isSet bool
}

func (v NullableServiceUserUuidDto) Get() *ServiceUserUuidDto {
	return v.value
}

func (v *NullableServiceUserUuidDto) Set(val *ServiceUserUuidDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUserUuidDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUserUuidDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUserUuidDto(val *ServiceUserUuidDto) *NullableServiceUserUuidDto {
	return &NullableServiceUserUuidDto{value: val, isSet: true}
}

func (v NullableServiceUserUuidDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUserUuidDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
