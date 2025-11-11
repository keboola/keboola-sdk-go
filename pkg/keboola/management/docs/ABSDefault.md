# ABSDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** |  | 
**ContainerName** | Pointer to **string** | of-container (required) | [optional] 
**Owner** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **string** |  | [optional] 

## Methods

### NewABSDefault

`func NewABSDefault(accountName string, owner string, ) *ABSDefault`

NewABSDefault instantiates a new ABSDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewABSDefaultWithDefaults

`func NewABSDefaultWithDefaults() *ABSDefault`

NewABSDefaultWithDefaults instantiates a new ABSDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ABSDefault) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ABSDefault) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ABSDefault) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetContainerName

`func (o *ABSDefault) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ABSDefault) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ABSDefault) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ABSDefault) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetOwner

`func (o *ABSDefault) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ABSDefault) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ABSDefault) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetId

`func (o *ABSDefault) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ABSDefault) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ABSDefault) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ABSDefault) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *ABSDefault) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ABSDefault) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ABSDefault) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ABSDefault) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


