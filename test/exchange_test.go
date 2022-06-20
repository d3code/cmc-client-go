package test

import (
	"fmt"
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

	fmt.Println(response)
}

func TestGetExchangeMap(t *testing.T) {
	response, err := testClient.GetExchangeMap(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetExchangeListingsLatest(t *testing.T) {
	response, err := testClient.GetExchangeListingsLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetExchangeMarketPairsLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeMarketPairsLatest(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetExchangeQuotesHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeQuotesHistorical(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetExchangeQuotesLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetExchangeQuotesLatest(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}
