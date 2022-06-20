package cmc

import "time"

type CryptocurrencyAirdrop struct {
	Data struct {
		Id          string `json:"id"`
		ProjectName string `json:"project_name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Coin        struct {
			Id     int    `json:"id"`
			Name   string `json:"name"`
			Slug   string `json:"slug"`
			Symbol string `json:"symbol"`
		} `json:"coin"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		TotalPrize  int64     `json:"total_prize"`
		WinnerCount int       `json:"winner_count"`
		Link        string    `json:"link"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyAirdrops struct {
	Data struct {
		Data []struct {
			Id          string `json:"id"`
			ProjectName string `json:"project_name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Coin        struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Slug   string `json:"slug"`
				Symbol string `json:"symbol"`
			} `json:"coin"`
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			TotalPrize  int       `json:"total_prize"`
			WinnerCount int       `json:"winner_count"`
			Link        string    `json:"link"`
		} `json:"data"`
		Status Status `json:"status"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyCategories struct {
	Status Status `json:"status"`
	Data   struct {
		Data []struct {
			Id              string  `json:"id"`
			Name            string  `json:"name"`
			Title           string  `json:"title"`
			Description     string  `json:"description"`
			NumTokens       int     `json:"num_tokens"`
			AvgPriceChange  float64 `json:"avg_price_change"`
			MarketCap       float64 `json:"market_cap"`
			MarketCapChange float64 `json:"market_cap_change"`
			Volume          float64 `json:"volume"`
			VolumeChange    float64 `json:"volume_change"`
			LastUpdated     int     `json:"last_updated"`
		} `json:"data"`
		Status Status `json:"status"`
	} `json:"data"`
}

type CryptocurrencyCategory struct {
	Data struct {
		Id              string  `json:"id"`
		Name            string  `json:"name"`
		Title           string  `json:"title"`
		Description     string  `json:"description"`
		NumTokens       int     `json:"num_tokens"`
		AvgPriceChange  float64 `json:"avg_price_change"`
		MarketCap       float64 `json:"market_cap"`
		MarketCapChange float64 `json:"market_cap_change"`
		Volume          float64 `json:"volume"`
		VolumeChange    float64 `json:"volume_change"`
		Coins           []struct {
			Id                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			CmcRank           int         `json:"cmc_rank,omitempty"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			CirculatingSupply int         `json:"circulating_supply"`
			TotalSupply       int         `json:"total_supply"`
			MaxSupply         int         `json:"max_supply"`
			LastUpdated       time.Time   `json:"last_updated"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			Platform          interface{} `json:"platform"`
			Quote             map[string]struct {
				Price            float64   `json:"price"`
				Volume24H        int64     `json:"volume_24h"`
				PercentChange1H  float64   `json:"percent_change_1h"`
				PercentChange24H float64   `json:"percent_change_24h"`
				PercentChange7D  float64   `json:"percent_change_7d"`
				MarketCap        int64     `json:"market_cap"`
				LastUpdated      time.Time `json:"last_updated"`
			} `json:"quote"`
		} `json:"coins"`
		LastUpdated int64 `json:"last_updated"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyMap struct {
	Data []struct {
		Id                  int       `json:"id"`
		Rank                int       `json:"rank"`
		Name                string    `json:"name"`
		Symbol              string    `json:"symbol"`
		Slug                string    `json:"slug"`
		IsActive            int       `json:"is_active"`
		FirstHistoricalData time.Time `json:"first_historical_data"`
		LastHistoricalData  time.Time `json:"last_historical_data"`
		Platform            *struct {
			Id           int    `json:"id"`
			Name         string `json:"name"`
			Symbol       string `json:"symbol"`
			Slug         string `json:"slug"`
			TokenAddress string `json:"token_address"`
		} `json:"platform"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyInfo struct {
	Data map[string]struct {
		Urls struct {
			Website      []string      `json:"website"`
			TechnicalDoc []string      `json:"technical_doc"`
			Twitter      []interface{} `json:"twitter"`
			Reddit       []string      `json:"reddit"`
			MessageBoard []string      `json:"message_board"`
			Announcement []interface{} `json:"announcement"`
			Chat         []interface{} `json:"chat"`
			Explorer     []string      `json:"explorer"`
			SourceCode   []string      `json:"source_code"`
		} `json:"urls"`
		Logo         string      `json:"logo"`
		Id           int         `json:"id"`
		Name         string      `json:"name"`
		Symbol       string      `json:"symbol"`
		Slug         string      `json:"slug"`
		Description  string      `json:"description"`
		DateAdded    time.Time   `json:"date_added"`
		DateLaunched time.Time   `json:"date_launched"`
		Tags         []string    `json:"tags"`
		Platform     interface{} `json:"platform"`
		Category     string      `json:"category"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyListingsHistorical struct {
	Data []struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		CmcRank           int         `json:"cmc_rank,omitempty"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		CirculatingSupply int         `json:"circulating_supply"`
		TotalSupply       int         `json:"total_supply"`
		MaxSupply         int         `json:"max_supply"`
		LastUpdated       time.Time   `json:"last_updated"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		Platform          interface{} `json:"platform"`
		Quote             map[string]struct {
			Price            float64   `json:"price"`
			Volume24H        int64     `json:"volume_24h"`
			PercentChange1H  float64   `json:"percent_change_1h"`
			PercentChange24H float64   `json:"percent_change_24h"`
			PercentChange7D  float64   `json:"percent_change_7d"`
			MarketCap        int64     `json:"market_cap"`
			LastUpdated      time.Time `json:"last_updated"`
		} `json:"quote"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyListingsLatest struct {
	Data []struct {
		Id                            int         `json:"id"`
		Name                          string      `json:"name"`
		Symbol                        string      `json:"symbol"`
		Slug                          string      `json:"slug"`
		CmcRank                       int         `json:"cmc_rank,omitempty"`
		NumMarketPairs                int         `json:"num_market_pairs"`
		CirculatingSupply             int         `json:"circulating_supply"`
		TotalSupply                   int         `json:"total_supply"`
		MaxSupply                     int         `json:"max_supply"`
		LastUpdated                   time.Time   `json:"last_updated"`
		DateAdded                     time.Time   `json:"date_added"`
		Tags                          []string    `json:"tags"`
		Platform                      interface{} `json:"platform"`
		SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
		SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
		Quote                         map[string]struct {
			Price                 float64   `json:"price"`
			Volume24H             int64     `json:"volume_24h"`
			VolumeChange24H       float64   `json:"volume_change_24h"`
			PercentChange1H       float64   `json:"percent_change_1h"`
			PercentChange24H      float64   `json:"percent_change_24h"`
			PercentChange7D       float64   `json:"percent_change_7d"`
			MarketCap             float64   `json:"market_cap"`
			MarketCapDominance    int       `json:"market_cap_dominance"`
			FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
			LastUpdated           time.Time `json:"last_updated"`
		} `json:"quote"`
	} `json:"data"`
	Status Status `json:"status"`
}

type CryptocurrencyListingsNew struct {
	Data []struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		CmcRank           int         `json:"cmc_rank,omitempty"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		CirculatingSupply int         `json:"circulating_supply"`
		TotalSupply       int         `json:"total_supply"`
		MaxSupply         int         `json:"max_supply"`
		LastUpdated       time.Time   `json:"last_updated"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		Platform          interface{} `json:"platform"`
		Quote             map[string]struct {
			Price                 float64   `json:"price"`
			Volume24H             int64     `json:"volume_24h"`
			VolumeChange24H       float64   `json:"volume_change_24h"`
			PercentChange1H       float64   `json:"percent_change_1h"`
			PercentChange24H      float64   `json:"percent_change_24h"`
			PercentChange7D       float64   `json:"percent_change_7d"`
			MarketCap             float64   `json:"market_cap"`
			MarketCapDominance    int       `json:"market_cap_dominance"`
			FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
			LastUpdated           time.Time `json:"last_updated"`
		} `json:"quote"`
	} `json:"data"`
	Status Status `json:"status"`
}
