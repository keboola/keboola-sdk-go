# DeletedProjectDetail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**Features** | Pointer to **[]interface{}** |  | [optional] 
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
**Limits** | Pointer to **map[string]interface{}** |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**BilledMonthlyPrice** | Pointer to **interface{}** |  | [optional] 
**DataRetentionTimeInDays** | Pointer to **float32** |  | [optional] 
**IsPurged** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeletedTime** | Pointer to **string** |  | [optional] 
**PurgedTime** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDeletedProjectDetail200Response

`func NewDeletedProjectDetail200Response() *DeletedProjectDetail200Response`

NewDeletedProjectDetail200Response instantiates a new DeletedProjectDetail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedProjectDetail200ResponseWithDefaults

`func NewDeletedProjectDetail200ResponseWithDefaults() *DeletedProjectDetail200Response`

NewDeletedProjectDetail200ResponseWithDefaults instantiates a new DeletedProjectDetail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeletedProjectDetail200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedProjectDetail200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedProjectDetail200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedProjectDetail200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeletedProjectDetail200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeletedProjectDetail200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeletedProjectDetail200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeletedProjectDetail200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DeletedProjectDetail200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeletedProjectDetail200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeletedProjectDetail200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeletedProjectDetail200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *DeletedProjectDetail200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeletedProjectDetail200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeletedProjectDetail200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeletedProjectDetail200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *DeletedProjectDetail200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeletedProjectDetail200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeletedProjectDetail200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeletedProjectDetail200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *DeletedProjectDetail200Response) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DeletedProjectDetail200Response) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DeletedProjectDetail200Response) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *DeletedProjectDetail200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *DeletedProjectDetail200Response) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *DeletedProjectDetail200Response) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFeatures

`func (o *DeletedProjectDetail200Response) GetFeatures() []interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DeletedProjectDetail200Response) GetFeaturesOk() (*[]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DeletedProjectDetail200Response) SetFeatures(v []interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DeletedProjectDetail200Response) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDataSizeBytes

`func (o *DeletedProjectDetail200Response) GetDataSizeBytes() float32`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *DeletedProjectDetail200Response) GetDataSizeBytesOk() (*float32, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *DeletedProjectDetail200Response) SetDataSizeBytes(v float32)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *DeletedProjectDetail200Response) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.

### GetRowsCount

`func (o *DeletedProjectDetail200Response) GetRowsCount() float32`

GetRowsCount returns the RowsCount field if non-nil, zero value otherwise.

### GetRowsCountOk

`func (o *DeletedProjectDetail200Response) GetRowsCountOk() (*float32, bool)`

GetRowsCountOk returns a tuple with the RowsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsCount

`func (o *DeletedProjectDetail200Response) SetRowsCount(v float32)`

SetRowsCount sets RowsCount field to given value.

### HasRowsCount

`func (o *DeletedProjectDetail200Response) HasRowsCount() bool`

HasRowsCount returns a boolean if a field has been set.

### GetHasMysql

`func (o *DeletedProjectDetail200Response) GetHasMysql() bool`

GetHasMysql returns the HasMysql field if non-nil, zero value otherwise.

### GetHasMysqlOk

`func (o *DeletedProjectDetail200Response) GetHasMysqlOk() (*bool, bool)`

GetHasMysqlOk returns a tuple with the HasMysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMysql

`func (o *DeletedProjectDetail200Response) SetHasMysql(v bool)`

SetHasMysql sets HasMysql field to given value.

### HasHasMysql

`func (o *DeletedProjectDetail200Response) HasHasMysql() bool`

HasHasMysql returns a boolean if a field has been set.

### GetHasRedshift

`func (o *DeletedProjectDetail200Response) GetHasRedshift() bool`

GetHasRedshift returns the HasRedshift field if non-nil, zero value otherwise.

### GetHasRedshiftOk

`func (o *DeletedProjectDetail200Response) GetHasRedshiftOk() (*bool, bool)`

GetHasRedshiftOk returns a tuple with the HasRedshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRedshift

`func (o *DeletedProjectDetail200Response) SetHasRedshift(v bool)`

SetHasRedshift sets HasRedshift field to given value.

### HasHasRedshift

`func (o *DeletedProjectDetail200Response) HasHasRedshift() bool`

HasHasRedshift returns a boolean if a field has been set.

### GetHasSynapse

`func (o *DeletedProjectDetail200Response) GetHasSynapse() bool`

GetHasSynapse returns the HasSynapse field if non-nil, zero value otherwise.

### GetHasSynapseOk

`func (o *DeletedProjectDetail200Response) GetHasSynapseOk() (*bool, bool)`

GetHasSynapseOk returns a tuple with the HasSynapse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSynapse

`func (o *DeletedProjectDetail200Response) SetHasSynapse(v bool)`

SetHasSynapse sets HasSynapse field to given value.

### HasHasSynapse

`func (o *DeletedProjectDetail200Response) HasHasSynapse() bool`

HasHasSynapse returns a boolean if a field has been set.

### GetHasExasol

`func (o *DeletedProjectDetail200Response) GetHasExasol() bool`

GetHasExasol returns the HasExasol field if non-nil, zero value otherwise.

### GetHasExasolOk

`func (o *DeletedProjectDetail200Response) GetHasExasolOk() (*bool, bool)`

GetHasExasolOk returns a tuple with the HasExasol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExasol

`func (o *DeletedProjectDetail200Response) SetHasExasol(v bool)`

SetHasExasol sets HasExasol field to given value.

### HasHasExasol

`func (o *DeletedProjectDetail200Response) HasHasExasol() bool`

HasHasExasol returns a boolean if a field has been set.

### GetHasTeradata

`func (o *DeletedProjectDetail200Response) GetHasTeradata() bool`

GetHasTeradata returns the HasTeradata field if non-nil, zero value otherwise.

### GetHasTeradataOk

`func (o *DeletedProjectDetail200Response) GetHasTeradataOk() (*bool, bool)`

GetHasTeradataOk returns a tuple with the HasTeradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTeradata

`func (o *DeletedProjectDetail200Response) SetHasTeradata(v bool)`

SetHasTeradata sets HasTeradata field to given value.

### HasHasTeradata

`func (o *DeletedProjectDetail200Response) HasHasTeradata() bool`

HasHasTeradata returns a boolean if a field has been set.

### GetHasSnowflake

`func (o *DeletedProjectDetail200Response) GetHasSnowflake() bool`

GetHasSnowflake returns the HasSnowflake field if non-nil, zero value otherwise.

### GetHasSnowflakeOk

`func (o *DeletedProjectDetail200Response) GetHasSnowflakeOk() (*bool, bool)`

GetHasSnowflakeOk returns a tuple with the HasSnowflake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnowflake

`func (o *DeletedProjectDetail200Response) SetHasSnowflake(v bool)`

SetHasSnowflake sets HasSnowflake field to given value.

### HasHasSnowflake

`func (o *DeletedProjectDetail200Response) HasHasSnowflake() bool`

HasHasSnowflake returns a boolean if a field has been set.

### GetDefaultBackend

`func (o *DeletedProjectDetail200Response) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *DeletedProjectDetail200Response) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *DeletedProjectDetail200Response) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *DeletedProjectDetail200Response) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetHasTryModeOn

`func (o *DeletedProjectDetail200Response) GetHasTryModeOn() string`

GetHasTryModeOn returns the HasTryModeOn field if non-nil, zero value otherwise.

### GetHasTryModeOnOk

`func (o *DeletedProjectDetail200Response) GetHasTryModeOnOk() (*string, bool)`

GetHasTryModeOnOk returns a tuple with the HasTryModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTryModeOn

`func (o *DeletedProjectDetail200Response) SetHasTryModeOn(v string)`

SetHasTryModeOn sets HasTryModeOn field to given value.

### HasHasTryModeOn

`func (o *DeletedProjectDetail200Response) HasHasTryModeOn() bool`

HasHasTryModeOn returns a boolean if a field has been set.

### GetLimits

`func (o *DeletedProjectDetail200Response) GetLimits() map[string]interface{}`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *DeletedProjectDetail200Response) GetLimitsOk() (*map[string]interface{}, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *DeletedProjectDetail200Response) SetLimits(v map[string]interface{})`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *DeletedProjectDetail200Response) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMetrics

`func (o *DeletedProjectDetail200Response) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DeletedProjectDetail200Response) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DeletedProjectDetail200Response) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DeletedProjectDetail200Response) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetIsDisabled

`func (o *DeletedProjectDetail200Response) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *DeletedProjectDetail200Response) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *DeletedProjectDetail200Response) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *DeletedProjectDetail200Response) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetBilledMonthlyPrice

`func (o *DeletedProjectDetail200Response) GetBilledMonthlyPrice() interface{}`

GetBilledMonthlyPrice returns the BilledMonthlyPrice field if non-nil, zero value otherwise.

### GetBilledMonthlyPriceOk

`func (o *DeletedProjectDetail200Response) GetBilledMonthlyPriceOk() (*interface{}, bool)`

GetBilledMonthlyPriceOk returns a tuple with the BilledMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMonthlyPrice

`func (o *DeletedProjectDetail200Response) SetBilledMonthlyPrice(v interface{})`

SetBilledMonthlyPrice sets BilledMonthlyPrice field to given value.

### HasBilledMonthlyPrice

`func (o *DeletedProjectDetail200Response) HasBilledMonthlyPrice() bool`

HasBilledMonthlyPrice returns a boolean if a field has been set.

### SetBilledMonthlyPriceNil

`func (o *DeletedProjectDetail200Response) SetBilledMonthlyPriceNil(b bool)`

 SetBilledMonthlyPriceNil sets the value for BilledMonthlyPrice to be an explicit nil

### UnsetBilledMonthlyPrice
`func (o *DeletedProjectDetail200Response) UnsetBilledMonthlyPrice()`

UnsetBilledMonthlyPrice ensures that no value is present for BilledMonthlyPrice, not even an explicit nil
### GetDataRetentionTimeInDays

`func (o *DeletedProjectDetail200Response) GetDataRetentionTimeInDays() float32`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *DeletedProjectDetail200Response) GetDataRetentionTimeInDaysOk() (*float32, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *DeletedProjectDetail200Response) SetDataRetentionTimeInDays(v float32)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.

### HasDataRetentionTimeInDays

`func (o *DeletedProjectDetail200Response) HasDataRetentionTimeInDays() bool`

HasDataRetentionTimeInDays returns a boolean if a field has been set.

### GetIsPurged

`func (o *DeletedProjectDetail200Response) GetIsPurged() bool`

GetIsPurged returns the IsPurged field if non-nil, zero value otherwise.

### GetIsPurgedOk

`func (o *DeletedProjectDetail200Response) GetIsPurgedOk() (*bool, bool)`

GetIsPurgedOk returns a tuple with the IsPurged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurged

`func (o *DeletedProjectDetail200Response) SetIsPurged(v bool)`

SetIsPurged sets IsPurged field to given value.

### HasIsPurged

`func (o *DeletedProjectDetail200Response) HasIsPurged() bool`

HasIsPurged returns a boolean if a field has been set.

### GetIsDeleted

`func (o *DeletedProjectDetail200Response) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DeletedProjectDetail200Response) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DeletedProjectDetail200Response) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DeletedProjectDetail200Response) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletedTime

`func (o *DeletedProjectDetail200Response) GetDeletedTime() string`

GetDeletedTime returns the DeletedTime field if non-nil, zero value otherwise.

### GetDeletedTimeOk

`func (o *DeletedProjectDetail200Response) GetDeletedTimeOk() (*string, bool)`

GetDeletedTimeOk returns a tuple with the DeletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTime

`func (o *DeletedProjectDetail200Response) SetDeletedTime(v string)`

SetDeletedTime sets DeletedTime field to given value.

### HasDeletedTime

`func (o *DeletedProjectDetail200Response) HasDeletedTime() bool`

HasDeletedTime returns a boolean if a field has been set.

### GetPurgedTime

`func (o *DeletedProjectDetail200Response) GetPurgedTime() interface{}`

GetPurgedTime returns the PurgedTime field if non-nil, zero value otherwise.

### GetPurgedTimeOk

`func (o *DeletedProjectDetail200Response) GetPurgedTimeOk() (*interface{}, bool)`

GetPurgedTimeOk returns a tuple with the PurgedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgedTime

`func (o *DeletedProjectDetail200Response) SetPurgedTime(v interface{})`

SetPurgedTime sets PurgedTime field to given value.

### HasPurgedTime

`func (o *DeletedProjectDetail200Response) HasPurgedTime() bool`

HasPurgedTime returns a boolean if a field has been set.

### SetPurgedTimeNil

`func (o *DeletedProjectDetail200Response) SetPurgedTimeNil(b bool)`

 SetPurgedTimeNil sets the value for PurgedTime to be an explicit nil

### UnsetPurgedTime
`func (o *DeletedProjectDetail200Response) UnsetPurgedTime()`

UnsetPurgedTime ensures that no value is present for PurgedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


