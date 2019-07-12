package iex

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

// Constants
const (
	PeriodAnnual  = "annual"
	PeriodQuarter = "quarter"
	Range5Y       = "5y"
	Range2Y       = "2y"
	Range1Y       = "1y"
	RangeYTD      = "ytd"
	Range6M       = "6m"
	Range3M       = "3m"
	Range1M       = "1m"
	RangeNext     = "next"
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
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	v := AdvancedStats{}
	p := path.Join("stock", sym, "advanced-stats")
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// BalanceSheet pulls balance sheet data. Available quarterly or annually with
// the default being the last available quarter.
func (c *Client) BalanceSheet(sym string, opts *BalanceSheetOpts) (*BalanceSheet, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	v := BalanceSheet{}
	vs := url.Values{}
	if opts != nil && opts.Period != "" {
		vs.Set("period", opts.Period)
	}
	if opts != nil && opts.Last != "" {
		vs.Set("last", opts.Last)
	}
	p := path.Join("stock", sym, "balance-sheet")
	err := c.get(p, &vs, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// CashFlow pulls cash flow data. Available quarterly or annually, with the
// default being the last available quarter.
func (c *Client) CashFlow(sym string, opts *CashFlowOpts) (*CashFlow, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	v := CashFlow{}
	vs := url.Values{}
	if opts != nil && opts.Period != "" {
		vs.Set("period", opts.Period)
	}
	if opts != nil && opts.Last != "" {
		vs.Set("last", opts.Last)
	}
	p := path.Join("stock", sym, "cash-flow")
	err := c.get(p, &vs, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Company pulls cash flow data. Available quarterly or annually, with the
// default being the last available quarter.
func (c *Client) Company(sym string) (*Company, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	v := Company{}
	p := path.Join("stock", sym, "company")
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// DelayedQuote returns the 15 minute delayed market quote.
func (c *Client) DelayedQuote(sym string) (*DelayedQuote, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	v := DelayedQuote{}
	p := path.Join("stock", sym, "delayed-quote")
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Dividends returns the 15 minute delayed market quote.
func (c *Client) Dividends(sym string, rng string) (*Dividends, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	switch rng {
	case Range5Y, Range2Y, Range1Y, RangeYTD, Range6M, Range3M, Range1M, RangeNext:
	default:
		return nil, errors.New("invalid argument: rng")
	}
	v := Dividends{}
	p := path.Join("stock", sym, "dividends", rng)
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
