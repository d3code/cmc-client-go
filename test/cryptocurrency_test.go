package test

import (
	"fmt"
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

	fmt.Println(response)
}

func TestGetCryptocurrencyAirdrops(t *testing.T) {
	response, err := testClient.GetCryptocurrencyAirdrops(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyCategories(t *testing.T) {
	response, err := testClient.GetCryptocurrencyCategories(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyCategory(t *testing.T) {
	values := url.Values{}
	values.Add("id", "m55quflw0cj")

	response, err := testClient.GetCryptocurrencyCategory(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyMap(t *testing.T) {
	response, err := testClient.GetCryptocurrencyMap(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyInfo(t *testing.T) {
	values := url.Values{}
	values.Add("slug", "eth")

	response, err := testClient.GetCryptocurrencyInfo(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyListingsHistorical(t *testing.T) {
	values := url.Values{}
	values.Add("date", time.Now().String())

	response, err := testClient.GetCryptocurrencyListingsHistorical(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyListingsLatest(t *testing.T) {
	response, err := testClient.GetCryptocurrencyListingsLatest(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyListingsNew(t *testing.T) {
	response, err := testClient.GetCryptocurrencyListingsNew(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}
