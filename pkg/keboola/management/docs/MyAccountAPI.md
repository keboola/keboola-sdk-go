# \MyAccountAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAMaintainerInvitation**](MyAccountAPI.md#AcceptAMaintainerInvitation) | **Put** /manage/current-user/maintainers-invitations/{invitation_id} | Accept a maintainer invitation
[**AcceptAProjectInvitation**](MyAccountAPI.md#AcceptAProjectInvitation) | **Put** /manage/current-user/projects-invitations/{invitation_id} | Accept a project invitation
[**AcceptAnOrganizationInvitation**](MyAccountAPI.md#AcceptAnOrganizationInvitation) | **Put** /manage/current-user/organizations-invitations/{invitation_id} | Accept an organization invitation
[**CancelAJoinRequest**](MyAccountAPI.md#CancelAJoinRequest) | **Delete** /manage/current-user/projects-join-requests/{join_request_id} | Cancel a join request
[**CreateProject**](MyAccountAPI.md#CreateProject) | **Post** /manage/current-user/promo-codes | Create project
[**DeclineAInvitation**](MyAccountAPI.md#DeclineAInvitation) | **Delete** /manage/current-user/projects-invitations/{invitation_id} | Decline a invitation
[**ListMyJoinRequests**](MyAccountAPI.md#ListMyJoinRequests) | **Get** /manage/current-user/projects-join-requests | List my join requests
[**ListMyMaintainerInvitations**](MyAccountAPI.md#ListMyMaintainerInvitations) | **Get** /manage/current-user/maintainers-invitations | List my maintainer invitations
[**ListMyOrganizationInvitations**](MyAccountAPI.md#ListMyOrganizationInvitations) | **Get** /manage/current-user/organizations-invitations | List my organization invitations
[**ListMyProjectInvitations**](MyAccountAPI.md#ListMyProjectInvitations) | **Get** /manage/current-user/projects-invitations | List my project invitations
[**ListUsedPromoCodes**](MyAccountAPI.md#ListUsedPromoCodes) | **Get** /manage/current-user/promo-codes | List used promo codes
[**MyJoinRequestDetail**](MyAccountAPI.md#MyJoinRequestDetail) | **Get** /manage/current-user/projects-join-requests/{join_request_id} | My join request detail
[**MyMaintainerInvitationDetail**](MyAccountAPI.md#MyMaintainerInvitationDetail) | **Get** /manage/current-user/maintainers-invitations/{invitation_id} | My maintainer invitation detail
[**MyOrganizationInvitationDetail**](MyAccountAPI.md#MyOrganizationInvitationDetail) | **Get** /manage/current-user/organizations-invitations/{invitation_id} | My organization invitation detail
[**MyProjectInvitationDetail**](MyAccountAPI.md#MyProjectInvitationDetail) | **Get** /manage/current-user/projects-invitations/{invitation_id} | My project invitation detail



## AcceptAMaintainerInvitation

> AcceptAMaintainerInvitation(ctx, invitationId).Execute()

Accept a maintainer invitation



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.AcceptAMaintainerInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.AcceptAMaintainerInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptAMaintainerInvitationRequest struct via the builder pattern


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


## AcceptAProjectInvitation

> AcceptAProjectInvitation(ctx, invitationId).Execute()

Accept a project invitation



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.AcceptAProjectInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.AcceptAProjectInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptAProjectInvitationRequest struct via the builder pattern


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


## AcceptAnOrganizationInvitation

> AcceptAnOrganizationInvitation(ctx, invitationId).Execute()

Accept an organization invitation



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.AcceptAnOrganizationInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.AcceptAnOrganizationInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptAnOrganizationInvitationRequest struct via the builder pattern


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


## CancelAJoinRequest

> CancelAJoinRequest(ctx, joinRequestId).Execute()

Cancel a join request



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
	joinRequestId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.CancelAJoinRequest(context.Background(), joinRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.CancelAJoinRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**joinRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAJoinRequestRequest struct via the builder pattern


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


## CreateProject

> CreateProject201Response CreateProject(ctx).CreateProjectRequest(createProjectRequest).Execute()

Create project



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
	createProjectRequest := *openapiclient.NewCreateProjectRequest("TEST1568723795") // CreateProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.CreateProject(context.Background()).CreateProjectRequest(createProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: CreateProject201Response
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) |  | 

### Return type

[**CreateProject201Response**](CreateProject201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineAInvitation

> DeclineAInvitation(ctx, invitationId).Execute()

Decline a invitation



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.DeclineAInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.DeclineAInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineAInvitationRequest struct via the builder pattern


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


## ListMyJoinRequests

> []ListMyJoinRequests200ResponseInner ListMyJoinRequests(ctx).Execute()

List my join requests



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
	resp, r, err := apiClient.MyAccountAPI.ListMyJoinRequests(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.ListMyJoinRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyJoinRequests`: []ListMyJoinRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.ListMyJoinRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyJoinRequestsRequest struct via the builder pattern


### Return type

[**[]ListMyJoinRequests200ResponseInner**](ListMyJoinRequests200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyMaintainerInvitations

> []ListMyMaintainerInvitations200ResponseInner ListMyMaintainerInvitations(ctx).Execute()

List my maintainer invitations



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
	resp, r, err := apiClient.MyAccountAPI.ListMyMaintainerInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.ListMyMaintainerInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyMaintainerInvitations`: []ListMyMaintainerInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.ListMyMaintainerInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyMaintainerInvitationsRequest struct via the builder pattern


### Return type

[**[]ListMyMaintainerInvitations200ResponseInner**](ListMyMaintainerInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyOrganizationInvitations

> []ListMyOrganizationInvitations200ResponseInner ListMyOrganizationInvitations(ctx).Execute()

List my organization invitations



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
	resp, r, err := apiClient.MyAccountAPI.ListMyOrganizationInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.ListMyOrganizationInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyOrganizationInvitations`: []ListMyOrganizationInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.ListMyOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyOrganizationInvitationsRequest struct via the builder pattern


### Return type

[**[]ListMyOrganizationInvitations200ResponseInner**](ListMyOrganizationInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyProjectInvitations

> []ListMyProjectInvitations200ResponseInner ListMyProjectInvitations(ctx).Execute()

List my project invitations



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
	resp, r, err := apiClient.MyAccountAPI.ListMyProjectInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.ListMyProjectInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyProjectInvitations`: []ListMyProjectInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.ListMyProjectInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyProjectInvitationsRequest struct via the builder pattern


### Return type

[**[]ListMyProjectInvitations200ResponseInner**](ListMyProjectInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsedPromoCodes

> []ListUsedPromoCodes200ResponseInner ListUsedPromoCodes(ctx).Execute()

List used promo codes



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
	resp, r, err := apiClient.MyAccountAPI.ListUsedPromoCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.ListUsedPromoCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsedPromoCodes`: []ListUsedPromoCodes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.ListUsedPromoCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsedPromoCodesRequest struct via the builder pattern


### Return type

[**[]ListUsedPromoCodes200ResponseInner**](ListUsedPromoCodes200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyJoinRequestDetail

> RequestAccessToAProject201Response MyJoinRequestDetail(ctx, joinRequestId).Execute()

My join request detail



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
	joinRequestId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.MyJoinRequestDetail(context.Background(), joinRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.MyJoinRequestDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MyJoinRequestDetail`: RequestAccessToAProject201Response
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.MyJoinRequestDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**joinRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyJoinRequestDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestAccessToAProject201Response**](RequestAccessToAProject201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyMaintainerInvitationDetail

> ListMyMaintainerInvitations200ResponseInner MyMaintainerInvitationDetail(ctx, invitationId).Execute()

My maintainer invitation detail



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.MyMaintainerInvitationDetail(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.MyMaintainerInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MyMaintainerInvitationDetail`: ListMyMaintainerInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.MyMaintainerInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyMaintainerInvitationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMyMaintainerInvitations200ResponseInner**](ListMyMaintainerInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyOrganizationInvitationDetail

> ListMyOrganizationInvitations200ResponseInner MyOrganizationInvitationDetail(ctx, invitationId).Execute()

My organization invitation detail



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.MyOrganizationInvitationDetail(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.MyOrganizationInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MyOrganizationInvitationDetail`: ListMyOrganizationInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.MyOrganizationInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyOrganizationInvitationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMyOrganizationInvitations200ResponseInner**](ListMyOrganizationInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyProjectInvitationDetail

> ListMyProjectInvitations200ResponseInner MyProjectInvitationDetail(ctx, invitationId).Execute()

My project invitation detail



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
	invitationId := float32(8.14) // float32 | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.MyProjectInvitationDetail(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.MyProjectInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MyProjectInvitationDetail`: ListMyProjectInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.MyProjectInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **float32** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyProjectInvitationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMyProjectInvitations200ResponseInner**](ListMyProjectInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

