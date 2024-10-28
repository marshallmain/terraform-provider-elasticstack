# ErrorSchemaError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**StatusCode** | **int32** |  | 

## Methods

### NewErrorSchemaError

`func NewErrorSchemaError(message string, statusCode int32, ) *ErrorSchemaError`

NewErrorSchemaError instantiates a new ErrorSchemaError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorSchemaErrorWithDefaults

`func NewErrorSchemaErrorWithDefaults() *ErrorSchemaError`

NewErrorSchemaErrorWithDefaults instantiates a new ErrorSchemaError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorSchemaError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorSchemaError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorSchemaError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatusCode

`func (o *ErrorSchemaError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ErrorSchemaError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ErrorSchemaError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


