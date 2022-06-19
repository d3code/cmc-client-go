package cmc

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

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

	return resBody, nil
}
