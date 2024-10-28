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

// checks if the NewTermsRulePatchFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewTermsRulePatchFields{}

// NewTermsRulePatchFields struct for NewTermsRulePatchFields
type NewTermsRulePatchFields struct {
	AlertSuppression *AlertSuppression `json:"alert_suppression,omitempty"`
	DataViewId       *string           `json:"data_view_id,omitempty"`
	Filters          []interface{}     `json:"filters,omitempty"`
	Index            []string          `json:"index,omitempty"`
	ResponseActions  []ResponseAction  `json:"response_actions,omitempty"`
	Language         *KqlQueryLanguage `json:"language,omitempty"`
	// A string that is not empty and does not contain only whitespace
	HistoryWindowStart *string  `json:"history_window_start,omitempty"`
	NewTermsFields     []string `json:"new_terms_fields,omitempty"`
	Query              *string  `json:"query,omitempty"`
	// Rule type
	Type *string `json:"type,omitempty"`
}

// NewNewTermsRulePatchFields instantiates a new NewTermsRulePatchFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewTermsRulePatchFields() *NewTermsRulePatchFields {
	this := NewTermsRulePatchFields{}
	return &this
}

// NewNewTermsRulePatchFieldsWithDefaults instantiates a new NewTermsRulePatchFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewTermsRulePatchFieldsWithDefaults() *NewTermsRulePatchFields {
	this := NewTermsRulePatchFields{}
	return &this
}

// GetAlertSuppression returns the AlertSuppression field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetAlertSuppression() AlertSuppression {
	if o == nil || IsNil(o.AlertSuppression) {
		var ret AlertSuppression
		return ret
	}
	return *o.AlertSuppression
}

// GetAlertSuppressionOk returns a tuple with the AlertSuppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetAlertSuppressionOk() (*AlertSuppression, bool) {
	if o == nil || IsNil(o.AlertSuppression) {
		return nil, false
	}
	return o.AlertSuppression, true
}

// HasAlertSuppression returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasAlertSuppression() bool {
	if o != nil && !IsNil(o.AlertSuppression) {
		return true
	}

	return false
}

// SetAlertSuppression gets a reference to the given AlertSuppression and assigns it to the AlertSuppression field.
func (o *NewTermsRulePatchFields) SetAlertSuppression(v AlertSuppression) {
	o.AlertSuppression = &v
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *NewTermsRulePatchFields) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetFilters() []interface{} {
	if o == nil || IsNil(o.Filters) {
		var ret []interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetFiltersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []interface{} and assigns it to the Filters field.
func (o *NewTermsRulePatchFields) SetFilters(v []interface{}) {
	o.Filters = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetIndex() []string {
	if o == nil || IsNil(o.Index) {
		var ret []string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetIndexOk() ([]string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []string and assigns it to the Index field.
func (o *NewTermsRulePatchFields) SetIndex(v []string) {
	o.Index = v
}

// GetResponseActions returns the ResponseActions field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetResponseActions() []ResponseAction {
	if o == nil || IsNil(o.ResponseActions) {
		var ret []ResponseAction
		return ret
	}
	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetResponseActionsOk() ([]ResponseAction, bool) {
	if o == nil || IsNil(o.ResponseActions) {
		return nil, false
	}
	return o.ResponseActions, true
}

// HasResponseActions returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasResponseActions() bool {
	if o != nil && !IsNil(o.ResponseActions) {
		return true
	}

	return false
}

// SetResponseActions gets a reference to the given []ResponseAction and assigns it to the ResponseActions field.
func (o *NewTermsRulePatchFields) SetResponseActions(v []ResponseAction) {
	o.ResponseActions = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetLanguage() KqlQueryLanguage {
	if o == nil || IsNil(o.Language) {
		var ret KqlQueryLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetLanguageOk() (*KqlQueryLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given KqlQueryLanguage and assigns it to the Language field.
func (o *NewTermsRulePatchFields) SetLanguage(v KqlQueryLanguage) {
	o.Language = &v
}

// GetHistoryWindowStart returns the HistoryWindowStart field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetHistoryWindowStart() string {
	if o == nil || IsNil(o.HistoryWindowStart) {
		var ret string
		return ret
	}
	return *o.HistoryWindowStart
}

// GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetHistoryWindowStartOk() (*string, bool) {
	if o == nil || IsNil(o.HistoryWindowStart) {
		return nil, false
	}
	return o.HistoryWindowStart, true
}

// HasHistoryWindowStart returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasHistoryWindowStart() bool {
	if o != nil && !IsNil(o.HistoryWindowStart) {
		return true
	}

	return false
}

// SetHistoryWindowStart gets a reference to the given string and assigns it to the HistoryWindowStart field.
func (o *NewTermsRulePatchFields) SetHistoryWindowStart(v string) {
	o.HistoryWindowStart = &v
}

// GetNewTermsFields returns the NewTermsFields field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetNewTermsFields() []string {
	if o == nil || IsNil(o.NewTermsFields) {
		var ret []string
		return ret
	}
	return o.NewTermsFields
}

// GetNewTermsFieldsOk returns a tuple with the NewTermsFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetNewTermsFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.NewTermsFields) {
		return nil, false
	}
	return o.NewTermsFields, true
}

// HasNewTermsFields returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasNewTermsFields() bool {
	if o != nil && !IsNil(o.NewTermsFields) {
		return true
	}

	return false
}

// SetNewTermsFields gets a reference to the given []string and assigns it to the NewTermsFields field.
func (o *NewTermsRulePatchFields) SetNewTermsFields(v []string) {
	o.NewTermsFields = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *NewTermsRulePatchFields) SetQuery(v string) {
	o.Query = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewTermsRulePatchFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewTermsRulePatchFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewTermsRulePatchFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewTermsRulePatchFields) SetType(v string) {
	o.Type = &v
}

func (o NewTermsRulePatchFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewTermsRulePatchFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResponseActions) {
		toSerialize["response_actions"] = o.ResponseActions
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.HistoryWindowStart) {
		toSerialize["history_window_start"] = o.HistoryWindowStart
	}
	if !IsNil(o.NewTermsFields) {
		toSerialize["new_terms_fields"] = o.NewTermsFields
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNewTermsRulePatchFields struct {
	value *NewTermsRulePatchFields
	isSet bool
}

func (v NullableNewTermsRulePatchFields) Get() *NewTermsRulePatchFields {
	return v.value
}

func (v *NullableNewTermsRulePatchFields) Set(val *NewTermsRulePatchFields) {
	v.value = val
	v.isSet = true
}

func (v NullableNewTermsRulePatchFields) IsSet() bool {
	return v.isSet
}

func (v *NullableNewTermsRulePatchFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewTermsRulePatchFields(val *NewTermsRulePatchFields) *NullableNewTermsRulePatchFields {
	return &NullableNewTermsRulePatchFields{value: val, isSet: true}
}

func (v NullableNewTermsRulePatchFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewTermsRulePatchFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
