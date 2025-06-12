# GiveCreditsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Credits amount | 
**Description** | **string** | Reason or label of given credits | 

## Methods

### NewGiveCreditsRequest

`func NewGiveCreditsRequest(amount float32, description string, ) *GiveCreditsRequest`

NewGiveCreditsRequest instantiates a new GiveCreditsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveCreditsRequestWithDefaults

`func NewGiveCreditsRequestWithDefaults() *GiveCreditsRequest`

NewGiveCreditsRequestWithDefaults instantiates a new GiveCreditsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GiveCreditsRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GiveCreditsRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GiveCreditsRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *GiveCreditsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GiveCreditsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GiveCreditsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


