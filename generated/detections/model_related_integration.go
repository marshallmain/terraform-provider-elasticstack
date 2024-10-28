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

// checks if the RelatedIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedIntegration{}

// RelatedIntegration Related integration is a potential dependency of a rule. It's assumed that if the user installs one of the related integrations of a rule, the rule might start to work properly because it will have source events (generated by this integration) potentially matching the rule's query.  NOTE: Proper work is not guaranteed, because a related integration, if installed, can be configured differently or generate data that is not necessarily relevant for this rule.  Related integration is a combination of a Fleet package and (optionally) one of the package's \"integrations\" that this package contains. It is represented by 3 properties:  - `package`: name of the package (required, unique id) - `version`: version of the package (required, semver-compatible) - `integration`: name of the integration of this package (optional, id within the package)  There are Fleet packages like `windows` that contain only one integration; in this case, `integration` should be unspecified. There are also packages like `aws` and `azure` that contain several integrations; in this case, `integration` should be specified.  @example const x: RelatedIntegration = {   package: 'windows',   version: '1.5.x', };  @example const x: RelatedIntegration = {   package: 'azure',   version: '~1.1.6',   integration: 'activitylogs', };
type RelatedIntegration struct {
	// A string that is not empty and does not contain only whitespace
	Integration *string `json:"integration,omitempty"`
	// A string that is not empty and does not contain only whitespace
	Package string `json:"package"`
	// A string that is not empty and does not contain only whitespace
	Version string `json:"version"`
}

// NewRelatedIntegration instantiates a new RelatedIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedIntegration(package_ string, version string) *RelatedIntegration {
	this := RelatedIntegration{}
	this.Package = package_
	this.Version = version
	return &this
}

// NewRelatedIntegrationWithDefaults instantiates a new RelatedIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedIntegrationWithDefaults() *RelatedIntegration {
	this := RelatedIntegration{}
	return &this
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *RelatedIntegration) GetIntegration() string {
	if o == nil || IsNil(o.Integration) {
		var ret string
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedIntegration) GetIntegrationOk() (*string, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *RelatedIntegration) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given string and assigns it to the Integration field.
func (o *RelatedIntegration) SetIntegration(v string) {
	o.Integration = &v
}

// GetPackage returns the Package field value
func (o *RelatedIntegration) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *RelatedIntegration) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *RelatedIntegration) SetPackage(v string) {
	o.Package = v
}

// GetVersion returns the Version field value
func (o *RelatedIntegration) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RelatedIntegration) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RelatedIntegration) SetVersion(v string) {
	o.Version = v
}

func (o RelatedIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["package"] = o.Package
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableRelatedIntegration struct {
	value *RelatedIntegration
	isSet bool
}

func (v NullableRelatedIntegration) Get() *RelatedIntegration {
	return v.value
}

func (v *NullableRelatedIntegration) Set(val *RelatedIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedIntegration(val *RelatedIntegration) *NullableRelatedIntegration {
	return &NullableRelatedIntegration{value: val, isSet: true}
}

func (v NullableRelatedIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
