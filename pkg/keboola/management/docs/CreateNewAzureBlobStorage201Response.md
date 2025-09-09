# CreateNewAzureBlobStorage201Response

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
**Creator** | Pointer to [**CreateNewAWSS3Storage201ResponseCreator**](CreateNewAWSS3Storage201ResponseCreator.md) |  | [optional] 

## Methods

### NewCreateNewAzureBlobStorage201Response

`func NewCreateNewAzureBlobStorage201Response(accountName string, owner string, ) *CreateNewAzureBlobStorage201Response`

NewCreateNewAzureBlobStorage201Response instantiates a new CreateNewAzureBlobStorage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewAzureBlobStorage201ResponseWithDefaults

`func NewCreateNewAzureBlobStorage201ResponseWithDefaults() *CreateNewAzureBlobStorage201Response`

NewCreateNewAzureBlobStorage201ResponseWithDefaults instantiates a new CreateNewAzureBlobStorage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *CreateNewAzureBlobStorage201Response) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreateNewAzureBlobStorage201Response) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreateNewAzureBlobStorage201Response) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetContainerName

`func (o *CreateNewAzureBlobStorage201Response) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CreateNewAzureBlobStorage201Response) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CreateNewAzureBlobStorage201Response) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *CreateNewAzureBlobStorage201Response) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetOwner

`func (o *CreateNewAzureBlobStorage201Response) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNewAzureBlobStorage201Response) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNewAzureBlobStorage201Response) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetId

`func (o *CreateNewAzureBlobStorage201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNewAzureBlobStorage201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNewAzureBlobStorage201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNewAzureBlobStorage201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateNewAzureBlobStorage201Response) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateNewAzureBlobStorage201Response) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateNewAzureBlobStorage201Response) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateNewAzureBlobStorage201Response) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *CreateNewAzureBlobStorage201Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateNewAzureBlobStorage201Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateNewAzureBlobStorage201Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateNewAzureBlobStorage201Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreated

`func (o *CreateNewAzureBlobStorage201Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateNewAzureBlobStorage201Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateNewAzureBlobStorage201Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateNewAzureBlobStorage201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *CreateNewAzureBlobStorage201Response) GetCreator() CreateNewAWSS3Storage201ResponseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateNewAzureBlobStorage201Response) GetCreatorOk() (*CreateNewAWSS3Storage201ResponseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateNewAzureBlobStorage201Response) SetCreator(v CreateNewAWSS3Storage201ResponseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CreateNewAzureBlobStorage201Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


