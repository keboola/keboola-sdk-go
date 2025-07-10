# ChangeRoleOfAUserInAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | User role in the project. Roles &#x60;admin&#x60;, &#x60;guest&#x60;, &#x60;readOnly&#x60; and &#x60;share&#x60; are allowed (default &#x60;admin&#x60;). | [optional] 

## Methods

### NewChangeRoleOfAUserInAProjectRequest

`func NewChangeRoleOfAUserInAProjectRequest() *ChangeRoleOfAUserInAProjectRequest`

NewChangeRoleOfAUserInAProjectRequest instantiates a new ChangeRoleOfAUserInAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRoleOfAUserInAProjectRequestWithDefaults

`func NewChangeRoleOfAUserInAProjectRequestWithDefaults() *ChangeRoleOfAUserInAProjectRequest`

NewChangeRoleOfAUserInAProjectRequestWithDefaults instantiates a new ChangeRoleOfAUserInAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChangeRoleOfAUserInAProjectRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChangeRoleOfAUserInAProjectRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChangeRoleOfAUserInAProjectRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ChangeRoleOfAUserInAProjectRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


