# ThreatMatchRuleUpdateProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**ThreatIndex** | **[]string** |  | 
**ThreatMapping** | [**[]ThreatMappingInner**](ThreatMappingInner.md) |  | 
**ThreatQuery** | **string** | Query to run | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**ConcurrentSearches** | Pointer to **int32** |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ItemsPerSearch** | Pointer to **int32** |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**ThreatFilters** | Pointer to **[]interface{}** |  | [optional] 
**ThreatIndicatorPath** | Pointer to **string** | Defines the path to the threat indicator in the indicator documents (optional) | [optional] 
**ThreatLanguage** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Actions** | Pointer to [**[]RuleAction**](RuleAction.md) |  | [optional] 
**AliasPurpose** | Pointer to [**SavedObjectResolveAliasPurpose**](SavedObjectResolveAliasPurpose.md) |  | [optional] 
**AliasTargetId** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **[]string** |  | [optional] 
**BuildingBlockType** | Pointer to **string** | Determines if the rule acts as a building block. By default, building-block alerts are not displayed in the UI. These rules are used as a foundation for other rules that do generate alerts. Its value must be default. | [optional] 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** | Determines whether the rule is enabled. | [optional] 
**ExceptionsList** | Pointer to [**[]RuleExceptionList**](RuleExceptionList.md) |  | [optional] 
**FalsePositives** | Pointer to **[]string** |  | [optional] 
**From** | Pointer to **string** | Time from which data is analyzed each time the rule runs, using a date math range. For example, now-4200s means the rule analyzes data from 70 minutes before its start time. Defaults to now-6m (analyzes data from 6 minutes before the start time). | [optional] 
**Id** | Pointer to **string** | A universally unique identifier | [optional] 
**Interval** | Pointer to **string** | Frequency of rule execution, using a date math range. For example, \&quot;1h\&quot; means the rule runs every hour. Defaults to 5m (5 minutes). | [optional] 
**InvestigationFields** | Pointer to [**InvestigationFields**](InvestigationFields.md) |  | [optional] 
**License** | Pointer to **string** | The rule&#39;s license. | [optional] 
**MaxSignals** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** | Has no effect. | [optional] 
**Note** | Pointer to **string** | Notes to help investigate alerts produced by the rule. | [optional] 
**Outcome** | Pointer to [**SavedObjectResolveOutcome**](SavedObjectResolveOutcome.md) |  | [optional] 
**OutputIndex** | Pointer to **string** | (deprecated) Has no effect. | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**RelatedIntegrations** | Pointer to [**[]RelatedIntegration**](RelatedIntegration.md) |  | [optional] 
**RequiredFields** | Pointer to [**[]RequiredFieldInput**](RequiredFieldInput.md) |  | [optional] 
**RiskScore** | **int32** | Risk score (0 to 100) | 
**RiskScoreMapping** | Pointer to [**[]RiskScoreMappingInner**](RiskScoreMappingInner.md) | Overrides generated alerts&#39; risk_score with a value from the source event | [optional] 
**RuleId** | Pointer to **string** | Could be any string, not necessarily a UUID | [optional] 
**RuleNameOverride** | Pointer to **string** | Sets the source field for the alert&#39;s signal.rule.name value | [optional] 
**Setup** | Pointer to **string** |  | [optional] 
**Severity** | [**Severity**](Severity.md) |  | 
**SeverityMapping** | Pointer to [**[]SeverityMappingInner**](SeverityMappingInner.md) | Overrides generated alerts&#39; severity with values from the source event | [optional] 
**Tags** | Pointer to **[]string** | String array containing words and phrases to help categorize, filter, and search rules. Defaults to an empty array. | [optional] 
**Threat** | Pointer to [**[]Threat**](Threat.md) |  | [optional] 
**Throttle** | Pointer to [**RuleActionThrottle**](RuleActionThrottle.md) |  | [optional] 
**TimelineId** | Pointer to **string** | Timeline template ID | [optional] 
**TimelineTitle** | Pointer to **string** | Timeline template title | [optional] 
**TimestampOverride** | Pointer to **string** | Sets the time field used to query indices | [optional] 
**TimestampOverrideFallbackDisabled** | Pointer to **bool** | Disables the fallback to the event&#39;s @timestamp field | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** | The rule&#39;s version number. | [optional] 

## Methods

### NewThreatMatchRuleUpdateProps

`func NewThreatMatchRuleUpdateProps(query string, threatIndex []string, threatMapping []ThreatMappingInner, threatQuery string, type_ string, description string, name string, riskScore int32, severity Severity, ) *ThreatMatchRuleUpdateProps`

NewThreatMatchRuleUpdateProps instantiates a new ThreatMatchRuleUpdateProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatMatchRuleUpdatePropsWithDefaults

`func NewThreatMatchRuleUpdatePropsWithDefaults() *ThreatMatchRuleUpdateProps`

NewThreatMatchRuleUpdatePropsWithDefaults instantiates a new ThreatMatchRuleUpdateProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ThreatMatchRuleUpdateProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ThreatMatchRuleUpdateProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ThreatMatchRuleUpdateProps) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetThreatIndex

`func (o *ThreatMatchRuleUpdateProps) GetThreatIndex() []string`

GetThreatIndex returns the ThreatIndex field if non-nil, zero value otherwise.

### GetThreatIndexOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatIndexOk() (*[]string, bool)`

GetThreatIndexOk returns a tuple with the ThreatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndex

`func (o *ThreatMatchRuleUpdateProps) SetThreatIndex(v []string)`

SetThreatIndex sets ThreatIndex field to given value.


### GetThreatMapping

`func (o *ThreatMatchRuleUpdateProps) GetThreatMapping() []ThreatMappingInner`

GetThreatMapping returns the ThreatMapping field if non-nil, zero value otherwise.

### GetThreatMappingOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatMappingOk() (*[]ThreatMappingInner, bool)`

GetThreatMappingOk returns a tuple with the ThreatMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMapping

`func (o *ThreatMatchRuleUpdateProps) SetThreatMapping(v []ThreatMappingInner)`

SetThreatMapping sets ThreatMapping field to given value.


### GetThreatQuery

`func (o *ThreatMatchRuleUpdateProps) GetThreatQuery() string`

GetThreatQuery returns the ThreatQuery field if non-nil, zero value otherwise.

### GetThreatQueryOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatQueryOk() (*string, bool)`

GetThreatQueryOk returns a tuple with the ThreatQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatQuery

`func (o *ThreatMatchRuleUpdateProps) SetThreatQuery(v string)`

SetThreatQuery sets ThreatQuery field to given value.


### GetType

`func (o *ThreatMatchRuleUpdateProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreatMatchRuleUpdateProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreatMatchRuleUpdateProps) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *ThreatMatchRuleUpdateProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *ThreatMatchRuleUpdateProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *ThreatMatchRuleUpdateProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *ThreatMatchRuleUpdateProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetConcurrentSearches

`func (o *ThreatMatchRuleUpdateProps) GetConcurrentSearches() int32`

GetConcurrentSearches returns the ConcurrentSearches field if non-nil, zero value otherwise.

### GetConcurrentSearchesOk

`func (o *ThreatMatchRuleUpdateProps) GetConcurrentSearchesOk() (*int32, bool)`

GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSearches

`func (o *ThreatMatchRuleUpdateProps) SetConcurrentSearches(v int32)`

SetConcurrentSearches sets ConcurrentSearches field to given value.

### HasConcurrentSearches

`func (o *ThreatMatchRuleUpdateProps) HasConcurrentSearches() bool`

HasConcurrentSearches returns a boolean if a field has been set.

### GetDataViewId

`func (o *ThreatMatchRuleUpdateProps) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ThreatMatchRuleUpdateProps) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ThreatMatchRuleUpdateProps) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ThreatMatchRuleUpdateProps) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *ThreatMatchRuleUpdateProps) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ThreatMatchRuleUpdateProps) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ThreatMatchRuleUpdateProps) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ThreatMatchRuleUpdateProps) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *ThreatMatchRuleUpdateProps) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ThreatMatchRuleUpdateProps) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ThreatMatchRuleUpdateProps) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ThreatMatchRuleUpdateProps) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetItemsPerSearch

`func (o *ThreatMatchRuleUpdateProps) GetItemsPerSearch() int32`

GetItemsPerSearch returns the ItemsPerSearch field if non-nil, zero value otherwise.

### GetItemsPerSearchOk

`func (o *ThreatMatchRuleUpdateProps) GetItemsPerSearchOk() (*int32, bool)`

GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerSearch

`func (o *ThreatMatchRuleUpdateProps) SetItemsPerSearch(v int32)`

SetItemsPerSearch sets ItemsPerSearch field to given value.

### HasItemsPerSearch

`func (o *ThreatMatchRuleUpdateProps) HasItemsPerSearch() bool`

HasItemsPerSearch returns a boolean if a field has been set.

### GetSavedId

`func (o *ThreatMatchRuleUpdateProps) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *ThreatMatchRuleUpdateProps) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *ThreatMatchRuleUpdateProps) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *ThreatMatchRuleUpdateProps) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetThreatFilters

`func (o *ThreatMatchRuleUpdateProps) GetThreatFilters() []interface{}`

GetThreatFilters returns the ThreatFilters field if non-nil, zero value otherwise.

### GetThreatFiltersOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatFiltersOk() (*[]interface{}, bool)`

GetThreatFiltersOk returns a tuple with the ThreatFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFilters

`func (o *ThreatMatchRuleUpdateProps) SetThreatFilters(v []interface{})`

SetThreatFilters sets ThreatFilters field to given value.

### HasThreatFilters

`func (o *ThreatMatchRuleUpdateProps) HasThreatFilters() bool`

HasThreatFilters returns a boolean if a field has been set.

### GetThreatIndicatorPath

`func (o *ThreatMatchRuleUpdateProps) GetThreatIndicatorPath() string`

GetThreatIndicatorPath returns the ThreatIndicatorPath field if non-nil, zero value otherwise.

### GetThreatIndicatorPathOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatIndicatorPathOk() (*string, bool)`

GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndicatorPath

`func (o *ThreatMatchRuleUpdateProps) SetThreatIndicatorPath(v string)`

SetThreatIndicatorPath sets ThreatIndicatorPath field to given value.

### HasThreatIndicatorPath

`func (o *ThreatMatchRuleUpdateProps) HasThreatIndicatorPath() bool`

HasThreatIndicatorPath returns a boolean if a field has been set.

### GetThreatLanguage

`func (o *ThreatMatchRuleUpdateProps) GetThreatLanguage() KqlQueryLanguage`

GetThreatLanguage returns the ThreatLanguage field if non-nil, zero value otherwise.

### GetThreatLanguageOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatLanguageOk() (*KqlQueryLanguage, bool)`

GetThreatLanguageOk returns a tuple with the ThreatLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLanguage

`func (o *ThreatMatchRuleUpdateProps) SetThreatLanguage(v KqlQueryLanguage)`

SetThreatLanguage sets ThreatLanguage field to given value.

### HasThreatLanguage

`func (o *ThreatMatchRuleUpdateProps) HasThreatLanguage() bool`

HasThreatLanguage returns a boolean if a field has been set.

### GetLanguage

`func (o *ThreatMatchRuleUpdateProps) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThreatMatchRuleUpdateProps) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThreatMatchRuleUpdateProps) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ThreatMatchRuleUpdateProps) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetActions

`func (o *ThreatMatchRuleUpdateProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ThreatMatchRuleUpdateProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ThreatMatchRuleUpdateProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ThreatMatchRuleUpdateProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *ThreatMatchRuleUpdateProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *ThreatMatchRuleUpdateProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *ThreatMatchRuleUpdateProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *ThreatMatchRuleUpdateProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *ThreatMatchRuleUpdateProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *ThreatMatchRuleUpdateProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *ThreatMatchRuleUpdateProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *ThreatMatchRuleUpdateProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *ThreatMatchRuleUpdateProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ThreatMatchRuleUpdateProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ThreatMatchRuleUpdateProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ThreatMatchRuleUpdateProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *ThreatMatchRuleUpdateProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *ThreatMatchRuleUpdateProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *ThreatMatchRuleUpdateProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *ThreatMatchRuleUpdateProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatMatchRuleUpdateProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatMatchRuleUpdateProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatMatchRuleUpdateProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *ThreatMatchRuleUpdateProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThreatMatchRuleUpdateProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThreatMatchRuleUpdateProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ThreatMatchRuleUpdateProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *ThreatMatchRuleUpdateProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *ThreatMatchRuleUpdateProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *ThreatMatchRuleUpdateProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *ThreatMatchRuleUpdateProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *ThreatMatchRuleUpdateProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *ThreatMatchRuleUpdateProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *ThreatMatchRuleUpdateProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *ThreatMatchRuleUpdateProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *ThreatMatchRuleUpdateProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ThreatMatchRuleUpdateProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ThreatMatchRuleUpdateProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ThreatMatchRuleUpdateProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *ThreatMatchRuleUpdateProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreatMatchRuleUpdateProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreatMatchRuleUpdateProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreatMatchRuleUpdateProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *ThreatMatchRuleUpdateProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ThreatMatchRuleUpdateProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ThreatMatchRuleUpdateProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ThreatMatchRuleUpdateProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *ThreatMatchRuleUpdateProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *ThreatMatchRuleUpdateProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *ThreatMatchRuleUpdateProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *ThreatMatchRuleUpdateProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *ThreatMatchRuleUpdateProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ThreatMatchRuleUpdateProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ThreatMatchRuleUpdateProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ThreatMatchRuleUpdateProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *ThreatMatchRuleUpdateProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *ThreatMatchRuleUpdateProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *ThreatMatchRuleUpdateProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *ThreatMatchRuleUpdateProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *ThreatMatchRuleUpdateProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThreatMatchRuleUpdateProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThreatMatchRuleUpdateProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThreatMatchRuleUpdateProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *ThreatMatchRuleUpdateProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatMatchRuleUpdateProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatMatchRuleUpdateProps) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ThreatMatchRuleUpdateProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ThreatMatchRuleUpdateProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ThreatMatchRuleUpdateProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ThreatMatchRuleUpdateProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *ThreatMatchRuleUpdateProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ThreatMatchRuleUpdateProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ThreatMatchRuleUpdateProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ThreatMatchRuleUpdateProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *ThreatMatchRuleUpdateProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ThreatMatchRuleUpdateProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ThreatMatchRuleUpdateProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ThreatMatchRuleUpdateProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *ThreatMatchRuleUpdateProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ThreatMatchRuleUpdateProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ThreatMatchRuleUpdateProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *ThreatMatchRuleUpdateProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *ThreatMatchRuleUpdateProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ThreatMatchRuleUpdateProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ThreatMatchRuleUpdateProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ThreatMatchRuleUpdateProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *ThreatMatchRuleUpdateProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *ThreatMatchRuleUpdateProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *ThreatMatchRuleUpdateProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *ThreatMatchRuleUpdateProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *ThreatMatchRuleUpdateProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *ThreatMatchRuleUpdateProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *ThreatMatchRuleUpdateProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *ThreatMatchRuleUpdateProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *ThreatMatchRuleUpdateProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ThreatMatchRuleUpdateProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ThreatMatchRuleUpdateProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskScoreMapping

`func (o *ThreatMatchRuleUpdateProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *ThreatMatchRuleUpdateProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *ThreatMatchRuleUpdateProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *ThreatMatchRuleUpdateProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *ThreatMatchRuleUpdateProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ThreatMatchRuleUpdateProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ThreatMatchRuleUpdateProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ThreatMatchRuleUpdateProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *ThreatMatchRuleUpdateProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *ThreatMatchRuleUpdateProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *ThreatMatchRuleUpdateProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *ThreatMatchRuleUpdateProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *ThreatMatchRuleUpdateProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ThreatMatchRuleUpdateProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ThreatMatchRuleUpdateProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *ThreatMatchRuleUpdateProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *ThreatMatchRuleUpdateProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ThreatMatchRuleUpdateProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ThreatMatchRuleUpdateProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSeverityMapping

`func (o *ThreatMatchRuleUpdateProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *ThreatMatchRuleUpdateProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *ThreatMatchRuleUpdateProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *ThreatMatchRuleUpdateProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *ThreatMatchRuleUpdateProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ThreatMatchRuleUpdateProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ThreatMatchRuleUpdateProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ThreatMatchRuleUpdateProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *ThreatMatchRuleUpdateProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *ThreatMatchRuleUpdateProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *ThreatMatchRuleUpdateProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *ThreatMatchRuleUpdateProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *ThreatMatchRuleUpdateProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *ThreatMatchRuleUpdateProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *ThreatMatchRuleUpdateProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *ThreatMatchRuleUpdateProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *ThreatMatchRuleUpdateProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *ThreatMatchRuleUpdateProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *ThreatMatchRuleUpdateProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *ThreatMatchRuleUpdateProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *ThreatMatchRuleUpdateProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *ThreatMatchRuleUpdateProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *ThreatMatchRuleUpdateProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *ThreatMatchRuleUpdateProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *ThreatMatchRuleUpdateProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *ThreatMatchRuleUpdateProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *ThreatMatchRuleUpdateProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *ThreatMatchRuleUpdateProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRuleUpdateProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *ThreatMatchRuleUpdateProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRuleUpdateProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRuleUpdateProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *ThreatMatchRuleUpdateProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ThreatMatchRuleUpdateProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ThreatMatchRuleUpdateProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ThreatMatchRuleUpdateProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *ThreatMatchRuleUpdateProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreatMatchRuleUpdateProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreatMatchRuleUpdateProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreatMatchRuleUpdateProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


