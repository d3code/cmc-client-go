package cmc

type ExchangeMapResponse struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	IsActive            int    `json:"is_active"`
	FirstHistoricalData string `json:"first_historical_data"`
	LastHistoricalData  string `json:"last_historical_data"`
}

type ExchangeMapResponseWrapper struct {
	Status Status                `json:"status"`
	Data   []ExchangeMapResponse `json:"data"`
}
