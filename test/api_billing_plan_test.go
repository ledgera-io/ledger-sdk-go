/*
ledgera

Testing BillingPlanLedgeraService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_BillingPlanLedgeraService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BillingPlanLedgeraService V1BillingBillingPlanIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var billingPlanId string

		resp, httpRes, err := apiClient.BillingPlanLedgera.V1BillingBillingPlanIdGet(context.Background(), billingPlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingPlanLedgeraService V1BillingBillingPlanIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var billingPlanId string

		resp, httpRes, err := apiClient.BillingPlanLedgera.V1BillingBillingPlanIdPut(context.Background(), billingPlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingPlanLedgeraService V1BillingBulkPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BillingPlanLedgera.V1BillingBulkPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingPlanLedgeraService V1BillingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BillingPlanLedgera.V1BillingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingPlanLedgeraService V1BillingPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BillingPlanLedgera.V1BillingPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
