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

// checks if the ErrorSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorSchema{}

// ErrorSchema struct for ErrorSchema
type ErrorSchema struct {
	Error  ErrorSchemaError `json:"error"`
	Id     *string          `json:"id,omitempty"`
	ItemId *string          `json:"item_id,omitempty"`
	ListId *string          `json:"list_id,omitempty"`
	// Could be any string, not necessarily a UUID
	RuleId *string `json:"rule_id,omitempty"`
}

// NewErrorSchema instantiates a new ErrorSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSchema(error_ ErrorSchemaError) *ErrorSchema {
	this := ErrorSchema{}
	this.Error = error_
	return &this
}

// NewErrorSchemaWithDefaults instantiates a new ErrorSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSchemaWithDefaults() *ErrorSchema {
	this := ErrorSchema{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorSchema) GetError() ErrorSchemaError {
	if o == nil {
		var ret ErrorSchemaError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetErrorOk() (*ErrorSchemaError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorSchema) SetError(v ErrorSchemaError) {
	o.Error = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ErrorSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ErrorSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ErrorSchema) SetId(v string) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ErrorSchema) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ErrorSchema) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *ErrorSchema) SetItemId(v string) {
	o.ItemId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ErrorSchema) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ErrorSchema) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ErrorSchema) SetListId(v string) {
	o.ListId = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *ErrorSchema) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *ErrorSchema) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *ErrorSchema) SetRuleId(v string) {
	o.RuleId = &v
}

func (o ErrorSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	return toSerialize, nil
}

type NullableErrorSchema struct {
	value *ErrorSchema
	isSet bool
}

func (v NullableErrorSchema) Get() *ErrorSchema {
	return v.value
}

func (v *NullableErrorSchema) Set(val *ErrorSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSchema(val *ErrorSchema) *NullableErrorSchema {
	return &NullableErrorSchema{value: val, isSet: true}
}

func (v NullableErrorSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
