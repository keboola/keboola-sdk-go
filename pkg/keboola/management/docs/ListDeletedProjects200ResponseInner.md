# ListDeletedProjects200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Region** | **string** |  | 
**Created** | **string** |  | 
**Expires** | **interface{}** |  | 
**Features** | **[]interface{}** |  | 
**DataSizeBytes** | **float32** |  | 
**RowsCount** | **float32** |  | 
**HasMysql** | **bool** |  | 
**HasRedshift** | **bool** |  | 
**HasSnowflake** | **bool** |  | 
**HasSynapse** | **bool** |  | 
**HasTeradata** | **bool** |  | 
**HasExasol** | **bool** |  | 
**DefaultBackend** | **string** |  | 
**HasTryModeOn** | **string** |  | 
**Limits** | **map[string]interface{}** |  | 
**Metrics** | **map[string]interface{}** |  | 
**IsDisabled** | **bool** |  | 
**BilledMonthlyPrice** | **interface{}** |  | 
**DataRetentionTimeInDays** | **float32** |  | 
**IsPurged** | **bool** |  | 
**IsDeleted** | **bool** |  | 
**DeletedTime** | **string** |  | 
**PurgedTime** | **interface{}** |  | 
**Organization** | [**ListDeletedProjects200ResponseInnerOrganization**](ListDeletedProjects200ResponseInnerOrganization.md) |  | 

## Methods

### NewListDeletedProjects200ResponseInner

`func NewListDeletedProjects200ResponseInner(id float32, name string, type_ string, region string, created string, expires interface{}, features []interface{}, dataSizeBytes float32, rowsCount float32, hasMysql bool, hasRedshift bool, hasSnowflake bool, hasSynapse bool, hasTeradata bool, hasExasol bool, defaultBackend string, hasTryModeOn string, limits map[string]interface{}, metrics map[string]interface{}, isDisabled bool, billedMonthlyPrice interface{}, dataRetentionTimeInDays float32, isPurged bool, isDeleted bool, deletedTime string, purgedTime interface{}, organization ListDeletedProjects200ResponseInnerOrganization, ) *ListDeletedProjects200ResponseInner`

NewListDeletedProjects200ResponseInner instantiates a new ListDeletedProjects200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeletedProjects200ResponseInnerWithDefaults

`func NewListDeletedProjects200ResponseInnerWithDefaults() *ListDeletedProjects200ResponseInner`

NewListDeletedProjects200ResponseInnerWithDefaults instantiates a new ListDeletedProjects200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDeletedProjects200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDeletedProjects200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDeletedProjects200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ListDeletedProjects200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDeletedProjects200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDeletedProjects200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ListDeletedProjects200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListDeletedProjects200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListDeletedProjects200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetRegion

`func (o *ListDeletedProjects200ResponseInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListDeletedProjects200ResponseInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListDeletedProjects200ResponseInner) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCreated

`func (o *ListDeletedProjects200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListDeletedProjects200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListDeletedProjects200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *ListDeletedProjects200ResponseInner) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListDeletedProjects200ResponseInner) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListDeletedProjects200ResponseInner) SetExpires(v interface{})`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *ListDeletedProjects200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListDeletedProjects200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFeatures

`func (o *ListDeletedProjects200ResponseInner) GetFeatures() []interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ListDeletedProjects200ResponseInner) GetFeaturesOk() (*[]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ListDeletedProjects200ResponseInner) SetFeatures(v []interface{})`

SetFeatures sets Features field to given value.


### GetDataSizeBytes

`func (o *ListDeletedProjects200ResponseInner) GetDataSizeBytes() float32`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *ListDeletedProjects200ResponseInner) GetDataSizeBytesOk() (*float32, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *ListDeletedProjects200ResponseInner) SetDataSizeBytes(v float32)`

SetDataSizeBytes sets DataSizeBytes field to given value.


### GetRowsCount

`func (o *ListDeletedProjects200ResponseInner) GetRowsCount() float32`

GetRowsCount returns the RowsCount field if non-nil, zero value otherwise.

### GetRowsCountOk

`func (o *ListDeletedProjects200ResponseInner) GetRowsCountOk() (*float32, bool)`

GetRowsCountOk returns a tuple with the RowsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsCount

`func (o *ListDeletedProjects200ResponseInner) SetRowsCount(v float32)`

SetRowsCount sets RowsCount field to given value.


### GetHasMysql

`func (o *ListDeletedProjects200ResponseInner) GetHasMysql() bool`

GetHasMysql returns the HasMysql field if non-nil, zero value otherwise.

### GetHasMysqlOk

`func (o *ListDeletedProjects200ResponseInner) GetHasMysqlOk() (*bool, bool)`

GetHasMysqlOk returns a tuple with the HasMysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMysql

`func (o *ListDeletedProjects200ResponseInner) SetHasMysql(v bool)`

SetHasMysql sets HasMysql field to given value.


### GetHasRedshift

`func (o *ListDeletedProjects200ResponseInner) GetHasRedshift() bool`

GetHasRedshift returns the HasRedshift field if non-nil, zero value otherwise.

### GetHasRedshiftOk

`func (o *ListDeletedProjects200ResponseInner) GetHasRedshiftOk() (*bool, bool)`

GetHasRedshiftOk returns a tuple with the HasRedshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRedshift

`func (o *ListDeletedProjects200ResponseInner) SetHasRedshift(v bool)`

SetHasRedshift sets HasRedshift field to given value.


### GetHasSnowflake

`func (o *ListDeletedProjects200ResponseInner) GetHasSnowflake() bool`

GetHasSnowflake returns the HasSnowflake field if non-nil, zero value otherwise.

### GetHasSnowflakeOk

`func (o *ListDeletedProjects200ResponseInner) GetHasSnowflakeOk() (*bool, bool)`

GetHasSnowflakeOk returns a tuple with the HasSnowflake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnowflake

`func (o *ListDeletedProjects200ResponseInner) SetHasSnowflake(v bool)`

SetHasSnowflake sets HasSnowflake field to given value.


### GetHasSynapse

`func (o *ListDeletedProjects200ResponseInner) GetHasSynapse() bool`

GetHasSynapse returns the HasSynapse field if non-nil, zero value otherwise.

### GetHasSynapseOk

`func (o *ListDeletedProjects200ResponseInner) GetHasSynapseOk() (*bool, bool)`

GetHasSynapseOk returns a tuple with the HasSynapse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSynapse

`func (o *ListDeletedProjects200ResponseInner) SetHasSynapse(v bool)`

SetHasSynapse sets HasSynapse field to given value.


### GetHasTeradata

`func (o *ListDeletedProjects200ResponseInner) GetHasTeradata() bool`

GetHasTeradata returns the HasTeradata field if non-nil, zero value otherwise.

### GetHasTeradataOk

`func (o *ListDeletedProjects200ResponseInner) GetHasTeradataOk() (*bool, bool)`

GetHasTeradataOk returns a tuple with the HasTeradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTeradata

`func (o *ListDeletedProjects200ResponseInner) SetHasTeradata(v bool)`

SetHasTeradata sets HasTeradata field to given value.


### GetHasExasol

`func (o *ListDeletedProjects200ResponseInner) GetHasExasol() bool`

GetHasExasol returns the HasExasol field if non-nil, zero value otherwise.

### GetHasExasolOk

`func (o *ListDeletedProjects200ResponseInner) GetHasExasolOk() (*bool, bool)`

GetHasExasolOk returns a tuple with the HasExasol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExasol

`func (o *ListDeletedProjects200ResponseInner) SetHasExasol(v bool)`

SetHasExasol sets HasExasol field to given value.


### GetDefaultBackend

`func (o *ListDeletedProjects200ResponseInner) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *ListDeletedProjects200ResponseInner) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *ListDeletedProjects200ResponseInner) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.


### GetHasTryModeOn

`func (o *ListDeletedProjects200ResponseInner) GetHasTryModeOn() string`

GetHasTryModeOn returns the HasTryModeOn field if non-nil, zero value otherwise.

### GetHasTryModeOnOk

`func (o *ListDeletedProjects200ResponseInner) GetHasTryModeOnOk() (*string, bool)`

GetHasTryModeOnOk returns a tuple with the HasTryModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTryModeOn

`func (o *ListDeletedProjects200ResponseInner) SetHasTryModeOn(v string)`

SetHasTryModeOn sets HasTryModeOn field to given value.


### GetLimits

`func (o *ListDeletedProjects200ResponseInner) GetLimits() map[string]interface{}`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ListDeletedProjects200ResponseInner) GetLimitsOk() (*map[string]interface{}, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ListDeletedProjects200ResponseInner) SetLimits(v map[string]interface{})`

SetLimits sets Limits field to given value.


### GetMetrics

`func (o *ListDeletedProjects200ResponseInner) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ListDeletedProjects200ResponseInner) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ListDeletedProjects200ResponseInner) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.


### GetIsDisabled

`func (o *ListDeletedProjects200ResponseInner) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ListDeletedProjects200ResponseInner) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ListDeletedProjects200ResponseInner) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetBilledMonthlyPrice

`func (o *ListDeletedProjects200ResponseInner) GetBilledMonthlyPrice() interface{}`

GetBilledMonthlyPrice returns the BilledMonthlyPrice field if non-nil, zero value otherwise.

### GetBilledMonthlyPriceOk

`func (o *ListDeletedProjects200ResponseInner) GetBilledMonthlyPriceOk() (*interface{}, bool)`

GetBilledMonthlyPriceOk returns a tuple with the BilledMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMonthlyPrice

`func (o *ListDeletedProjects200ResponseInner) SetBilledMonthlyPrice(v interface{})`

SetBilledMonthlyPrice sets BilledMonthlyPrice field to given value.


### SetBilledMonthlyPriceNil

`func (o *ListDeletedProjects200ResponseInner) SetBilledMonthlyPriceNil(b bool)`

 SetBilledMonthlyPriceNil sets the value for BilledMonthlyPrice to be an explicit nil

### UnsetBilledMonthlyPrice
`func (o *ListDeletedProjects200ResponseInner) UnsetBilledMonthlyPrice()`

UnsetBilledMonthlyPrice ensures that no value is present for BilledMonthlyPrice, not even an explicit nil
### GetDataRetentionTimeInDays

`func (o *ListDeletedProjects200ResponseInner) GetDataRetentionTimeInDays() float32`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *ListDeletedProjects200ResponseInner) GetDataRetentionTimeInDaysOk() (*float32, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *ListDeletedProjects200ResponseInner) SetDataRetentionTimeInDays(v float32)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.


### GetIsPurged

`func (o *ListDeletedProjects200ResponseInner) GetIsPurged() bool`

GetIsPurged returns the IsPurged field if non-nil, zero value otherwise.

### GetIsPurgedOk

`func (o *ListDeletedProjects200ResponseInner) GetIsPurgedOk() (*bool, bool)`

GetIsPurgedOk returns a tuple with the IsPurged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurged

`func (o *ListDeletedProjects200ResponseInner) SetIsPurged(v bool)`

SetIsPurged sets IsPurged field to given value.


### GetIsDeleted

`func (o *ListDeletedProjects200ResponseInner) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ListDeletedProjects200ResponseInner) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ListDeletedProjects200ResponseInner) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetDeletedTime

`func (o *ListDeletedProjects200ResponseInner) GetDeletedTime() string`

GetDeletedTime returns the DeletedTime field if non-nil, zero value otherwise.

### GetDeletedTimeOk

`func (o *ListDeletedProjects200ResponseInner) GetDeletedTimeOk() (*string, bool)`

GetDeletedTimeOk returns a tuple with the DeletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTime

`func (o *ListDeletedProjects200ResponseInner) SetDeletedTime(v string)`

SetDeletedTime sets DeletedTime field to given value.


### GetPurgedTime

`func (o *ListDeletedProjects200ResponseInner) GetPurgedTime() interface{}`

GetPurgedTime returns the PurgedTime field if non-nil, zero value otherwise.

### GetPurgedTimeOk

`func (o *ListDeletedProjects200ResponseInner) GetPurgedTimeOk() (*interface{}, bool)`

GetPurgedTimeOk returns a tuple with the PurgedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgedTime

`func (o *ListDeletedProjects200ResponseInner) SetPurgedTime(v interface{})`

SetPurgedTime sets PurgedTime field to given value.


### SetPurgedTimeNil

`func (o *ListDeletedProjects200ResponseInner) SetPurgedTimeNil(b bool)`

 SetPurgedTimeNil sets the value for PurgedTime to be an explicit nil

### UnsetPurgedTime
`func (o *ListDeletedProjects200ResponseInner) UnsetPurgedTime()`

UnsetPurgedTime ensures that no value is present for PurgedTime, not even an explicit nil
### GetOrganization

`func (o *ListDeletedProjects200ResponseInner) GetOrganization() ListDeletedProjects200ResponseInnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ListDeletedProjects200ResponseInner) GetOrganizationOk() (*ListDeletedProjects200ResponseInnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ListDeletedProjects200ResponseInner) SetOrganization(v ListDeletedProjects200ResponseInnerOrganization)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


