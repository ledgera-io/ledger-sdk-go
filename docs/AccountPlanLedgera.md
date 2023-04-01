# \AccountPlanLedgera

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountPlanBulkPost**](AccountPlanLedgera.md#V1AccountPlanBulkPost) | **Post** /v1/account-plan/bulk | Bulk create account plan
[**V1AccountPlanGet**](AccountPlanLedgera.md#V1AccountPlanGet) | **Get** /v1/account-plan | List account plan
[**V1AccountPlanIdDelete**](AccountPlanLedgera.md#V1AccountPlanIdDelete) | **Delete** /v1/account-plan/{id} | Delete account plan
[**V1AccountPlanIdGet**](AccountPlanLedgera.md#V1AccountPlanIdGet) | **Get** /v1/account-plan/{id} | Get account plan
[**V1AccountPlanIdPut**](AccountPlanLedgera.md#V1AccountPlanIdPut) | **Put** /v1/account-plan/{id} | Update account plan
[**V1AccountPlanPost**](AccountPlanLedgera.md#V1AccountPlanPost) | **Post** /v1/account-plan | Create account plan



## V1AccountPlanBulkPost

> HttpBulkCreateAccountPlan V1AccountPlanBulkPost(ctx).Id(id).Execute()

Bulk create account plan



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanBulkPost(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanBulkPost`: HttpBulkCreateAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**HttpBulkCreateAccountPlan**](HttpBulkCreateAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPlanGet

> []AccountPlanAccountPlan V1AccountPlanGet(ctx).Execute()

List account plan



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
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanGet`: []AccountPlanAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanGetRequest struct via the builder pattern


### Return type

[**[]AccountPlanAccountPlan**](AccountPlanAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPlanIdDelete

> AccountPlanAccountPlan V1AccountPlanIdDelete(ctx, id).Execute()

Delete account plan



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
    id := "id_example" // string | Plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanIdDelete`: AccountPlanAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountPlanAccountPlan**](AccountPlanAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPlanIdGet

> AccountPlanAccountPlan V1AccountPlanIdGet(ctx, id).Execute()

Get account plan



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
    id := "id_example" // string | Plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanIdGet`: AccountPlanAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountPlanAccountPlan**](AccountPlanAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPlanIdPut

> AccountPlanAccountPlan V1AccountPlanIdPut(ctx, id).Execute()

Update account plan



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
    id := "id_example" // string | Plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanIdPut`: AccountPlanAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountPlanAccountPlan**](AccountPlanAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPlanPost

> AccountPlanAccountPlan V1AccountPlanPost(ctx).Plan(plan).Execute()

Create account plan



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
    plan := *openapiclient.NewAccountPlanCreateAccountPlan(openapiclient.account_plan.BalanceType("available"), "BalanceValidation_example", "SubType_example", "Type_example") // AccountPlanCreateAccountPlan | Account plan JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountPlanLedgera.V1AccountPlanPost(context.Background()).Plan(plan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPlanLedgera.V1AccountPlanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPlanPost`: AccountPlanAccountPlan
    fmt.Fprintf(os.Stdout, "Response from `AccountPlanLedgera.V1AccountPlanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPlanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | [**AccountPlanCreateAccountPlan**](AccountPlanCreateAccountPlan.md) | Account plan JSON | 

### Return type

[**AccountPlanAccountPlan**](AccountPlanAccountPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

