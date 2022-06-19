package test

import (
	"fmt"
	"github.com/d3code/cmc-client-go/pkg/cmc"
	"net/url"
	"testing"
)

var testCryptocurrencyClient = cmc.Client("").Test(true).CryptocurrencyClient()

func TestGetCryptocurrencyMap(t *testing.T) {
	response, err := testCryptocurrencyClient.GetCryptocurrencyMap(url.Values{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}
