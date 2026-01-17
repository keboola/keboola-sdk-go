# \ProjectsAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAProject**](ProjectsAPI.md#AddAProject) | **Post** /manage/organizations/{organization_id}/projects | Add a project
[**AddAUserToAProject**](ProjectsAPI.md#AddAUserToAProject) | **Post** /manage/projects/{project_id}/users | Add a user to a project
[**ApproveJoinRequest**](ProjectsAPI.md#ApproveJoinRequest) | **Put** /manage/projects/{project_id}/join-requests/{join_request_id} | Approve join request
[**AssignProjectFileStorage**](ProjectsAPI.md#AssignProjectFileStorage) | **Post** /manage/projects/{project_id}/file-storage | Assign project file Storage
[**AssignProjectStorageBackend**](ProjectsAPI.md#AssignProjectStorageBackend) | **Post** /manage/projects/{project_id}/storage-backend | Assign project Storage backend
[**CancelProjectInvitation**](ProjectsAPI.md#CancelProjectInvitation) | **Delete** /manage/projects/{project_id}/invitations/{invitation_id} | Cancel project invitation
[**ChangeProjectDisabledStatus**](ProjectsAPI.md#ChangeProjectDisabledStatus) | **Post** /manage/projects/{project_id}/disabled | Change project disabled status
[**ChangeRoleOfAUserInAProject**](ProjectsAPI.md#ChangeRoleOfAUserInAProject) | **Patch** /manage/projects/{project_id}/users/{user_id} | Change role of a user in a project
[**CreateStorageToken**](ProjectsAPI.md#CreateStorageToken) | **Post** /manage/projects/{project_id}/tokens | Create Storage token
[**DeleteAProject**](ProjectsAPI.md#DeleteAProject) | **Delete** /manage/projects/{project_id} | Delete a project
[**DeleteAUserFromAProject**](ProjectsAPI.md#DeleteAUserFromAProject) | **Delete** /manage/projects/{project_id}/users/{user_id} | Delete a user from a project
[**GiveCredits**](ProjectsAPI.md#GiveCredits) | **Post** /manage/projects/{project_id}/credits | Give credits
[**InviteAUserToAProject**](ProjectsAPI.md#InviteAUserToAProject) | **Post** /manage/projects/{project_id}/invitations | Invite a user to a project
[**JoinAProject**](ProjectsAPI.md#JoinAProject) | **Post** /manage/projects/{project_id}/join-project | Join a project
[**JoinRequestDetail**](ProjectsAPI.md#JoinRequestDetail) | **Get** /manage/projects/{project_id}/join-requests/{join_request_id} | Join request detail
[**ListJoinRequests**](ProjectsAPI.md#ListJoinRequests) | **Get** /manage/projects/{project_id}/join-requests | List join requests
[**ListProjectInvitations**](ProjectsAPI.md#ListProjectInvitations) | **Get** /manage/projects/{project_id}/invitations | List project invitations
[**ListProjectMetadata**](ProjectsAPI.md#ListProjectMetadata) | **Get** /manage/projects/{project_id}/metadata | List project Metadata
[**ListProjectUsers**](ProjectsAPI.md#ListProjectUsers) | **Get** /manage/projects/{project_id}/users | List project users
[**ListProjectsForAnOrganization**](ProjectsAPI.md#ListProjectsForAnOrganization) | **Get** /manage/organizations/{organization_id}/projects | List projects for an organization
[**ListTemplates**](ProjectsAPI.md#ListTemplates) | **Get** /manage/project-templates/ | List templates
[**MoveAProject**](ProjectsAPI.md#MoveAProject) | **Post** /manage/projects/{project_id}/organizations | Move a project
[**ProjectDetail**](ProjectsAPI.md#ProjectDetail) | **Get** /manage/projects/{project_id} | Project detail
[**ProjectInvitationDetail**](ProjectsAPI.md#ProjectInvitationDetail) | **Get** /manage/projects/{project_id}/invitations/{invitation_id} | Project invitation detail
[**ProjectTemplateDetail**](ProjectsAPI.md#ProjectTemplateDetail) | **Get** /manage/project-templates/{template_string_id} | Project template detail
[**RejectJoinRequest**](ProjectsAPI.md#RejectJoinRequest) | **Delete** /manage/projects/{project_id}/join-requests/{join_request_id} | Reject join request
[**RemoveProjectLimit**](ProjectsAPI.md#RemoveProjectLimit) | **Delete** /manage/projects/{project_id}/limits/{limit_name} | Remove project limit
[**RemoveProjectMetadata**](ProjectsAPI.md#RemoveProjectMetadata) | **Delete** /manage/projects/{project_id}/metadata/{metadata_id} | Remove project metadata
[**RemoveProjectStorageBackend**](ProjectsAPI.md#RemoveProjectStorageBackend) | **Delete** /manage/projects/{project_id}/storage-backend/{storage_backend_id} | Remove project Storage backend
[**RequestAccessToAProject**](ProjectsAPI.md#RequestAccessToAProject) | **Post** /manage/projects/{project_id}/request-access | Request access to a project
[**SetProjectLimits**](ProjectsAPI.md#SetProjectLimits) | **Post** /manage/projects/{project_id}/limits | Set project limits
[**SetProjectMetadata**](ProjectsAPI.md#SetProjectMetadata) | **Post** /manage/projects/{project_id}/metadata | Set project metadata
[**UpdateAProject**](ProjectsAPI.md#UpdateAProject) | **Put** /manage/projects/{project_id} | Update a project



## AddAProject

> ProjectModel AddAProject(ctx, organizationId).AddAProjectRequest(addAProjectRequest).Execute()

Add a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID; projects have to be created in an organization.
	addAProjectRequest := *openapiclient.NewAddAProjectRequest("My Demo", "demo") // AddAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.AddAProject(context.Background(), organizationId).AddAProjectRequest(addAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAProject`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID; projects have to be created in an organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAProjectRequest** | [**AddAProjectRequest**](AddAProjectRequest.md) |  | 

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


## AddAUserToAProject

> AddAUserToAProject(ctx, projectId).AddAUserToAProjectRequest(addAUserToAProjectRequest).Execute()

Add a user to a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	addAUserToAProjectRequest := *openapiclient.NewAddAUserToAProjectRequest("martin@keboola.com") // AddAUserToAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.AddAUserToAProject(context.Background(), projectId).AddAUserToAProjectRequest(addAUserToAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddAUserToAProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddAUserToAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAUserToAProjectRequest** | [**AddAUserToAProjectRequest**](AddAUserToAProjectRequest.md) |  | 

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


## ApproveJoinRequest

> ApproveJoinRequest(ctx, projectId, joinRequestId).Execute()

Approve join request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	joinRequestId := "joinRequestId_example" // string | Join Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApproveJoinRequest(context.Background(), projectId, joinRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApproveJoinRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**joinRequestId** | **string** | Join Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveJoinRequestRequest struct via the builder pattern


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


## AssignProjectFileStorage

> ProjectModel AssignProjectFileStorage(ctx, projectId).AssignProjectFileStorageRequest(assignProjectFileStorageRequest).Execute()

Assign project file Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	assignProjectFileStorageRequest := *openapiclient.NewAssignProjectFileStorageRequest("32") // AssignProjectFileStorageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.AssignProjectFileStorage(context.Background(), projectId).AssignProjectFileStorageRequest(assignProjectFileStorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AssignProjectFileStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignProjectFileStorage`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AssignProjectFileStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignProjectFileStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignProjectFileStorageRequest** | [**AssignProjectFileStorageRequest**](AssignProjectFileStorageRequest.md) |  | 

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


## AssignProjectStorageBackend

> ProjectModel AssignProjectStorageBackend(ctx, projectId).AssignProjectStorageBackendRequest(assignProjectStorageBackendRequest).Execute()

Assign project Storage backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	assignProjectStorageBackendRequest := *openapiclient.NewAssignProjectStorageBackendRequest("32") // AssignProjectStorageBackendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.AssignProjectStorageBackend(context.Background(), projectId).AssignProjectStorageBackendRequest(assignProjectStorageBackendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AssignProjectStorageBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignProjectStorageBackend`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AssignProjectStorageBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignProjectStorageBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignProjectStorageBackendRequest** | [**AssignProjectStorageBackendRequest**](AssignProjectStorageBackendRequest.md) |  | 

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


## CancelProjectInvitation

> CancelProjectInvitation(ctx, projectId, invitationId).Execute()

Cancel project invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.CancelProjectInvitation(context.Background(), projectId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CancelProjectInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelProjectInvitationRequest struct via the builder pattern


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


## ChangeProjectDisabledStatus

> ProjectModel ChangeProjectDisabledStatus(ctx, projectId).ChangeProjectDisabledStatusRequest(changeProjectDisabledStatusRequest).Execute()

Change project disabled status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	changeProjectDisabledStatusRequest := *openapiclient.NewChangeProjectDisabledStatusRequest() // ChangeProjectDisabledStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ChangeProjectDisabledStatus(context.Background(), projectId).ChangeProjectDisabledStatusRequest(changeProjectDisabledStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ChangeProjectDisabledStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeProjectDisabledStatus`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ChangeProjectDisabledStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeProjectDisabledStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeProjectDisabledStatusRequest** | [**ChangeProjectDisabledStatusRequest**](ChangeProjectDisabledStatusRequest.md) |  | 

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


## ChangeRoleOfAUserInAProject

> ChangeRoleOfAUserInAProject200Response ChangeRoleOfAUserInAProject(ctx, projectId, userId).ChangeRoleOfAUserInAProjectRequest(changeRoleOfAUserInAProjectRequest).Execute()

Change role of a user in a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	userId := "userId_example" // string | User ID
	changeRoleOfAUserInAProjectRequest := *openapiclient.NewChangeRoleOfAUserInAProjectRequest() // ChangeRoleOfAUserInAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ChangeRoleOfAUserInAProject(context.Background(), projectId, userId).ChangeRoleOfAUserInAProjectRequest(changeRoleOfAUserInAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ChangeRoleOfAUserInAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRoleOfAUserInAProject`: ChangeRoleOfAUserInAProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ChangeRoleOfAUserInAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRoleOfAUserInAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeRoleOfAUserInAProjectRequest** | [**ChangeRoleOfAUserInAProjectRequest**](ChangeRoleOfAUserInAProjectRequest.md) |  | 

### Return type

[**ChangeRoleOfAUserInAProject200Response**](ChangeRoleOfAUserInAProject200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageToken

> CreateStorageToken201Response CreateStorageToken(ctx, projectId).CreateStorageTokenRequest(createStorageTokenRequest).Execute()

Create Storage token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	createStorageTokenRequest := *openapiclient.NewCreateStorageTokenRequest("Test Token") // CreateStorageTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateStorageToken(context.Background(), projectId).CreateStorageTokenRequest(createStorageTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateStorageToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorageToken`: CreateStorageToken201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateStorageToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStorageTokenRequest** | [**CreateStorageTokenRequest**](CreateStorageTokenRequest.md) |  | 

### Return type

[**CreateStorageToken201Response**](CreateStorageToken201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAProject

> DeleteAProject(ctx, projectId).Execute()

Delete a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteAProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteAProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAProjectRequest struct via the builder pattern


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


## DeleteAUserFromAProject

> DeleteAUserFromAProject(ctx, projectId, userId).Execute()

Delete a user from a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteAUserFromAProject(context.Background(), projectId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteAUserFromAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAUserFromAProjectRequest struct via the builder pattern


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


## GiveCredits

> GiveCredits201Response GiveCredits(ctx, projectId).GiveCreditsRequest(giveCreditsRequest).Execute()

Give credits



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	giveCreditsRequest := *openapiclient.NewGiveCreditsRequest(float32(10), "Promo credits") // GiveCreditsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GiveCredits(context.Background(), projectId).GiveCreditsRequest(giveCreditsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GiveCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GiveCredits`: GiveCredits201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GiveCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGiveCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **giveCreditsRequest** | [**GiveCreditsRequest**](GiveCreditsRequest.md) |  | 

### Return type

[**GiveCredits201Response**](GiveCredits201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteAUserToAProject

> InviteAUserToAProject201Response InviteAUserToAProject(ctx, projectId).InviteAUserToAProjectRequest(inviteAUserToAProjectRequest).Execute()

Invite a user to a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	inviteAUserToAProjectRequest := *openapiclient.NewInviteAUserToAProjectRequest("martin@keboola.com") // InviteAUserToAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.InviteAUserToAProject(context.Background(), projectId).InviteAUserToAProjectRequest(inviteAUserToAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.InviteAUserToAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteAUserToAProject`: InviteAUserToAProject201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.InviteAUserToAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteAUserToAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteAUserToAProjectRequest** | [**InviteAUserToAProjectRequest**](InviteAUserToAProjectRequest.md) |  | 

### Return type

[**InviteAUserToAProject201Response**](InviteAUserToAProject201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinAProject

> JoinAProject(ctx, projectId).JoinAProjectRequest(joinAProjectRequest).Execute()

Join a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	joinAProjectRequest := *openapiclient.NewJoinAProjectRequest() // JoinAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.JoinAProject(context.Background(), projectId).JoinAProjectRequest(joinAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.JoinAProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiJoinAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **joinAProjectRequest** | [**JoinAProjectRequest**](JoinAProjectRequest.md) |  | 

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


## JoinRequestDetail

> JoinRequestDetail200Response JoinRequestDetail(ctx, projectId, joinRequestId).Execute()

Join request detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	joinRequestId := "joinRequestId_example" // string | Join Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.JoinRequestDetail(context.Background(), projectId, joinRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.JoinRequestDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoinRequestDetail`: JoinRequestDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.JoinRequestDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**joinRequestId** | **string** | Join Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinRequestDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JoinRequestDetail200Response**](JoinRequestDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJoinRequests

> []ListJoinRequests200ResponseInner ListJoinRequests(ctx, projectId).Execute()

List join requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListJoinRequests(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListJoinRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJoinRequests`: []ListJoinRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListJoinRequests200ResponseInner**](ListJoinRequests200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectInvitations

> []ListProjectInvitations200ResponseInner ListProjectInvitations(ctx, projectId).Execute()

List project invitations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListProjectInvitations(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListProjectInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectInvitations`: []ListProjectInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListProjectInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListProjectInvitations200ResponseInner**](ListProjectInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectMetadata

> []ListOrganizationMetadata200ResponseInner ListProjectMetadata(ctx, projectId).Execute()

List project Metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListProjectMetadata(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListProjectMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListProjectMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListOrganizationMetadata200ResponseInner**](ListOrganizationMetadata200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectUsers

> []ListProjectUsers200ResponseInner ListProjectUsers(ctx, projectId).Execute()

List project users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListProjectUsers(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListProjectUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectUsers`: []ListProjectUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListProjectUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListProjectUsers200ResponseInner**](ListProjectUsers200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectsForAnOrganization

> ListProjectsForAnOrganization200Response ListProjectsForAnOrganization(ctx, organizationId).Execute()

List projects for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID; projects have to be created in an organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListProjectsForAnOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListProjectsForAnOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectsForAnOrganization`: ListProjectsForAnOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListProjectsForAnOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID; projects have to be created in an organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsForAnOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListProjectsForAnOrganization200Response**](ListProjectsForAnOrganization200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> []ListTemplates200ResponseInner ListTemplates(ctx).Execute()

List templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: []ListTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


### Return type

[**[]ListTemplates200ResponseInner**](ListTemplates200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveAProject

> ProjectModel MoveAProject(ctx, projectId).MoveAProjectRequest(moveAProjectRequest).Execute()

Move a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	moveAProjectRequest := *openapiclient.NewMoveAProjectRequest("523") // MoveAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.MoveAProject(context.Background(), projectId).MoveAProjectRequest(moveAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.MoveAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveAProject`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.MoveAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveAProjectRequest** | [**MoveAProjectRequest**](MoveAProjectRequest.md) |  | 

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


## ProjectDetail

> ProjectModel ProjectDetail(ctx, projectId).Execute()

Project detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectDetail(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectDetail`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectInvitationDetail

> InviteAUserToAProject201Response ProjectInvitationDetail(ctx, projectId, invitationId).Execute()

Project invitation detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectInvitationDetail(context.Background(), projectId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectInvitationDetail`: InviteAUserToAProject201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectInvitationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InviteAUserToAProject201Response**](InviteAUserToAProject201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectTemplateDetail

> ProjectTemplateDetail200Response ProjectTemplateDetail(ctx, templateStringId).Execute()

Project template detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	templateStringId := "templateStringId_example" // string | Project string ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectTemplateDetail(context.Background(), templateStringId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectTemplateDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectTemplateDetail`: ProjectTemplateDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectTemplateDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateStringId** | **string** | Project string ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectTemplateDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectTemplateDetail200Response**](ProjectTemplateDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectJoinRequest

> RejectJoinRequest(ctx, projectId, joinRequestId).Execute()

Reject join request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	joinRequestId := "joinRequestId_example" // string | Join Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.RejectJoinRequest(context.Background(), projectId, joinRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RejectJoinRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**joinRequestId** | **string** | Join Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectJoinRequestRequest struct via the builder pattern


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


## RemoveProjectLimit

> RemoveProjectLimit(ctx, projectId, limitName).Execute()

Remove project limit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	limitName := "limitName_example" // string | Limit name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.RemoveProjectLimit(context.Background(), projectId, limitName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RemoveProjectLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**limitName** | **string** | Limit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectLimitRequest struct via the builder pattern


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


## RemoveProjectMetadata

> RemoveProjectMetadata(ctx, projectId, metadataId).Execute()

Remove project metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	metadataId := "metadataId_example" // string | Metadata ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.RemoveProjectMetadata(context.Background(), projectId, metadataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RemoveProjectMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**metadataId** | **string** | Metadata ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectMetadataRequest struct via the builder pattern


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


## RemoveProjectStorageBackend

> ProjectModel RemoveProjectStorageBackend(ctx, projectId, storageBackendId).Execute()

Remove project Storage backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	storageBackendId := "storageBackendId_example" // string | Storage backend ID; must be one of Storage backends assigned to the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.RemoveProjectStorageBackend(context.Background(), projectId, storageBackendId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RemoveProjectStorageBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProjectStorageBackend`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.RemoveProjectStorageBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**storageBackendId** | **string** | Storage backend ID; must be one of Storage backends assigned to the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectStorageBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestAccessToAProject

> RequestAccessToAProject201Response RequestAccessToAProject(ctx, projectId).JoinAProjectRequest(joinAProjectRequest).Execute()

Request access to a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	joinAProjectRequest := *openapiclient.NewJoinAProjectRequest() // JoinAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.RequestAccessToAProject(context.Background(), projectId).JoinAProjectRequest(joinAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RequestAccessToAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestAccessToAProject`: RequestAccessToAProject201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.RequestAccessToAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestAccessToAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **joinAProjectRequest** | [**JoinAProjectRequest**](JoinAProjectRequest.md) |  | 

### Return type

[**RequestAccessToAProject201Response**](RequestAccessToAProject201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProjectLimits

> ProjectModel SetProjectLimits(ctx, projectId).SetProjectLimitsRequest(setProjectLimitsRequest).Execute()

Set project limits



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	setProjectLimitsRequest := *openapiclient.NewSetProjectLimitsRequest() // SetProjectLimitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.SetProjectLimits(context.Background(), projectId).SetProjectLimitsRequest(setProjectLimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.SetProjectLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProjectLimits`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.SetProjectLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setProjectLimitsRequest** | [**SetProjectLimitsRequest**](SetProjectLimitsRequest.md) |  | 

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


## SetProjectMetadata

> []ListOrganizationMetadata200ResponseInner SetProjectMetadata(ctx, projectId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()

Set project metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	setMaintainerMetadataRequest := *openapiclient.NewSetMaintainerMetadataRequest("Provider_example", []interface{}{nil}) // SetMaintainerMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.SetProjectMetadata(context.Background(), projectId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.SetProjectMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProjectMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.SetProjectMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMaintainerMetadataRequest** | [**SetMaintainerMetadataRequest**](SetMaintainerMetadataRequest.md) |  | 

### Return type

[**[]ListOrganizationMetadata200ResponseInner**](ListOrganizationMetadata200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAProject

> ProjectModel UpdateAProject(ctx, projectId).UpdateAProjectRequest(updateAProjectRequest).Execute()

Update a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func main() {
	projectId := "projectId_example" // string | Project ID
	updateAProjectRequest := *openapiclient.NewUpdateAProjectRequest() // UpdateAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateAProject(context.Background(), projectId).UpdateAProjectRequest(updateAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateAProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAProject`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateAProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAProjectRequest** | [**UpdateAProjectRequest**](UpdateAProjectRequest.md) |  | 

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

