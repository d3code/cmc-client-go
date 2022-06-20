package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetExchangeInfo(q url.Values) (*ExchangeInfo, error) {
	requestURL := "/v1/exchange/info"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeInfo

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetExchangeMap(q url.Values) (*ExchangeMap, error) {
	requestURL := "/v1/exchange/map"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse ExchangeMap

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetExchangeListingsLatest(q url.Values) (*ExchangeListingsLatest, error) {
	requestURL := "/v1/exchange/listings/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeListingsLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetExchangeMarketPairsLatest(q url.Values) (*ExchangeMarketPairsLatest, error) {
	requestURL := "/v1/exchange/market-pairs/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeMarketPairsLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetExchangeQuotesHistorical(q url.Values) (*ExchangeQuotesHistorical, error) {
	requestURL := "/v1/exchange/quotes/historical"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeQuotesHistorical

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetExchangeQuotesLatest(q url.Values) (*ExchangeQuotesLatest, error) {
	requestURL := "/v1/exchange/quotes/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeQuotesLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
