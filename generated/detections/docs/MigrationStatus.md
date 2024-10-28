# MigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that is not empty and does not contain only whitespace | 
**Status** | **string** |  | 
**Updated** | **time.Time** |  | 
**Version** | **int32** |  | 

## Methods

### NewMigrationStatus

`func NewMigrationStatus(id string, status string, updated time.Time, version int32, ) *MigrationStatus`

NewMigrationStatus instantiates a new MigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationStatusWithDefaults

`func NewMigrationStatusWithDefaults() *MigrationStatus`

NewMigrationStatusWithDefaults instantiates a new MigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationStatus) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *MigrationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *MigrationStatus) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MigrationStatus) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MigrationStatus) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetVersion

`func (o *MigrationStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MigrationStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MigrationStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


