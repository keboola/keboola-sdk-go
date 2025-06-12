# ListMyOrganizationInvitations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | [optional] 
**Creator** | Pointer to [**ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md) |  | [optional] 

## Methods

### NewListMyOrganizationInvitations200ResponseInner

`func NewListMyOrganizationInvitations200ResponseInner() *ListMyOrganizationInvitations200ResponseInner`

NewListMyOrganizationInvitations200ResponseInner instantiates a new ListMyOrganizationInvitations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMyOrganizationInvitations200ResponseInnerWithDefaults

`func NewListMyOrganizationInvitations200ResponseInnerWithDefaults() *ListMyOrganizationInvitations200ResponseInner`

NewListMyOrganizationInvitations200ResponseInnerWithDefaults instantiates a new ListMyOrganizationInvitations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMyOrganizationInvitations200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMyOrganizationInvitations200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMyOrganizationInvitations200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ListMyOrganizationInvitations200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *ListMyOrganizationInvitations200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListMyOrganizationInvitations200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListMyOrganizationInvitations200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListMyOrganizationInvitations200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOrganization

`func (o *ListMyOrganizationInvitations200ResponseInner) GetOrganization() TokenVerification200ResponseCreator`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ListMyOrganizationInvitations200ResponseInner) GetOrganizationOk() (*TokenVerification200ResponseCreator, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ListMyOrganizationInvitations200ResponseInner) SetOrganization(v TokenVerification200ResponseCreator)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ListMyOrganizationInvitations200ResponseInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCreator

`func (o *ListMyOrganizationInvitations200ResponseInner) GetCreator() ListMaintainersInvitations200ResponseInnerUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ListMyOrganizationInvitations200ResponseInner) GetCreatorOk() (*ListMaintainersInvitations200ResponseInnerUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ListMyOrganizationInvitations200ResponseInner) SetCreator(v ListMaintainersInvitations200ResponseInnerUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ListMyOrganizationInvitations200ResponseInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


