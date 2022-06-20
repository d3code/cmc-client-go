package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestFiatMap(t *testing.T) {
	response, err := testClient.GetFiatMap(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}
