# ChangeProjectDisabledStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDisabled** | Pointer to **bool** | Enable or disable project | [optional] 
**DisableReason** | Pointer to **string** | Why the project is disabled | [optional] 
**EstimatedEndTime** | Pointer to **string** | When the project will be enabled | [optional] 

## Methods

### NewChangeProjectDisabledStatusRequest

`func NewChangeProjectDisabledStatusRequest() *ChangeProjectDisabledStatusRequest`

NewChangeProjectDisabledStatusRequest instantiates a new ChangeProjectDisabledStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeProjectDisabledStatusRequestWithDefaults

`func NewChangeProjectDisabledStatusRequestWithDefaults() *ChangeProjectDisabledStatusRequest`

NewChangeProjectDisabledStatusRequestWithDefaults instantiates a new ChangeProjectDisabledStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDisabled

`func (o *ChangeProjectDisabledStatusRequest) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ChangeProjectDisabledStatusRequest) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ChangeProjectDisabledStatusRequest) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ChangeProjectDisabledStatusRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetDisableReason

`func (o *ChangeProjectDisabledStatusRequest) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *ChangeProjectDisabledStatusRequest) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *ChangeProjectDisabledStatusRequest) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *ChangeProjectDisabledStatusRequest) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetEstimatedEndTime

`func (o *ChangeProjectDisabledStatusRequest) GetEstimatedEndTime() string`

GetEstimatedEndTime returns the EstimatedEndTime field if non-nil, zero value otherwise.

### GetEstimatedEndTimeOk

`func (o *ChangeProjectDisabledStatusRequest) GetEstimatedEndTimeOk() (*string, bool)`

GetEstimatedEndTimeOk returns a tuple with the EstimatedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEndTime

`func (o *ChangeProjectDisabledStatusRequest) SetEstimatedEndTime(v string)`

SetEstimatedEndTime sets EstimatedEndTime field to given value.

### HasEstimatedEndTime

`func (o *ChangeProjectDisabledStatusRequest) HasEstimatedEndTime() bool`

HasEstimatedEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


