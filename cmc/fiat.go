package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetFiatMap(q url.Values) (*Fiat, error) {
	requestURL := c.baseURL + "/v1/fiat/map"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse Fiat

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
