# TokenVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**LastUsed** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 
**IsSessionToken** | Pointer to **bool** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**Scopes** | Pointer to **[]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | [optional] 
**User** | Pointer to [**TokenVerification200ResponseUser**](TokenVerification200ResponseUser.md) |  | [optional] 

## Methods

### NewTokenVerification200Response

`func NewTokenVerification200Response() *TokenVerification200Response`

NewTokenVerification200Response instantiates a new TokenVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenVerification200ResponseWithDefaults

`func NewTokenVerification200ResponseWithDefaults() *TokenVerification200Response`

NewTokenVerification200ResponseWithDefaults instantiates a new TokenVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenVerification200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenVerification200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenVerification200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *TokenVerification200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *TokenVerification200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenVerification200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenVerification200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenVerification200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *TokenVerification200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TokenVerification200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TokenVerification200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TokenVerification200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUsed

`func (o *TokenVerification200Response) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *TokenVerification200Response) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *TokenVerification200Response) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *TokenVerification200Response) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetExpires

`func (o *TokenVerification200Response) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenVerification200Response) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenVerification200Response) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenVerification200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsSessionToken

`func (o *TokenVerification200Response) GetIsSessionToken() bool`

GetIsSessionToken returns the IsSessionToken field if non-nil, zero value otherwise.

### GetIsSessionTokenOk

`func (o *TokenVerification200Response) GetIsSessionTokenOk() (*bool, bool)`

GetIsSessionTokenOk returns a tuple with the IsSessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSessionToken

`func (o *TokenVerification200Response) SetIsSessionToken(v bool)`

SetIsSessionToken sets IsSessionToken field to given value.

### HasIsSessionToken

`func (o *TokenVerification200Response) HasIsSessionToken() bool`

HasIsSessionToken returns a boolean if a field has been set.

### GetIsExpired

`func (o *TokenVerification200Response) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *TokenVerification200Response) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *TokenVerification200Response) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *TokenVerification200Response) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetIsDisabled

`func (o *TokenVerification200Response) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *TokenVerification200Response) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *TokenVerification200Response) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *TokenVerification200Response) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetScopes

`func (o *TokenVerification200Response) GetScopes() []interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenVerification200Response) GetScopesOk() (*[]interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenVerification200Response) SetScopes(v []interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenVerification200Response) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetType

`func (o *TokenVerification200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenVerification200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenVerification200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenVerification200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreator

`func (o *TokenVerification200Response) GetCreator() TokenVerification200ResponseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TokenVerification200Response) GetCreatorOk() (*TokenVerification200ResponseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TokenVerification200Response) SetCreator(v TokenVerification200ResponseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *TokenVerification200Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetUser

`func (o *TokenVerification200Response) GetUser() TokenVerification200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TokenVerification200Response) GetUserOk() (*TokenVerification200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TokenVerification200Response) SetUser(v TokenVerification200ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *TokenVerification200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


