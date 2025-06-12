# \SUPERCommandsAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunCommand**](SUPERCommandsAPI.md#RunCommand) | **Post** /manage/commands | Run Command



## RunCommand

> PurgeDeletedProject200Response RunCommand(ctx).RunCommandRequest(runCommandRequest).Execute()

Run Command



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	runCommandRequest := *openapiclient.NewRunCommandRequest("Command_example") // RunCommandRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERCommandsAPI.RunCommand(context.Background()).RunCommandRequest(runCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERCommandsAPI.RunCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunCommand`: PurgeDeletedProject200Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERCommandsAPI.RunCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCommandRequest** | [**RunCommandRequest**](RunCommandRequest.md) |  | 

### Return type

[**PurgeDeletedProject200Response**](PurgeDeletedProject200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

