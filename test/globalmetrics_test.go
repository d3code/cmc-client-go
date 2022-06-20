package test

import (
	"net/url"
	"testing"
)

func TestGlobalMetricsQuotesHistorical(t *testing.T) {
	response, err := testClient.GetGlobalMetricsQuotesHistorical(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGlobalMetricsQuotesLatest(t *testing.T) {
	response, err := testClient.GetGlobalMetricsQuotesLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}
