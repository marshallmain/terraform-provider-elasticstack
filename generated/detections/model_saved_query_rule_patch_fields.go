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

// checks if the SavedQueryRulePatchFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedQueryRulePatchFields{}

// SavedQueryRulePatchFields struct for SavedQueryRulePatchFields
type SavedQueryRulePatchFields struct {
	AlertSuppression *AlertSuppression `json:"alert_suppression,omitempty"`
	DataViewId       *string           `json:"data_view_id,omitempty"`
	Filters          []interface{}     `json:"filters,omitempty"`
	Index            []string          `json:"index,omitempty"`
	Query            *string           `json:"query,omitempty"`
	ResponseActions  []ResponseAction  `json:"response_actions,omitempty"`
	Language         *KqlQueryLanguage `json:"language,omitempty"`
	SavedId          *string           `json:"saved_id,omitempty"`
	// Rule type
	Type *string `json:"type,omitempty"`
}

// NewSavedQueryRulePatchFields instantiates a new SavedQueryRulePatchFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedQueryRulePatchFields() *SavedQueryRulePatchFields {
	this := SavedQueryRulePatchFields{}
	return &this
}

// NewSavedQueryRulePatchFieldsWithDefaults instantiates a new SavedQueryRulePatchFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedQueryRulePatchFieldsWithDefaults() *SavedQueryRulePatchFields {
	this := SavedQueryRulePatchFields{}
	return &this
}

// GetAlertSuppression returns the AlertSuppression field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetAlertSuppression() AlertSuppression {
	if o == nil || IsNil(o.AlertSuppression) {
		var ret AlertSuppression
		return ret
	}
	return *o.AlertSuppression
}

// GetAlertSuppressionOk returns a tuple with the AlertSuppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetAlertSuppressionOk() (*AlertSuppression, bool) {
	if o == nil || IsNil(o.AlertSuppression) {
		return nil, false
	}
	return o.AlertSuppression, true
}

// HasAlertSuppression returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasAlertSuppression() bool {
	if o != nil && !IsNil(o.AlertSuppression) {
		return true
	}

	return false
}

// SetAlertSuppression gets a reference to the given AlertSuppression and assigns it to the AlertSuppression field.
func (o *SavedQueryRulePatchFields) SetAlertSuppression(v AlertSuppression) {
	o.AlertSuppression = &v
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *SavedQueryRulePatchFields) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetFilters() []interface{} {
	if o == nil || IsNil(o.Filters) {
		var ret []interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetFiltersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []interface{} and assigns it to the Filters field.
func (o *SavedQueryRulePatchFields) SetFilters(v []interface{}) {
	o.Filters = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetIndex() []string {
	if o == nil || IsNil(o.Index) {
		var ret []string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetIndexOk() ([]string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []string and assigns it to the Index field.
func (o *SavedQueryRulePatchFields) SetIndex(v []string) {
	o.Index = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SavedQueryRulePatchFields) SetQuery(v string) {
	o.Query = &v
}

// GetResponseActions returns the ResponseActions field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetResponseActions() []ResponseAction {
	if o == nil || IsNil(o.ResponseActions) {
		var ret []ResponseAction
		return ret
	}
	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetResponseActionsOk() ([]ResponseAction, bool) {
	if o == nil || IsNil(o.ResponseActions) {
		return nil, false
	}
	return o.ResponseActions, true
}

// HasResponseActions returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasResponseActions() bool {
	if o != nil && !IsNil(o.ResponseActions) {
		return true
	}

	return false
}

// SetResponseActions gets a reference to the given []ResponseAction and assigns it to the ResponseActions field.
func (o *SavedQueryRulePatchFields) SetResponseActions(v []ResponseAction) {
	o.ResponseActions = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetLanguage() KqlQueryLanguage {
	if o == nil || IsNil(o.Language) {
		var ret KqlQueryLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetLanguageOk() (*KqlQueryLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given KqlQueryLanguage and assigns it to the Language field.
func (o *SavedQueryRulePatchFields) SetLanguage(v KqlQueryLanguage) {
	o.Language = &v
}

// GetSavedId returns the SavedId field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetSavedId() string {
	if o == nil || IsNil(o.SavedId) {
		var ret string
		return ret
	}
	return *o.SavedId
}

// GetSavedIdOk returns a tuple with the SavedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetSavedIdOk() (*string, bool) {
	if o == nil || IsNil(o.SavedId) {
		return nil, false
	}
	return o.SavedId, true
}

// HasSavedId returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasSavedId() bool {
	if o != nil && !IsNil(o.SavedId) {
		return true
	}

	return false
}

// SetSavedId gets a reference to the given string and assigns it to the SavedId field.
func (o *SavedQueryRulePatchFields) SetSavedId(v string) {
	o.SavedId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SavedQueryRulePatchFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQueryRulePatchFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SavedQueryRulePatchFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SavedQueryRulePatchFields) SetType(v string) {
	o.Type = &v
}

func (o SavedQueryRulePatchFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedQueryRulePatchFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.ResponseActions) {
		toSerialize["response_actions"] = o.ResponseActions
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.SavedId) {
		toSerialize["saved_id"] = o.SavedId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSavedQueryRulePatchFields struct {
	value *SavedQueryRulePatchFields
	isSet bool
}

func (v NullableSavedQueryRulePatchFields) Get() *SavedQueryRulePatchFields {
	return v.value
}

func (v *NullableSavedQueryRulePatchFields) Set(val *SavedQueryRulePatchFields) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedQueryRulePatchFields) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedQueryRulePatchFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedQueryRulePatchFields(val *SavedQueryRulePatchFields) *NullableSavedQueryRulePatchFields {
	return &NullableSavedQueryRulePatchFields{value: val, isSet: true}
}

func (v NullableSavedQueryRulePatchFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedQueryRulePatchFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
