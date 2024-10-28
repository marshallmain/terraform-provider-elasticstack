# OsqueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsMapping** | Pointer to [**map[string]EcsMappingValue**](EcsMappingValue.md) |  | [optional] 
**PackId** | Pointer to **string** |  | [optional] 
**Queries** | Pointer to [**[]OsqueryQuery**](OsqueryQuery.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SavedQueryId** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **float32** |  | [optional] 

## Methods

### NewOsqueryParams

`func NewOsqueryParams() *OsqueryParams`

NewOsqueryParams instantiates a new OsqueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsqueryParamsWithDefaults

`func NewOsqueryParamsWithDefaults() *OsqueryParams`

NewOsqueryParamsWithDefaults instantiates a new OsqueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsMapping

`func (o *OsqueryParams) GetEcsMapping() map[string]EcsMappingValue`

GetEcsMapping returns the EcsMapping field if non-nil, zero value otherwise.

### GetEcsMappingOk

`func (o *OsqueryParams) GetEcsMappingOk() (*map[string]EcsMappingValue, bool)`

GetEcsMappingOk returns a tuple with the EcsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsMapping

`func (o *OsqueryParams) SetEcsMapping(v map[string]EcsMappingValue)`

SetEcsMapping sets EcsMapping field to given value.

### HasEcsMapping

`func (o *OsqueryParams) HasEcsMapping() bool`

HasEcsMapping returns a boolean if a field has been set.

### GetPackId

`func (o *OsqueryParams) GetPackId() string`

GetPackId returns the PackId field if non-nil, zero value otherwise.

### GetPackIdOk

`func (o *OsqueryParams) GetPackIdOk() (*string, bool)`

GetPackIdOk returns a tuple with the PackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackId

`func (o *OsqueryParams) SetPackId(v string)`

SetPackId sets PackId field to given value.

### HasPackId

`func (o *OsqueryParams) HasPackId() bool`

HasPackId returns a boolean if a field has been set.

### GetQueries

`func (o *OsqueryParams) GetQueries() []OsqueryQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *OsqueryParams) GetQueriesOk() (*[]OsqueryQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *OsqueryParams) SetQueries(v []OsqueryQuery)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *OsqueryParams) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetQuery

`func (o *OsqueryParams) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OsqueryParams) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OsqueryParams) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *OsqueryParams) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSavedQueryId

`func (o *OsqueryParams) GetSavedQueryId() string`

GetSavedQueryId returns the SavedQueryId field if non-nil, zero value otherwise.

### GetSavedQueryIdOk

`func (o *OsqueryParams) GetSavedQueryIdOk() (*string, bool)`

GetSavedQueryIdOk returns a tuple with the SavedQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedQueryId

`func (o *OsqueryParams) SetSavedQueryId(v string)`

SetSavedQueryId sets SavedQueryId field to given value.

### HasSavedQueryId

`func (o *OsqueryParams) HasSavedQueryId() bool`

HasSavedQueryId returns a boolean if a field has been set.

### GetTimeout

`func (o *OsqueryParams) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OsqueryParams) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OsqueryParams) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OsqueryParams) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


