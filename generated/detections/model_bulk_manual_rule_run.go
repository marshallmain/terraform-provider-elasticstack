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

// checks if the BulkManualRuleRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkManualRuleRun{}

// BulkManualRuleRun struct for BulkManualRuleRun
type BulkManualRuleRun struct {
	Action string `json:"action"`
	// Array of rule IDs
	Ids []string `json:"ids,omitempty"`
	// Query to filter rules
	Query *string              `json:"query,omitempty"`
	Run   BulkManualRuleRunRun `json:"run"`
}

// NewBulkManualRuleRun instantiates a new BulkManualRuleRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkManualRuleRun(action string, run BulkManualRuleRunRun) *BulkManualRuleRun {
	this := BulkManualRuleRun{}
	this.Action = action
	this.Run = run
	return &this
}

// NewBulkManualRuleRunWithDefaults instantiates a new BulkManualRuleRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkManualRuleRunWithDefaults() *BulkManualRuleRun {
	this := BulkManualRuleRun{}
	return &this
}

// GetAction returns the Action field value
func (o *BulkManualRuleRun) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRun) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BulkManualRuleRun) SetAction(v string) {
	o.Action = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BulkManualRuleRun) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRun) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BulkManualRuleRun) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BulkManualRuleRun) SetIds(v []string) {
	o.Ids = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BulkManualRuleRun) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRun) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BulkManualRuleRun) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *BulkManualRuleRun) SetQuery(v string) {
	o.Query = &v
}

// GetRun returns the Run field value
func (o *BulkManualRuleRun) GetRun() BulkManualRuleRunRun {
	if o == nil {
		var ret BulkManualRuleRunRun
		return ret
	}

	return o.Run
}

// GetRunOk returns a tuple with the Run field value
// and a boolean to check if the value has been set.
func (o *BulkManualRuleRun) GetRunOk() (*BulkManualRuleRunRun, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Run, true
}

// SetRun sets field value
func (o *BulkManualRuleRun) SetRun(v BulkManualRuleRunRun) {
	o.Run = v
}

func (o BulkManualRuleRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkManualRuleRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	toSerialize["run"] = o.Run
	return toSerialize, nil
}

type NullableBulkManualRuleRun struct {
	value *BulkManualRuleRun
	isSet bool
}

func (v NullableBulkManualRuleRun) Get() *BulkManualRuleRun {
	return v.value
}

func (v *NullableBulkManualRuleRun) Set(val *BulkManualRuleRun) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkManualRuleRun) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkManualRuleRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkManualRuleRun(val *BulkManualRuleRun) *NullableBulkManualRuleRun {
	return &NullableBulkManualRuleRun{value: val, isSet: true}
}

func (v NullableBulkManualRuleRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkManualRuleRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
