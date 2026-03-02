# \SUPERUIManagementAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplication**](SUPERUIManagementAPI.md#DeleteApplication) | **Delete** /manage/ui-apps/{id} | Delete Application
[**ListApplications**](SUPERUIManagementAPI.md#ListApplications) | **Get** /manage/ui-apps | List Applications
[**RegisterNewApplicationVersion**](SUPERUIManagementAPI.md#RegisterNewApplicationVersion) | **Post** /manage/ui-apps | Register New Application/Version



## DeleteApplication

> DeleteApplication(ctx, id).Execute()

Delete Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | stringId/name of application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERUIManagementAPI.DeleteApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERUIManagementAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | stringId/name of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> []ListApplications200ResponseInner ListApplications(ctx).Execute()

List Applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERUIManagementAPI.ListApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERUIManagementAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SUPERUIManagementAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNewApplicationVersion

> RegisterNewApplicationVersion201Response RegisterNewApplicationVersion(ctx).RegisterNewApplicationVersionRequest(registerNewApplicationVersionRequest).Execute()

Register New Application/Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	registerNewApplicationVersionRequest := *openapiclient.NewRegisterNewApplicationVersionRequest("ManifestUrl_example") // RegisterNewApplicationVersionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERUIManagementAPI.RegisterNewApplicationVersion(context.Background()).RegisterNewApplicationVersionRequest(registerNewApplicationVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERUIManagementAPI.RegisterNewApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterNewApplicationVersion`: RegisterNewApplicationVersion201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERUIManagementAPI.RegisterNewApplicationVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNewApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerNewApplicationVersionRequest** | [**RegisterNewApplicationVersionRequest**](RegisterNewApplicationVersionRequest.md) |  | 

### Return type

[**RegisterNewApplicationVersion201Response**](RegisterNewApplicationVersion201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

