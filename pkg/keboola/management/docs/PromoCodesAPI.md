# \PromoCodesAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewPromoCodes**](PromoCodesAPI.md#CreateNewPromoCodes) | **Post** /manage/maintainers/{maintainer_id}/promo-codes | Create new promo codes
[**RetrieveAllPromoCodes**](PromoCodesAPI.md#RetrieveAllPromoCodes) | **Get** /manage/maintainers/{maintainer_id}/promo-codes | Retrieve all promo codes



## CreateNewPromoCodes

> CreateNewPromoCodes201Response CreateNewPromoCodes(ctx, maintainerId).CreateNewPromoCodesRequest(createNewPromoCodesRequest).Execute()

Create new promo codes



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
	maintainerId := int32(1) // int32 | 
	createNewPromoCodesRequest := *openapiclient.NewCreateNewPromoCodesRequest("TEST", float32(10), float32(1), "poc15DaysGuideMode") // CreateNewPromoCodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCodesAPI.CreateNewPromoCodes(context.Background(), maintainerId).CreateNewPromoCodesRequest(createNewPromoCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCodesAPI.CreateNewPromoCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewPromoCodes`: CreateNewPromoCodes201Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCodesAPI.CreateNewPromoCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewPromoCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNewPromoCodesRequest** | [**CreateNewPromoCodesRequest**](CreateNewPromoCodesRequest.md) |  | 

### Return type

[**CreateNewPromoCodes201Response**](CreateNewPromoCodes201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllPromoCodes

> []RetrieveAllPromoCodes200ResponseInner RetrieveAllPromoCodes(ctx, maintainerId).Execute()

Retrieve all promo codes



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
	maintainerId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCodesAPI.RetrieveAllPromoCodes(context.Background(), maintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCodesAPI.RetrieveAllPromoCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllPromoCodes`: []RetrieveAllPromoCodes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PromoCodesAPI.RetrieveAllPromoCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maintainerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllPromoCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RetrieveAllPromoCodes200ResponseInner**](RetrieveAllPromoCodes200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

