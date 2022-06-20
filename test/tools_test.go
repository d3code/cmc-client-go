package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestToolsPriceConversion(t *testing.T) {
	values := url.Values{}
	values.Add("amount", "100")
	values.Add("symbol", "BTC")

	response, err := testClient.GetToolsPriceConversion(values)
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}
