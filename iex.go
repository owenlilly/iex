package iex

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// Constants
const (
	IPOCalTypeToday    = "today-ipos"
	IPOCalTypeUpcoming = "upcoming-ipos"
	ListMostActive     = "mostactive"
	ListGainers        = "gainers"
	ListLosers         = "losers"
	ListIEXVolume      = "iexvolume"
	ListIEXPercent     = "iexpercent"
	PeriodAnnual       = "annual"
	PeriodQuarter      = "quarter"
	Range5Y            = "5y"
	Range2Y            = "2y"
	Range1Y            = "1y"
	RangeYTD           = "ytd"
	Range6M            = "6m"
	Range3M            = "3m"
	Range1M            = "1m"
	RangeNext          = "next"
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
	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(string(body))
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

// Dividends ...
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

// Earnings returns the earnings data for a given company including the actual
// EPS, consensus, and fiscal period. Earnings are available quarterly (last
// 4 quarters).
func (c *Client) Earnings(sym string, opts *EarningsOpts) (*Earnings, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "earnings")
	if opts != nil && opts.Last > 0 {
		p = path.Join(p, strconv.FormatUint(opts.Last, 10))
	}
	if opts != nil && opts.Field != "" {
		p = path.Join(p, opts.Field)
	}
	v := Earnings{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// EarningsToday returns the earnings data for a given company including the actual
// EPS, consensus, and fiscal period. Earnings are available quarterly (last
// 4 quarters).
func (c *Client) EarningsToday() (*EarningsToday, error) {
	p := path.Join("stock", "market", "today-earnings")
	v := EarningsToday{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Estimates provides the latest consensus estimate for the next fiscal period.
func (c *Client) Estimates(sym string) (*Estimates, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "estimates")
	v := Estimates{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// FundOwnership returns the top 10 fund holders, meaning any firm not defined
// as buy-side or sell-side such as mutual funds, pension funds, endowments,
// investment firms, and other large entities that manage funds on behalf of
// others.
func (c *Client) FundOwnership(sym string) (*FundOwnership, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "fund-ownership")
	v := FundOwnership{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// IncomeStmt pulls income statement data. Available quarterly or
// annually with the default being the last available quarter.
func (c *Client) IncomeStmt(sym string, opts *IncomeStmtOpts) (*IncomeStmt, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "income")
	vs := url.Values{}
	if opts != nil && opts.Period != "" {
		vs.Set("period", opts.Period)
	}
	if opts != nil && opts.Last != "" {
		vs.Set("last", opts.Last)
	}
	v := IncomeStmt{}
	err := c.get(p, &vs, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// InsiderRoster returns the top 10 insiders, with the most recent information.
func (c *Client) InsiderRoster(sym string) (*InsiderRoster, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "insider-roster")
	v := InsiderRoster{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// InsiderSum returns aggregated insiders summary data for the last 6 months.
func (c *Client) InsiderSum(sym string) (*InsiderSum, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "insider-summary")
	v := InsiderSum{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// InsiderTXN returns aggregated insiders summary data for the last 6 months.
func (c *Client) InsiderTXN(sym string) (*InsiderTXN, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "insider-transactions")
	v := InsiderTXN{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// InstlOwnership returns the top 10 institutional holders, defined as buy-
// side or sell-side firms.
func (c *Client) InstlOwnership(sym string) (*InstlOwnership, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "institutional-ownership")
	v := InstlOwnership{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// IntradayPrices returns aggregated intraday prices in one minute buckets.
func (c *Client) IntradayPrices(sym string) (*IntradayPrices, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "intraday-prices")
	v := IntradayPrices{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// IPOCal returns a list of upcoming or today IPOs scheduled for the current
// and next month. The response is split into two structures: rawData and
// viewData. rawData represents all available data for an IPO. viewData
// represents data structured for display to a user.
func (c *Client) IPOCal(t string) (*IPOCal, error) {
	p := path.Join("stock", "market")
	switch t {
	case IPOCalTypeToday:
		p = path.Join(p, IPOCalTypeToday)
	case IPOCalTypeUpcoming:
		p = path.Join(p, IPOCalTypeUpcoming)
	default:
		return nil, errors.New("invalid argument: t")
	}
	v := IPOCal{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// KeyStats ...
func (c *Client) KeyStats(sym string) (*KeyStats, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "stats")
	v := KeyStats{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// LgstTrades returns 15 minute delayed, last sale eligible trades.
func (c *Client) LgstTrades(sym string) (*LgstTrades, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "largest-trades")
	v := LgstTrades{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// List ...
func (c *Client) List(t string, opts *ListOpts) (*List, error) {
	p := path.Join("stock", t)
	switch t {
	case ListMostActive, ListGainers, ListLosers, ListIEXPercent, ListIEXVolume:
		p = path.Join(p, t)
	default:
		return nil, errors.New("invalid argument: t")
	}
	vs := url.Values{}
	if opts != nil && opts.DisplayPercent {
		vs.Set("displayPercent", "true")
	}
	if opts != nil && opts.ListLimit > 0 {
		vs.Set("listLimit", strconv.FormatUint(opts.ListLimit, 10))
	}
	v := List{}
	err := c.get(p, &vs, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Logo ...
func (c *Client) Logo(sym string) (*Logo, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "logo")
	v := Logo{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// MarketVol ...
func (c *Client) MarketVol() (*MarketVol, error) {
	p := path.Join("stock", "market", "volume")
	v := MarketVol{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// News ...
func (c *Client) News(sym string) (*News, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "volume")
	v := News{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// OHLC ...
func (c *Client) OHLC(sym string) (*OHLC, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "volume")
	v := OHLC{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Peers ...
func (c *Client) Peers(sym string) (*Peers, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "peers")
	v := Peers{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// PrevDayPrice ...
func (c *Client) PrevDayPrice(sym string) (*PrevDayPrice, error) {
	if sym == "" {
		sym = "market"
	}
	p := path.Join("stock", sym, "previous")
	v := PrevDayPrice{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Price ...
func (c *Client) Price(sym string) (*Price, error) {
	if sym == "" {
		sym = "market"
	}
	p := path.Join("stock", sym, "previous")
	var v Price
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// PriceTgt ...
func (c *Client) PriceTgt(sym string) (*PriceTgt, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "previous")
	v := PriceTgt{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Quote ...
func (c *Client) Quote(sym string) (*Quote, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "previous")
	v := Quote{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// SecPerf ...
func (c *Client) SecPerf() (*SecPerf, error) {
	p := path.Join("stock", "market", "sector-performance")
	v := SecPerf{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// Splits ...
func (c *Client) Splits(sym string) (*Splits, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "splits")
	v := Splits{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// VolByVenue ...
func (c *Client) VolByVenue(sym string) (*VolByVenue, error) {
	if sym == "" {
		return nil, errors.New("invalid argument: sym")
	}
	p := path.Join("stock", sym, "volume-by-venue")
	v := VolByVenue{}
	err := c.get(p, nil, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
