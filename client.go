package xscan

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// predefined endpoints
const (
	etherscanUrl   = "https://api.etherscan.io/api"
	bscscanUrl     = "https://api.bscscan.com/api"
	polygonscanUrl = "https://api.polygonscan.com/api"
	ftmscanUrl     = "https://api.ftmscan.com/api"
)

type networkOpts struct {
	endpoint string
}

var (
	EthOpts = &networkOpts{
		endpoint: etherscanUrl,
	}
	BscOpts = &networkOpts{
		endpoint: bscscanUrl,
	}
	PolygonOpts = &networkOpts{
		endpoint: polygonscanUrl,
	}
	FtmOpts = &networkOpts{
		endpoint: ftmscanUrl,
	}
)

// NewClient is a constructor for the Client struct
func NewClient(netOpts *networkOpts, apiKey string) *Client {
	return &Client{
		endpoint: netOpts.endpoint,
		apiKey:   apiKey,
		http:     new(http.Client),
	}
}

// setQueryParam is a helper function to set query parameters to a url including the apikey
func setQueryParam(endpoint *string, module, action, apiKey string, params []map[string]interface{}) {
	*endpoint = fmt.Sprintf("%s?module=%s&action=%s", *endpoint, module, action)
	for _, param := range params {
		for i := range param {
			*endpoint = fmt.Sprintf("%s&%s=%v", *endpoint, i, param[i])
		}
	}
	*endpoint = fmt.Sprintf("%s&apikey=%s", *endpoint, apiKey)
}

// doRequest executes the request, the result get returned as a pointer to the caller. It also returns a possible error and the http status code
func (c *Client) doRequest(ctx context.Context, module, action string, expRes interface{}, opts ...map[string]interface{}) (int, error) {
	callURL := c.endpoint

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&callURL, module, action, c.apiKey, opts)
	}
	fmt.Println(callURL)
	req, err := http.NewRequestWithContext(ctx, "GET", callURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil

	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}
