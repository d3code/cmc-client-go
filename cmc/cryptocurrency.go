package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetCryptocurrencyAirdrop(q url.Values) (*CryptocurrencyAirdrop, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/airdrop"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyAirdrop

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyAirdrops(q url.Values) (*CryptocurrencyAirdrops, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/airdrops"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyAirdrops

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyMap(q url.Values) (*CryptocurrencyMap, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/map"

	resBody, err := doGetRequest(requestURL, q, c)
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
