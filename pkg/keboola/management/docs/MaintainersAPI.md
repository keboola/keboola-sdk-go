# \MaintainersAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAUserToMaintainer**](MaintainersAPI.md#AddAUserToMaintainer) | **Post** /manage/maintainers/{maintainer_id}/users | Add a User to maintainer
[**CancelMaintainerInvitation**](MaintainersAPI.md#CancelMaintainerInvitation) | **Delete** /manage/maintainers/{maintainer_id}/invitations/{invitation_id} | Cancel maintainer invitation
[**CreateAMaintainer**](MaintainersAPI.md#CreateAMaintainer) | **Post** /manage/maintainers | Create a maintainer
[**DeleteAMaintainer**](MaintainersAPI.md#DeleteAMaintainer) | **Delete** /manage/maintainers/{maintainer_id} | Delete a maintainer
[**InviteAUserToAMaintainer**](MaintainersAPI.md#InviteAUserToAMaintainer) | **Post** /manage/maintainers/{maintainer_id}/invitations | Invite a user to a maintainer
[**ListMaintainerMetadata**](MaintainersAPI.md#ListMaintainerMetadata) | **Get** /manage/maintainers/{maintainer_id}/metadata | List maintainer metadata
[**ListMaintainerUsers**](MaintainersAPI.md#ListMaintainerUsers) | **Get** /manage/maintainers/{maintainer_id}/users | List maintainer users
[**ListMaintainers**](MaintainersAPI.md#ListMaintainers) | **Get** /manage/maintainers | List maintainers
[**ListMaintainersInvitations**](MaintainersAPI.md#ListMaintainersInvitations) | **Get** /manage/maintainers/{maintainer_id}/invitations | List maintainers invitations
[**MaintainerInvitationDetail**](MaintainersAPI.md#MaintainerInvitationDetail) | **Get** /manage/maintainers/{maintainer_id}/invitations/{invitation_id} | Maintainer invitation detail
[**RemoveAUserFromMaintainer**](MaintainersAPI.md#RemoveAUserFromMaintainer) | **Delete** /manage/maintainers/{maintainer_id}/users/{user_id} | Remove a user from maintainer
[**RemoveMaintainerMetadata**](MaintainersAPI.md#RemoveMaintainerMetadata) | **Delete** /manage/maintainers/{maintainer_id}/metadata/{metadata_id} | Remove maintainer metadata
[**RetrieveAMaintainer**](MaintainersAPI.md#RetrieveAMaintainer) | **Get** /manage/maintainers/{maintainer_id} | Retrieve a maintainer
[**SetMaintainerMetadata**](MaintainersAPI.md#SetMaintainerMetadata) | **Post** /manage/maintainers/{maintainer_id}/metadata | Set maintainer metadata
[**UpdateAMaintainer**](MaintainersAPI.md#UpdateAMaintainer) | **Patch** /manage/maintainers/{maintainer_id} | Update a maintainer



## AddAUserToMaintainer

> AddAUserToMaintainer(ctx, maintainerId).AddAUserToMaintainerRequest(addAUserToMaintainerRequest).Execute()

Add a User to maintainer



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
	maintainerId := float32(1) // float32 | ID of the maintainer
	addAUserToMaintainerRequest := *openapiclient.NewAddAUserToMaintainerRequest() // AddAUserToMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.AddAUserToMaintainer(context.Background(), maintainerId).AddAUserToMaintainerRequest(addAUserToMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.AddAUserToMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **float32** | ID of the maintainer | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAUserToMaintainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAUserToMaintainerRequest** | [**AddAUserToMaintainerRequest**](AddAUserToMaintainerRequest.md) |  | 

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


## CancelMaintainerInvitation

> CancelMaintainerInvitation(ctx, maintainerId, invitationId).Execute()

Cancel maintainer invitation



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
	maintainerId := "maintainerId_example" // string | Maintainer ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.CancelMaintainerInvitation(context.Background(), maintainerId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.CancelMaintainerInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMaintainerInvitationRequest struct via the builder pattern


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


## CreateAMaintainer

> MaintainerModel CreateAMaintainer(ctx).CreateAMaintainerRequest(createAMaintainerRequest).Execute()

Create a maintainer



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
	createAMaintainerRequest := *openapiclient.NewCreateAMaintainerRequest("Example") // CreateAMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.CreateAMaintainer(context.Background()).CreateAMaintainerRequest(createAMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.CreateAMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAMaintainer`: MaintainerModel
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.CreateAMaintainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMaintainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAMaintainerRequest** | [**CreateAMaintainerRequest**](CreateAMaintainerRequest.md) |  | 

### Return type

[**MaintainerModel**](MaintainerModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAMaintainer

> DeleteAMaintainer(ctx, maintainerId).Execute()

Delete a maintainer



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
	maintainerId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.DeleteAMaintainer(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.DeleteAMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMaintainerRequest struct via the builder pattern


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


## InviteAUserToAMaintainer

> InviteAUserToAMaintainer201Response InviteAUserToAMaintainer(ctx, maintainerId).InviteAUserToAMaintainerRequest(inviteAUserToAMaintainerRequest).Execute()

Invite a user to a maintainer



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
	maintainerId := "maintainerId_example" // string | Maintainer ID
	inviteAUserToAMaintainerRequest := *openapiclient.NewInviteAUserToAMaintainerRequest("martin@keboola.com") // InviteAUserToAMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.InviteAUserToAMaintainer(context.Background(), maintainerId).InviteAUserToAMaintainerRequest(inviteAUserToAMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.InviteAUserToAMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteAUserToAMaintainer`: InviteAUserToAMaintainer201Response
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.InviteAUserToAMaintainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteAUserToAMaintainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteAUserToAMaintainerRequest** | [**InviteAUserToAMaintainerRequest**](InviteAUserToAMaintainerRequest.md) |  | 

### Return type

[**InviteAUserToAMaintainer201Response**](InviteAUserToAMaintainer201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMaintainerMetadata

> ListMaintainerMetadata(ctx, maintainerId).Execute()

List maintainer metadata



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
	maintainerId := "maintainerId_example" // string | Maintainer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.ListMaintainerMetadata(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.ListMaintainerMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintainerMetadataRequest struct via the builder pattern


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


## ListMaintainerUsers

> []ListMaintainersInvitations200ResponseInnerUser ListMaintainerUsers(ctx, maintainerId).Execute()

List maintainer users



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
	maintainerId := float32(1) // float32 | ID of the maintainer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.ListMaintainerUsers(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.ListMaintainerUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMaintainerUsers`: []ListMaintainersInvitations200ResponseInnerUser
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.ListMaintainerUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **float32** | ID of the maintainer | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintainerUsersRequest struct via the builder pattern


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


## ListMaintainers

> ListMaintainers(ctx).Execute()

List maintainers



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
	r, err := apiClient.MaintainersAPI.ListMaintainers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.ListMaintainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintainersRequest struct via the builder pattern


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


## ListMaintainersInvitations

> []ListMaintainersInvitations200ResponseInner ListMaintainersInvitations(ctx, maintainerId).Execute()

List maintainers invitations



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
	maintainerId := "maintainerId_example" // string | Maintainer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.ListMaintainersInvitations(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.ListMaintainersInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMaintainersInvitations`: []ListMaintainersInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.ListMaintainersInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintainersInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMaintainersInvitations200ResponseInner**](ListMaintainersInvitations200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintainerInvitationDetail

> InviteAUserToAMaintainer201Response MaintainerInvitationDetail(ctx, maintainerId, invitationId).Execute()

Maintainer invitation detail



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
	maintainerId := "maintainerId_example" // string | Maintainer ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.MaintainerInvitationDetail(context.Background(), maintainerId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.MaintainerInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintainerInvitationDetail`: InviteAUserToAMaintainer201Response
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.MaintainerInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintainerInvitationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InviteAUserToAMaintainer201Response**](InviteAUserToAMaintainer201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAUserFromMaintainer

> RemoveAUserFromMaintainer(ctx, maintainerId, userId).Execute()

Remove a user from maintainer



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
	maintainerId := float32(1) // float32 | ID of the maintainer
	userId := "1" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.RemoveAUserFromMaintainer(context.Background(), maintainerId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.RemoveAUserFromMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **float32** | ID of the maintainer | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAUserFromMaintainerRequest struct via the builder pattern


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


## RemoveMaintainerMetadata

> RemoveMaintainerMetadata(ctx, maintainerId, metadataId).Execute()

Remove maintainer metadata



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
	maintainerId := "maintainerId_example" // string | Maintainer ID
	metadataId := "metadataId_example" // string | Metadata ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.RemoveMaintainerMetadata(context.Background(), maintainerId, metadataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.RemoveMaintainerMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 
**metadataId** | **string** | Metadata ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaintainerMetadataRequest struct via the builder pattern


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


## RetrieveAMaintainer

> MaintainerModel RetrieveAMaintainer(ctx, maintainerId).Execute()

Retrieve a maintainer



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
	maintainerId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.RetrieveAMaintainer(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.RetrieveAMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAMaintainer`: MaintainerModel
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.RetrieveAMaintainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAMaintainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintainerModel**](MaintainerModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaintainerMetadata

> SetMaintainerMetadata(ctx, maintainerId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()

Set maintainer metadata



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
	maintainerId := "maintainerId_example" // string | Maintainer ID
	setMaintainerMetadataRequest := *openapiclient.NewSetMaintainerMetadataRequest("Provider_example", []interface{}{nil}) // SetMaintainerMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintainersAPI.SetMaintainerMetadata(context.Background(), maintainerId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.SetMaintainerMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **string** | Maintainer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaintainerMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMaintainerMetadataRequest** | [**SetMaintainerMetadataRequest**](SetMaintainerMetadataRequest.md) |  | 

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


## UpdateAMaintainer

> MaintainerModel UpdateAMaintainer(ctx, maintainerId).UpdateAMaintainerRequest(updateAMaintainerRequest).Execute()

Update a maintainer



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
	maintainerId := int32(1) // int32 | 
	updateAMaintainerRequest := *openapiclient.NewUpdateAMaintainerRequest() // UpdateAMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintainersAPI.UpdateAMaintainer(context.Background(), maintainerId).UpdateAMaintainerRequest(updateAMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintainersAPI.UpdateAMaintainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAMaintainer`: MaintainerModel
	fmt.Fprintf(os.Stdout, "Response from `MaintainersAPI.UpdateAMaintainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAMaintainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAMaintainerRequest** | [**UpdateAMaintainerRequest**](UpdateAMaintainerRequest.md) |  | 

### Return type

[**MaintainerModel**](MaintainerModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

