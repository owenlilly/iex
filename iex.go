package iex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// Constants
const (
	PeriodAnnual  = "annual"
	PeriodQuarter = "quarter"
)

// Client is an IEX client.
type Client struct {
	publishable string
	secret      string
	version     string
	httpClient  *http.Client
}

// NewClient returns IEX client.
func NewClient(publishable string, secret string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		publishable: publishable,
		secret:      secret,
		httpClient:  httpClient,
	}
}

func (c *Client) url(p string, vs *url.Values) *url.URL {
	if vs == nil {
		vs = &url.Values{}
	}
	vs.Set("token", c.secret)
	return &url.URL{
		Scheme:   "https",
		Host:     "sandbox.iexapis.com",
		Path:     path.Join("stable", p),
		RawQuery: vs.Encode(),
	}
}

func (c *Client) get(p string, vs *url.Values, v interface{}) error {
	url := c.url(p, vs)
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// AdvancedStats returns everything in key stats plus additional advanced stats
// such as EBITDA, ratios, key financial data, and more.
func (c *Client) AdvancedStats(sym string) (*AdvancedStats, error) {
	v := AdvancedStats{}
	p := fmt.Sprintf("/stock/%v/advanced-stats", sym)
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// BalanceSheet pulls balance sheet data. Available quarterly or annually with
// the default being the last available quarter.
func (c *Client) BalanceSheet(sym string, opts *BalanceSheetOpts) (*BalanceSheet, error) {
	v := BalanceSheet{}
	vs := url.Values{}
	if opts != nil && opts.Period != "" {
		vs.Set("period", opts.Period)
	}
	if opts != nil && opts.Last != "" {
		vs.Set("last", opts.Last)
	}
	p := fmt.Sprintf("/stock/%v/balance-sheet", sym)
	err := c.get(p, &vs, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
