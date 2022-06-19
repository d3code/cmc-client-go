package cmc

import (
	"encoding/json"
	"net/url"
)

type cryptocurrencyClient struct {
	client
}

func (c *client) CryptocurrencyClient() *cryptocurrencyClient {
	return &cryptocurrencyClient{*c}
}

func (c *cryptocurrencyClient) GetCryptocurrencyMap(q url.Values) (*CryptocurrencyMap, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/map"

	resBody, err := doGetRequest(requestURL, q, &c.client)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyMap

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
