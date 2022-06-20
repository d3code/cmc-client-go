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

func (c *client) GetCryptocurrencyCategories(q url.Values) (*CryptocurrencyCategories, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/categories"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyCategories

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyCategory(q url.Values) (*CryptocurrencyCategory, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/category"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyCategory

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

func (c *client) GetCryptocurrencyInfo(q url.Values) (*CryptocurrencyInfo, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/info"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyInfo

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyListingsHistorical(q url.Values) (*CryptocurrencyListingsHistorical, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/listings/historical"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyListingsHistorical

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyListingsLatest(q url.Values) (*CryptocurrencyListingsLatest, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/listings/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyListingsLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyListingsNew(q url.Values) (*CryptocurrencyListingsNew, error) {
	requestURL := c.baseURL + "/v1/cryptocurrency/listings/new"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyListingsNew

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
