# \AccountLedgera

All URIs are relative to *https://api.ledgera*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountIdBalancesGet**](AccountLedgera.md#V1AccountIdBalancesGet) | **Get** /v1/account/{id}/balances | Get a account balances
[**V1AccountIdGet**](AccountLedgera.md#V1AccountIdGet) | **Get** /v1/account/{id} | Get a account from id
[**V1AccountIdJournalGet**](AccountLedgera.md#V1AccountIdJournalGet) | **Get** /v1/account/{id}/journal | Get a account journal
[**V1AccountPost**](AccountLedgera.md#V1AccountPost) | **Post** /v1/account | Create a account
[**V1CurrencyGet**](AccountLedgera.md#V1CurrencyGet) | **Get** /v1/currency | Get a ledger currencies
[**V1CurrencyPost**](AccountLedgera.md#V1CurrencyPost) | **Post** /v1/currency | Get a ledger currencies



## V1AccountIdBalancesGet

> HelperPaginatedArrayLedgerBalance V1AccountIdBalancesGet(ctx, id).Execute()

Get a account balances



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
    id := "id_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountLedgera.V1AccountIdBalancesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1AccountIdBalancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountIdBalancesGet`: HelperPaginatedArrayLedgerBalance
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1AccountIdBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountIdBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelperPaginatedArrayLedgerBalance**](HelperPaginatedArrayLedgerBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountIdGet

> HelperPaginatedLedgerAccount V1AccountIdGet(ctx, id).Execute()

Get a account from id



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
    id := "id_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountLedgera.V1AccountIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1AccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountIdGet`: HelperPaginatedLedgerAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1AccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelperPaginatedLedgerAccount**](HelperPaginatedLedgerAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountIdJournalGet

> HelperPaginatedArrayLedgerJournalEntry V1AccountIdJournalGet(ctx, id).Cursor(cursor).DateFrom(dateFrom).DateTo(dateTo).Type_(type_).Limit(limit).Execute()

Get a account journal



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
    id := "id_example" // string | Account ID
    cursor := "cursor_example" // string |  (optional)
    dateFrom := "dateFrom_example" // string |  (optional)
    dateTo := "dateTo_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountLedgera.V1AccountIdJournalGet(context.Background(), id).Cursor(cursor).DateFrom(dateFrom).DateTo(dateTo).Type_(type_).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1AccountIdJournalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountIdJournalGet`: HelperPaginatedArrayLedgerJournalEntry
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1AccountIdJournalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountIdJournalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **type_** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**HelperPaginatedArrayLedgerJournalEntry**](HelperPaginatedArrayLedgerJournalEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPost

> LedgerAccount V1AccountPost(ctx).Account(account).Execute()

Create a account



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
    account := *openapiclient.NewHttpRegisterAccount([]string{"Currencies_example"}, "SubType_example", "Type_example") // HttpRegisterAccount | Account JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountLedgera.V1AccountPost(context.Background()).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1AccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountPost`: LedgerAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1AccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**HttpRegisterAccount**](HttpRegisterAccount.md) | Account JSON | 

### Return type

[**LedgerAccount**](LedgerAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CurrencyGet

> HelperPaginatedArrayLedgerCurrency V1CurrencyGet(ctx).Execute()

Get a ledger currencies



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
    resp, r, err := apiClient.AccountLedgera.V1CurrencyGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1CurrencyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CurrencyGet`: HelperPaginatedArrayLedgerCurrency
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1CurrencyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CurrencyGetRequest struct via the builder pattern


### Return type

[**HelperPaginatedArrayLedgerCurrency**](HelperPaginatedArrayLedgerCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CurrencyPost

> LedgerCurrency V1CurrencyPost(ctx).Currency(currency).Execute()

Get a ledger currencies



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
    currency := *openapiclient.NewHttpCurrencyParams("Code_example", "Name_example") // HttpCurrencyParams | Currency JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountLedgera.V1CurrencyPost(context.Background()).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountLedgera.V1CurrencyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CurrencyPost`: LedgerCurrency
    fmt.Fprintf(os.Stdout, "Response from `AccountLedgera.V1CurrencyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CurrencyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**HttpCurrencyParams**](HttpCurrencyParams.md) | Currency JSON | 

### Return type

[**LedgerCurrency**](LedgerCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

