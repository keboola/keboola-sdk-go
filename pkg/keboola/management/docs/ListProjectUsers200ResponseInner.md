# ListProjectUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**Features** | **[]interface{}** |  | 
**Expires** | **NullableString** |  | 
**Created** | **string** |  | 
**Reason** | **string** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**Invitor** | [**NullableListProjectUsers200ResponseInnerInvitor**](ListProjectUsers200ResponseInnerInvitor.md) |  | 
**Approver** | [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | 

## Methods

### NewListProjectUsers200ResponseInner

`func NewListProjectUsers200ResponseInner(id float32, name string, email string, features []interface{}, expires NullableString, created string, reason string, role string, status string, invitor NullableListProjectUsers200ResponseInnerInvitor, approver ListMaintainersInvitations200ResponseInnerUser, ) *ListProjectUsers200ResponseInner`

NewListProjectUsers200ResponseInner instantiates a new ListProjectUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProjectUsers200ResponseInnerWithDefaults

`func NewListProjectUsers200ResponseInnerWithDefaults() *ListProjectUsers200ResponseInner`

NewListProjectUsers200ResponseInnerWithDefaults instantiates a new ListProjectUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProjectUsers200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProjectUsers200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProjectUsers200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ListProjectUsers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProjectUsers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProjectUsers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ListProjectUsers200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListProjectUsers200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListProjectUsers200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFeatures

`func (o *ListProjectUsers200ResponseInner) GetFeatures() []interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ListProjectUsers200ResponseInner) GetFeaturesOk() (*[]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ListProjectUsers200ResponseInner) SetFeatures(v []interface{})`

SetFeatures sets Features field to given value.


### GetExpires

`func (o *ListProjectUsers200ResponseInner) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListProjectUsers200ResponseInner) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListProjectUsers200ResponseInner) SetExpires(v string)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *ListProjectUsers200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListProjectUsers200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetCreated

`func (o *ListProjectUsers200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListProjectUsers200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListProjectUsers200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetReason

`func (o *ListProjectUsers200ResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListProjectUsers200ResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListProjectUsers200ResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRole

`func (o *ListProjectUsers200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListProjectUsers200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListProjectUsers200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ListProjectUsers200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListProjectUsers200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListProjectUsers200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInvitor

`func (o *ListProjectUsers200ResponseInner) GetInvitor() ListProjectUsers200ResponseInnerInvitor`

GetInvitor returns the Invitor field if non-nil, zero value otherwise.

### GetInvitorOk

`func (o *ListProjectUsers200ResponseInner) GetInvitorOk() (*ListProjectUsers200ResponseInnerInvitor, bool)`

GetInvitorOk returns a tuple with the Invitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitor

`func (o *ListProjectUsers200ResponseInner) SetInvitor(v ListProjectUsers200ResponseInnerInvitor)`

SetInvitor sets Invitor field to given value.


### SetInvitorNil

`func (o *ListProjectUsers200ResponseInner) SetInvitorNil(b bool)`

 SetInvitorNil sets the value for Invitor to be an explicit nil

### UnsetInvitor
`func (o *ListProjectUsers200ResponseInner) UnsetInvitor()`

UnsetInvitor ensures that no value is present for Invitor, not even an explicit nil
### GetApprover

`func (o *ListProjectUsers200ResponseInner) GetApprover() ListMaintainersInvitations200ResponseInnerUser`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ListProjectUsers200ResponseInner) GetApproverOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ListProjectUsers200ResponseInner) SetApprover(v ListMaintainersInvitations200ResponseInnerUser)`

SetApprover sets Approver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


