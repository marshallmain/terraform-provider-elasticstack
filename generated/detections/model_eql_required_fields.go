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

// checks if the EqlRequiredFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EqlRequiredFields{}

// EqlRequiredFields struct for EqlRequiredFields
type EqlRequiredFields struct {
	Language EqlQueryLanguage `json:"language"`
	Query    string           `json:"query"`
	// Rule type
	Type string `json:"type"`
}

// NewEqlRequiredFields instantiates a new EqlRequiredFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEqlRequiredFields(language EqlQueryLanguage, query string, type_ string) *EqlRequiredFields {
	this := EqlRequiredFields{}
	this.Language = language
	this.Query = query
	this.Type = type_
	return &this
}

// NewEqlRequiredFieldsWithDefaults instantiates a new EqlRequiredFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEqlRequiredFieldsWithDefaults() *EqlRequiredFields {
	this := EqlRequiredFields{}
	return &this
}

// GetLanguage returns the Language field value
func (o *EqlRequiredFields) GetLanguage() EqlQueryLanguage {
	if o == nil {
		var ret EqlQueryLanguage
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *EqlRequiredFields) GetLanguageOk() (*EqlQueryLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *EqlRequiredFields) SetLanguage(v EqlQueryLanguage) {
	o.Language = v
}

// GetQuery returns the Query field value
func (o *EqlRequiredFields) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EqlRequiredFields) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *EqlRequiredFields) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *EqlRequiredFields) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EqlRequiredFields) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EqlRequiredFields) SetType(v string) {
	o.Type = v
}

func (o EqlRequiredFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EqlRequiredFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["query"] = o.Query
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableEqlRequiredFields struct {
	value *EqlRequiredFields
	isSet bool
}

func (v NullableEqlRequiredFields) Get() *EqlRequiredFields {
	return v.value
}

func (v *NullableEqlRequiredFields) Set(val *EqlRequiredFields) {
	v.value = val
	v.isSet = true
}

func (v NullableEqlRequiredFields) IsSet() bool {
	return v.isSet
}

func (v *NullableEqlRequiredFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEqlRequiredFields(val *EqlRequiredFields) *NullableEqlRequiredFields {
	return &NullableEqlRequiredFields{value: val, isSet: true}
}

func (v NullableEqlRequiredFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEqlRequiredFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
