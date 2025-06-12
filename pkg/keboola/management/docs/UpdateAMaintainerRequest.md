# UpdateAMaintainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Maintainer name | [optional] 
**DefaultConnectionRedshiftId** | Pointer to **string** | Default Redshift Connection ID | [optional] 
**DefaultConnectionSnowflakeId** | Pointer to **string** | Default Snowflake Connection ID | [optional] 
**DefaultConnectionSynapseId** | Pointer to **string** | Default Synapse Connection ID | [optional] 
**DefaultConnectionExasolId** | Pointer to **string** | Default Exasol Connection ID | [optional] 
**DefaultConnectionTeradataId** | Pointer to **string** | Default Teradata Connection ID | [optional] 
**DefaultFileStorageId** | Pointer to **string** | Default File Storage ID | [optional] 
**ZendeskUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAMaintainerRequest

`func NewUpdateAMaintainerRequest() *UpdateAMaintainerRequest`

NewUpdateAMaintainerRequest instantiates a new UpdateAMaintainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAMaintainerRequestWithDefaults

`func NewUpdateAMaintainerRequestWithDefaults() *UpdateAMaintainerRequest`

NewUpdateAMaintainerRequestWithDefaults instantiates a new UpdateAMaintainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAMaintainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAMaintainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAMaintainerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAMaintainerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultConnectionRedshiftId

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionRedshiftId() string`

GetDefaultConnectionRedshiftId returns the DefaultConnectionRedshiftId field if non-nil, zero value otherwise.

### GetDefaultConnectionRedshiftIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionRedshiftIdOk() (*string, bool)`

GetDefaultConnectionRedshiftIdOk returns a tuple with the DefaultConnectionRedshiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionRedshiftId

`func (o *UpdateAMaintainerRequest) SetDefaultConnectionRedshiftId(v string)`

SetDefaultConnectionRedshiftId sets DefaultConnectionRedshiftId field to given value.

### HasDefaultConnectionRedshiftId

`func (o *UpdateAMaintainerRequest) HasDefaultConnectionRedshiftId() bool`

HasDefaultConnectionRedshiftId returns a boolean if a field has been set.

### GetDefaultConnectionSnowflakeId

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionSnowflakeId() string`

GetDefaultConnectionSnowflakeId returns the DefaultConnectionSnowflakeId field if non-nil, zero value otherwise.

### GetDefaultConnectionSnowflakeIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionSnowflakeIdOk() (*string, bool)`

GetDefaultConnectionSnowflakeIdOk returns a tuple with the DefaultConnectionSnowflakeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionSnowflakeId

`func (o *UpdateAMaintainerRequest) SetDefaultConnectionSnowflakeId(v string)`

SetDefaultConnectionSnowflakeId sets DefaultConnectionSnowflakeId field to given value.

### HasDefaultConnectionSnowflakeId

`func (o *UpdateAMaintainerRequest) HasDefaultConnectionSnowflakeId() bool`

HasDefaultConnectionSnowflakeId returns a boolean if a field has been set.

### GetDefaultConnectionSynapseId

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionSynapseId() string`

GetDefaultConnectionSynapseId returns the DefaultConnectionSynapseId field if non-nil, zero value otherwise.

### GetDefaultConnectionSynapseIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionSynapseIdOk() (*string, bool)`

GetDefaultConnectionSynapseIdOk returns a tuple with the DefaultConnectionSynapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionSynapseId

`func (o *UpdateAMaintainerRequest) SetDefaultConnectionSynapseId(v string)`

SetDefaultConnectionSynapseId sets DefaultConnectionSynapseId field to given value.

### HasDefaultConnectionSynapseId

`func (o *UpdateAMaintainerRequest) HasDefaultConnectionSynapseId() bool`

HasDefaultConnectionSynapseId returns a boolean if a field has been set.

### GetDefaultConnectionExasolId

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionExasolId() string`

GetDefaultConnectionExasolId returns the DefaultConnectionExasolId field if non-nil, zero value otherwise.

### GetDefaultConnectionExasolIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionExasolIdOk() (*string, bool)`

GetDefaultConnectionExasolIdOk returns a tuple with the DefaultConnectionExasolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionExasolId

`func (o *UpdateAMaintainerRequest) SetDefaultConnectionExasolId(v string)`

SetDefaultConnectionExasolId sets DefaultConnectionExasolId field to given value.

### HasDefaultConnectionExasolId

`func (o *UpdateAMaintainerRequest) HasDefaultConnectionExasolId() bool`

HasDefaultConnectionExasolId returns a boolean if a field has been set.

### GetDefaultConnectionTeradataId

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionTeradataId() string`

GetDefaultConnectionTeradataId returns the DefaultConnectionTeradataId field if non-nil, zero value otherwise.

### GetDefaultConnectionTeradataIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultConnectionTeradataIdOk() (*string, bool)`

GetDefaultConnectionTeradataIdOk returns a tuple with the DefaultConnectionTeradataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionTeradataId

`func (o *UpdateAMaintainerRequest) SetDefaultConnectionTeradataId(v string)`

SetDefaultConnectionTeradataId sets DefaultConnectionTeradataId field to given value.

### HasDefaultConnectionTeradataId

`func (o *UpdateAMaintainerRequest) HasDefaultConnectionTeradataId() bool`

HasDefaultConnectionTeradataId returns a boolean if a field has been set.

### GetDefaultFileStorageId

`func (o *UpdateAMaintainerRequest) GetDefaultFileStorageId() string`

GetDefaultFileStorageId returns the DefaultFileStorageId field if non-nil, zero value otherwise.

### GetDefaultFileStorageIdOk

`func (o *UpdateAMaintainerRequest) GetDefaultFileStorageIdOk() (*string, bool)`

GetDefaultFileStorageIdOk returns a tuple with the DefaultFileStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFileStorageId

`func (o *UpdateAMaintainerRequest) SetDefaultFileStorageId(v string)`

SetDefaultFileStorageId sets DefaultFileStorageId field to given value.

### HasDefaultFileStorageId

`func (o *UpdateAMaintainerRequest) HasDefaultFileStorageId() bool`

HasDefaultFileStorageId returns a boolean if a field has been set.

### GetZendeskUrl

`func (o *UpdateAMaintainerRequest) GetZendeskUrl() string`

GetZendeskUrl returns the ZendeskUrl field if non-nil, zero value otherwise.

### GetZendeskUrlOk

`func (o *UpdateAMaintainerRequest) GetZendeskUrlOk() (*string, bool)`

GetZendeskUrlOk returns a tuple with the ZendeskUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskUrl

`func (o *UpdateAMaintainerRequest) SetZendeskUrl(v string)`

SetZendeskUrl sets ZendeskUrl field to given value.

### HasZendeskUrl

`func (o *UpdateAMaintainerRequest) HasZendeskUrl() bool`

HasZendeskUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


