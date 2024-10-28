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

// checks if the BulkActionEditPayloadRuleActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkActionEditPayloadRuleActions{}

// BulkActionEditPayloadRuleActions struct for BulkActionEditPayloadRuleActions
type BulkActionEditPayloadRuleActions struct {
	Type  string                                `json:"type"`
	Value BulkActionEditPayloadRuleActionsValue `json:"value"`
}

// NewBulkActionEditPayloadRuleActions instantiates a new BulkActionEditPayloadRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkActionEditPayloadRuleActions(type_ string, value BulkActionEditPayloadRuleActionsValue) *BulkActionEditPayloadRuleActions {
	this := BulkActionEditPayloadRuleActions{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewBulkActionEditPayloadRuleActionsWithDefaults instantiates a new BulkActionEditPayloadRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkActionEditPayloadRuleActionsWithDefaults() *BulkActionEditPayloadRuleActions {
	this := BulkActionEditPayloadRuleActions{}
	return &this
}

// GetType returns the Type field value
func (o *BulkActionEditPayloadRuleActions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkActionEditPayloadRuleActions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BulkActionEditPayloadRuleActions) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *BulkActionEditPayloadRuleActions) GetValue() BulkActionEditPayloadRuleActionsValue {
	if o == nil {
		var ret BulkActionEditPayloadRuleActionsValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BulkActionEditPayloadRuleActions) GetValueOk() (*BulkActionEditPayloadRuleActionsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BulkActionEditPayloadRuleActions) SetValue(v BulkActionEditPayloadRuleActionsValue) {
	o.Value = v
}

func (o BulkActionEditPayloadRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkActionEditPayloadRuleActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableBulkActionEditPayloadRuleActions struct {
	value *BulkActionEditPayloadRuleActions
	isSet bool
}

func (v NullableBulkActionEditPayloadRuleActions) Get() *BulkActionEditPayloadRuleActions {
	return v.value
}

func (v *NullableBulkActionEditPayloadRuleActions) Set(val *BulkActionEditPayloadRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkActionEditPayloadRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkActionEditPayloadRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkActionEditPayloadRuleActions(val *BulkActionEditPayloadRuleActions) *NullableBulkActionEditPayloadRuleActions {
	return &NullableBulkActionEditPayloadRuleActions{value: val, isSet: true}
}

func (v NullableBulkActionEditPayloadRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkActionEditPayloadRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
