# ListUsedPromoCodes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**UsedAt** | **string** |  | 
**Project** | [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | 

## Methods

### NewListUsedPromoCodes200ResponseInner

`func NewListUsedPromoCodes200ResponseInner(code string, usedAt string, project TokenVerification200ResponseCreator, ) *ListUsedPromoCodes200ResponseInner`

NewListUsedPromoCodes200ResponseInner instantiates a new ListUsedPromoCodes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsedPromoCodes200ResponseInnerWithDefaults

`func NewListUsedPromoCodes200ResponseInnerWithDefaults() *ListUsedPromoCodes200ResponseInner`

NewListUsedPromoCodes200ResponseInnerWithDefaults instantiates a new ListUsedPromoCodes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ListUsedPromoCodes200ResponseInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListUsedPromoCodes200ResponseInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListUsedPromoCodes200ResponseInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetUsedAt

`func (o *ListUsedPromoCodes200ResponseInner) GetUsedAt() string`

GetUsedAt returns the UsedAt field if non-nil, zero value otherwise.

### GetUsedAtOk

`func (o *ListUsedPromoCodes200ResponseInner) GetUsedAtOk() (*string, bool)`

GetUsedAtOk returns a tuple with the UsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAt

`func (o *ListUsedPromoCodes200ResponseInner) SetUsedAt(v string)`

SetUsedAt sets UsedAt field to given value.


### GetProject

`func (o *ListUsedPromoCodes200ResponseInner) GetProject() TokenVerification200ResponseCreator`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ListUsedPromoCodes200ResponseInner) GetProjectOk() (*TokenVerification200ResponseCreator, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ListUsedPromoCodes200ResponseInner) SetProject(v TokenVerification200ResponseCreator)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


