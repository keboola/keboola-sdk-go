# ListProjectsForAnOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**DataSizeBytes** | Pointer to **float32** |  | [optional] 
**RowsCount** | Pointer to **float32** |  | [optional] 
**HasMysql** | Pointer to **bool** |  | [optional] 
**HasRedshift** | Pointer to **bool** |  | [optional] 
**HasSynapse** | Pointer to **bool** |  | [optional] 
**HasExasol** | Pointer to **bool** |  | [optional] 
**HasTeradata** | Pointer to **bool** |  | [optional] 
**HasSnowflake** | Pointer to **bool** |  | [optional] 
**DefaultBackend** | Pointer to **string** |  | [optional] 
**HasTryModeOn** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**ListProjectsForAnOrganization200ResponseLimits**](ListProjectsForAnOrganization200ResponseLimits.md) |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**BilledMonthlyPrice** | Pointer to **interface{}** |  | [optional] 
**DataRetentionTimeInDays** | Pointer to **float32** |  | [optional] 
**FileStorageProvider** | Pointer to **string** |  | [optional] 

## Methods

### NewListProjectsForAnOrganization200Response

`func NewListProjectsForAnOrganization200Response() *ListProjectsForAnOrganization200Response`

NewListProjectsForAnOrganization200Response instantiates a new ListProjectsForAnOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProjectsForAnOrganization200ResponseWithDefaults

`func NewListProjectsForAnOrganization200ResponseWithDefaults() *ListProjectsForAnOrganization200Response`

NewListProjectsForAnOrganization200ResponseWithDefaults instantiates a new ListProjectsForAnOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProjectsForAnOrganization200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProjectsForAnOrganization200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProjectsForAnOrganization200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ListProjectsForAnOrganization200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListProjectsForAnOrganization200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProjectsForAnOrganization200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProjectsForAnOrganization200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListProjectsForAnOrganization200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListProjectsForAnOrganization200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListProjectsForAnOrganization200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListProjectsForAnOrganization200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListProjectsForAnOrganization200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *ListProjectsForAnOrganization200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListProjectsForAnOrganization200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListProjectsForAnOrganization200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ListProjectsForAnOrganization200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *ListProjectsForAnOrganization200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListProjectsForAnOrganization200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListProjectsForAnOrganization200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListProjectsForAnOrganization200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *ListProjectsForAnOrganization200Response) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListProjectsForAnOrganization200Response) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListProjectsForAnOrganization200Response) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ListProjectsForAnOrganization200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *ListProjectsForAnOrganization200Response) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListProjectsForAnOrganization200Response) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFeatures

`func (o *ListProjectsForAnOrganization200Response) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ListProjectsForAnOrganization200Response) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ListProjectsForAnOrganization200Response) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ListProjectsForAnOrganization200Response) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDataSizeBytes

`func (o *ListProjectsForAnOrganization200Response) GetDataSizeBytes() float32`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *ListProjectsForAnOrganization200Response) GetDataSizeBytesOk() (*float32, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *ListProjectsForAnOrganization200Response) SetDataSizeBytes(v float32)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *ListProjectsForAnOrganization200Response) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.

### GetRowsCount

`func (o *ListProjectsForAnOrganization200Response) GetRowsCount() float32`

GetRowsCount returns the RowsCount field if non-nil, zero value otherwise.

### GetRowsCountOk

`func (o *ListProjectsForAnOrganization200Response) GetRowsCountOk() (*float32, bool)`

GetRowsCountOk returns a tuple with the RowsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsCount

`func (o *ListProjectsForAnOrganization200Response) SetRowsCount(v float32)`

SetRowsCount sets RowsCount field to given value.

### HasRowsCount

`func (o *ListProjectsForAnOrganization200Response) HasRowsCount() bool`

HasRowsCount returns a boolean if a field has been set.

### GetHasMysql

`func (o *ListProjectsForAnOrganization200Response) GetHasMysql() bool`

GetHasMysql returns the HasMysql field if non-nil, zero value otherwise.

### GetHasMysqlOk

`func (o *ListProjectsForAnOrganization200Response) GetHasMysqlOk() (*bool, bool)`

GetHasMysqlOk returns a tuple with the HasMysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMysql

`func (o *ListProjectsForAnOrganization200Response) SetHasMysql(v bool)`

SetHasMysql sets HasMysql field to given value.

### HasHasMysql

`func (o *ListProjectsForAnOrganization200Response) HasHasMysql() bool`

HasHasMysql returns a boolean if a field has been set.

### GetHasRedshift

`func (o *ListProjectsForAnOrganization200Response) GetHasRedshift() bool`

GetHasRedshift returns the HasRedshift field if non-nil, zero value otherwise.

### GetHasRedshiftOk

`func (o *ListProjectsForAnOrganization200Response) GetHasRedshiftOk() (*bool, bool)`

GetHasRedshiftOk returns a tuple with the HasRedshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRedshift

`func (o *ListProjectsForAnOrganization200Response) SetHasRedshift(v bool)`

SetHasRedshift sets HasRedshift field to given value.

### HasHasRedshift

`func (o *ListProjectsForAnOrganization200Response) HasHasRedshift() bool`

HasHasRedshift returns a boolean if a field has been set.

### GetHasSynapse

`func (o *ListProjectsForAnOrganization200Response) GetHasSynapse() bool`

GetHasSynapse returns the HasSynapse field if non-nil, zero value otherwise.

### GetHasSynapseOk

`func (o *ListProjectsForAnOrganization200Response) GetHasSynapseOk() (*bool, bool)`

GetHasSynapseOk returns a tuple with the HasSynapse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSynapse

`func (o *ListProjectsForAnOrganization200Response) SetHasSynapse(v bool)`

SetHasSynapse sets HasSynapse field to given value.

### HasHasSynapse

`func (o *ListProjectsForAnOrganization200Response) HasHasSynapse() bool`

HasHasSynapse returns a boolean if a field has been set.

### GetHasExasol

`func (o *ListProjectsForAnOrganization200Response) GetHasExasol() bool`

GetHasExasol returns the HasExasol field if non-nil, zero value otherwise.

### GetHasExasolOk

`func (o *ListProjectsForAnOrganization200Response) GetHasExasolOk() (*bool, bool)`

GetHasExasolOk returns a tuple with the HasExasol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExasol

`func (o *ListProjectsForAnOrganization200Response) SetHasExasol(v bool)`

SetHasExasol sets HasExasol field to given value.

### HasHasExasol

`func (o *ListProjectsForAnOrganization200Response) HasHasExasol() bool`

HasHasExasol returns a boolean if a field has been set.

### GetHasTeradata

`func (o *ListProjectsForAnOrganization200Response) GetHasTeradata() bool`

GetHasTeradata returns the HasTeradata field if non-nil, zero value otherwise.

### GetHasTeradataOk

`func (o *ListProjectsForAnOrganization200Response) GetHasTeradataOk() (*bool, bool)`

GetHasTeradataOk returns a tuple with the HasTeradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTeradata

`func (o *ListProjectsForAnOrganization200Response) SetHasTeradata(v bool)`

SetHasTeradata sets HasTeradata field to given value.

### HasHasTeradata

`func (o *ListProjectsForAnOrganization200Response) HasHasTeradata() bool`

HasHasTeradata returns a boolean if a field has been set.

### GetHasSnowflake

`func (o *ListProjectsForAnOrganization200Response) GetHasSnowflake() bool`

GetHasSnowflake returns the HasSnowflake field if non-nil, zero value otherwise.

### GetHasSnowflakeOk

`func (o *ListProjectsForAnOrganization200Response) GetHasSnowflakeOk() (*bool, bool)`

GetHasSnowflakeOk returns a tuple with the HasSnowflake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnowflake

`func (o *ListProjectsForAnOrganization200Response) SetHasSnowflake(v bool)`

SetHasSnowflake sets HasSnowflake field to given value.

### HasHasSnowflake

`func (o *ListProjectsForAnOrganization200Response) HasHasSnowflake() bool`

HasHasSnowflake returns a boolean if a field has been set.

### GetDefaultBackend

`func (o *ListProjectsForAnOrganization200Response) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *ListProjectsForAnOrganization200Response) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *ListProjectsForAnOrganization200Response) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *ListProjectsForAnOrganization200Response) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetHasTryModeOn

`func (o *ListProjectsForAnOrganization200Response) GetHasTryModeOn() string`

GetHasTryModeOn returns the HasTryModeOn field if non-nil, zero value otherwise.

### GetHasTryModeOnOk

`func (o *ListProjectsForAnOrganization200Response) GetHasTryModeOnOk() (*string, bool)`

GetHasTryModeOnOk returns a tuple with the HasTryModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTryModeOn

`func (o *ListProjectsForAnOrganization200Response) SetHasTryModeOn(v string)`

SetHasTryModeOn sets HasTryModeOn field to given value.

### HasHasTryModeOn

`func (o *ListProjectsForAnOrganization200Response) HasHasTryModeOn() bool`

HasHasTryModeOn returns a boolean if a field has been set.

### GetLimits

`func (o *ListProjectsForAnOrganization200Response) GetLimits() ListProjectsForAnOrganization200ResponseLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ListProjectsForAnOrganization200Response) GetLimitsOk() (*ListProjectsForAnOrganization200ResponseLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ListProjectsForAnOrganization200Response) SetLimits(v ListProjectsForAnOrganization200ResponseLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ListProjectsForAnOrganization200Response) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMetrics

`func (o *ListProjectsForAnOrganization200Response) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ListProjectsForAnOrganization200Response) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ListProjectsForAnOrganization200Response) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ListProjectsForAnOrganization200Response) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ListProjectsForAnOrganization200Response) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ListProjectsForAnOrganization200Response) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ListProjectsForAnOrganization200Response) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ListProjectsForAnOrganization200Response) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetBilledMonthlyPrice

`func (o *ListProjectsForAnOrganization200Response) GetBilledMonthlyPrice() interface{}`

GetBilledMonthlyPrice returns the BilledMonthlyPrice field if non-nil, zero value otherwise.

### GetBilledMonthlyPriceOk

`func (o *ListProjectsForAnOrganization200Response) GetBilledMonthlyPriceOk() (*interface{}, bool)`

GetBilledMonthlyPriceOk returns a tuple with the BilledMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMonthlyPrice

`func (o *ListProjectsForAnOrganization200Response) SetBilledMonthlyPrice(v interface{})`

SetBilledMonthlyPrice sets BilledMonthlyPrice field to given value.

### HasBilledMonthlyPrice

`func (o *ListProjectsForAnOrganization200Response) HasBilledMonthlyPrice() bool`

HasBilledMonthlyPrice returns a boolean if a field has been set.

### SetBilledMonthlyPriceNil

`func (o *ListProjectsForAnOrganization200Response) SetBilledMonthlyPriceNil(b bool)`

 SetBilledMonthlyPriceNil sets the value for BilledMonthlyPrice to be an explicit nil

### UnsetBilledMonthlyPrice
`func (o *ListProjectsForAnOrganization200Response) UnsetBilledMonthlyPrice()`

UnsetBilledMonthlyPrice ensures that no value is present for BilledMonthlyPrice, not even an explicit nil
### GetDataRetentionTimeInDays

`func (o *ListProjectsForAnOrganization200Response) GetDataRetentionTimeInDays() float32`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *ListProjectsForAnOrganization200Response) GetDataRetentionTimeInDaysOk() (*float32, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *ListProjectsForAnOrganization200Response) SetDataRetentionTimeInDays(v float32)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.

### HasDataRetentionTimeInDays

`func (o *ListProjectsForAnOrganization200Response) HasDataRetentionTimeInDays() bool`

HasDataRetentionTimeInDays returns a boolean if a field has been set.

### GetFileStorageProvider

`func (o *ListProjectsForAnOrganization200Response) GetFileStorageProvider() string`

GetFileStorageProvider returns the FileStorageProvider field if non-nil, zero value otherwise.

### GetFileStorageProviderOk

`func (o *ListProjectsForAnOrganization200Response) GetFileStorageProviderOk() (*string, bool)`

GetFileStorageProviderOk returns a tuple with the FileStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStorageProvider

`func (o *ListProjectsForAnOrganization200Response) SetFileStorageProvider(v string)`

SetFileStorageProvider sets FileStorageProvider field to given value.

### HasFileStorageProvider

`func (o *ListProjectsForAnOrganization200Response) HasFileStorageProvider() bool`

HasFileStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


