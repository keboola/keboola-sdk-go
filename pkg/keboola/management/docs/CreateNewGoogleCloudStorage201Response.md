# CreateNewGoogleCloudStorage201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | (ISO8601 format) &#x60;Y-m-d\\TH:i:sO&#x60; | [optional] 
**Creator** | Pointer to [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | [optional] 
**GcsCredentials** | Pointer to [**CreateNewGoogleCloudStorage201ResponseGcsCredentials**](CreateNewGoogleCloudStorage201ResponseGcsCredentials.md) |  | [optional] 
**GcsSnowflakeIntegrationName** | Pointer to **string** |  | [optional] 
**FilesBucket** | Pointer to **string** | name of the bucket | [optional] 

## Methods

### NewCreateNewGoogleCloudStorage201Response

`func NewCreateNewGoogleCloudStorage201Response() *CreateNewGoogleCloudStorage201Response`

NewCreateNewGoogleCloudStorage201Response instantiates a new CreateNewGoogleCloudStorage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewGoogleCloudStorage201ResponseWithDefaults

`func NewCreateNewGoogleCloudStorage201ResponseWithDefaults() *CreateNewGoogleCloudStorage201Response`

NewCreateNewGoogleCloudStorage201ResponseWithDefaults instantiates a new CreateNewGoogleCloudStorage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNewGoogleCloudStorage201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNewGoogleCloudStorage201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNewGoogleCloudStorage201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNewGoogleCloudStorage201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateNewGoogleCloudStorage201Response) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateNewGoogleCloudStorage201Response) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateNewGoogleCloudStorage201Response) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateNewGoogleCloudStorage201Response) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *CreateNewGoogleCloudStorage201Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateNewGoogleCloudStorage201Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateNewGoogleCloudStorage201Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateNewGoogleCloudStorage201Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *CreateNewGoogleCloudStorage201Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateNewGoogleCloudStorage201Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateNewGoogleCloudStorage201Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateNewGoogleCloudStorage201Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *CreateNewGoogleCloudStorage201Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateNewGoogleCloudStorage201Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateNewGoogleCloudStorage201Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateNewGoogleCloudStorage201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *CreateNewGoogleCloudStorage201Response) GetCreator() TokenVerification200ResponseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateNewGoogleCloudStorage201Response) GetCreatorOk() (*TokenVerification200ResponseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateNewGoogleCloudStorage201Response) SetCreator(v TokenVerification200ResponseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CreateNewGoogleCloudStorage201Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetGcsCredentials

`func (o *CreateNewGoogleCloudStorage201Response) GetGcsCredentials() CreateNewGoogleCloudStorage201ResponseGcsCredentials`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *CreateNewGoogleCloudStorage201Response) GetGcsCredentialsOk() (*CreateNewGoogleCloudStorage201ResponseGcsCredentials, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *CreateNewGoogleCloudStorage201Response) SetGcsCredentials(v CreateNewGoogleCloudStorage201ResponseGcsCredentials)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *CreateNewGoogleCloudStorage201Response) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetGcsSnowflakeIntegrationName

`func (o *CreateNewGoogleCloudStorage201Response) GetGcsSnowflakeIntegrationName() string`

GetGcsSnowflakeIntegrationName returns the GcsSnowflakeIntegrationName field if non-nil, zero value otherwise.

### GetGcsSnowflakeIntegrationNameOk

`func (o *CreateNewGoogleCloudStorage201Response) GetGcsSnowflakeIntegrationNameOk() (*string, bool)`

GetGcsSnowflakeIntegrationNameOk returns a tuple with the GcsSnowflakeIntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsSnowflakeIntegrationName

`func (o *CreateNewGoogleCloudStorage201Response) SetGcsSnowflakeIntegrationName(v string)`

SetGcsSnowflakeIntegrationName sets GcsSnowflakeIntegrationName field to given value.

### HasGcsSnowflakeIntegrationName

`func (o *CreateNewGoogleCloudStorage201Response) HasGcsSnowflakeIntegrationName() bool`

HasGcsSnowflakeIntegrationName returns a boolean if a field has been set.

### GetFilesBucket

`func (o *CreateNewGoogleCloudStorage201Response) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *CreateNewGoogleCloudStorage201Response) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *CreateNewGoogleCloudStorage201Response) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.

### HasFilesBucket

`func (o *CreateNewGoogleCloudStorage201Response) HasFilesBucket() bool`

HasFilesBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


