# CreateAMaintainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Maintainer name | 
**DefaultConnectionRedshiftId** | Pointer to **string** | Default Redshift Connection ID | [optional] 
**DefaultConnectionSnowflakeId** | Pointer to **string** | Default Snowflake Connection ID | [optional] 
**DefaultConnectionSynapseId** | Pointer to **string** | Default Synapse Connection ID | [optional] 
**DefaultConnectionExasolId** | Pointer to **string** | Default Exasol Connection ID | [optional] 
**DefaultConnectionTeradataId** | Pointer to **string** | Default Teradata Connection ID | [optional] 
**DefaultFileStorageId** | Pointer to **string** | Default File Storage ID | [optional] 
**ZendeskUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateAMaintainerRequest

`func NewCreateAMaintainerRequest(name string, ) *CreateAMaintainerRequest`

NewCreateAMaintainerRequest instantiates a new CreateAMaintainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAMaintainerRequestWithDefaults

`func NewCreateAMaintainerRequestWithDefaults() *CreateAMaintainerRequest`

NewCreateAMaintainerRequestWithDefaults instantiates a new CreateAMaintainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAMaintainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAMaintainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAMaintainerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultConnectionRedshiftId

`func (o *CreateAMaintainerRequest) GetDefaultConnectionRedshiftId() string`

GetDefaultConnectionRedshiftId returns the DefaultConnectionRedshiftId field if non-nil, zero value otherwise.

### GetDefaultConnectionRedshiftIdOk

`func (o *CreateAMaintainerRequest) GetDefaultConnectionRedshiftIdOk() (*string, bool)`

GetDefaultConnectionRedshiftIdOk returns a tuple with the DefaultConnectionRedshiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionRedshiftId

`func (o *CreateAMaintainerRequest) SetDefaultConnectionRedshiftId(v string)`

SetDefaultConnectionRedshiftId sets DefaultConnectionRedshiftId field to given value.

### HasDefaultConnectionRedshiftId

`func (o *CreateAMaintainerRequest) HasDefaultConnectionRedshiftId() bool`

HasDefaultConnectionRedshiftId returns a boolean if a field has been set.

### GetDefaultConnectionSnowflakeId

`func (o *CreateAMaintainerRequest) GetDefaultConnectionSnowflakeId() string`

GetDefaultConnectionSnowflakeId returns the DefaultConnectionSnowflakeId field if non-nil, zero value otherwise.

### GetDefaultConnectionSnowflakeIdOk

`func (o *CreateAMaintainerRequest) GetDefaultConnectionSnowflakeIdOk() (*string, bool)`

GetDefaultConnectionSnowflakeIdOk returns a tuple with the DefaultConnectionSnowflakeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionSnowflakeId

`func (o *CreateAMaintainerRequest) SetDefaultConnectionSnowflakeId(v string)`

SetDefaultConnectionSnowflakeId sets DefaultConnectionSnowflakeId field to given value.

### HasDefaultConnectionSnowflakeId

`func (o *CreateAMaintainerRequest) HasDefaultConnectionSnowflakeId() bool`

HasDefaultConnectionSnowflakeId returns a boolean if a field has been set.

### GetDefaultConnectionSynapseId

`func (o *CreateAMaintainerRequest) GetDefaultConnectionSynapseId() string`

GetDefaultConnectionSynapseId returns the DefaultConnectionSynapseId field if non-nil, zero value otherwise.

### GetDefaultConnectionSynapseIdOk

`func (o *CreateAMaintainerRequest) GetDefaultConnectionSynapseIdOk() (*string, bool)`

GetDefaultConnectionSynapseIdOk returns a tuple with the DefaultConnectionSynapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionSynapseId

`func (o *CreateAMaintainerRequest) SetDefaultConnectionSynapseId(v string)`

SetDefaultConnectionSynapseId sets DefaultConnectionSynapseId field to given value.

### HasDefaultConnectionSynapseId

`func (o *CreateAMaintainerRequest) HasDefaultConnectionSynapseId() bool`

HasDefaultConnectionSynapseId returns a boolean if a field has been set.

### GetDefaultConnectionExasolId

`func (o *CreateAMaintainerRequest) GetDefaultConnectionExasolId() string`

GetDefaultConnectionExasolId returns the DefaultConnectionExasolId field if non-nil, zero value otherwise.

### GetDefaultConnectionExasolIdOk

`func (o *CreateAMaintainerRequest) GetDefaultConnectionExasolIdOk() (*string, bool)`

GetDefaultConnectionExasolIdOk returns a tuple with the DefaultConnectionExasolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionExasolId

`func (o *CreateAMaintainerRequest) SetDefaultConnectionExasolId(v string)`

SetDefaultConnectionExasolId sets DefaultConnectionExasolId field to given value.

### HasDefaultConnectionExasolId

`func (o *CreateAMaintainerRequest) HasDefaultConnectionExasolId() bool`

HasDefaultConnectionExasolId returns a boolean if a field has been set.

### GetDefaultConnectionTeradataId

`func (o *CreateAMaintainerRequest) GetDefaultConnectionTeradataId() string`

GetDefaultConnectionTeradataId returns the DefaultConnectionTeradataId field if non-nil, zero value otherwise.

### GetDefaultConnectionTeradataIdOk

`func (o *CreateAMaintainerRequest) GetDefaultConnectionTeradataIdOk() (*string, bool)`

GetDefaultConnectionTeradataIdOk returns a tuple with the DefaultConnectionTeradataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionTeradataId

`func (o *CreateAMaintainerRequest) SetDefaultConnectionTeradataId(v string)`

SetDefaultConnectionTeradataId sets DefaultConnectionTeradataId field to given value.

### HasDefaultConnectionTeradataId

`func (o *CreateAMaintainerRequest) HasDefaultConnectionTeradataId() bool`

HasDefaultConnectionTeradataId returns a boolean if a field has been set.

### GetDefaultFileStorageId

`func (o *CreateAMaintainerRequest) GetDefaultFileStorageId() string`

GetDefaultFileStorageId returns the DefaultFileStorageId field if non-nil, zero value otherwise.

### GetDefaultFileStorageIdOk

`func (o *CreateAMaintainerRequest) GetDefaultFileStorageIdOk() (*string, bool)`

GetDefaultFileStorageIdOk returns a tuple with the DefaultFileStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFileStorageId

`func (o *CreateAMaintainerRequest) SetDefaultFileStorageId(v string)`

SetDefaultFileStorageId sets DefaultFileStorageId field to given value.

### HasDefaultFileStorageId

`func (o *CreateAMaintainerRequest) HasDefaultFileStorageId() bool`

HasDefaultFileStorageId returns a boolean if a field has been set.

### GetZendeskUrl

`func (o *CreateAMaintainerRequest) GetZendeskUrl() string`

GetZendeskUrl returns the ZendeskUrl field if non-nil, zero value otherwise.

### GetZendeskUrlOk

`func (o *CreateAMaintainerRequest) GetZendeskUrlOk() (*string, bool)`

GetZendeskUrlOk returns a tuple with the ZendeskUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskUrl

`func (o *CreateAMaintainerRequest) SetZendeskUrl(v string)`

SetZendeskUrl sets ZendeskUrl field to given value.

### HasZendeskUrl

`func (o *CreateAMaintainerRequest) HasZendeskUrl() bool`

HasZendeskUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


