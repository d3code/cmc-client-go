package test

import (
	"net/url"
	"testing"
)

func TestGetKeyInfo(t *testing.T) {
    response, err := testClient.GetKeyInfo(url.Values{})
    if err != nil {
        t.Error(err)
    }
    checkResponse(t, response.Status)
}
