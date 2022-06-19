package test

import (
	"fmt"
	"github.com/d3code/cmc-client-go/pkg/cmc"
	"net/url"
	"testing"
)

var testExchangeClient = cmc.Client("").Test(true).ExchangeClient()

func TestGetExchangeMap(t *testing.T) {
	response, err := testExchangeClient.GetExchangeMap(url.Values{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}

func TestGetExchangeInfo(t *testing.T) {
	values := url.Values{}
	values.Add("slug", "eth")

	response, err := testExchangeClient.GetExchangeInfo(values)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}
