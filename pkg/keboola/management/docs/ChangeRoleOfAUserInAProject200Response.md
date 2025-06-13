# ChangeRoleOfAUserInAProject200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]interface{}** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Invitor** | Pointer to [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | [optional] 
**Approver** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChangeRoleOfAUserInAProject200Response

`func NewChangeRoleOfAUserInAProject200Response() *ChangeRoleOfAUserInAProject200Response`

NewChangeRoleOfAUserInAProject200Response instantiates a new ChangeRoleOfAUserInAProject200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRoleOfAUserInAProject200ResponseWithDefaults

`func NewChangeRoleOfAUserInAProject200ResponseWithDefaults() *ChangeRoleOfAUserInAProject200Response`

NewChangeRoleOfAUserInAProject200ResponseWithDefaults instantiates a new ChangeRoleOfAUserInAProject200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRoleOfAUserInAProject200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRoleOfAUserInAProject200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeRoleOfAUserInAProject200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChangeRoleOfAUserInAProject200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeRoleOfAUserInAProject200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChangeRoleOfAUserInAProject200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ChangeRoleOfAUserInAProject200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChangeRoleOfAUserInAProject200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ChangeRoleOfAUserInAProject200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFeatures

`func (o *ChangeRoleOfAUserInAProject200Response) GetFeatures() []interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetFeaturesOk() (*[]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ChangeRoleOfAUserInAProject200Response) SetFeatures(v []interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ChangeRoleOfAUserInAProject200Response) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetExpires

`func (o *ChangeRoleOfAUserInAProject200Response) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ChangeRoleOfAUserInAProject200Response) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ChangeRoleOfAUserInAProject200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *ChangeRoleOfAUserInAProject200Response) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ChangeRoleOfAUserInAProject200Response) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetCreated

`func (o *ChangeRoleOfAUserInAProject200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChangeRoleOfAUserInAProject200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChangeRoleOfAUserInAProject200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetReason

`func (o *ChangeRoleOfAUserInAProject200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChangeRoleOfAUserInAProject200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChangeRoleOfAUserInAProject200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *ChangeRoleOfAUserInAProject200Response) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChangeRoleOfAUserInAProject200Response) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ChangeRoleOfAUserInAProject200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ChangeRoleOfAUserInAProject200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeRoleOfAUserInAProject200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChangeRoleOfAUserInAProject200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInvitor

`func (o *ChangeRoleOfAUserInAProject200Response) GetInvitor() ListMaintainersInvitations200ResponseInnerUser`

GetInvitor returns the Invitor field if non-nil, zero value otherwise.

### GetInvitorOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetInvitorOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetInvitorOk returns a tuple with the Invitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitor

`func (o *ChangeRoleOfAUserInAProject200Response) SetInvitor(v ListMaintainersInvitations200ResponseInnerUser)`

SetInvitor sets Invitor field to given value.

### HasInvitor

`func (o *ChangeRoleOfAUserInAProject200Response) HasInvitor() bool`

HasInvitor returns a boolean if a field has been set.

### GetApprover

`func (o *ChangeRoleOfAUserInAProject200Response) GetApprover() interface{}`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ChangeRoleOfAUserInAProject200Response) GetApproverOk() (*interface{}, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ChangeRoleOfAUserInAProject200Response) SetApprover(v interface{})`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *ChangeRoleOfAUserInAProject200Response) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *ChangeRoleOfAUserInAProject200Response) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *ChangeRoleOfAUserInAProject200Response) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


