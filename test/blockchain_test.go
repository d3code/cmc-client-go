package test

import (
    "net/url"
    "testing"
)

func TestBlockchainStatisticsLatest(t *testing.T) {
    values := url.Values{}
    values.Add("id", "1")

    response, err := testClient.GetBlockchainStatisticsLatest(values)
    if err != nil {
        t.Error(err)
    }

    checkResponse(t, response.Status)
}
