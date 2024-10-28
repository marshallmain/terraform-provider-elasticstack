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

// checks if the MachineLearningRuleOptionalFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineLearningRuleOptionalFields{}

// MachineLearningRuleOptionalFields struct for MachineLearningRuleOptionalFields
type MachineLearningRuleOptionalFields struct {
	AlertSuppression *AlertSuppression `json:"alert_suppression,omitempty"`
}

// NewMachineLearningRuleOptionalFields instantiates a new MachineLearningRuleOptionalFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineLearningRuleOptionalFields() *MachineLearningRuleOptionalFields {
	this := MachineLearningRuleOptionalFields{}
	return &this
}

// NewMachineLearningRuleOptionalFieldsWithDefaults instantiates a new MachineLearningRuleOptionalFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineLearningRuleOptionalFieldsWithDefaults() *MachineLearningRuleOptionalFields {
	this := MachineLearningRuleOptionalFields{}
	return &this
}

// GetAlertSuppression returns the AlertSuppression field value if set, zero value otherwise.
func (o *MachineLearningRuleOptionalFields) GetAlertSuppression() AlertSuppression {
	if o == nil || IsNil(o.AlertSuppression) {
		var ret AlertSuppression
		return ret
	}
	return *o.AlertSuppression
}

// GetAlertSuppressionOk returns a tuple with the AlertSuppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineLearningRuleOptionalFields) GetAlertSuppressionOk() (*AlertSuppression, bool) {
	if o == nil || IsNil(o.AlertSuppression) {
		return nil, false
	}
	return o.AlertSuppression, true
}

// HasAlertSuppression returns a boolean if a field has been set.
func (o *MachineLearningRuleOptionalFields) HasAlertSuppression() bool {
	if o != nil && !IsNil(o.AlertSuppression) {
		return true
	}

	return false
}

// SetAlertSuppression gets a reference to the given AlertSuppression and assigns it to the AlertSuppression field.
func (o *MachineLearningRuleOptionalFields) SetAlertSuppression(v AlertSuppression) {
	o.AlertSuppression = &v
}

func (o MachineLearningRuleOptionalFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineLearningRuleOptionalFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertSuppression) {
		toSerialize["alert_suppression"] = o.AlertSuppression
	}
	return toSerialize, nil
}

type NullableMachineLearningRuleOptionalFields struct {
	value *MachineLearningRuleOptionalFields
	isSet bool
}

func (v NullableMachineLearningRuleOptionalFields) Get() *MachineLearningRuleOptionalFields {
	return v.value
}

func (v *NullableMachineLearningRuleOptionalFields) Set(val *MachineLearningRuleOptionalFields) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineLearningRuleOptionalFields) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineLearningRuleOptionalFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineLearningRuleOptionalFields(val *MachineLearningRuleOptionalFields) *NullableMachineLearningRuleOptionalFields {
	return &NullableMachineLearningRuleOptionalFields{value: val, isSet: true}
}

func (v NullableMachineLearningRuleOptionalFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineLearningRuleOptionalFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
