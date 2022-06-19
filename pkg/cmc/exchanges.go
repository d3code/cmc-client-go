package cmc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) GetExchanges() (*ExchangeResponseWrapper, error) {
	requestURL := c.baseURL + "/exchange/map"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("Could not create request: %s\n", err)
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

	respBodyString := string(resBody)
	fmt.Println(respBodyString)

	var cmcResponse ExchangeResponseWrapper
	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

type ExchangeResponseWrapper struct {
	Status map[string]interface{} `json:"status"`
	Data   []ExchangeResponse     `json:"data"`
}

type ExchangeResponse struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	IsActive            int    `json:"is_active"`
	FirstHistoricalData string `json:"first_historical_data"`
	LastHistoricalData  string `json:"last_historical_data"`
}
