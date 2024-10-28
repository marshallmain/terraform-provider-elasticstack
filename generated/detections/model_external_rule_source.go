/*
Security Detections (Elastic Cloud and self-hosted)

You can create rules that automatically turn events and external alerts sent to Elastic Security into detection alerts. These alerts are displayed on the Detections page.

API version: 2023-10-31
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package detections

import (
	"encoding/json"
)

// checks if the ExternalRuleSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalRuleSource{}

// ExternalRuleSource Type of rule source for externally sourced rules, i.e. rules that have an external source, such as the Elastic Prebuilt rules repo.
type ExternalRuleSource struct {
	// Determines whether an external/prebuilt rule has been customized by the user (i.e. any of its fields have been modified and diverged from the base value).
	IsCustomized bool   `json:"is_customized"`
	Type         string `json:"type"`
}

// NewExternalRuleSource instantiates a new ExternalRuleSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalRuleSource(isCustomized bool, type_ string) *ExternalRuleSource {
	this := ExternalRuleSource{}
	this.IsCustomized = isCustomized
	this.Type = type_
	return &this
}

// NewExternalRuleSourceWithDefaults instantiates a new ExternalRuleSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalRuleSourceWithDefaults() *ExternalRuleSource {
	this := ExternalRuleSource{}
	return &this
}

// GetIsCustomized returns the IsCustomized field value
func (o *ExternalRuleSource) GetIsCustomized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustomized
}

// GetIsCustomizedOk returns a tuple with the IsCustomized field value
// and a boolean to check if the value has been set.
func (o *ExternalRuleSource) GetIsCustomizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomized, true
}

// SetIsCustomized sets field value
func (o *ExternalRuleSource) SetIsCustomized(v bool) {
	o.IsCustomized = v
}

// GetType returns the Type field value
func (o *ExternalRuleSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalRuleSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalRuleSource) SetType(v string) {
	o.Type = v
}

func (o ExternalRuleSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalRuleSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_customized"] = o.IsCustomized
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableExternalRuleSource struct {
	value *ExternalRuleSource
	isSet bool
}

func (v NullableExternalRuleSource) Get() *ExternalRuleSource {
	return v.value
}

func (v *NullableExternalRuleSource) Set(val *ExternalRuleSource) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalRuleSource) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalRuleSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalRuleSource(val *ExternalRuleSource) *NullableExternalRuleSource {
	return &NullableExternalRuleSource{value: val, isSet: true}
}

func (v NullableExternalRuleSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalRuleSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
