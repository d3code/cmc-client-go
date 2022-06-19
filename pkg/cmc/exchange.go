package cmc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type exchangeClient struct {
	client
}

func (c *client) ExchangeClient() *exchangeClient {
	exchanges := exchangeClient{*c}
	return &exchanges
}

func (c *exchangeClient) GetExchangeMap() (*ExchangeMapResponseWrapper, error) {
	requestURL := c.baseURL + "/v1/exchange/map"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, requestError := client.Do(req)
	if requestError != nil {
		return nil, requestError
	}

	resBody, responseError := ioutil.ReadAll(res.Body)
	if responseError != nil {
		return nil, responseError
	}

	var cmcResponse ExchangeMapResponseWrapper
	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
