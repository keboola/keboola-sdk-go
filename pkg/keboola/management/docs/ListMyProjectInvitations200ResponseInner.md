# ListMyProjectInvitations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | [optional] 
**Creator** | Pointer to [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | [optional] 

## Methods

### NewListMyProjectInvitations200ResponseInner

`func NewListMyProjectInvitations200ResponseInner() *ListMyProjectInvitations200ResponseInner`

NewListMyProjectInvitations200ResponseInner instantiates a new ListMyProjectInvitations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMyProjectInvitations200ResponseInnerWithDefaults

`func NewListMyProjectInvitations200ResponseInnerWithDefaults() *ListMyProjectInvitations200ResponseInner`

NewListMyProjectInvitations200ResponseInnerWithDefaults instantiates a new ListMyProjectInvitations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMyProjectInvitations200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMyProjectInvitations200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMyProjectInvitations200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ListMyProjectInvitations200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *ListMyProjectInvitations200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListMyProjectInvitations200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListMyProjectInvitations200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListMyProjectInvitations200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *ListMyProjectInvitations200ResponseInner) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListMyProjectInvitations200ResponseInner) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListMyProjectInvitations200ResponseInner) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ListMyProjectInvitations200ResponseInner) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *ListMyProjectInvitations200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListMyProjectInvitations200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetReason

`func (o *ListMyProjectInvitations200ResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListMyProjectInvitations200ResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListMyProjectInvitations200ResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListMyProjectInvitations200ResponseInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *ListMyProjectInvitations200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListMyProjectInvitations200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListMyProjectInvitations200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListMyProjectInvitations200ResponseInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetProject

`func (o *ListMyProjectInvitations200ResponseInner) GetProject() TokenVerification200ResponseCreator`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ListMyProjectInvitations200ResponseInner) GetProjectOk() (*TokenVerification200ResponseCreator, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ListMyProjectInvitations200ResponseInner) SetProject(v TokenVerification200ResponseCreator)`

SetProject sets Project field to given value.

### HasProject

`func (o *ListMyProjectInvitations200ResponseInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetCreator

`func (o *ListMyProjectInvitations200ResponseInner) GetCreator() ListMaintainersInvitations200ResponseInnerUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ListMyProjectInvitations200ResponseInner) GetCreatorOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ListMyProjectInvitations200ResponseInner) SetCreator(v ListMaintainersInvitations200ResponseInnerUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ListMyProjectInvitations200ResponseInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


