# SetMaintainerMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Metadata** | **[]interface{}** | Array of metadata objects | 

## Methods

### NewSetMaintainerMetadataRequest

`func NewSetMaintainerMetadataRequest(provider string, metadata []interface{}, ) *SetMaintainerMetadataRequest`

NewSetMaintainerMetadataRequest instantiates a new SetMaintainerMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMaintainerMetadataRequestWithDefaults

`func NewSetMaintainerMetadataRequestWithDefaults() *SetMaintainerMetadataRequest`

NewSetMaintainerMetadataRequestWithDefaults instantiates a new SetMaintainerMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *SetMaintainerMetadataRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SetMaintainerMetadataRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SetMaintainerMetadataRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMetadata

`func (o *SetMaintainerMetadataRequest) GetMetadata() []interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SetMaintainerMetadataRequest) GetMetadataOk() (*[]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SetMaintainerMetadataRequest) SetMetadata(v []interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


