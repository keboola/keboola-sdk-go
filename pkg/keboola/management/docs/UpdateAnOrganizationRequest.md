# UpdateAnOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Organization name | [optional] 
**MaintainerId** | Pointer to **string** | Assign the organization to another maintainer. | [optional] 
**AllowAutoJoin** | Pointer to **string** | Set whether superAdmins need approval to join the organization&#39;s projects (default &#x60;true&#x60;). | [optional] 
**CrmId** | Pointer to **string** | Set CRM ID. Only maintainer members and superadmins can change this. | [optional] 
**ActivityCenterProjectId** | Pointer to **string** | Set ActivityCenter ProjectId. Only maintainer members and superadmins can change this. | [optional] 
**MfaRequired** | Pointer to **string** | Toggle whether all members of or organization and its projects must have enabled multi-factor authentication (default &#x60;false&#x60;). | [optional] 

## Methods

### NewUpdateAnOrganizationRequest

`func NewUpdateAnOrganizationRequest() *UpdateAnOrganizationRequest`

NewUpdateAnOrganizationRequest instantiates a new UpdateAnOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnOrganizationRequestWithDefaults

`func NewUpdateAnOrganizationRequestWithDefaults() *UpdateAnOrganizationRequest`

NewUpdateAnOrganizationRequestWithDefaults instantiates a new UpdateAnOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAnOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAnOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAnOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAnOrganizationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaintainerId

`func (o *UpdateAnOrganizationRequest) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *UpdateAnOrganizationRequest) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *UpdateAnOrganizationRequest) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *UpdateAnOrganizationRequest) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetAllowAutoJoin

`func (o *UpdateAnOrganizationRequest) GetAllowAutoJoin() string`

GetAllowAutoJoin returns the AllowAutoJoin field if non-nil, zero value otherwise.

### GetAllowAutoJoinOk

`func (o *UpdateAnOrganizationRequest) GetAllowAutoJoinOk() (*string, bool)`

GetAllowAutoJoinOk returns a tuple with the AllowAutoJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoJoin

`func (o *UpdateAnOrganizationRequest) SetAllowAutoJoin(v string)`

SetAllowAutoJoin sets AllowAutoJoin field to given value.

### HasAllowAutoJoin

`func (o *UpdateAnOrganizationRequest) HasAllowAutoJoin() bool`

HasAllowAutoJoin returns a boolean if a field has been set.

### GetCrmId

`func (o *UpdateAnOrganizationRequest) GetCrmId() string`

GetCrmId returns the CrmId field if non-nil, zero value otherwise.

### GetCrmIdOk

`func (o *UpdateAnOrganizationRequest) GetCrmIdOk() (*string, bool)`

GetCrmIdOk returns a tuple with the CrmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmId

`func (o *UpdateAnOrganizationRequest) SetCrmId(v string)`

SetCrmId sets CrmId field to given value.

### HasCrmId

`func (o *UpdateAnOrganizationRequest) HasCrmId() bool`

HasCrmId returns a boolean if a field has been set.

### GetActivityCenterProjectId

`func (o *UpdateAnOrganizationRequest) GetActivityCenterProjectId() string`

GetActivityCenterProjectId returns the ActivityCenterProjectId field if non-nil, zero value otherwise.

### GetActivityCenterProjectIdOk

`func (o *UpdateAnOrganizationRequest) GetActivityCenterProjectIdOk() (*string, bool)`

GetActivityCenterProjectIdOk returns a tuple with the ActivityCenterProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCenterProjectId

`func (o *UpdateAnOrganizationRequest) SetActivityCenterProjectId(v string)`

SetActivityCenterProjectId sets ActivityCenterProjectId field to given value.

### HasActivityCenterProjectId

`func (o *UpdateAnOrganizationRequest) HasActivityCenterProjectId() bool`

HasActivityCenterProjectId returns a boolean if a field has been set.

### GetMfaRequired

`func (o *UpdateAnOrganizationRequest) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *UpdateAnOrganizationRequest) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *UpdateAnOrganizationRequest) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *UpdateAnOrganizationRequest) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


