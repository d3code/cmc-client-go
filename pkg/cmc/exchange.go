package cmc

import (
	"encoding/json"
	"net/url"
)

type exchangeClient struct {
	client
}

func (c *client) ExchangeClient() *exchangeClient {
	exchanges := exchangeClient{*c}
	return &exchanges
}

func (c *exchangeClient) GetExchangeMap(q url.Values) (*ExchangeMapResponse, error) {
	requestURL := c.baseURL + "/v1/exchange/map"

	resBody, err := doGetRequest(requestURL, q, &c.client)
	if err != nil {
		return nil, err
	}

	var cmcResponse ExchangeMapResponse

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *exchangeClient) GetExchangeInfo(q url.Values) (*ExchangeInfoResponseWrapper, error) {
	requestURL := c.baseURL + "/v1/exchange/info"

	resBody, err := doGetRequest(requestURL, q, &c.client)
	if err != nil {
		return nil, err
	}
	var cmcResponse ExchangeInfoResponseWrapper

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
