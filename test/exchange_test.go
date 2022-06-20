package test

import (
	"net/url"
	"testing"
)

func TestGetExchangeInfo(t *testing.T) {
	values := url.Values{}
	values.Add("slug", "eth")

	response, err := testClient.GetExchangeInfo(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetExchangeMap(t *testing.T) {
	response, err := testClient.GetExchangeMap(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetExchangeListingsLatest(t *testing.T) {
	response, err := testClient.GetExchangeListingsLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetExchangeMarketPairsLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeMarketPairsLatest(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetExchangeQuotesHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeQuotesHistorical(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetExchangeQuotesLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeQuotesLatest(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}
