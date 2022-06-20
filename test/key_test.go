package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetKeyInfo(t *testing.T) {
	response, err := testClient.GetKeyInfo(url.Values{})
	checkResponse(t, err, response.Status)

	fmt.Println(response)
}
