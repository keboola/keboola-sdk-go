# \DeletedProjectsAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelProjectDeletion**](DeletedProjectsAPI.md#CancelProjectDeletion) | **Delete** /manage/deleted-projects/{project_id} | Cancel project deletion
[**DeletedProjectDetail**](DeletedProjectsAPI.md#DeletedProjectDetail) | **Get** /manage/deleted-projects/{project_id} | Deleted project detail
[**ListDeletedProjects**](DeletedProjectsAPI.md#ListDeletedProjects) | **Get** /manage/deleted-projects | List deleted projects
[**PurgeDeletedProject**](DeletedProjectsAPI.md#PurgeDeletedProject) | **Post** /manage/deleted-projects/{project_id}/purge | Purge deleted project



## CancelProjectDeletion

> CancelProjectDeletion(ctx, projectId).ExpirationDays(expirationDays).Execute()

Cancel project deletion



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
	projectId := "projectId_example" // string | Project ID
	expirationDays := float32(8.14) // float32 | Project expiration in days (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeletedProjectsAPI.CancelProjectDeletion(context.Background(), projectId).ExpirationDays(expirationDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeletedProjectsAPI.CancelProjectDeletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelProjectDeletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expirationDays** | **float32** | Project expiration in days | 

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


## DeletedProjectDetail

> DeletedProjectDetail200Response DeletedProjectDetail(ctx, projectId).Execute()

Deleted project detail



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
	projectId := "projectId_example" // string | Deleted project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeletedProjectsAPI.DeletedProjectDetail(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeletedProjectsAPI.DeletedProjectDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletedProjectDetail`: DeletedProjectDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `DeletedProjectsAPI.DeletedProjectDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Deleted project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletedProjectDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedProjectDetail200Response**](DeletedProjectDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeletedProjects

> []ListDeletedProjects200ResponseInner ListDeletedProjects(ctx).Execute()

List deleted projects



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
	resp, r, err := apiClient.DeletedProjectsAPI.ListDeletedProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeletedProjectsAPI.ListDeletedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeletedProjects`: []ListDeletedProjects200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DeletedProjectsAPI.ListDeletedProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeletedProjectsRequest struct via the builder pattern


### Return type

[**[]ListDeletedProjects200ResponseInner**](ListDeletedProjects200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeDeletedProject

> PurgeDeletedProject200Response PurgeDeletedProject(ctx, projectId).PurgeDeletedProjectRequest(purgeDeletedProjectRequest).Execute()

Purge deleted project



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
	projectId := "projectId_example" // string | Deleted project ID
	purgeDeletedProjectRequest := *openapiclient.NewPurgeDeletedProjectRequest() // PurgeDeletedProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeletedProjectsAPI.PurgeDeletedProject(context.Background(), projectId).PurgeDeletedProjectRequest(purgeDeletedProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeletedProjectsAPI.PurgeDeletedProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurgeDeletedProject`: PurgeDeletedProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DeletedProjectsAPI.PurgeDeletedProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Deleted project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeDeletedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purgeDeletedProjectRequest** | [**PurgeDeletedProjectRequest**](PurgeDeletedProjectRequest.md) |  | 

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

