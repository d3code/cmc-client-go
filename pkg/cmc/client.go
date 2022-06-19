package cmc

const (
	cmcBaseUrl        = "https://pro-api.coinmarketcap.com/v1"
	cmcBaseTestnetUrl = "https://pro-api.coinmarketcap.com/v1"
)

type Client struct {
	ApiKey  string
	BaseURL string
}

func NewClient(apiKey string, testnet bool) *Client {
	c := &Client{
		ApiKey:  apiKey,
		BaseURL: cmcBaseUrl,
	}

	if testnet {
		c.BaseURL = cmcBaseTestnetUrl
	}

	return c
}
