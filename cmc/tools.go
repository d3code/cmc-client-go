package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetToolsPriceConversion(q url.Values) (*ToolsPriceConversion, error) {
	requestURL := "/v2/tools/price-conversion"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse ToolsPriceConversion

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
