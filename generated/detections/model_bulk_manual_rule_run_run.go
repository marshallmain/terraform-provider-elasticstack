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

// checks if the BulkManualRuleRunRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkManualRuleRunRun{}

// BulkManualRuleRunRun struct for BulkManualRuleRunRun
type BulkManualRuleRunRun struct {
	// End date of the manual rule run
	EndDate *string `json:"end_date,omitempty"`
	// Start date of the manual rule run
	StartDate string `json:"start_date"`
}

// NewBulkManualRuleRunRun instantiates a new BulkManualRuleRunRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkManualRuleRunRun(startDate string) *BulkManualRuleRunRun {
	this := BulkManualRuleRunRun{}
	this.StartDate = startDate
	return &this
}

// NewBulkManualRuleRunRunWithDefaults instantiates a new BulkManualRuleRunRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkManualRuleRunRunWithDefaults() *BulkManualRuleRunRun {
	this := BulkManualRuleRunRun{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BulkManualRuleRunRun) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRunRun) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BulkManualRuleRunRun) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *BulkManualRuleRunRun) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value
func (o *BulkManualRuleRunRun) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRunRun) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *BulkManualRuleRunRun) SetStartDate(v string) {
	o.StartDate = v
}

func (o BulkManualRuleRunRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkManualRuleRunRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	toSerialize["start_date"] = o.StartDate
	return toSerialize, nil
}

type NullableBulkManualRuleRunRun struct {
	value *BulkManualRuleRunRun
	isSet bool
}

func (v NullableBulkManualRuleRunRun) Get() *BulkManualRuleRunRun {
	return v.value
}

func (v *NullableBulkManualRuleRunRun) Set(val *BulkManualRuleRunRun) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkManualRuleRunRun) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkManualRuleRunRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkManualRuleRunRun(val *BulkManualRuleRunRun) *NullableBulkManualRuleRunRun {
	return &NullableBulkManualRuleRunRun{value: val, isSet: true}
}

func (v NullableBulkManualRuleRunRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkManualRuleRunRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
