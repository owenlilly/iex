package iex

import "testing"

const (
	publishable = "REPLACE_WITH_PUBLISHABLE_TOKEN"
	secret      = "REPLACE_WITH_SECRET_TOKEN"
)

func TestNewClient(t *testing.T) {
	got := NewClient(publishable, secret, nil)
	if got != nil {
		t.Logf("got:\t%v", got)
	}
}

func TestAdvancedStats(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.AdvancedStats("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestBalanceSheet(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.BalanceSheet("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestCashFlow(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.CashFlow("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestCompany(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Company("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestDelayedQuote(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.DelayedQuote("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestDividends(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Dividends("aapl", Range5Y)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEarnings(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Earnings("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEarningsToday(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.EarningsToday()
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEstimates(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Estimates("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestFundOwnership(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.FundOwnership("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestIncomeStmt(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.IncomeStmt("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderRoster(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderRoster("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderSum(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderSum("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderTXN(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderTXN("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInstlOwnership(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InstlOwnership("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestIPOCal(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.IPOCal(IPOCalTypeToday)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestKeyStats(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.KeyStats("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestLgstTrades(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.LgstTrades("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestList(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.List(ListGainers, nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestLogo(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Logo("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestMarketVol(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.MarketVol()
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestNews(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.News("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestOHLC(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.OHLC("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}
func TestPeers(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Peers("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestPrevDayPrice(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.PrevDayPrice("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestPrice(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Price("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestPriceTgt(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.PriceTgt("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestQuote(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Quote("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestSecPerf(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.SecPerf()
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestSplits(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Splits("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestVolByVenue(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.VolByVenue("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}
