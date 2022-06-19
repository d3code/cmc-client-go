package cmc

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(apiKey string, testnet bool) *Client {
	cmcBaseUrl := "https://pro-api.coinmarketcap.com/v1"
	cmcBaseTestnetUrl := "https://sandbox-api.coinmarketcap.com/v1"

	c := &Client{
		apiKey:  apiKey,
		baseURL: cmcBaseUrl,
	}

	if testnet {
		c.apiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
		c.baseURL = cmcBaseTestnetUrl
	}

	return c
}
