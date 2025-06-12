# JoinRequestDetail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | [optional] 

## Methods

### NewJoinRequestDetail200Response

`func NewJoinRequestDetail200Response() *JoinRequestDetail200Response`

NewJoinRequestDetail200Response instantiates a new JoinRequestDetail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinRequestDetail200ResponseWithDefaults

`func NewJoinRequestDetail200ResponseWithDefaults() *JoinRequestDetail200Response`

NewJoinRequestDetail200ResponseWithDefaults instantiates a new JoinRequestDetail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JoinRequestDetail200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JoinRequestDetail200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JoinRequestDetail200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *JoinRequestDetail200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *JoinRequestDetail200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JoinRequestDetail200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JoinRequestDetail200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JoinRequestDetail200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *JoinRequestDetail200Response) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *JoinRequestDetail200Response) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *JoinRequestDetail200Response) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *JoinRequestDetail200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *JoinRequestDetail200Response) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *JoinRequestDetail200Response) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetReason

`func (o *JoinRequestDetail200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JoinRequestDetail200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JoinRequestDetail200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JoinRequestDetail200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUser

`func (o *JoinRequestDetail200Response) GetUser() ListMaintainersInvitations200ResponseInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JoinRequestDetail200Response) GetUserOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JoinRequestDetail200Response) SetUser(v ListMaintainersInvitations200ResponseInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *JoinRequestDetail200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


