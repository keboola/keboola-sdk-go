# ABSCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** |  | 
**ContainerName** | Pointer to **string** | of-container (required) | [optional] 
**Owner** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | 06-17T15:07:48+01:00 | [optional] 
**Creator** | Pointer to [**ABSCreatedAllOfCreator**](ABSCreatedAllOfCreator.md) |  | [optional] 

## Methods

### NewABSCreated

`func NewABSCreated(accountName string, owner string, ) *ABSCreated`

NewABSCreated instantiates a new ABSCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewABSCreatedWithDefaults

`func NewABSCreatedWithDefaults() *ABSCreated`

NewABSCreatedWithDefaults instantiates a new ABSCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ABSCreated) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ABSCreated) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ABSCreated) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetContainerName

`func (o *ABSCreated) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ABSCreated) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ABSCreated) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ABSCreated) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetOwner

`func (o *ABSCreated) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ABSCreated) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ABSCreated) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetId

`func (o *ABSCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ABSCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ABSCreated) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ABSCreated) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *ABSCreated) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ABSCreated) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ABSCreated) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ABSCreated) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *ABSCreated) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ABSCreated) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ABSCreated) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ABSCreated) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreated

`func (o *ABSCreated) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ABSCreated) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ABSCreated) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ABSCreated) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *ABSCreated) GetCreator() ABSCreatedAllOfCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ABSCreated) GetCreatorOk() (*ABSCreatedAllOfCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ABSCreated) SetCreator(v ABSCreatedAllOfCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ABSCreated) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


