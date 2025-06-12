# AddAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Project name | 
**Type** | **string** | One of &#x60;production&#x60;, &#x60;poc&#x60;, &#x60;demo&#x60;; default is &#x60;production&#x60; | 
**DefaultBackend** | Pointer to **string** | Project default backend &#x60;snowflake&#x60; or &#x60;redshift&#x60;; default is &#x60;snowflake&#x60; | [optional] 
**DataRetentionTimeInDays** | Pointer to **string** | Data retention in days for Time Travel | [optional] 

## Methods

### NewAddAProjectRequest

`func NewAddAProjectRequest(name string, type_ string, ) *AddAProjectRequest`

NewAddAProjectRequest instantiates a new AddAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAProjectRequestWithDefaults

`func NewAddAProjectRequestWithDefaults() *AddAProjectRequest`

NewAddAProjectRequestWithDefaults instantiates a new AddAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddAProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddAProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddAProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddAProjectRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddAProjectRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddAProjectRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultBackend

`func (o *AddAProjectRequest) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *AddAProjectRequest) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *AddAProjectRequest) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *AddAProjectRequest) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetDataRetentionTimeInDays

`func (o *AddAProjectRequest) GetDataRetentionTimeInDays() string`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *AddAProjectRequest) GetDataRetentionTimeInDaysOk() (*string, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *AddAProjectRequest) SetDataRetentionTimeInDays(v string)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.

### HasDataRetentionTimeInDays

`func (o *AddAProjectRequest) HasDataRetentionTimeInDays() bool`

HasDataRetentionTimeInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


