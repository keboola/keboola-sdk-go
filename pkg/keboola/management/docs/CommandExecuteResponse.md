# CommandExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandExecutionId** | **string** | execution ID; this ID is appended to all logs provided by the command | 

## Methods

### NewCommandExecuteResponse

`func NewCommandExecuteResponse(commandExecutionId string, ) *CommandExecuteResponse`

NewCommandExecuteResponse instantiates a new CommandExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandExecuteResponseWithDefaults

`func NewCommandExecuteResponseWithDefaults() *CommandExecuteResponse`

NewCommandExecuteResponseWithDefaults instantiates a new CommandExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandExecutionId

`func (o *CommandExecuteResponse) GetCommandExecutionId() string`

GetCommandExecutionId returns the CommandExecutionId field if non-nil, zero value otherwise.

### GetCommandExecutionIdOk

`func (o *CommandExecuteResponse) GetCommandExecutionIdOk() (*string, bool)`

GetCommandExecutionIdOk returns a tuple with the CommandExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandExecutionId

`func (o *CommandExecuteResponse) SetCommandExecutionId(v string)`

SetCommandExecutionId sets CommandExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


