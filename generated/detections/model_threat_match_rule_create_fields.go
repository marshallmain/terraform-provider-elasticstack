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

// checks if the ThreatMatchRuleCreateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatMatchRuleCreateFields{}

// ThreatMatchRuleCreateFields struct for ThreatMatchRuleCreateFields
type ThreatMatchRuleCreateFields struct {
	Query         string               `json:"query"`
	ThreatIndex   []string             `json:"threat_index"`
	ThreatMapping []ThreatMappingInner `json:"threat_mapping"`
	// Query to run
	ThreatQuery string `json:"threat_query"`
	// Rule type
	Type               string            `json:"type"`
	AlertSuppression   *AlertSuppression `json:"alert_suppression,omitempty"`
	ConcurrentSearches *int32            `json:"concurrent_searches,omitempty"`
	DataViewId         *string           `json:"data_view_id,omitempty"`
	Filters            []interface{}     `json:"filters,omitempty"`
	Index              []string          `json:"index,omitempty"`
	ItemsPerSearch     *int32            `json:"items_per_search,omitempty"`
	SavedId            *string           `json:"saved_id,omitempty"`
	ThreatFilters      []interface{}     `json:"threat_filters,omitempty"`
	// Defines the path to the threat indicator in the indicator documents (optional)
	ThreatIndicatorPath *string           `json:"threat_indicator_path,omitempty"`
	ThreatLanguage      *KqlQueryLanguage `json:"threat_language,omitempty"`
	Language            *KqlQueryLanguage `json:"language,omitempty"`
}

// NewThreatMatchRuleCreateFields instantiates a new ThreatMatchRuleCreateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatMatchRuleCreateFields(query string, threatIndex []string, threatMapping []ThreatMappingInner, threatQuery string, type_ string) *ThreatMatchRuleCreateFields {
	this := ThreatMatchRuleCreateFields{}
	this.Query = query
	this.ThreatIndex = threatIndex
	this.ThreatMapping = threatMapping
	this.ThreatQuery = threatQuery
	this.Type = type_
	return &this
}

// NewThreatMatchRuleCreateFieldsWithDefaults instantiates a new ThreatMatchRuleCreateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatMatchRuleCreateFieldsWithDefaults() *ThreatMatchRuleCreateFields {
	this := ThreatMatchRuleCreateFields{}
	return &this
}

// GetQuery returns the Query field value
func (o *ThreatMatchRuleCreateFields) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ThreatMatchRuleCreateFields) SetQuery(v string) {
	o.Query = v
}

// GetThreatIndex returns the ThreatIndex field value
func (o *ThreatMatchRuleCreateFields) GetThreatIndex() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ThreatIndex
}

// GetThreatIndexOk returns a tuple with the ThreatIndex field value
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatIndexOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreatIndex, true
}

// SetThreatIndex sets field value
func (o *ThreatMatchRuleCreateFields) SetThreatIndex(v []string) {
	o.ThreatIndex = v
}

// GetThreatMapping returns the ThreatMapping field value
func (o *ThreatMatchRuleCreateFields) GetThreatMapping() []ThreatMappingInner {
	if o == nil {
		var ret []ThreatMappingInner
		return ret
	}

	return o.ThreatMapping
}

// GetThreatMappingOk returns a tuple with the ThreatMapping field value
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatMappingOk() ([]ThreatMappingInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreatMapping, true
}

// SetThreatMapping sets field value
func (o *ThreatMatchRuleCreateFields) SetThreatMapping(v []ThreatMappingInner) {
	o.ThreatMapping = v
}

// GetThreatQuery returns the ThreatQuery field value
func (o *ThreatMatchRuleCreateFields) GetThreatQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreatQuery
}

// GetThreatQueryOk returns a tuple with the ThreatQuery field value
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreatQuery, true
}

// SetThreatQuery sets field value
func (o *ThreatMatchRuleCreateFields) SetThreatQuery(v string) {
	o.ThreatQuery = v
}

// GetType returns the Type field value
func (o *ThreatMatchRuleCreateFields) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThreatMatchRuleCreateFields) SetType(v string) {
	o.Type = v
}

// GetAlertSuppression returns the AlertSuppression field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetAlertSuppression() AlertSuppression {
	if o == nil || IsNil(o.AlertSuppression) {
		var ret AlertSuppression
		return ret
	}
	return *o.AlertSuppression
}

// GetAlertSuppressionOk returns a tuple with the AlertSuppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetAlertSuppressionOk() (*AlertSuppression, bool) {
	if o == nil || IsNil(o.AlertSuppression) {
		return nil, false
	}
	return o.AlertSuppression, true
}

// HasAlertSuppression returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasAlertSuppression() bool {
	if o != nil && !IsNil(o.AlertSuppression) {
		return true
	}

	return false
}

// SetAlertSuppression gets a reference to the given AlertSuppression and assigns it to the AlertSuppression field.
func (o *ThreatMatchRuleCreateFields) SetAlertSuppression(v AlertSuppression) {
	o.AlertSuppression = &v
}

// GetConcurrentSearches returns the ConcurrentSearches field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetConcurrentSearches() int32 {
	if o == nil || IsNil(o.ConcurrentSearches) {
		var ret int32
		return ret
	}
	return *o.ConcurrentSearches
}

// GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetConcurrentSearchesOk() (*int32, bool) {
	if o == nil || IsNil(o.ConcurrentSearches) {
		return nil, false
	}
	return o.ConcurrentSearches, true
}

// HasConcurrentSearches returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasConcurrentSearches() bool {
	if o != nil && !IsNil(o.ConcurrentSearches) {
		return true
	}

	return false
}

// SetConcurrentSearches gets a reference to the given int32 and assigns it to the ConcurrentSearches field.
func (o *ThreatMatchRuleCreateFields) SetConcurrentSearches(v int32) {
	o.ConcurrentSearches = &v
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *ThreatMatchRuleCreateFields) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetFilters() []interface{} {
	if o == nil || IsNil(o.Filters) {
		var ret []interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetFiltersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []interface{} and assigns it to the Filters field.
func (o *ThreatMatchRuleCreateFields) SetFilters(v []interface{}) {
	o.Filters = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetIndex() []string {
	if o == nil || IsNil(o.Index) {
		var ret []string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetIndexOk() ([]string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []string and assigns it to the Index field.
func (o *ThreatMatchRuleCreateFields) SetIndex(v []string) {
	o.Index = v
}

// GetItemsPerSearch returns the ItemsPerSearch field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetItemsPerSearch() int32 {
	if o == nil || IsNil(o.ItemsPerSearch) {
		var ret int32
		return ret
	}
	return *o.ItemsPerSearch
}

// GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetItemsPerSearchOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerSearch) {
		return nil, false
	}
	return o.ItemsPerSearch, true
}

// HasItemsPerSearch returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasItemsPerSearch() bool {
	if o != nil && !IsNil(o.ItemsPerSearch) {
		return true
	}

	return false
}

// SetItemsPerSearch gets a reference to the given int32 and assigns it to the ItemsPerSearch field.
func (o *ThreatMatchRuleCreateFields) SetItemsPerSearch(v int32) {
	o.ItemsPerSearch = &v
}

// GetSavedId returns the SavedId field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetSavedId() string {
	if o == nil || IsNil(o.SavedId) {
		var ret string
		return ret
	}
	return *o.SavedId
}

// GetSavedIdOk returns a tuple with the SavedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetSavedIdOk() (*string, bool) {
	if o == nil || IsNil(o.SavedId) {
		return nil, false
	}
	return o.SavedId, true
}

// HasSavedId returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasSavedId() bool {
	if o != nil && !IsNil(o.SavedId) {
		return true
	}

	return false
}

// SetSavedId gets a reference to the given string and assigns it to the SavedId field.
func (o *ThreatMatchRuleCreateFields) SetSavedId(v string) {
	o.SavedId = &v
}

// GetThreatFilters returns the ThreatFilters field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetThreatFilters() []interface{} {
	if o == nil || IsNil(o.ThreatFilters) {
		var ret []interface{}
		return ret
	}
	return o.ThreatFilters
}

// GetThreatFiltersOk returns a tuple with the ThreatFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatFiltersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ThreatFilters) {
		return nil, false
	}
	return o.ThreatFilters, true
}

// HasThreatFilters returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasThreatFilters() bool {
	if o != nil && !IsNil(o.ThreatFilters) {
		return true
	}

	return false
}

// SetThreatFilters gets a reference to the given []interface{} and assigns it to the ThreatFilters field.
func (o *ThreatMatchRuleCreateFields) SetThreatFilters(v []interface{}) {
	o.ThreatFilters = v
}

// GetThreatIndicatorPath returns the ThreatIndicatorPath field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetThreatIndicatorPath() string {
	if o == nil || IsNil(o.ThreatIndicatorPath) {
		var ret string
		return ret
	}
	return *o.ThreatIndicatorPath
}

// GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatIndicatorPathOk() (*string, bool) {
	if o == nil || IsNil(o.ThreatIndicatorPath) {
		return nil, false
	}
	return o.ThreatIndicatorPath, true
}

// HasThreatIndicatorPath returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasThreatIndicatorPath() bool {
	if o != nil && !IsNil(o.ThreatIndicatorPath) {
		return true
	}

	return false
}

// SetThreatIndicatorPath gets a reference to the given string and assigns it to the ThreatIndicatorPath field.
func (o *ThreatMatchRuleCreateFields) SetThreatIndicatorPath(v string) {
	o.ThreatIndicatorPath = &v
}

// GetThreatLanguage returns the ThreatLanguage field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetThreatLanguage() KqlQueryLanguage {
	if o == nil || IsNil(o.ThreatLanguage) {
		var ret KqlQueryLanguage
		return ret
	}
	return *o.ThreatLanguage
}

// GetThreatLanguageOk returns a tuple with the ThreatLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetThreatLanguageOk() (*KqlQueryLanguage, bool) {
	if o == nil || IsNil(o.ThreatLanguage) {
		return nil, false
	}
	return o.ThreatLanguage, true
}

// HasThreatLanguage returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasThreatLanguage() bool {
	if o != nil && !IsNil(o.ThreatLanguage) {
		return true
	}

	return false
}

// SetThreatLanguage gets a reference to the given KqlQueryLanguage and assigns it to the ThreatLanguage field.
func (o *ThreatMatchRuleCreateFields) SetThreatLanguage(v KqlQueryLanguage) {
	o.ThreatLanguage = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ThreatMatchRuleCreateFields) GetLanguage() KqlQueryLanguage {
	if o == nil || IsNil(o.Language) {
		var ret KqlQueryLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatMatchRuleCreateFields) GetLanguageOk() (*KqlQueryLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ThreatMatchRuleCreateFields) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given KqlQueryLanguage and assigns it to the Language field.
func (o *ThreatMatchRuleCreateFields) SetLanguage(v KqlQueryLanguage) {
	o.Language = &v
}

func (o ThreatMatchRuleCreateFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatMatchRuleCreateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["threat_index"] = o.ThreatIndex
	toSerialize["threat_mapping"] = o.ThreatMapping
	toSerialize["threat_query"] = o.ThreatQuery
	toSerialize["type"] = o.Type
	if !IsNil(o.AlertSuppression) {
		toSerialize["alert_suppression"] = o.AlertSuppression
	}
	if !IsNil(o.ConcurrentSearches) {
		toSerialize["concurrent_searches"] = o.ConcurrentSearches
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
	if !IsNil(o.ItemsPerSearch) {
		toSerialize["items_per_search"] = o.ItemsPerSearch
	}
	if !IsNil(o.SavedId) {
		toSerialize["saved_id"] = o.SavedId
	}
	if !IsNil(o.ThreatFilters) {
		toSerialize["threat_filters"] = o.ThreatFilters
	}
	if !IsNil(o.ThreatIndicatorPath) {
		toSerialize["threat_indicator_path"] = o.ThreatIndicatorPath
	}
	if !IsNil(o.ThreatLanguage) {
		toSerialize["threat_language"] = o.ThreatLanguage
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableThreatMatchRuleCreateFields struct {
	value *ThreatMatchRuleCreateFields
	isSet bool
}

func (v NullableThreatMatchRuleCreateFields) Get() *ThreatMatchRuleCreateFields {
	return v.value
}

func (v *NullableThreatMatchRuleCreateFields) Set(val *ThreatMatchRuleCreateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatMatchRuleCreateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatMatchRuleCreateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatMatchRuleCreateFields(val *ThreatMatchRuleCreateFields) *NullableThreatMatchRuleCreateFields {
	return &NullableThreatMatchRuleCreateFields{value: val, isSet: true}
}

func (v NullableThreatMatchRuleCreateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatMatchRuleCreateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
