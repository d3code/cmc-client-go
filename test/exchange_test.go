package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetExchangeMap(t *testing.T) {
	response, err := testClient.GetExchangeMap(url.Values{})
	checkResponse(t, err, response.Status)

	fmt.Println(response)
}

func TestGetExchangeInfo(t *testing.T) {
	values := url.Values{}
	values.Add("slug", "eth")

	response, err := testClient.GetExchangeInfo(values)
	checkResponse(t, err, response.Status)

	fmt.Println(response)
}
