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

// checks if the ThresholdRuleOptionalFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdRuleOptionalFields{}

// ThresholdRuleOptionalFields struct for ThresholdRuleOptionalFields
type ThresholdRuleOptionalFields struct {
	AlertSuppression *ThresholdAlertSuppression `json:"alert_suppression,omitempty"`
	DataViewId       *string                    `json:"data_view_id,omitempty"`
	Filters          []interface{}              `json:"filters,omitempty"`
	Index            []string                   `json:"index,omitempty"`
	SavedId          *string                    `json:"saved_id,omitempty"`
}

// NewThresholdRuleOptionalFields instantiates a new ThresholdRuleOptionalFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdRuleOptionalFields() *ThresholdRuleOptionalFields {
	this := ThresholdRuleOptionalFields{}
	return &this
}

// NewThresholdRuleOptionalFieldsWithDefaults instantiates a new ThresholdRuleOptionalFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdRuleOptionalFieldsWithDefaults() *ThresholdRuleOptionalFields {
	this := ThresholdRuleOptionalFields{}
	return &this
}

// GetAlertSuppression returns the AlertSuppression field value if set, zero value otherwise.
func (o *ThresholdRuleOptionalFields) GetAlertSuppression() ThresholdAlertSuppression {
	if o == nil || IsNil(o.AlertSuppression) {
		var ret ThresholdAlertSuppression
		return ret
	}
	return *o.AlertSuppression
}

// GetAlertSuppressionOk returns a tuple with the AlertSuppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRuleOptionalFields) GetAlertSuppressionOk() (*ThresholdAlertSuppression, bool) {
	if o == nil || IsNil(o.AlertSuppression) {
		return nil, false
	}
	return o.AlertSuppression, true
}

// HasAlertSuppression returns a boolean if a field has been set.
func (o *ThresholdRuleOptionalFields) HasAlertSuppression() bool {
	if o != nil && !IsNil(o.AlertSuppression) {
		return true
	}

	return false
}

// SetAlertSuppression gets a reference to the given ThresholdAlertSuppression and assigns it to the AlertSuppression field.
func (o *ThresholdRuleOptionalFields) SetAlertSuppression(v ThresholdAlertSuppression) {
	o.AlertSuppression = &v
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *ThresholdRuleOptionalFields) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRuleOptionalFields) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *ThresholdRuleOptionalFields) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *ThresholdRuleOptionalFields) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ThresholdRuleOptionalFields) GetFilters() []interface{} {
	if o == nil || IsNil(o.Filters) {
		var ret []interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRuleOptionalFields) GetFiltersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ThresholdRuleOptionalFields) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []interface{} and assigns it to the Filters field.
func (o *ThresholdRuleOptionalFields) SetFilters(v []interface{}) {
	o.Filters = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ThresholdRuleOptionalFields) GetIndex() []string {
	if o == nil || IsNil(o.Index) {
		var ret []string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRuleOptionalFields) GetIndexOk() ([]string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ThresholdRuleOptionalFields) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []string and assigns it to the Index field.
func (o *ThresholdRuleOptionalFields) SetIndex(v []string) {
	o.Index = v
}

// GetSavedId returns the SavedId field value if set, zero value otherwise.
func (o *ThresholdRuleOptionalFields) GetSavedId() string {
	if o == nil || IsNil(o.SavedId) {
		var ret string
		return ret
	}
	return *o.SavedId
}

// GetSavedIdOk returns a tuple with the SavedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRuleOptionalFields) GetSavedIdOk() (*string, bool) {
	if o == nil || IsNil(o.SavedId) {
		return nil, false
	}
	return o.SavedId, true
}

// HasSavedId returns a boolean if a field has been set.
func (o *ThresholdRuleOptionalFields) HasSavedId() bool {
	if o != nil && !IsNil(o.SavedId) {
		return true
	}

	return false
}

// SetSavedId gets a reference to the given string and assigns it to the SavedId field.
func (o *ThresholdRuleOptionalFields) SetSavedId(v string) {
	o.SavedId = &v
}

func (o ThresholdRuleOptionalFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdRuleOptionalFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertSuppression) {
		toSerialize["alert_suppression"] = o.AlertSuppression
	}
	if !IsNil(o.DataViewId) {
		toSerialize["data_view_id"] = o.DataViewId
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.SavedId) {
		toSerialize["saved_id"] = o.SavedId
	}
	return toSerialize, nil
}

type NullableThresholdRuleOptionalFields struct {
	value *ThresholdRuleOptionalFields
	isSet bool
}

func (v NullableThresholdRuleOptionalFields) Get() *ThresholdRuleOptionalFields {
	return v.value
}

func (v *NullableThresholdRuleOptionalFields) Set(val *ThresholdRuleOptionalFields) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdRuleOptionalFields) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdRuleOptionalFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdRuleOptionalFields(val *ThresholdRuleOptionalFields) *NullableThresholdRuleOptionalFields {
	return &NullableThresholdRuleOptionalFields{value: val, isSet: true}
}

func (v NullableThresholdRuleOptionalFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdRuleOptionalFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
