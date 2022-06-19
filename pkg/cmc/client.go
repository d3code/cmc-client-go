package cmc

const (
	cmcBaseUrl        = "https://pro-api.coinmarketcap.com"
	cmcBaseTestnetUrl = "https://sandbox-api.coinmarketcap.com"
)

type client struct {
	apiKey  string
	baseURL string
}

func Client(apiKey string) *client {
	c := &client{
		apiKey:  apiKey,
		baseURL: cmcBaseUrl,
	}

	return c
}

func (c *client) Test(test bool) *client {
	if test {
		c.apiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
		c.baseURL = cmcBaseTestnetUrl
	}

	return c
}

type ResponseWrapper struct {
	Status map[string]interface{} `json:"status"`
	Data   interface{}            `json:"data"`
}
