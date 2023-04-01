# \BillingPlanLedgera

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BillingBillingPlanIdGet**](BillingPlanLedgera.md#V1BillingBillingPlanIdGet) | **Get** /v1/billing/{billingPlanId} | Get billing plan by ID
[**V1BillingBillingPlanIdPut**](BillingPlanLedgera.md#V1BillingBillingPlanIdPut) | **Put** /v1/billing/{billingPlanId} | Update billing plan
[**V1BillingBulkPost**](BillingPlanLedgera.md#V1BillingBulkPost) | **Post** /v1/billing/bulk | Create billing plan
[**V1BillingGet**](BillingPlanLedgera.md#V1BillingGet) | **Get** /v1/billing | List billing
[**V1BillingPost**](BillingPlanLedgera.md#V1BillingPost) | **Post** /v1/billing | Create billing plan



## V1BillingBillingPlanIdGet

> BillingBillingPlan V1BillingBillingPlanIdGet(ctx, billingPlanId).Execute()

Get billing plan by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    billingPlanId := "billingPlanId_example" // string | Billing plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPlanLedgera.V1BillingBillingPlanIdGet(context.Background(), billingPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPlanLedgera.V1BillingBillingPlanIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingBillingPlanIdGet`: BillingBillingPlan
    fmt.Fprintf(os.Stdout, "Response from `BillingPlanLedgera.V1BillingBillingPlanIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingPlanId** | **string** | Billing plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingBillingPlanIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingBillingPlan**](BillingBillingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingBillingPlanIdPut

> BillingBillingPlan V1BillingBillingPlanIdPut(ctx, billingPlanId).Execute()

Update billing plan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    billingPlanId := "billingPlanId_example" // string | Billing plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPlanLedgera.V1BillingBillingPlanIdPut(context.Background(), billingPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPlanLedgera.V1BillingBillingPlanIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingBillingPlanIdPut`: BillingBillingPlan
    fmt.Fprintf(os.Stdout, "Response from `BillingPlanLedgera.V1BillingBillingPlanIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingPlanId** | **string** | Billing plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingBillingPlanIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingBillingPlan**](BillingBillingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingBulkPost

> []BillingBillingPlan V1BillingBulkPost(ctx).Currency(currency).Execute()

Create billing plan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    currency := *openapiclient.NewBillingCreateBulkBilling([]openapiclient.BillingCreateBilling{*openapiclient.NewBillingCreateBilling([]openapiclient.BillingTransactionEntry{*openapiclient.NewBillingTransactionEntry()}, "Expiration_example", "Name_example")}) // BillingCreateBulkBilling | Billing JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPlanLedgera.V1BillingBulkPost(context.Background()).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPlanLedgera.V1BillingBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingBulkPost`: []BillingBillingPlan
    fmt.Fprintf(os.Stdout, "Response from `BillingPlanLedgera.V1BillingBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**BillingCreateBulkBilling**](BillingCreateBulkBilling.md) | Billing JSON | 

### Return type

[**[]BillingBillingPlan**](BillingBillingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingGet

> []BillingBillingPlan V1BillingGet(ctx).Execute()

List billing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPlanLedgera.V1BillingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPlanLedgera.V1BillingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingGet`: []BillingBillingPlan
    fmt.Fprintf(os.Stdout, "Response from `BillingPlanLedgera.V1BillingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingGetRequest struct via the builder pattern


### Return type

[**[]BillingBillingPlan**](BillingBillingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPost

> BillingBillingPlan V1BillingPost(ctx).Currency(currency).Execute()

Create billing plan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    currency := *openapiclient.NewBillingCreateBilling([]openapiclient.BillingTransactionEntry{*openapiclient.NewBillingTransactionEntry()}, "Expiration_example", "Name_example") // BillingCreateBilling | Billing JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPlanLedgera.V1BillingPost(context.Background()).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPlanLedgera.V1BillingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPost`: BillingBillingPlan
    fmt.Fprintf(os.Stdout, "Response from `BillingPlanLedgera.V1BillingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**BillingCreateBilling**](BillingCreateBilling.md) | Billing JSON | 

### Return type

[**BillingBillingPlan**](BillingBillingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

