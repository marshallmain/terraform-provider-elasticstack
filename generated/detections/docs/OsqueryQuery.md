# OsqueryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsMapping** | Pointer to [**map[string]EcsMappingValue**](EcsMappingValue.md) |  | [optional] 
**Id** | **string** | Query ID | 
**Platform** | Pointer to **string** |  | [optional] 
**Query** | **string** | Query to run | 
**Removed** | Pointer to **bool** |  | [optional] 
**Snapshot** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** | Query version | [optional] 

## Methods

### NewOsqueryQuery

`func NewOsqueryQuery(id string, query string, ) *OsqueryQuery`

NewOsqueryQuery instantiates a new OsqueryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsqueryQueryWithDefaults

`func NewOsqueryQueryWithDefaults() *OsqueryQuery`

NewOsqueryQueryWithDefaults instantiates a new OsqueryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsMapping

`func (o *OsqueryQuery) GetEcsMapping() map[string]EcsMappingValue`

GetEcsMapping returns the EcsMapping field if non-nil, zero value otherwise.

### GetEcsMappingOk

`func (o *OsqueryQuery) GetEcsMappingOk() (*map[string]EcsMappingValue, bool)`

GetEcsMappingOk returns a tuple with the EcsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsMapping

`func (o *OsqueryQuery) SetEcsMapping(v map[string]EcsMappingValue)`

SetEcsMapping sets EcsMapping field to given value.

### HasEcsMapping

`func (o *OsqueryQuery) HasEcsMapping() bool`

HasEcsMapping returns a boolean if a field has been set.

### GetId

`func (o *OsqueryQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsqueryQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsqueryQuery) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *OsqueryQuery) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OsqueryQuery) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OsqueryQuery) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *OsqueryQuery) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetQuery

`func (o *OsqueryQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OsqueryQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OsqueryQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRemoved

`func (o *OsqueryQuery) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *OsqueryQuery) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *OsqueryQuery) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *OsqueryQuery) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetSnapshot

`func (o *OsqueryQuery) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *OsqueryQuery) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *OsqueryQuery) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *OsqueryQuery) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetVersion

`func (o *OsqueryQuery) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OsqueryQuery) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OsqueryQuery) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OsqueryQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


