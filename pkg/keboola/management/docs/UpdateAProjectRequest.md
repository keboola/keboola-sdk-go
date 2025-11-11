# UpdateAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rename project | [optional] 
**DefaultBackend** | Pointer to **string** | Change project default backend type | [optional] 
**Type** | Pointer to **string** | Change project type (demo, production, poc, ...) - allowed only for a super admin | [optional] 
**ExpirationDays** | Pointer to **float32** | Change project expiration - allowed only for a super admin | [optional] 
**BilledMonthlyPrice** | Pointer to **float32** | Change project monthly fee - allowed only for a super admin | [optional] 
**DataRetentionTimeInDays** | Pointer to **float32** | (Snowflake only) Change the data retention period for this project - allowed only for a super admin | [optional] 

## Methods

### NewUpdateAProjectRequest

`func NewUpdateAProjectRequest() *UpdateAProjectRequest`

NewUpdateAProjectRequest instantiates a new UpdateAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAProjectRequestWithDefaults

`func NewUpdateAProjectRequestWithDefaults() *UpdateAProjectRequest`

NewUpdateAProjectRequestWithDefaults instantiates a new UpdateAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultBackend

`func (o *UpdateAProjectRequest) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *UpdateAProjectRequest) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *UpdateAProjectRequest) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *UpdateAProjectRequest) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetType

`func (o *UpdateAProjectRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAProjectRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAProjectRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateAProjectRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpirationDays

`func (o *UpdateAProjectRequest) GetExpirationDays() float32`

GetExpirationDays returns the ExpirationDays field if non-nil, zero value otherwise.

### GetExpirationDaysOk

`func (o *UpdateAProjectRequest) GetExpirationDaysOk() (*float32, bool)`

GetExpirationDaysOk returns a tuple with the ExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDays

`func (o *UpdateAProjectRequest) SetExpirationDays(v float32)`

SetExpirationDays sets ExpirationDays field to given value.

### HasExpirationDays

`func (o *UpdateAProjectRequest) HasExpirationDays() bool`

HasExpirationDays returns a boolean if a field has been set.

### GetBilledMonthlyPrice

`func (o *UpdateAProjectRequest) GetBilledMonthlyPrice() float32`

GetBilledMonthlyPrice returns the BilledMonthlyPrice field if non-nil, zero value otherwise.

### GetBilledMonthlyPriceOk

`func (o *UpdateAProjectRequest) GetBilledMonthlyPriceOk() (*float32, bool)`

GetBilledMonthlyPriceOk returns a tuple with the BilledMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMonthlyPrice

`func (o *UpdateAProjectRequest) SetBilledMonthlyPrice(v float32)`

SetBilledMonthlyPrice sets BilledMonthlyPrice field to given value.

### HasBilledMonthlyPrice

`func (o *UpdateAProjectRequest) HasBilledMonthlyPrice() bool`

HasBilledMonthlyPrice returns a boolean if a field has been set.

### GetDataRetentionTimeInDays

`func (o *UpdateAProjectRequest) GetDataRetentionTimeInDays() float32`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *UpdateAProjectRequest) GetDataRetentionTimeInDaysOk() (*float32, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *UpdateAProjectRequest) SetDataRetentionTimeInDays(v float32)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.

### HasDataRetentionTimeInDays

`func (o *UpdateAProjectRequest) HasDataRetentionTimeInDays() bool`

HasDataRetentionTimeInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


