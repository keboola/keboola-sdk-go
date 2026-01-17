# PurgeDeletedProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreBackendErrors** | Pointer to **bool** | Ignore all errors from the backend, e.g. deleted Redshift cluster. &#x60;false&#x60; is default. | [optional] 

## Methods

### NewPurgeDeletedProjectRequest

`func NewPurgeDeletedProjectRequest() *PurgeDeletedProjectRequest`

NewPurgeDeletedProjectRequest instantiates a new PurgeDeletedProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeDeletedProjectRequestWithDefaults

`func NewPurgeDeletedProjectRequestWithDefaults() *PurgeDeletedProjectRequest`

NewPurgeDeletedProjectRequestWithDefaults instantiates a new PurgeDeletedProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreBackendErrors

`func (o *PurgeDeletedProjectRequest) GetIgnoreBackendErrors() bool`

GetIgnoreBackendErrors returns the IgnoreBackendErrors field if non-nil, zero value otherwise.

### GetIgnoreBackendErrorsOk

`func (o *PurgeDeletedProjectRequest) GetIgnoreBackendErrorsOk() (*bool, bool)`

GetIgnoreBackendErrorsOk returns a tuple with the IgnoreBackendErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreBackendErrors

`func (o *PurgeDeletedProjectRequest) SetIgnoreBackendErrors(v bool)`

SetIgnoreBackendErrors sets IgnoreBackendErrors field to given value.

### HasIgnoreBackendErrors

`func (o *PurgeDeletedProjectRequest) HasIgnoreBackendErrors() bool`

HasIgnoreBackendErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


