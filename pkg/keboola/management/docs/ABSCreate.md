# ABSCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** |  | 
**ContainerName** | Pointer to **string** | of-container (required) | [optional] 
**Owner** | **string** |  | 
**AccountKey** | **string** | new key to be rotated to | 

## Methods

### NewABSCreate

`func NewABSCreate(accountName string, owner string, accountKey string, ) *ABSCreate`

NewABSCreate instantiates a new ABSCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewABSCreateWithDefaults

`func NewABSCreateWithDefaults() *ABSCreate`

NewABSCreateWithDefaults instantiates a new ABSCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ABSCreate) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ABSCreate) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ABSCreate) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetContainerName

`func (o *ABSCreate) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ABSCreate) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ABSCreate) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ABSCreate) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetOwner

`func (o *ABSCreate) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ABSCreate) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ABSCreate) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAccountKey

`func (o *ABSCreate) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *ABSCreate) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *ABSCreate) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


