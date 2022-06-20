package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetKeyInfo(t *testing.T) {
	response, err := testClient.GetKeyInfo(url.Values{})
	if err != nil {
		t.Error(err)
	}
	checkResponse(t, response.Status)

	fmt.Println(response)
}
