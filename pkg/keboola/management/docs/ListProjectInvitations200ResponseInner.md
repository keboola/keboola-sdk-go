# ListProjectInvitations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Created** | **string** |  | 
**Expires** | **NullableString** |  | 
**Reason** | **string** |  | 
**User** | [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | 
**Creator** | [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | 

## Methods

### NewListProjectInvitations200ResponseInner

`func NewListProjectInvitations200ResponseInner(id float32, created string, expires NullableString, reason string, user ListMaintainersInvitations200ResponseInnerUser, creator ListMaintainersInvitations200ResponseInnerUser, ) *ListProjectInvitations200ResponseInner`

NewListProjectInvitations200ResponseInner instantiates a new ListProjectInvitations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProjectInvitations200ResponseInnerWithDefaults

`func NewListProjectInvitations200ResponseInnerWithDefaults() *ListProjectInvitations200ResponseInner`

NewListProjectInvitations200ResponseInnerWithDefaults instantiates a new ListProjectInvitations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProjectInvitations200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProjectInvitations200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProjectInvitations200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetCreated

`func (o *ListProjectInvitations200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListProjectInvitations200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListProjectInvitations200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *ListProjectInvitations200ResponseInner) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListProjectInvitations200ResponseInner) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListProjectInvitations200ResponseInner) SetExpires(v string)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *ListProjectInvitations200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListProjectInvitations200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetReason

`func (o *ListProjectInvitations200ResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListProjectInvitations200ResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListProjectInvitations200ResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.


### GetUser

`func (o *ListProjectInvitations200ResponseInner) GetUser() ListMaintainersInvitations200ResponseInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListProjectInvitations200ResponseInner) GetUserOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListProjectInvitations200ResponseInner) SetUser(v ListMaintainersInvitations200ResponseInnerUser)`

SetUser sets User field to given value.


### GetCreator

`func (o *ListProjectInvitations200ResponseInner) GetCreator() ListMaintainersInvitations200ResponseInnerUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ListProjectInvitations200ResponseInner) GetCreatorOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ListProjectInvitations200ResponseInner) SetCreator(v ListMaintainersInvitations200ResponseInnerUser)`

SetCreator sets Creator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


