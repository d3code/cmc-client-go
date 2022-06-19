package cmc

import "time"

type ExchangeMapResponse struct {
	Data []struct {
		Id                  int       `json:"id"`
		Name                string    `json:"name"`
		Slug                string    `json:"slug"`
		IsActive            int       `json:"is_active"`
		Status              string    `json:"status"`
		FirstHistoricalData time.Time `json:"first_historical_data"`
		LastHistoricalData  time.Time `json:"last_historical_data"`
	} `json:"data"`
	Status Status `json:"status"`
}

type ExchangeInfoResponseWrapper struct {
	Data map[string]struct {
		Id                    int           `json:"id"`
		Name                  string        `json:"name"`
		Slug                  string        `json:"slug"`
		Logo                  string        `json:"logo"`
		Description           string        `json:"description"`
		DateLaunched          time.Time     `json:"date_launched"`
		Notice                interface{}   `json:"notice"`
		Countries             []interface{} `json:"countries"`
		Fiats                 []string      `json:"fiats"`
		Tags                  interface{}   `json:"tags"`
		Type                  string        `json:"type"`
		MakerFee              float64       `json:"maker_fee"`
		TakerFee              float64       `json:"taker_fee"`
		WeeklyVisits          int           `json:"weekly_visits"`
		SpotVolumeUsd         float64       `json:"spot_volume_usd"`
		SpotVolumeLastUpdated time.Time     `json:"spot_volume_last_updated"`
		Urls                  struct {
			Website []string      `json:"website"`
			Twitter []string      `json:"twitter"`
			Blog    []interface{} `json:"blog"`
			Chat    []string      `json:"chat"`
			Fee     []string      `json:"fee"`
		} `json:"urls"`
	} `json:"data"`
	Status Status `json:"status"`
}
