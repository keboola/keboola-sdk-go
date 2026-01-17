# ProjectModel

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
**HasSnowflake** | Pointer to **bool** |  | [optional] 
**HasSynapse** | Pointer to **bool** |  | [optional] 
**HasExasol** | Pointer to **bool** |  | [optional] 
**HasTeradata** | Pointer to **bool** |  | [optional] 
**DefaultBackend** | Pointer to **string** |  | [optional] 
**HasTryModeOn** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**ProjectModelLimits**](ProjectModelLimits.md) |  | [optional] 
**Metrics** | Pointer to [**ProjectModelMetrics**](ProjectModelMetrics.md) |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**BilledMonthlyPrice** | Pointer to **interface{}** |  | [optional] 
**DataRetentionTimeInDays** | Pointer to **float32** |  | [optional] 
**Organization** | Pointer to [**ProjectModelOrganization**](ProjectModelOrganization.md) |  | [optional] 
**FileStorage** | Pointer to [**ProjectModelFileStorage**](ProjectModelFileStorage.md) |  | [optional] 
**Backends** | Pointer to [**ProjectModelBackends**](ProjectModelBackends.md) |  | [optional] 
**PayAsYouGo** | Pointer to [**ProjectModelPayAsYouGo**](ProjectModelPayAsYouGo.md) |  | [optional] 

## Methods

### NewProjectModel

`func NewProjectModel() *ProjectModel`

NewProjectModel instantiates a new ProjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectModelWithDefaults

`func NewProjectModelWithDefaults() *ProjectModel`

NewProjectModelWithDefaults instantiates a new ProjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectModel) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectModel) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectModel) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ProjectModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *ProjectModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProjectModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProjectModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProjectModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *ProjectModel) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProjectModel) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProjectModel) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ProjectModel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *ProjectModel) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ProjectModel) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ProjectModel) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ProjectModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *ProjectModel) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ProjectModel) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFeatures

`func (o *ProjectModel) GetFeatures() []interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProjectModel) GetFeaturesOk() (*[]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProjectModel) SetFeatures(v []interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ProjectModel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDataSizeBytes

`func (o *ProjectModel) GetDataSizeBytes() float32`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *ProjectModel) GetDataSizeBytesOk() (*float32, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *ProjectModel) SetDataSizeBytes(v float32)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *ProjectModel) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.

### GetRowsCount

`func (o *ProjectModel) GetRowsCount() float32`

GetRowsCount returns the RowsCount field if non-nil, zero value otherwise.

### GetRowsCountOk

`func (o *ProjectModel) GetRowsCountOk() (*float32, bool)`

GetRowsCountOk returns a tuple with the RowsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsCount

`func (o *ProjectModel) SetRowsCount(v float32)`

SetRowsCount sets RowsCount field to given value.

### HasRowsCount

`func (o *ProjectModel) HasRowsCount() bool`

HasRowsCount returns a boolean if a field has been set.

### GetHasMysql

`func (o *ProjectModel) GetHasMysql() bool`

GetHasMysql returns the HasMysql field if non-nil, zero value otherwise.

### GetHasMysqlOk

`func (o *ProjectModel) GetHasMysqlOk() (*bool, bool)`

GetHasMysqlOk returns a tuple with the HasMysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMysql

`func (o *ProjectModel) SetHasMysql(v bool)`

SetHasMysql sets HasMysql field to given value.

### HasHasMysql

`func (o *ProjectModel) HasHasMysql() bool`

HasHasMysql returns a boolean if a field has been set.

### GetHasRedshift

`func (o *ProjectModel) GetHasRedshift() bool`

GetHasRedshift returns the HasRedshift field if non-nil, zero value otherwise.

### GetHasRedshiftOk

`func (o *ProjectModel) GetHasRedshiftOk() (*bool, bool)`

GetHasRedshiftOk returns a tuple with the HasRedshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRedshift

`func (o *ProjectModel) SetHasRedshift(v bool)`

SetHasRedshift sets HasRedshift field to given value.

### HasHasRedshift

`func (o *ProjectModel) HasHasRedshift() bool`

HasHasRedshift returns a boolean if a field has been set.

### GetHasSnowflake

`func (o *ProjectModel) GetHasSnowflake() bool`

GetHasSnowflake returns the HasSnowflake field if non-nil, zero value otherwise.

### GetHasSnowflakeOk

`func (o *ProjectModel) GetHasSnowflakeOk() (*bool, bool)`

GetHasSnowflakeOk returns a tuple with the HasSnowflake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnowflake

`func (o *ProjectModel) SetHasSnowflake(v bool)`

SetHasSnowflake sets HasSnowflake field to given value.

### HasHasSnowflake

`func (o *ProjectModel) HasHasSnowflake() bool`

HasHasSnowflake returns a boolean if a field has been set.

### GetHasSynapse

`func (o *ProjectModel) GetHasSynapse() bool`

GetHasSynapse returns the HasSynapse field if non-nil, zero value otherwise.

### GetHasSynapseOk

`func (o *ProjectModel) GetHasSynapseOk() (*bool, bool)`

GetHasSynapseOk returns a tuple with the HasSynapse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSynapse

`func (o *ProjectModel) SetHasSynapse(v bool)`

SetHasSynapse sets HasSynapse field to given value.

### HasHasSynapse

`func (o *ProjectModel) HasHasSynapse() bool`

HasHasSynapse returns a boolean if a field has been set.

### GetHasExasol

`func (o *ProjectModel) GetHasExasol() bool`

GetHasExasol returns the HasExasol field if non-nil, zero value otherwise.

### GetHasExasolOk

`func (o *ProjectModel) GetHasExasolOk() (*bool, bool)`

GetHasExasolOk returns a tuple with the HasExasol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExasol

`func (o *ProjectModel) SetHasExasol(v bool)`

SetHasExasol sets HasExasol field to given value.

### HasHasExasol

`func (o *ProjectModel) HasHasExasol() bool`

HasHasExasol returns a boolean if a field has been set.

### GetHasTeradata

`func (o *ProjectModel) GetHasTeradata() bool`

GetHasTeradata returns the HasTeradata field if non-nil, zero value otherwise.

### GetHasTeradataOk

`func (o *ProjectModel) GetHasTeradataOk() (*bool, bool)`

GetHasTeradataOk returns a tuple with the HasTeradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTeradata

`func (o *ProjectModel) SetHasTeradata(v bool)`

SetHasTeradata sets HasTeradata field to given value.

### HasHasTeradata

`func (o *ProjectModel) HasHasTeradata() bool`

HasHasTeradata returns a boolean if a field has been set.

### GetDefaultBackend

`func (o *ProjectModel) GetDefaultBackend() string`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *ProjectModel) GetDefaultBackendOk() (*string, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *ProjectModel) SetDefaultBackend(v string)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *ProjectModel) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetHasTryModeOn

`func (o *ProjectModel) GetHasTryModeOn() string`

GetHasTryModeOn returns the HasTryModeOn field if non-nil, zero value otherwise.

### GetHasTryModeOnOk

`func (o *ProjectModel) GetHasTryModeOnOk() (*string, bool)`

GetHasTryModeOnOk returns a tuple with the HasTryModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTryModeOn

`func (o *ProjectModel) SetHasTryModeOn(v string)`

SetHasTryModeOn sets HasTryModeOn field to given value.

### HasHasTryModeOn

`func (o *ProjectModel) HasHasTryModeOn() bool`

HasHasTryModeOn returns a boolean if a field has been set.

### GetLimits

`func (o *ProjectModel) GetLimits() ProjectModelLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ProjectModel) GetLimitsOk() (*ProjectModelLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ProjectModel) SetLimits(v ProjectModelLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ProjectModel) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMetrics

`func (o *ProjectModel) GetMetrics() ProjectModelMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ProjectModel) GetMetricsOk() (*ProjectModelMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ProjectModel) SetMetrics(v ProjectModelMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ProjectModel) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ProjectModel) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ProjectModel) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ProjectModel) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ProjectModel) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetBilledMonthlyPrice

`func (o *ProjectModel) GetBilledMonthlyPrice() interface{}`

GetBilledMonthlyPrice returns the BilledMonthlyPrice field if non-nil, zero value otherwise.

### GetBilledMonthlyPriceOk

`func (o *ProjectModel) GetBilledMonthlyPriceOk() (*interface{}, bool)`

GetBilledMonthlyPriceOk returns a tuple with the BilledMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMonthlyPrice

`func (o *ProjectModel) SetBilledMonthlyPrice(v interface{})`

SetBilledMonthlyPrice sets BilledMonthlyPrice field to given value.

### HasBilledMonthlyPrice

`func (o *ProjectModel) HasBilledMonthlyPrice() bool`

HasBilledMonthlyPrice returns a boolean if a field has been set.

### SetBilledMonthlyPriceNil

`func (o *ProjectModel) SetBilledMonthlyPriceNil(b bool)`

 SetBilledMonthlyPriceNil sets the value for BilledMonthlyPrice to be an explicit nil

### UnsetBilledMonthlyPrice
`func (o *ProjectModel) UnsetBilledMonthlyPrice()`

UnsetBilledMonthlyPrice ensures that no value is present for BilledMonthlyPrice, not even an explicit nil
### GetDataRetentionTimeInDays

`func (o *ProjectModel) GetDataRetentionTimeInDays() float32`

GetDataRetentionTimeInDays returns the DataRetentionTimeInDays field if non-nil, zero value otherwise.

### GetDataRetentionTimeInDaysOk

`func (o *ProjectModel) GetDataRetentionTimeInDaysOk() (*float32, bool)`

GetDataRetentionTimeInDaysOk returns a tuple with the DataRetentionTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionTimeInDays

`func (o *ProjectModel) SetDataRetentionTimeInDays(v float32)`

SetDataRetentionTimeInDays sets DataRetentionTimeInDays field to given value.

### HasDataRetentionTimeInDays

`func (o *ProjectModel) HasDataRetentionTimeInDays() bool`

HasDataRetentionTimeInDays returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectModel) GetOrganization() ProjectModelOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectModel) GetOrganizationOk() (*ProjectModelOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectModel) SetOrganization(v ProjectModelOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectModel) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetFileStorage

`func (o *ProjectModel) GetFileStorage() ProjectModelFileStorage`

GetFileStorage returns the FileStorage field if non-nil, zero value otherwise.

### GetFileStorageOk

`func (o *ProjectModel) GetFileStorageOk() (*ProjectModelFileStorage, bool)`

GetFileStorageOk returns a tuple with the FileStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStorage

`func (o *ProjectModel) SetFileStorage(v ProjectModelFileStorage)`

SetFileStorage sets FileStorage field to given value.

### HasFileStorage

`func (o *ProjectModel) HasFileStorage() bool`

HasFileStorage returns a boolean if a field has been set.

### GetBackends

`func (o *ProjectModel) GetBackends() ProjectModelBackends`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *ProjectModel) GetBackendsOk() (*ProjectModelBackends, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *ProjectModel) SetBackends(v ProjectModelBackends)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *ProjectModel) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetPayAsYouGo

`func (o *ProjectModel) GetPayAsYouGo() ProjectModelPayAsYouGo`

GetPayAsYouGo returns the PayAsYouGo field if non-nil, zero value otherwise.

### GetPayAsYouGoOk

`func (o *ProjectModel) GetPayAsYouGoOk() (*ProjectModelPayAsYouGo, bool)`

GetPayAsYouGoOk returns a tuple with the PayAsYouGo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAsYouGo

`func (o *ProjectModel) SetPayAsYouGo(v ProjectModelPayAsYouGo)`

SetPayAsYouGo sets PayAsYouGo field to given value.

### HasPayAsYouGo

`func (o *ProjectModel) HasPayAsYouGo() bool`

HasPayAsYouGo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


