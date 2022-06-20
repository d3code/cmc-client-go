package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetBlockchainStatisticsLatest(q url.Values) (*BlockchainStatisticsLatest, error) {
	requestURL := "/v1/blockchain/statistics/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse BlockchainStatisticsLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
