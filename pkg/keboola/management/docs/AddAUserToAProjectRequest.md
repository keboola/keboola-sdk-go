# AddAUserToAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of an invited user | 
**Role** | Pointer to **string** | User role in the project. Roles &#x60;admin&#x60;, &#x60;guest&#x60;, &#x60;readOnly&#x60; and &#x60;share&#x60; are allowed (default &#x60;admin&#x60;). | [optional] 
**ExpirationSeconds** | Pointer to **float32** | After how many seconds the membership of a user will expire | [optional] 

## Methods

### NewAddAUserToAProjectRequest

`func NewAddAUserToAProjectRequest(email string, ) *AddAUserToAProjectRequest`

NewAddAUserToAProjectRequest instantiates a new AddAUserToAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAUserToAProjectRequestWithDefaults

`func NewAddAUserToAProjectRequestWithDefaults() *AddAUserToAProjectRequest`

NewAddAUserToAProjectRequestWithDefaults instantiates a new AddAUserToAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AddAUserToAProjectRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddAUserToAProjectRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddAUserToAProjectRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *AddAUserToAProjectRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddAUserToAProjectRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddAUserToAProjectRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddAUserToAProjectRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetExpirationSeconds

`func (o *AddAUserToAProjectRequest) GetExpirationSeconds() float32`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *AddAUserToAProjectRequest) GetExpirationSecondsOk() (*float32, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *AddAUserToAProjectRequest) SetExpirationSeconds(v float32)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *AddAUserToAProjectRequest) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


