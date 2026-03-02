# \UsersAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableMFAForUser**](UsersAPI.md#DisableMFAForUser) | **Delete** /manage/users/{user_id_or_email}/mfa | Disable MFA for User
[**ListUserMetadata**](UsersAPI.md#ListUserMetadata) | **Get** /manage/users/{user_id_or_email}/metadata | List user Metadata
[**RemoveSuperAdminPrivilegeFromUser**](UsersAPI.md#RemoveSuperAdminPrivilegeFromUser) | **Delete** /manage/users/{user_id_or_email}/super-admin | Remove super admin privilege from User
[**RemoveUser**](UsersAPI.md#RemoveUser) | **Delete** /manage/users/{user_id_or_email} | Remove User
[**RemoveUserMetadata**](UsersAPI.md#RemoveUserMetadata) | **Delete** /manage/users/{user_id_or_email}/metadata/{metadata_id} | Remove User Metadata
[**SetUserMetadata**](UsersAPI.md#SetUserMetadata) | **Post** /manage/users/{user_id_or_email}/metadata | Set user metadata
[**UpdateAUser**](UsersAPI.md#UpdateAUser) | **Put** /manage/users/{user_id_or_email} | Update a user
[**UserDetail**](UsersAPI.md#UserDetail) | **Get** /manage/users/{user_id_or_email} | User detail



## DisableMFAForUser

> DisableMFAForUser(ctx, userIdOrEmail).Execute()

Disable MFA for User



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DisableMFAForUser(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DisableMFAForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableMFAForUserRequest struct via the builder pattern


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


## ListUserMetadata

> []ListOrganizationMetadata200ResponseInner ListUserMetadata(ctx, userIdOrEmail).Execute()

List user Metadata



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
	userIdOrEmail := "userIdOrEmail_example" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserMetadata(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMetadataRequest struct via the builder pattern


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


## RemoveSuperAdminPrivilegeFromUser

> UserDetail200Response RemoveSuperAdminPrivilegeFromUser(ctx, userIdOrEmail).Execute()

Remove super admin privilege from User



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RemoveSuperAdminPrivilegeFromUser(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveSuperAdminPrivilegeFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSuperAdminPrivilegeFromUser`: UserDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveSuperAdminPrivilegeFromUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSuperAdminPrivilegeFromUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDetail200Response**](UserDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUser

> RemoveUser(ctx, userIdOrEmail).Execute()

Remove User



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.RemoveUser(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserRequest struct via the builder pattern


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


## RemoveUserMetadata

> RemoveUserMetadata(ctx, userIdOrEmail, metadataId).Execute()

Remove User Metadata



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
	userIdOrEmail := "userIdOrEmail_example" // string | User ID
	metadataId := "metadataId_example" // string | Metadata ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.RemoveUserMetadata(context.Background(), userIdOrEmail, metadataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveUserMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID | 
**metadataId** | **string** | Metadata ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserMetadataRequest struct via the builder pattern


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


## SetUserMetadata

> []ListOrganizationMetadata200ResponseInner SetUserMetadata(ctx, userIdOrEmail).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()

Set user metadata



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
	userIdOrEmail := "userIdOrEmail_example" // string | Project ID
	setMaintainerMetadataRequest := *openapiclient.NewSetMaintainerMetadataRequest("Provider_example", []interface{}{nil}) // SetMaintainerMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SetUserMetadata(context.Background(), userIdOrEmail).SetMaintainerMetadataRequest(setMaintainerMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SetUserMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserMetadata`: []ListOrganizationMetadata200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SetUserMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserMetadataRequest struct via the builder pattern


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


## UpdateAUser

> UserDetail200Response UpdateAUser(ctx, userIdOrEmail).UpdateAUserRequest(updateAUserRequest).Execute()

Update a user



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
	updateAUserRequest := *openapiclient.NewUpdateAUserRequest() // UpdateAUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateAUser(context.Background(), userIdOrEmail).UpdateAUserRequest(updateAUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateAUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAUser`: UserDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateAUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAUserRequest** | [**UpdateAUserRequest**](UpdateAUserRequest.md) |  | 

### Return type

[**UserDetail200Response**](UserDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDetail

> UserDetail200Response UserDetail(ctx, userIdOrEmail).Execute()

User detail



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserDetail(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDetail`: UserDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | User ID or email | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDetail200Response**](UserDetail200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

