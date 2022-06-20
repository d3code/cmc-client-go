package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetExchangeMap(t *testing.T) {
	response, err := testClient.GetExchangeMap(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}

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
