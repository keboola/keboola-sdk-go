# CreateNewPromoCodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Promo-code name | 
**ExpirationDays** | **float32** | Promo-code expiration days | 
**OrganizationId** | **float32** | Add to organization | 
**ProjectTemplateStringId** | **string** | Use project template | 

## Methods

### NewCreateNewPromoCodesRequest

`func NewCreateNewPromoCodesRequest(code string, expirationDays float32, organizationId float32, projectTemplateStringId string, ) *CreateNewPromoCodesRequest`

NewCreateNewPromoCodesRequest instantiates a new CreateNewPromoCodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewPromoCodesRequestWithDefaults

`func NewCreateNewPromoCodesRequestWithDefaults() *CreateNewPromoCodesRequest`

NewCreateNewPromoCodesRequestWithDefaults instantiates a new CreateNewPromoCodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateNewPromoCodesRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateNewPromoCodesRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateNewPromoCodesRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetExpirationDays

`func (o *CreateNewPromoCodesRequest) GetExpirationDays() float32`

GetExpirationDays returns the ExpirationDays field if non-nil, zero value otherwise.

### GetExpirationDaysOk

`func (o *CreateNewPromoCodesRequest) GetExpirationDaysOk() (*float32, bool)`

GetExpirationDaysOk returns a tuple with the ExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDays

`func (o *CreateNewPromoCodesRequest) SetExpirationDays(v float32)`

SetExpirationDays sets ExpirationDays field to given value.


### GetOrganizationId

`func (o *CreateNewPromoCodesRequest) GetOrganizationId() float32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateNewPromoCodesRequest) GetOrganizationIdOk() (*float32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateNewPromoCodesRequest) SetOrganizationId(v float32)`

SetOrganizationId sets OrganizationId field to given value.


### GetProjectTemplateStringId

`func (o *CreateNewPromoCodesRequest) GetProjectTemplateStringId() string`

GetProjectTemplateStringId returns the ProjectTemplateStringId field if non-nil, zero value otherwise.

### GetProjectTemplateStringIdOk

`func (o *CreateNewPromoCodesRequest) GetProjectTemplateStringIdOk() (*string, bool)`

GetProjectTemplateStringIdOk returns a tuple with the ProjectTemplateStringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplateStringId

`func (o *CreateNewPromoCodesRequest) SetProjectTemplateStringId(v string)`

SetProjectTemplateStringId sets ProjectTemplateStringId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


