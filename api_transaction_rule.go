/*
Ledgera core API

Ledgera servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// TransactionRuleLedgeraService TransactionRuleLedgera service
type TransactionRuleLedgeraService service

type ApiV1TransactionRuleBulkPostRequest struct {
	ctx context.Context
	ApiService *TransactionRuleLedgeraService
	transaction *[]LedgerTransactionRule
}

// Transaction rule JSON
func (r ApiV1TransactionRuleBulkPostRequest) Transaction(transaction []LedgerTransactionRule) ApiV1TransactionRuleBulkPostRequest {
	r.transaction = &transaction
	return r
}

func (r ApiV1TransactionRuleBulkPostRequest) Execute() ([]LedgerTransactionRule, *http.Response, error) {
	return r.ApiService.V1TransactionRuleBulkPostExecute(r)
}

/*
V1TransactionRuleBulkPost Create transaction rule

Create transaction rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1TransactionRuleBulkPostRequest
*/
func (a *TransactionRuleLedgeraService) V1TransactionRuleBulkPost(ctx context.Context) ApiV1TransactionRuleBulkPostRequest {
	return ApiV1TransactionRuleBulkPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LedgerTransactionRule
func (a *TransactionRuleLedgeraService) V1TransactionRuleBulkPostExecute(r ApiV1TransactionRuleBulkPostRequest) ([]LedgerTransactionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LedgerTransactionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionRuleLedgeraService.V1TransactionRuleBulkPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transaction/rule/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transaction == nil {
		return localVarReturnValue, nil, reportError("transaction is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transaction
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1TransactionRuleGetRequest struct {
	ctx context.Context
	ApiService *TransactionRuleLedgeraService
}

func (r ApiV1TransactionRuleGetRequest) Execute() ([]LedgerTransactionRule, *http.Response, error) {
	return r.ApiService.V1TransactionRuleGetExecute(r)
}

/*
V1TransactionRuleGet List transaction rule

List transaction rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1TransactionRuleGetRequest
*/
func (a *TransactionRuleLedgeraService) V1TransactionRuleGet(ctx context.Context) ApiV1TransactionRuleGetRequest {
	return ApiV1TransactionRuleGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LedgerTransactionRule
func (a *TransactionRuleLedgeraService) V1TransactionRuleGetExecute(r ApiV1TransactionRuleGetRequest) ([]LedgerTransactionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LedgerTransactionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionRuleLedgeraService.V1TransactionRuleGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transaction/rule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1TransactionRulePostRequest struct {
	ctx context.Context
	ApiService *TransactionRuleLedgeraService
	transaction *LedgerTransactionRuleParams
}

// Transaction rule JSON
func (r ApiV1TransactionRulePostRequest) Transaction(transaction LedgerTransactionRuleParams) ApiV1TransactionRulePostRequest {
	r.transaction = &transaction
	return r
}

func (r ApiV1TransactionRulePostRequest) Execute() (*LedgerTransactionRule, *http.Response, error) {
	return r.ApiService.V1TransactionRulePostExecute(r)
}

/*
V1TransactionRulePost Create transaction rule

Create transaction rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1TransactionRulePostRequest
*/
func (a *TransactionRuleLedgeraService) V1TransactionRulePost(ctx context.Context) ApiV1TransactionRulePostRequest {
	return ApiV1TransactionRulePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LedgerTransactionRule
func (a *TransactionRuleLedgeraService) V1TransactionRulePostExecute(r ApiV1TransactionRulePostRequest) (*LedgerTransactionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LedgerTransactionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionRuleLedgeraService.V1TransactionRulePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transaction/rule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transaction == nil {
		return localVarReturnValue, nil, reportError("transaction is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transaction
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
