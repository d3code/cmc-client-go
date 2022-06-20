package test

import (
	"net/url"
	"testing"
	"time"
)

func TestGetCryptocurrencyAirdrop(t *testing.T) {
	values := url.Values{}
	values.Add("id", "60e59b99c8ca1d58514a2322")

	response, err := testClient.GetCryptocurrencyAirdrop(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyAirdrops(t *testing.T) {
	response, err := testClient.GetCryptocurrencyAirdrops(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyCategories(t *testing.T) {
	response, err := testClient.GetCryptocurrencyCategories(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyCategory(t *testing.T) {
	values := url.Values{}
	values.Add("id", "m55quflw0cj")

	response, err := testClient.GetCryptocurrencyCategory(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyMap(t *testing.T) {
	response, err := testClient.GetCryptocurrencyMap(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyInfo(t *testing.T) {
	values := url.Values{}
	values.Add("slug", "eth")

	response, err := testClient.GetCryptocurrencyInfo(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyListingsHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("date", time.Now().String())

	response, err := testClient.GetCryptocurrencyListingsHistorical(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyListingsLatest(t *testing.T) {
	response, err := testClient.GetCryptocurrencyListingsLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyListingsNew(t *testing.T) {
	response, err := testClient.GetCryptocurrencyListingsNew(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyTrendingGainersLosers(t *testing.T) {
	response, err := testClient.GetCryptocurrencyTrendingGainersLosers(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyTrendingLatest(t *testing.T) {
	response, err := testClient.GetCryptocurrencyTrendingLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyTrendingMostVisited(t *testing.T) {
	response, err := testClient.GetCryptocurrencyTrendingMostVisited(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyMarketPairs(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyMarketPairs(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyOHLCVHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyOHLCVHistorical(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyOHLCVLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyOHLCVLatest(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyPricePerformanceStats(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyPricePerformanceStats(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyQuotesHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyQuotesHistorical(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}

func TestGetCryptocurrencyQuotesLatest(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")

	response, err := testClient.GetCryptocurrencyQuotesLatest(values)
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}
