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

// FindRulesSortField the model 'FindRulesSortField'
type FindRulesSortField string

// List of FindRulesSortField
const (
	CREATED_AT                                                          FindRulesSortField = "created_at"
	CREATED_AT2                                                         FindRulesSortField = "createdAt"
	ENABLED                                                             FindRulesSortField = "enabled"
	EXECUTION_SUMMARY_LAST_EXECUTION_DATE                               FindRulesSortField = "execution_summary.last_execution.date"
	EXECUTION_SUMMARY_LAST_EXECUTION_METRICS_EXECUTION_GAP_DURATION_S   FindRulesSortField = "execution_summary.last_execution.metrics.execution_gap_duration_s"
	EXECUTION_SUMMARY_LAST_EXECUTION_METRICS_TOTAL_INDEXING_DURATION_MS FindRulesSortField = "execution_summary.last_execution.metrics.total_indexing_duration_ms"
	EXECUTION_SUMMARY_LAST_EXECUTION_METRICS_TOTAL_SEARCH_DURATION_MS   FindRulesSortField = "execution_summary.last_execution.metrics.total_search_duration_ms"
	EXECUTION_SUMMARY_LAST_EXECUTION_STATUS                             FindRulesSortField = "execution_summary.last_execution.status"
	NAME                                                                FindRulesSortField = "name"
	RISK_SCORE                                                          FindRulesSortField = "risk_score"
	RISK_SCORE2                                                         FindRulesSortField = "riskScore"
	SEVERITY                                                            FindRulesSortField = "severity"
	UPDATED_AT                                                          FindRulesSortField = "updated_at"
	UPDATED_AT2                                                         FindRulesSortField = "updatedAt"
)

// All allowed values of FindRulesSortField enum
var AllowedFindRulesSortFieldEnumValues = []FindRulesSortField{
	"created_at",
	"createdAt",
	"enabled",
	"execution_summary.last_execution.date",
	"execution_summary.last_execution.metrics.execution_gap_duration_s",
	"execution_summary.last_execution.metrics.total_indexing_duration_ms",
	"execution_summary.last_execution.metrics.total_search_duration_ms",
	"execution_summary.last_execution.status",
	"name",
	"risk_score",
	"riskScore",
	"severity",
	"updated_at",
	"updatedAt",
}

func (v *FindRulesSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FindRulesSortField(value)
	for _, existing := range AllowedFindRulesSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FindRulesSortField", value)
}

// NewFindRulesSortFieldFromValue returns a pointer to a valid FindRulesSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFindRulesSortFieldFromValue(v string) (*FindRulesSortField, error) {
	ev := FindRulesSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FindRulesSortField: valid values are %v", v, AllowedFindRulesSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FindRulesSortField) IsValid() bool {
	for _, existing := range AllowedFindRulesSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindRulesSortField value
func (v FindRulesSortField) Ptr() *FindRulesSortField {
	return &v
}

type NullableFindRulesSortField struct {
	value *FindRulesSortField
	isSet bool
}

func (v NullableFindRulesSortField) Get() *FindRulesSortField {
	return v.value
}

func (v *NullableFindRulesSortField) Set(val *FindRulesSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableFindRulesSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableFindRulesSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindRulesSortField(val *FindRulesSortField) *NullableFindRulesSortField {
	return &NullableFindRulesSortField{value: val, isSet: true}
}

func (v NullableFindRulesSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindRulesSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
