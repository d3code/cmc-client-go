package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetCryptocurrencyAirdrop(t *testing.T) {
	values := url.Values{}
	values.Add("id", "60e59b99c8ca1d58514a2322")

	response, err := testClient.GetCryptocurrencyAirdrop(values)
	checkResponse(t, err, response.Status)

	fmt.Println(response)
}

func TestGetCryptocurrencyMap(t *testing.T) {
	response, err := testClient.GetCryptocurrencyMap(url.Values{})
	checkResponse(t, err, response.Status)

	fmt.Println(response)
}
