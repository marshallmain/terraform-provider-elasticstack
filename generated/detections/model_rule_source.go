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

// RuleSource Discriminated union that determines whether the rule is internally sourced (created within the Kibana app) or has an external source, such as the Elastic Prebuilt rules repo.
type RuleSource struct {
	ExternalRuleSource *ExternalRuleSource
	InternalRuleSource *InternalRuleSource
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RuleSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ExternalRuleSource'
	if jsonDict["type"] == "ExternalRuleSource" {
		// try to unmarshal JSON data into ExternalRuleSource
		err = json.Unmarshal(data, &dst.ExternalRuleSource)
		if err == nil {
			jsonExternalRuleSource, _ := json.Marshal(dst.ExternalRuleSource)
			if string(jsonExternalRuleSource) == "{}" { // empty struct
				dst.ExternalRuleSource = nil
			} else {
				return nil // data stored in dst.ExternalRuleSource, return on the first match
			}
		} else {
			dst.ExternalRuleSource = nil
		}
	}

	// check if the discriminator value is 'InternalRuleSource'
	if jsonDict["type"] == "InternalRuleSource" {
		// try to unmarshal JSON data into InternalRuleSource
		err = json.Unmarshal(data, &dst.InternalRuleSource)
		if err == nil {
			jsonInternalRuleSource, _ := json.Marshal(dst.InternalRuleSource)
			if string(jsonInternalRuleSource) == "{}" { // empty struct
				dst.InternalRuleSource = nil
			} else {
				return nil // data stored in dst.InternalRuleSource, return on the first match
			}
		} else {
			dst.InternalRuleSource = nil
		}
	}

	// try to unmarshal JSON data into ExternalRuleSource
	err = json.Unmarshal(data, &dst.ExternalRuleSource)
	if err == nil {
		jsonExternalRuleSource, _ := json.Marshal(dst.ExternalRuleSource)
		if string(jsonExternalRuleSource) == "{}" { // empty struct
			dst.ExternalRuleSource = nil
		} else {
			return nil // data stored in dst.ExternalRuleSource, return on the first match
		}
	} else {
		dst.ExternalRuleSource = nil
	}

	// try to unmarshal JSON data into InternalRuleSource
	err = json.Unmarshal(data, &dst.InternalRuleSource)
	if err == nil {
		jsonInternalRuleSource, _ := json.Marshal(dst.InternalRuleSource)
		if string(jsonInternalRuleSource) == "{}" { // empty struct
			dst.InternalRuleSource = nil
		} else {
			return nil // data stored in dst.InternalRuleSource, return on the first match
		}
	} else {
		dst.InternalRuleSource = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RuleSource)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RuleSource) MarshalJSON() ([]byte, error) {
	if src.ExternalRuleSource != nil {
		return json.Marshal(&src.ExternalRuleSource)
	}

	if src.InternalRuleSource != nil {
		return json.Marshal(&src.InternalRuleSource)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRuleSource struct {
	value *RuleSource
	isSet bool
}

func (v NullableRuleSource) Get() *RuleSource {
	return v.value
}

func (v *NullableRuleSource) Set(val *RuleSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSource(val *RuleSource) *NullableRuleSource {
	return &NullableRuleSource{value: val, isSet: true}
}

func (v NullableRuleSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
