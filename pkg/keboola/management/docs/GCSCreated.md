# GCSCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | (ISO8601 format) &#x60;Y-m-d\\TH:i:sO&#x60; | [optional] 
**Creator** | Pointer to [**AdminIdAndName**](AdminIdAndName.md) |  | [optional] 
**GcsCredentials** | Pointer to [**GCPCredentialsNoPK**](GCPCredentialsNoPK.md) |  | [optional] 
**GcsSnowflakeIntegrationName** | Pointer to **string** |  | [optional] 
**FilesBucket** | Pointer to **string** | name of the bucket | [optional] 

## Methods

### NewGCSCreated

`func NewGCSCreated() *GCSCreated`

NewGCSCreated instantiates a new GCSCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSCreatedWithDefaults

`func NewGCSCreatedWithDefaults() *GCSCreated`

NewGCSCreatedWithDefaults instantiates a new GCSCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GCSCreated) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCSCreated) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCSCreated) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GCSCreated) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *GCSCreated) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GCSCreated) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GCSCreated) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GCSCreated) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *GCSCreated) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GCSCreated) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GCSCreated) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GCSCreated) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *GCSCreated) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GCSCreated) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GCSCreated) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GCSCreated) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *GCSCreated) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GCSCreated) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GCSCreated) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GCSCreated) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *GCSCreated) GetCreator() AdminIdAndName`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GCSCreated) GetCreatorOk() (*AdminIdAndName, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GCSCreated) SetCreator(v AdminIdAndName)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GCSCreated) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetGcsCredentials

`func (o *GCSCreated) GetGcsCredentials() GCPCredentialsNoPK`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *GCSCreated) GetGcsCredentialsOk() (*GCPCredentialsNoPK, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *GCSCreated) SetGcsCredentials(v GCPCredentialsNoPK)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *GCSCreated) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetGcsSnowflakeIntegrationName

`func (o *GCSCreated) GetGcsSnowflakeIntegrationName() string`

GetGcsSnowflakeIntegrationName returns the GcsSnowflakeIntegrationName field if non-nil, zero value otherwise.

### GetGcsSnowflakeIntegrationNameOk

`func (o *GCSCreated) GetGcsSnowflakeIntegrationNameOk() (*string, bool)`

GetGcsSnowflakeIntegrationNameOk returns a tuple with the GcsSnowflakeIntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsSnowflakeIntegrationName

`func (o *GCSCreated) SetGcsSnowflakeIntegrationName(v string)`

SetGcsSnowflakeIntegrationName sets GcsSnowflakeIntegrationName field to given value.

### HasGcsSnowflakeIntegrationName

`func (o *GCSCreated) HasGcsSnowflakeIntegrationName() bool`

HasGcsSnowflakeIntegrationName returns a boolean if a field has been set.

### GetFilesBucket

`func (o *GCSCreated) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *GCSCreated) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *GCSCreated) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.

### HasFilesBucket

`func (o *GCSCreated) HasFilesBucket() bool`

HasFilesBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


