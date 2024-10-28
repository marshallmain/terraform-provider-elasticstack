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

// checks if the RulePreviewLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulePreviewLogs{}

// RulePreviewLogs struct for RulePreviewLogs
type RulePreviewLogs struct {
	// Execution duration in milliseconds
	Duration int32                      `json:"duration"`
	Errors   []string                   `json:"errors"`
	Requests []RulePreviewLoggedRequest `json:"requests,omitempty"`
	// A string that is not empty and does not contain only whitespace
	StartedAt *string  `json:"startedAt,omitempty"`
	Warnings  []string `json:"warnings"`
}

// NewRulePreviewLogs instantiates a new RulePreviewLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulePreviewLogs(duration int32, errors []string, warnings []string) *RulePreviewLogs {
	this := RulePreviewLogs{}
	this.Duration = duration
	this.Errors = errors
	this.Warnings = warnings
	return &this
}

// NewRulePreviewLogsWithDefaults instantiates a new RulePreviewLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulePreviewLogsWithDefaults() *RulePreviewLogs {
	this := RulePreviewLogs{}
	return &this
}

// GetDuration returns the Duration field value
func (o *RulePreviewLogs) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *RulePreviewLogs) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *RulePreviewLogs) SetDuration(v int32) {
	o.Duration = v
}

// GetErrors returns the Errors field value
func (o *RulePreviewLogs) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *RulePreviewLogs) GetErrorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *RulePreviewLogs) SetErrors(v []string) {
	o.Errors = v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *RulePreviewLogs) GetRequests() []RulePreviewLoggedRequest {
	if o == nil || IsNil(o.Requests) {
		var ret []RulePreviewLoggedRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulePreviewLogs) GetRequestsOk() ([]RulePreviewLoggedRequest, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *RulePreviewLogs) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []RulePreviewLoggedRequest and assigns it to the Requests field.
func (o *RulePreviewLogs) SetRequests(v []RulePreviewLoggedRequest) {
	o.Requests = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *RulePreviewLogs) GetStartedAt() string {
	if o == nil || IsNil(o.StartedAt) {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulePreviewLogs) GetStartedAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *RulePreviewLogs) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *RulePreviewLogs) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetWarnings returns the Warnings field value
func (o *RulePreviewLogs) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *RulePreviewLogs) GetWarningsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *RulePreviewLogs) SetWarnings(v []string) {
	o.Warnings = v
}

func (o RulePreviewLogs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulePreviewLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["errors"] = o.Errors
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	toSerialize["warnings"] = o.Warnings
	return toSerialize, nil
}

type NullableRulePreviewLogs struct {
	value *RulePreviewLogs
	isSet bool
}

func (v NullableRulePreviewLogs) Get() *RulePreviewLogs {
	return v.value
}

func (v *NullableRulePreviewLogs) Set(val *RulePreviewLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableRulePreviewLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableRulePreviewLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulePreviewLogs(val *RulePreviewLogs) *NullableRulePreviewLogs {
	return &NullableRulePreviewLogs{value: val, isSet: true}
}

func (v NullableRulePreviewLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulePreviewLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
