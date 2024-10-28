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

// checks if the ProcessesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessesParams{}

// ProcessesParams struct for ProcessesParams
type ProcessesParams struct {
	Command string                `json:"command"`
	Comment *string               `json:"comment,omitempty"`
	Config  ProcessesParamsConfig `json:"config"`
}

// NewProcessesParams instantiates a new ProcessesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessesParams(command string, config ProcessesParamsConfig) *ProcessesParams {
	this := ProcessesParams{}
	this.Command = command
	this.Config = config
	return &this
}

// NewProcessesParamsWithDefaults instantiates a new ProcessesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessesParamsWithDefaults() *ProcessesParams {
	this := ProcessesParams{}
	return &this
}

// GetCommand returns the Command field value
func (o *ProcessesParams) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ProcessesParams) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *ProcessesParams) SetCommand(v string) {
	o.Command = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ProcessesParams) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesParams) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ProcessesParams) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ProcessesParams) SetComment(v string) {
	o.Comment = &v
}

// GetConfig returns the Config field value
func (o *ProcessesParams) GetConfig() ProcessesParamsConfig {
	if o == nil {
		var ret ProcessesParamsConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ProcessesParams) GetConfigOk() (*ProcessesParamsConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ProcessesParams) SetConfig(v ProcessesParamsConfig) {
	o.Config = v
}

func (o ProcessesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["command"] = o.Command
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableProcessesParams struct {
	value *ProcessesParams
	isSet bool
}

func (v NullableProcessesParams) Get() *ProcessesParams {
	return v.value
}

func (v *NullableProcessesParams) Set(val *ProcessesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessesParams(val *ProcessesParams) *NullableProcessesParams {
	return &NullableProcessesParams{value: val, isSet: true}
}

func (v NullableProcessesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
