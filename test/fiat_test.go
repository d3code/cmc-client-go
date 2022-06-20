package test

import (
	"net/url"
	"testing"
)

func TestFiatMap(t *testing.T) {
	response, err := testClient.GetFiatMap(url.Values{})
	if err != nil {
		t.Error(err)
	}

	checkResponse(t, response.Status)
}
