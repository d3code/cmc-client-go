package cmc

import (
	"encoding/json"
	"net/url"
)

func (c *client) GetCryptocurrencyAirdrop(q url.Values) (*CryptocurrencyAirdrop, error) {
	requestURL := "/v1/cryptocurrency/airdrop"

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
	requestURL := "/v1/cryptocurrency/airdrops"

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
	requestURL := "/v1/cryptocurrency/categories"

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
	requestURL := "/v1/cryptocurrency/category"

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
	requestURL := "/v1/cryptocurrency/map"

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
	requestURL := "/v1/cryptocurrency/info"

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
	requestURL := "/v1/cryptocurrency/listings/historical"

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
	requestURL := "/v1/cryptocurrency/listings/latest"

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
	requestURL := "/v1/cryptocurrency/listings/new"

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

func (c *client) GetCryptocurrencyTrendingGainersLosers(q url.Values) (*CryptocurrencyTrendingGainersLosers, error) {
	requestURL := "/v1/cryptocurrency/trending/gainers-losers"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyTrendingGainersLosers

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyTrendingLatest(q url.Values) (*CryptocurrencyTrendingLatest, error) {
	requestURL := "/v1/cryptocurrency/trending/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyTrendingLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyTrendingMostVisited(q url.Values) (*CryptocurrencyTrendingMostVisited, error) {
	requestURL := "/v1/cryptocurrency/trending/most-visited"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyTrendingMostVisited

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyMarketPairs(q url.Values) (*CryptocurrencyMarketPairsLatest, error) {
	requestURL := "/v2/cryptocurrency/market-pairs/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyMarketPairsLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyOHLCVHistorical(q url.Values) (*CryptocurrencyOHLCVHistorical, error) {
	requestURL := "/v2/cryptocurrency/ohlcv/historical"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyOHLCVHistorical

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyOHLCVLatest(q url.Values) (*CryptocurrencyOHLCVLatest, error) {
	requestURL := "/v2/cryptocurrency/ohlcv/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyOHLCVLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyPricePerformanceStats(q url.Values) (*CryptocurrencyPricePerformanceStats, error) {
	requestURL := "/v2/cryptocurrency/price-performance-stats/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyPricePerformanceStats

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyQuotesHistorical(q url.Values) (*CryptocurrencyQuotesHistorical, error) {
	requestURL := "/v2/cryptocurrency/quotes/historical"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyQuotesHistorical

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCryptocurrencyQuotesLatest(q url.Values) (*CryptocurrencyQuotesLatest, error) {
	requestURL := "/v2/cryptocurrency/quotes/latest"

	resBody, err := doGetRequest(requestURL, q, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse CryptocurrencyQuotesLatest

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
