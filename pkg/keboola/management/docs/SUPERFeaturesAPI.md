# \SUPERFeaturesAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAFeature**](SUPERFeaturesAPI.md#AddAFeature) | **Post** /manage/project-templates/{template}/features | Add a feature
[**AddAProjectFeature**](SUPERFeaturesAPI.md#AddAProjectFeature) | **Post** /manage/projects/{project_id}/features | Add a project feature
[**AddAUserFeature**](SUPERFeaturesAPI.md#AddAUserFeature) | **Post** /manage/users/{user_id_or_email}/features | Add a user feature
[**CreateAFeature**](SUPERFeaturesAPI.md#CreateAFeature) | **Post** /manage/features | Create a feature
[**DeleteAFeature**](SUPERFeaturesAPI.md#DeleteAFeature) | **Delete** /manage/features/{id} | Delete a feature
[**ListFeatures**](SUPERFeaturesAPI.md#ListFeatures) | **Get** /manage/project-templates/{template}/features | List features
[**RemoveAFeature**](SUPERFeaturesAPI.md#RemoveAFeature) | **Delete** /manage/project-templates/{template}/features/{feature} | Remove a feature
[**RemoveAProjectFeature**](SUPERFeaturesAPI.md#RemoveAProjectFeature) | **Delete** /manage/projects/{project_id}/features/{feature} | Remove a project feature
[**RemoveAUserFeature**](SUPERFeaturesAPI.md#RemoveAUserFeature) | **Delete** /manage/users/{user_id_or_email}/features/{feature} | Remove a user feature
[**RetrieveAllFeatures**](SUPERFeaturesAPI.md#RetrieveAllFeatures) | **Get** /manage/features?type&#x3D;{type} | Retrieve all features
[**RetrieveFeatureProjects**](SUPERFeaturesAPI.md#RetrieveFeatureProjects) | **Get** /manage/features/{id}/projects | Retrieve feature projects
[**RetrieveFeatureUsers**](SUPERFeaturesAPI.md#RetrieveFeatureUsers) | **Get** /manage/features/{id}/admins | Retrieve feature users
[**RetrieveOneFeature**](SUPERFeaturesAPI.md#RetrieveOneFeature) | **Get** /manage/features/{id} | Retrieve one feature
[**UpdateAFeature**](SUPERFeaturesAPI.md#UpdateAFeature) | **Patch** /manage/features/{id} | Update a feature



## AddAFeature

> AddAFeature(ctx, template).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()

Add a feature



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
	template := "template_example" // string | StringId of project template
	addAProjectFeatureRequest := *openapiclient.NewAddAProjectFeatureRequest("show-new-design") // AddAProjectFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERFeaturesAPI.AddAFeature(context.Background(), template).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.AddAFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | StringId of project template | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAProjectFeatureRequest** | [**AddAProjectFeatureRequest**](AddAProjectFeatureRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAProjectFeature

> ProjectModel AddAProjectFeature(ctx, projectId).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()

Add a project feature



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
	addAProjectFeatureRequest := *openapiclient.NewAddAProjectFeatureRequest("show-new-design") // AddAProjectFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.AddAProjectFeature(context.Background(), projectId).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.AddAProjectFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAProjectFeature`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.AddAProjectFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAProjectFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAProjectFeatureRequest** | [**AddAProjectFeatureRequest**](AddAProjectFeatureRequest.md) |  | 

### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAUserFeature

> AddAUserFeature200Response AddAUserFeature(ctx, userIdOrEmail).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()

Add a user feature



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
	userIdOrEmail := "userIdOrEmail_example" // string | User ID or email
	addAProjectFeatureRequest := *openapiclient.NewAddAProjectFeatureRequest("show-new-design") // AddAProjectFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.AddAUserFeature(context.Background(), userIdOrEmail).AddAProjectFeatureRequest(addAProjectFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.AddAUserFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAUserFeature`: AddAUserFeature200Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.AddAUserFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAUserFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAProjectFeatureRequest** | [**AddAProjectFeatureRequest**](AddAProjectFeatureRequest.md) |  | 

### Return type

[**AddAUserFeature200Response**](AddAUserFeature200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAFeature

> CreateAFeature201Response CreateAFeature(ctx).CreateAFeatureRequest(createAFeatureRequest).Execute()

Create a feature



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
	createAFeatureRequest := *openapiclient.NewCreateAFeatureRequest("show-new-design", "Show new design", "Users with this feature will see new UI") // CreateAFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.CreateAFeature(context.Background()).CreateAFeatureRequest(createAFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.CreateAFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAFeature`: CreateAFeature201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.CreateAFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAFeatureRequest** | [**CreateAFeatureRequest**](CreateAFeatureRequest.md) |  | 

### Return type

[**CreateAFeature201Response**](CreateAFeature201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAFeature

> DeleteAFeature(ctx, id).Execute()

Delete a feature



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
	id := float32(1) // float32 | ID of the feature to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERFeaturesAPI.DeleteAFeature(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.DeleteAFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID of the feature to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAFeatureRequest struct via the builder pattern


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


## ListFeatures

> []RetrieveAllFeatures200ResponseInner ListFeatures(ctx, template).Execute()

List features



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
	template := "template_example" // string | StringId of project template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.ListFeatures(context.Background(), template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.ListFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatures`: []RetrieveAllFeatures200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.ListFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | StringId of project template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RetrieveAllFeatures200ResponseInner**](RetrieveAllFeatures200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAFeature

> RemoveAFeature(ctx, template, feature).Execute()

Remove a feature



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
	template := "template_example" // string | StringId of project template
	feature := "show-new-design" // string | Feature name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERFeaturesAPI.RemoveAFeature(context.Background(), template, feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RemoveAFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | StringId of project template | 
**feature** | **string** | Feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAFeatureRequest struct via the builder pattern


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


## RemoveAProjectFeature

> RemoveAProjectFeature(ctx, projectId, feature).Execute()

Remove a project feature



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
	feature := "show-new-design" // string | Feature string ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERFeaturesAPI.RemoveAProjectFeature(context.Background(), projectId, feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RemoveAProjectFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**feature** | **string** | Feature string ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAProjectFeatureRequest struct via the builder pattern


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


## RemoveAUserFeature

> AddAUserFeature200Response RemoveAUserFeature(ctx, userIdOrEmail, feature).Execute()

Remove a user feature



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
	userIdOrEmail := "userIdOrEmail_example" // string | User ID or email
	feature := "show-new-design" // string | Feature name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.RemoveAUserFeature(context.Background(), userIdOrEmail, feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RemoveAUserFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAUserFeature`: AddAUserFeature200Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.RemoveAUserFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 
**feature** | **string** | Feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAUserFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddAUserFeature200Response**](AddAUserFeature200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllFeatures

> []RetrieveAllFeatures200ResponseInner RetrieveAllFeatures(ctx, type_).Execute()

Retrieve all features



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
	type_ := "admin" // string | Type of features you want to filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.RetrieveAllFeatures(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RetrieveAllFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllFeatures`: []RetrieveAllFeatures200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.RetrieveAllFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type of features you want to filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RetrieveAllFeatures200ResponseInner**](RetrieveAllFeatures200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFeatureProjects

> []TokenVerification200ResponseCreator RetrieveFeatureProjects(ctx, id).Execute()

Retrieve feature projects



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
	id := float32(1) // float32 | ID of the feature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.RetrieveFeatureProjects(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RetrieveFeatureProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFeatureProjects`: []TokenVerification200ResponseCreator
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.RetrieveFeatureProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID of the feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFeatureProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFeatureUsers

> []ListMaintainersInvitations200ResponseInnerUser RetrieveFeatureUsers(ctx, id).Execute()

Retrieve feature users



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
	id := float32(1) // float32 | ID of the feature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.RetrieveFeatureUsers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RetrieveFeatureUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFeatureUsers`: []ListMaintainersInvitations200ResponseInnerUser
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.RetrieveFeatureUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID of the feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFeatureUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMaintainersInvitations200ResponseInnerUser**](ListMaintainersInvitations200ResponseInnerUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOneFeature

> RetrieveAllFeatures200ResponseInner RetrieveOneFeature(ctx, id).Execute()

Retrieve one feature



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
	id := float32(1) // float32 | ID of the feature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFeaturesAPI.RetrieveOneFeature(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.RetrieveOneFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveOneFeature`: RetrieveAllFeatures200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SUPERFeaturesAPI.RetrieveOneFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID of the feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOneFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveAllFeatures200ResponseInner**](RetrieveAllFeatures200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAFeature

> UpdateAFeature(ctx, id).UpdateAFeatureRequest(updateAFeatureRequest).Execute()

Update a feature



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
	id := float32(1) // float32 | ID of the feature to update
	updateAFeatureRequest := *openapiclient.NewUpdateAFeatureRequest() // UpdateAFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERFeaturesAPI.UpdateAFeature(context.Background(), id).UpdateAFeatureRequest(updateAFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFeaturesAPI.UpdateAFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID of the feature to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAFeatureRequest** | [**UpdateAFeatureRequest**](UpdateAFeatureRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

