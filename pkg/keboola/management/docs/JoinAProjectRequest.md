# JoinAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason for joining the project | [optional] 
**ExpirationSeconds** | Pointer to **float32** | The number of seconds until the user&#39;s membership expires | [optional] 

## Methods

### NewJoinAProjectRequest

`func NewJoinAProjectRequest() *JoinAProjectRequest`

NewJoinAProjectRequest instantiates a new JoinAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinAProjectRequestWithDefaults

`func NewJoinAProjectRequestWithDefaults() *JoinAProjectRequest`

NewJoinAProjectRequestWithDefaults instantiates a new JoinAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *JoinAProjectRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JoinAProjectRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JoinAProjectRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JoinAProjectRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExpirationSeconds

`func (o *JoinAProjectRequest) GetExpirationSeconds() float32`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *JoinAProjectRequest) GetExpirationSecondsOk() (*float32, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *JoinAProjectRequest) SetExpirationSeconds(v float32)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *JoinAProjectRequest) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


