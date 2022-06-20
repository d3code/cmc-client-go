package cmc

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
)

const (
    cmcBaseUrl        = "https://pro-api.coinmarketcap.com"
    cmcBaseTestnetUrl = "https://sandbox-api.coinmarketcap.com"
)

type client struct {
    apiKey        string
    baseURL       string
    printResponse bool
}

func Client(apiKey string) *client {
    c := &client{
        apiKey:        apiKey,
        baseURL:       cmcBaseUrl,
        printResponse: false,
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

func (c *client) PrintResponse(test bool) *client {
    if test {
        c.printResponse = true
        return c
    }

    c.printResponse = false
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
    req, err := http.NewRequest(http.MethodGet, requestURL, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accepts", "application/json")
    req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
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
