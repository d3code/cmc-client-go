package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetExchangeMap(q url.Values) (*ExchangeMapResponse, error) {
	requestURL := "/v1/exchange/map"

	resBody, err := doGetRequest(requestURL, q, c)
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

func (c *client) GetExchangeInfo(q url.Values) (*ExchangeInfoResponseWrapper, error) {
	requestURL := "/v1/exchange/info"

	resBody, err := doGetRequest(requestURL, q, c)
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
