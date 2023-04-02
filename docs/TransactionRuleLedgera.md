# \TransactionRuleLedgera

All URIs are relative to *https://api.ledgera*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TransactionRuleBulkPost**](TransactionRuleLedgera.md#V1TransactionRuleBulkPost) | **Post** /v1/transaction/rule/bulk | Create transaction rule
[**V1TransactionRuleGet**](TransactionRuleLedgera.md#V1TransactionRuleGet) | **Get** /v1/transaction/rule | List transaction rule
[**V1TransactionRulePost**](TransactionRuleLedgera.md#V1TransactionRulePost) | **Post** /v1/transaction/rule | Create transaction rule



## V1TransactionRuleBulkPost

> []LedgerTransactionRule V1TransactionRuleBulkPost(ctx).Transaction(transaction).Execute()

Create transaction rule



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
    transaction := []openapiclient.LedgerTransactionRule{*openapiclient.NewLedgerTransactionRule()} // []LedgerTransactionRule | Transaction rule JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRuleLedgera.V1TransactionRuleBulkPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.V1TransactionRuleBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionRuleBulkPost`: []LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.V1TransactionRuleBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionRuleBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**[]LedgerTransactionRule**](LedgerTransactionRule.md) | Transaction rule JSON | 

### Return type

[**[]LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransactionRuleGet

> []LedgerTransactionRule V1TransactionRuleGet(ctx).Execute()

List transaction rule



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
    resp, r, err := apiClient.TransactionRuleLedgera.V1TransactionRuleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.V1TransactionRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionRuleGet`: []LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.V1TransactionRuleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionRuleGetRequest struct via the builder pattern


### Return type

[**[]LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransactionRulePost

> LedgerTransactionRule V1TransactionRulePost(ctx).Transaction(transaction).Execute()

Create transaction rule



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
    transaction := *openapiclient.NewLedgerTransactionRuleParams("CreditAccountBalanceType_example", "CreditAccountSubType_example", "CreditAccountType_example", "DebitAccountBalanceType_example", "DebitAccountSubType_example", "DebitAccountType_example", int32(123), "EntryType_example", openapiclient.ledger.MessageType("single"), openapiclient.ledger.ProcessType("execution"), "TransactionType_example") // LedgerTransactionRuleParams | Transaction rule JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRuleLedgera.V1TransactionRulePost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.V1TransactionRulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionRulePost`: LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.V1TransactionRulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionRulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**LedgerTransactionRuleParams**](LedgerTransactionRuleParams.md) | Transaction rule JSON | 

### Return type

[**LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

