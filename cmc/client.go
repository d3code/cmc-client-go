package cmc

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	cmcBaseUrl     = "https://pro-api.coinmarketcap.com"
	cmcBaseTestUrl = "https://sandbox-api.coinmarketcap.com"
	testApiKey     = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
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

func doGetRequest(requestURL string, q url.Values, c *client) ([]byte, error) {
	var request string

	if c.test {
		request = cmcBaseTestUrl + requestURL
	} else {
		request = cmcBaseUrl + requestURL
	}

	req, err := http.NewRequest(http.MethodGet, request, nil)
	if err != nil {
		return nil, err
	}

	if c.test {
		req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
	} else {
		req.Header.Set("X-CMC_PRO_API_KEY", testApiKey)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")
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
		println(responseString)
	}

	return resBody, nil
}
