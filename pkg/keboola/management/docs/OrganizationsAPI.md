# \OrganizationsAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAUserToOrganization**](OrganizationsAPI.md#AddAUserToOrganization) | **Post** /manage/organizations/{organization_id}/users | Add a User to organization
[**CancelOrganizationInvitation**](OrganizationsAPI.md#CancelOrganizationInvitation) | **Delete** /manage/organizations/{organization_id}/invitations/{invitation_id} | Cancel organization invitation
[**CreateAnOrganization**](OrganizationsAPI.md#CreateAnOrganization) | **Post** /manage/maintainers/{maintainer_id}/organizations | Create an organization
[**DeleteAnOrganization**](OrganizationsAPI.md#DeleteAnOrganization) | **Delete** /manage/organizations/{organization_id} | Delete an organization
[**InviteAUserToAOrganization**](OrganizationsAPI.md#InviteAUserToAOrganization) | **Post** /manage/organizations/{organization_id}/invitations | Invite a user to a organization
[**JoinAOrganization**](OrganizationsAPI.md#JoinAOrganization) | **Post** /manage/organizations/{organization_id}/join-organization | Join a organization
[**ListMaintainersOrganizations**](OrganizationsAPI.md#ListMaintainersOrganizations) | **Get** /manage/maintainers/{maintainer_id}/organizations | List maintainer&#39;s organizations
[**ListOrganizationInvitations**](OrganizationsAPI.md#ListOrganizationInvitations) | **Get** /manage/organizations/{organization_id}/invitations | List organization invitations
[**ListOrganizationMetadata**](OrganizationsAPI.md#ListOrganizationMetadata) | **Get** /manage/organizations/{organization_id}/metadata | List organization Metadata
[**ListOrganizationUsers**](OrganizationsAPI.md#ListOrganizationUsers) | **Get** /manage/organizations/{organization_id}/users | List organization users
[**ListProjectUsersInOrganization**](OrganizationsAPI.md#ListProjectUsersInOrganization) | **Get** /manage/organizations/{organization_id}/projects-users | List project users in organization
[**OrganizationInvitationDetail**](OrganizationsAPI.md#OrganizationInvitationDetail) | **Get** /manage/organizations/{organization_id}/invitations/{invitation_id} | Organization invitation detail
[**RemoveAUserFromOrganization**](OrganizationsAPI.md#RemoveAUserFromOrganization) | **Delete** /manage/organizations/{organization_id}/users/{user_id} | Remove a user from organization
[**RemoveOrganizationMetadata**](OrganizationsAPI.md#RemoveOrganizationMetadata) | **Delete** /manage/organizations/{organization_id}/metadata/{metadata_id} | Remove organization metadata
[**RequireMultiFactorAuthentication**](OrganizationsAPI.md#RequireMultiFactorAuthentication) | **Post** /manage/organizations/{organization_id}/force-mfa | Require Multi-Factor Authentication
[**RetrieveAnOrganization**](OrganizationsAPI.md#RetrieveAnOrganization) | **Get** /manage/organizations/{organization_id} | Retrieve an organization
[**SetOrganizationMetadata**](OrganizationsAPI.md#SetOrganizationMetadata) | **Post** /manage/organizations/{organization_id}/metadata | Set organization metadata
[**UpdateAnOrganization**](OrganizationsAPI.md#UpdateAnOrganization) | **Patch** /manage/organizations/{organization_id} | Update an organization



## AddAUserToOrganization

> AddAUserToOrganization(ctx, organizationId).AddAUserToMaintainerRequest(addAUserToMaintainerRequest).Execute()

Add a User to organization



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
	organizationId := float32(1) // float32 | ID of the organization
	addAUserToMaintainerRequest := *openapiclient.NewAddAUserToMaintainerRequest() // AddAUserToMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AddAUserToOrganization(context.Background(), organizationId).AddAUserToMaintainerRequest(addAUserToMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddAUserToOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAUserToOrganizationRequest struct via the builder pattern


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


## CancelOrganizationInvitation

> CancelOrganizationInvitation(ctx, organizationId, invitationId).Execute()

Cancel organization invitation



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
	organizationId := "organizationId_example" // string | Organization ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.CancelOrganizationInvitation(context.Background(), organizationId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CancelOrganizationInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrganizationInvitationRequest struct via the builder pattern


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


## CreateAnOrganization

> OrganizationModel CreateAnOrganization(ctx, maintainerId).CreateAnOrganizationRequest(createAnOrganizationRequest).Execute()

Create an organization



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
	maintainerId := float32(1) // float32 | ID of the parent maintainer. The created organization will be assigned to this maintainer.
	createAnOrganizationRequest := *openapiclient.NewCreateAnOrganizationRequest() // CreateAnOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateAnOrganization(context.Background(), maintainerId).CreateAnOrganizationRequest(createAnOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateAnOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnOrganization`: OrganizationModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateAnOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **float32** | ID of the parent maintainer. The created organization will be assigned to this maintainer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAnOrganizationRequest** | [**CreateAnOrganizationRequest**](CreateAnOrganizationRequest.md) |  | 

### Return type

[**OrganizationModel**](OrganizationModel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnOrganization

> DeleteAnOrganization(ctx, organizationId).Execute()

Delete an organization



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
	organizationId := float32(1) // float32 | ID of the organization in form of an integer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteAnOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteAnOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization in form of an integer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnOrganizationRequest struct via the builder pattern


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


## InviteAUserToAOrganization

> InviteAUserToAMaintainer201Response InviteAUserToAOrganization(ctx, organizationId).InviteAUserToAMaintainerRequest(inviteAUserToAMaintainerRequest).Execute()

Invite a user to a organization



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
	organizationId := "organizationId_example" // string | Organization ID
	inviteAUserToAMaintainerRequest := *openapiclient.NewInviteAUserToAMaintainerRequest("martin@keboola.com") // InviteAUserToAMaintainerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.InviteAUserToAOrganization(context.Background(), organizationId).InviteAUserToAMaintainerRequest(inviteAUserToAMaintainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InviteAUserToAOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteAUserToAOrganization`: InviteAUserToAMaintainer201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.InviteAUserToAOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteAUserToAOrganizationRequest struct via the builder pattern


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


## JoinAOrganization

> JoinAOrganization(ctx, organizationId).Execute()

Join a organization



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
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.JoinAOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.JoinAOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinAOrganizationRequest struct via the builder pattern


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


## ListMaintainersOrganizations

> ListMaintainersOrganizations(ctx, maintainerId).Execute()

List maintainer's organizations



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
	maintainerId := float32(1) // float32 | ID of the parent maintainer. The created organization will be assigned to this maintainer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ListMaintainersOrganizations(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListMaintainersOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **float32** | ID of the parent maintainer. The created organization will be assigned to this maintainer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintainersOrganizationsRequest struct via the builder pattern


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


## ListOrganizationInvitations

> []ListMaintainersInvitations200ResponseInner ListOrganizationInvitations(ctx, organizationId).Execute()

List organization invitations



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
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationInvitations(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationInvitations`: []ListMaintainersInvitations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationInvitationsRequest struct via the builder pattern


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


## ListOrganizationMetadata

> []ListOrganizationMetadata200ResponseInner ListOrganizationMetadata(ctx, organizationId).Execute()

List organization Metadata



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
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationMetadata(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationMetadataRequest struct via the builder pattern


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


## ListOrganizationUsers

> []ListMaintainersInvitations200ResponseInnerUser ListOrganizationUsers(ctx, organizationId).Execute()

List organization users



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
	organizationId := float32(1) // float32 | ID of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationUsers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationUsers`: []ListMaintainersInvitations200ResponseInnerUser
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUsersRequest struct via the builder pattern


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


## ListProjectUsersInOrganization

> []ListProjectUsersInOrganization200ResponseInner ListProjectUsersInOrganization(ctx, organizationId).Execute()

List project users in organization



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
	organizationId := float32(1) // float32 | ID of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListProjectUsersInOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListProjectUsersInOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectUsersInOrganization`: []ListProjectUsersInOrganization200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListProjectUsersInOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectUsersInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListProjectUsersInOrganization200ResponseInner**](ListProjectUsersInOrganization200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationInvitationDetail

> InviteAUserToAMaintainer201Response OrganizationInvitationDetail(ctx, organizationId, invitationId).Execute()

Organization invitation detail



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
	organizationId := "organizationId_example" // string | Organization ID
	invitationId := "invitationId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationInvitationDetail(context.Background(), organizationId, invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationInvitationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationInvitationDetail`: InviteAUserToAMaintainer201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationInvitationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**invitationId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationInvitationDetailRequest struct via the builder pattern


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


## RemoveAUserFromOrganization

> RemoveAUserFromOrganization(ctx, organizationId, userId).Execute()

Remove a user from organization



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
	organizationId := float32(1) // float32 | ID of the organization
	userId := "1" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.RemoveAUserFromOrganization(context.Background(), organizationId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RemoveAUserFromOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAUserFromOrganizationRequest struct via the builder pattern


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


## RemoveOrganizationMetadata

> RemoveOrganizationMetadata(ctx, organizationId, metadataId).Execute()

Remove organization metadata



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
	organizationId := "organizationId_example" // string | Organization ID
	metadataId := "metadataId_example" // string | Metadata ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.RemoveOrganizationMetadata(context.Background(), organizationId, metadataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RemoveOrganizationMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**metadataId** | **string** | Metadata ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationMetadataRequest struct via the builder pattern


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


## RequireMultiFactorAuthentication

> RequireMultiFactorAuthentication(ctx, organizationId).Execute()

Require Multi-Factor Authentication



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
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.RequireMultiFactorAuthentication(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RequireMultiFactorAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequireMultiFactorAuthenticationRequest struct via the builder pattern


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


## RetrieveAnOrganization

> RetrieveAnOrganization200Response RetrieveAnOrganization(ctx, organizationId).Execute()

Retrieve an organization



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
	organizationId := float32(1) // float32 | ID of the organization in form of an integer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RetrieveAnOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RetrieveAnOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAnOrganization`: RetrieveAnOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RetrieveAnOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization in form of an integer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAnOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveAnOrganization200Response**](RetrieveAnOrganization200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrganizationMetadata

> []ListOrganizationMetadata200ResponseInner SetOrganizationMetadata(ctx, organizationId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()

Set organization metadata



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
	organizationId := "organizationId_example" // string | Organization ID
	setMaintainerMetadataRequest := *openapiclient.NewSetMaintainerMetadataRequest("Provider_example", []interface{}{nil}) // SetMaintainerMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.SetOrganizationMetadata(context.Background(), organizationId).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.SetOrganizationMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrganizationMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.SetOrganizationMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizationMetadataRequest struct via the builder pattern


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


## UpdateAnOrganization

> RetrieveAnOrganization200Response UpdateAnOrganization(ctx, organizationId).UpdateAnOrganizationRequest(updateAnOrganizationRequest).Execute()

Update an organization



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
	organizationId := float32(1) // float32 | ID of the organization in form of an integer
	updateAnOrganizationRequest := *openapiclient.NewUpdateAnOrganizationRequest() // UpdateAnOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateAnOrganization(context.Background(), organizationId).UpdateAnOrganizationRequest(updateAnOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateAnOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnOrganization`: RetrieveAnOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateAnOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **float32** | ID of the organization in form of an integer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAnOrganizationRequest** | [**UpdateAnOrganizationRequest**](UpdateAnOrganizationRequest.md) |  | 

### Return type

[**RetrieveAnOrganization200Response**](RetrieveAnOrganization200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

