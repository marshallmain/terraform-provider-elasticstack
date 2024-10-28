/*
Security Detections (Elastic Cloud and self-hosted)

You can create rules that automatically turn events and external alerts sent to Elastic Security into detection alerts. These alerts are displayed on the Detections page.

API version: 2023-10-31
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package detections

import (
	"encoding/json"
	"fmt"
)

// KqlQueryLanguage the model 'KqlQueryLanguage'
type KqlQueryLanguage string

// List of KqlQueryLanguage
const (
	KUERY  KqlQueryLanguage = "kuery"
	LUCENE KqlQueryLanguage = "lucene"
)

// All allowed values of KqlQueryLanguage enum
var AllowedKqlQueryLanguageEnumValues = []KqlQueryLanguage{
	"kuery",
	"lucene",
}

func (v *KqlQueryLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KqlQueryLanguage(value)
	for _, existing := range AllowedKqlQueryLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KqlQueryLanguage", value)
}

// NewKqlQueryLanguageFromValue returns a pointer to a valid KqlQueryLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKqlQueryLanguageFromValue(v string) (*KqlQueryLanguage, error) {
	ev := KqlQueryLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KqlQueryLanguage: valid values are %v", v, AllowedKqlQueryLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KqlQueryLanguage) IsValid() bool {
	for _, existing := range AllowedKqlQueryLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KqlQueryLanguage value
func (v KqlQueryLanguage) Ptr() *KqlQueryLanguage {
	return &v
}

type NullableKqlQueryLanguage struct {
	value *KqlQueryLanguage
	isSet bool
}

func (v NullableKqlQueryLanguage) Get() *KqlQueryLanguage {
	return v.value
}

func (v *NullableKqlQueryLanguage) Set(val *KqlQueryLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableKqlQueryLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableKqlQueryLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKqlQueryLanguage(val *KqlQueryLanguage) *NullableKqlQueryLanguage {
	return &NullableKqlQueryLanguage{value: val, isSet: true}
}

func (v NullableKqlQueryLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKqlQueryLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
