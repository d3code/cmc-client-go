package cmc

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	cmcBaseUrl      = "https://pro-api.coinmarketcap.com"
	cmcBaseTestUrl  = "https://sandbox-api.coinmarketcap.com"
	cmcApiKeyHeader = "X-CMC_PRO_API_KEY"
	testApiKey      = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
)

type client struct {
	apiKey        string
	test          bool
	printResponse bool
}

func Client(apiKey string) *client {
	c := &client{
		apiKey:        apiKey,
		test:          false,
		printResponse: false,
	}

	return c
}

func (c *client) Test(test bool) *client {
	c.test = test
	return c
}

func (c *client) PrintResponse(printResponse bool) *client {
	c.printResponse = printResponse
	return c
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
	Notice       string    `json:"notice"`
}

func doGetRequest(requestUrl string, q url.Values, c *client) ([]byte, error) {
	var baseUrl string

	if c.test {
		baseUrl = cmcBaseTestUrl + requestUrl
	} else {
		baseUrl = cmcBaseUrl + requestUrl
	}

	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return nil, err
	}

	if c.test {
		req.Header.Set(cmcApiKeyHeader, testApiKey)
	} else {
		req.Header.Set(cmcApiKeyHeader, c.apiKey)
	}

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	res, requestError := httpClient.Do(req)

	if requestError != nil {
		return nil, requestError
	}

	resBody, responseError := ioutil.ReadAll(res.Body)
	if responseError != nil {
		return nil, responseError
	}

	if c.printResponse && resBody != nil {
		responseString := string(resBody)
		println("Response: ", responseString)
	}

	return resBody, nil
}
